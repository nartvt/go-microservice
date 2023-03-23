package models

import "time"

type UserRepo struct {
	Id          int
	UserName    string
	Password    string
	Email       string
	PhoneNumber string
	FullName    string
	Role        int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
