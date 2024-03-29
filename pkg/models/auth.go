package model

import (
	"regexp"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (lo *Login) Validate() bool {
	if lo.Email == "" || lo.Password == "" {
		return false
	}
	return true
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

func (reg *Register) Validate() bool {
	if reg.Email == "" || reg.Password == "" {
		return false
	}
	return true
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

func (r *Reset) Validate() bool {
	if r.Email == "" || r.OldPassword == "" || r.NewPassword == "" {
		return false
	}
	return true
}

type ForceReset struct {
	Email string `json:"email" binding:"required"`
}

func (fr *ForceReset) Validate() bool {
	if fr.Email == "" {
		return false
	}
	return true
}
