package models

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

type UserRole string

// These are the valid role of a User.
const (
	RoleAdmin UserRole = "Admin"
	RoleUser  UserRole = "User"
)

type User struct {
	UUID     string   `bson:"uuid" json:"uuid"`
	Email    string   `bson:"email" json:"email"`
	Name     string   `bson:"name" json:"name"`
	Agency   string   `bson:"agency" json:"agency"`
	Unit     string   `bson:"unit" json:"unit"`
	JobTitle string   `bson:"jobTitle" json:"jobTitle"`
	Phone    string   `bson:"phone" json:"phone"`
	Active   bool     `bson:"active" json:"active"`
	Role     UserRole `bson:"role" json:"role"`
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

type Password struct {
	UserUUID string `bson:"userUUID" json:"userUUID"`
	Email    string `bson:"email" json:"email"`
	Secret   string `bson:"secret" json:"secret"`
}
