package file

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"go-mail-sender/config"
	"go-mail-sender/services/apiCore/internal/models"
	"go-mail-sender/services/apiCore/internal/repository"
	"go-mail-sender/services/apiCore/internal/utils"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ProcessTask struct {
	Line   []string
	FileID uuid.UUID
}

type Semaphore struct {
	C chan struct{}
}

func (s *Semaphore) Acquire() {
	s.C <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.C
}

type FileService struct {
	fileRepository repository.FileRepository
	log            *logrus.Logger
	cfg            *config.Config
}

func NewFileService(fileRepository repository.FileRepository, log *logrus.Logger, cfg *config.Config) *FileService {
	return &FileService{
		fileRepository: fileRepository,
		log:            log,
		cfg:            cfg,
	}
}

func (s *FileService) GetFiles(ctx context.Context, userID uuid.UUID) ([]*models.File, error) {
	files, err := s.fileRepository.GetFiles(ctx, userID)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (s *FileService) GetFileByID(fileID, userID uuid.UUID) (*models.File, error) {
	file, err := s.fileRepository.FindFileByID(fileID, userID)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *FileService) CreateFile(file *multipart.FileHeader, userID uuid.UUID) (*models.FileResponse, error) {
	tempFileName := utils.GenerateUniqueFileName(file.Filename)
	tempFilePath := filepath.Join(s.cfg.UploadFolder, tempFileName)

	if err := s.saveFile(file, tempFilePath); err != nil {
		return nil, err
	}

	rows, err := utils.CountLinesInFile(tempFilePath)
	if err != nil {
		return nil, err
	}

	newFile, err := s.fileRepository.CreateFile(tempFileName, userID, rows)
	if err != nil {
		return nil, err
	}
	gCount, err := strconv.Atoi(s.cfg.GoroutineCount)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	bufferSize, err := strconv.Atoi(s.cfg.BufferSize)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	sem := s.createSemaphore(gCount)

	taskChannel := make(chan ProcessTask, bufferSize)
	defer close(taskChannel)

	reader, err := s.openCSVFile(tempFilePath)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	go func() {
		skipHeader := true
		for {
			line, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				s.handleReadError(err, newFile)
			}

			if skipHeader {
				skipHeader = false
				continue // pass the first line
			}

			sem.Acquire()
			wg.Add(1)
			go func(line []string, newFile *models.File) {
				defer wg.Done()
				defer sem.Release()
				s.processTask(line, newFile, userID)
			}(line, newFile)
		}

		wg.Wait()

		s.updateFileStatus(newFile)
	}()

	return &models.FileResponse{
		Name: tempFileName,
		ID:   newFile.ID,
	}, nil
}

func (s *FileService) createSemaphore(gCount int) *Semaphore {
	return &Semaphore{
		C: make(chan struct{}, gCount),
	}
}

func (s *FileService) saveFile(file *multipart.FileHeader, filePath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	return err
}

func (s *FileService) openCSVFile(filePath string) (*csv.Reader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return csv.NewReader(file), nil
}

func (s *FileService) handleReadError(err error, newFile *models.File) {
	if newFile == nil {
		s.log.Errorf("Error reading line: %v", err)
		return
	}

	s.log.Errorf("Error reading line: %v", err)
	newFile.LoadingAccounts--
	newFile.FailAccounts++

	if _, err := s.fileRepository.UpdateFile(newFile); err != nil {
		s.log.Errorf("Failed to update file: %v", err)
	}
}

func (s *FileService) processTask(line []string, newFile *models.File, userID uuid.UUID) {
	if newFile == nil {
		s.log.Errorf("Nil newFile encountered")
		return
	}

	if !utils.ValidateLine(line) {
		s.handleInvalidLine(line, newFile)
		return
	}

	accountEmail := line[2]

	exists, err := s.getAccountByEmail(accountEmail, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			s.createAccount(line, accountEmail, newFile, userID)
			return
		}
		s.log.Errorf("Error finding subscriber: %v", err)
		return
	}
	if exists != nil {
		s.handleAccountExists(accountEmail, newFile)
		return
	}
	s.createAccount(line, accountEmail, newFile, userID)
}

func (s *FileService) updateFileStatus(newFile *models.File) {
	if newFile.LoadingAccounts == 0 {
		newFile.Status = models.SuccessStatus
		newFile.EndTime = time.Now()
		if _, err := s.fileRepository.UpdateFile(newFile); err != nil {
			s.log.Errorf("Failed to update file: %v", err)
		}
	}
}

func (s *FileService) handleInvalidLine(line []string, newFile *models.File) {
	s.log.Errorf("Invalid line: %v", line)
	newFile.FailAccounts++
	newFile.LoadingAccounts--
	if _, err := s.fileRepository.UpdateFile(newFile); err != nil {
		s.log.Errorf("Failed to update file: %v", err)
	}
}

func (s *FileService) handleAccountExists(accountEmail string, newFile *models.File) {
	s.log.Errorf("Account already exists: %s", accountEmail)
	newFile.FailAccounts++
	newFile.LoadingAccounts--
	if _, err := s.fileRepository.UpdateFile(newFile); err != nil {
		s.log.Errorf("Failed to update file: %v", err)
	}
}

func (s *FileService) getAccountByEmail(accountEmail string, userID uuid.UUID) (*models.Subscriber, error) {
	url := fmt.Sprintf("%s:%s/api/v1/subscribers/%s?user_id=%s", s.cfg.AppDividerURL, s.cfg.AppDividerPort, accountEmail, userID)

	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var subscriber models.Subscriber
	err = json.NewDecoder(resp.Body).Decode(&subscriber)
	if err != nil {
		return nil, err
	}

	return &subscriber, nil
}

func (s *FileService) createAccount(line []string, accountEmail string, newFile *models.File, userID uuid.UUID) {
	subscriber := &models.Subscriber{
		Email:     accountEmail,
		FirstName: line[0],
		LastName:  line[1],
		UserID:    userID,
	}

	data, err := json.Marshal(subscriber)
	if err != nil {
		s.log.Errorf("Failed to marshal subscriber data: %v", err)
		return
	}
	url := fmt.Sprintf("%s:%s/api/v1/subscribers", s.cfg.AppDividerURL, s.cfg.AppDividerPort)
	client := &http.Client{}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		s.log.Errorf("Failed to send POST request to divider: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		s.log.Errorf("Failed to create subscriber in divider. Status code: %d", resp.StatusCode)
		return
	}

	s.log.Infof("Account created in divider: %s", accountEmail)

	newFile.SuccessAccounts++
	newFile.LoadingAccounts--

	if _, err := s.fileRepository.UpdateFile(newFile); err != nil {
		s.log.Errorf("Failed to update file: %v", err)
	}
}
