package models

import(
	"time"
)

type Company struct {
	CompanyID			int				`json:"company_id"`
	Email				string			`json:"email" binding:"required,email"`
	Password			string			`josn:"password" binding:"required"`
	JobPosition			[]JobPosition	`json:"job_positions"`
	IsVerify			bool			`json:"is_verify"`
	VerificationFile	string			`json:"verification_file"`
}

type JobPosition struct {
	JobPositionID	int			`json:"job_position_id"`
	Detail			string		`json:"detail"`
	Role			string		`json:"role"`
	StartDateTime	time.Time	`json:"start_date_time"`
	EndDateTime		time.Time	`json:"end_date_time"`
}

