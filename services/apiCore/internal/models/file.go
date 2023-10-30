package models

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID              uuid.UUID  `json:"id"`
	UserID          uuid.UUID  `json:"user_id"`
	Name            string     `json:"name"`
	SuccessAccounts int        `json:"success_accounts"`
	FailAccounts    int        `json:"fail_accounts"`
	LoadingAccounts int        `json:"loading_accounts"`
	CreatedAt       time.Time  `json:"created_at"`
	EndTime         time.Time  `json:"end_time"`
	Status          FileStatus `json:"status"`
}

type FileStatus string

const SuccessStatus FileStatus = "success"
const FailStatus FileStatus = "fail"
const LoadingStatus FileStatus = "loading"

type FileResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
