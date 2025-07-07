package models

import (

)

//* User model 
type User struct {
	ID       int        `json:"userId"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Resume   string     `json:"resume"`
	Files    []string   `json:"files"`
	Detail   UserDetail `json:"detail"`
}

type UserDetail struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phoneNumber"`
	University  string `json:"university"`
}

//* company model
type Company struct {
	ID           int			`json:"id"`
	Email        string			`json:"email"`
	Password     string			`json:"password"`
	JobPositions []JobPosition	`json:"jobPositions"`
	Verifications[]string		`json:"verifications"`
}

type JobPosition struct {
	ID     int    `json:"id"`
	Detail string `json:"detail"`
	Role   string `json:"role"`
}




