package models

type User struct {
	UUID     string   `bson:"uuid" json:"uuid"`
	Email    string   `bson:"email" json:"email"`
	Name     string   `bson:"name" json:"name"`
	Agency   string   `bson:"agency" json:"agency"`
	Unit     string   `bson:"unit" json:"unit"`
	JobTitle string   `bson:"jobTitle" json:"jobTitle"`
	Phone    string   `bson:"phone" json:"phone"`
	Active   bool     `bson:"active,omitempty" json:"active"`
	Block    bool     `bson:"block,omitempty" json:"block"`
	Role     RoleType `bson:"role,omitempty" json:"role"`
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

type UserRole struct {
	UserUUID string   `bson:"userUUID" json:"userUUID" binding:"required"`
	Name     RoleType `bson:"name" json:"name" binding:"required"`
}

type RoleType string

// These are the valid role name of a User.
const (
	RoleAdmin RoleType = "Admin"
	RoleUser  RoleType = "User"
)

func (role *UserRole) Validate() {
	if role.Name != RoleAdmin && role.Name != RoleUser {
		role.Name = RoleUser
	}
}

type UserStatus struct {
	UserUUID string `bson:"userUUID" json:"userUUID" binding:"required"`
	Active   bool   `bson:"active" json:"active" binding:"required"`
	Block    bool   `bson:"block" json:"block" binding:"required"`
}

type Password struct {
	UserUUID string `bson:"userUUID" json:"userUUID"`
	Secret   string `bson:"secret" json:"secret"`
}
