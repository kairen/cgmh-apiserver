package model

type User struct {
	UUID     string   `bson:"uuid" json:"uuid"`
	Email    string   `bson:"email" json:"email"`
	Name     string   `bson:"name" json:"name"`
	Agency   string   `bson:"agency" json:"agency"`
	Unit     string   `bson:"unit" json:"unit"`
	JobTitle string   `bson:"jobTitle" json:"jobTitle"`
	Phone    string   `bson:"phone" json:"phone"`
	Level    string   `bson:"level,omitempty" json:"level"`
	Point    int      `bson:"point,omitempty" json:"point"`
	Active   bool     `bson:"active,omitempty" json:"active"`
	Block    bool     `bson:"block,omitempty" json:"block"`
	Role     RoleType `bson:"role,omitempty" json:"role"`
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) Validate() bool {
	if u.UUID == "" || u.Email == "" || u.Name == "" {
		return false
	}
	return true
}

type RoleType string

// These are the valid role name of a User.
const (
	RoleAdmin RoleType = "Admin"
	RoleUser  RoleType = "User"
)

type UserRole struct {
	UserUUID string   `bson:"userUUID" json:"userUUID" binding:"required"`
	Name     RoleType `bson:"name" json:"name" binding:"required"`
}

func (ur *UserRole) Validate() bool {
	if ur.UserUUID == "" || ur.Name == "" {
		return false
	}
	return true
}

func (role *UserRole) ValidateRole() {
	if role.Name != RoleAdmin && role.Name != RoleUser {
		role.Name = RoleUser
	}
}

type UserStatus struct {
	UserUUID string `bson:"userUUID" json:"userUUID" binding:"required"`
	Active   bool   `bson:"active" json:"active" binding:"required"`
	Block    bool   `bson:"block" json:"block" binding:"required"`
}

func (us *UserStatus) Validate() bool {
	if us.UserUUID == "" {
		return false
	}
	return true
}

const LevelNone = "None"

type UserLevel struct {
	UserUUID string `bson:"userUUID" json:"userUUID" binding:"required"`
	Name     string `bson:"name" json:"name" binding:"required"`
}

func (ul *UserLevel) Validate() bool {
	if ul.UserUUID == "" || ul.Name == "" {
		return false
	}
	return true
}

type UserPassword struct {
	UserUUID string `bson:"userUUID" json:"userUUID"`
	Secret   string `bson:"secret" json:"secret"`
}
