package models

import (

)

type User struct {
	UserID		int		`json:"user_id"`
	Email		string	`json:"email" binding:"required,email"`
	Password	string	`josn:"password" binding:"required"`
	Resume		string	`json:"resume"`
	Files		[]string`json:"files"`
	Detail 		Detail	`json:"detail" binding:"required"`
}

type Detail struct {
	Name	string		`json:"name" binding:"required"`
	Role	[]string	`json:"role" binding:"required"`
	PhoneNumber	string	`json:"phone_number" binding:"required"`
	University	string	`json:"university" binding:"required"`
}