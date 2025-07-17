package models

import (
	"time"
)

type JobInterview struct {
	JobInterviewID	int
	JobPositionID	int
	UserID			int
	CompanyID		int
	Detail			string
	DateTime		time.Time
	CreatedAt		time.Time
}