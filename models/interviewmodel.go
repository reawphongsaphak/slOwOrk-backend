package models

import (
	"time"
)

type JobInterview struct {
	ID            int       `json:"id"`
	JobPositionID int       `json:"jobPositionId"`
	UserID        int       `json:"userId"`
	CompanyID     int       `json:"companyId"`
	Detail        string    `json:"detail"`
	Date          time.Time `json:"date"`
}

