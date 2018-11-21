package models

import (
	"regexp"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Agency   string `json:"agency" binding:"required"`
	Unit     string `json:"unit" binding:"required"`
	JobTitle string `json:"jobTitle" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

func (reg *Register) ValidateEmail() bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(reg.Email)
}

func (reg *Register) ToUser() *User {
	return &User{
		Email:    reg.Email,
		Name:     reg.Name,
		Agency:   reg.Agency,
		Unit:     reg.Unit,
		JobTitle: reg.JobTitle,
		Phone:    reg.Phone,
	}
}

type Reset struct {
	Email       string `json:"email" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type ForceReset struct {
	Email string `json:"email" binding:"required"`
}
