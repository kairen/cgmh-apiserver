package model

type User struct {
	UUID     string   `bson:"uuid" json:"uuid"`
	Email    string   `bson:"email" json:"email"`
	Name     string   `bson:"name" json:"name"`
	Agency   string   `bson:"agency" json:"agency"`
	Unit     string   `bson:"unit" json:"unit"`
	JobTitle string   `bson:"jobTitle" json:"jobTitle"`
	Phone    string   `bson:"phone" json:"phone"`
	LevelID  string   `bson:"levelID" json:"levelID"`
	Level    string   `bson:"level,omitempty" json:"level,omitempty"`
	Point    int      `bson:"point" json:"point"`
	Active   bool     `bson:"active" json:"active"`
	Block    bool     `bson:"block" json:"block"`
	Role     RoleType `bson:"role" json:"role"`
}

func (u *User) Default() {
	u.Point = 0
	u.Active = false
	u.Block = false
	u.Role = RoleUser
	u.LevelID = ""
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

type UserPost struct {
	UUID     string `bson:"uuid" json:"uuid"`
	Name     string `bson:"name" json:"name"`
	Agency   string `bson:"agency" json:"agency"`
	Unit     string `bson:"unit" json:"unit"`
	JobTitle string `bson:"jobTitle" json:"jobTitle"`
	Phone    string `bson:"phone" json:"phone"`
}

func (up *UserPost) Validate() bool {
	return !(up.Name == "")
}

type RoleType string

// These are the valid role name of a User.
const (
	RoleAdmin RoleType = "Admin"
	RoleUser  RoleType = "User"
)

type UserRole struct {
	UserUUID string   `bson:"userUUID,omitempty" json:"userUUID,omitempty" binding:"required"`
	Role     RoleType `bson:"role" json:"role" binding:"required"`
}

func (ur *UserRole) Validate() bool {
	if ur.Role == "" {
		return false
	}
	return true
}

func (role *UserRole) ValidateRole() {
	if role.Role != RoleAdmin && role.Role != RoleUser {
		role.Role = RoleUser
	}
}

type UserStatus struct {
	UserUUID string `bson:"userUUID,omitempty" json:"userUUID,omitempty" binding:"required"`
	Active   bool   `bson:"active" json:"active" binding:"required"`
	Block    bool   `bson:"block" json:"block" binding:"required"`
}

const LevelNone = "None"

type UserLevel struct {
	UserUUID string `bson:"userUUID,omitempty" json:"userUUID,omitempty" binding:"required"`
	LevelID  string `bson:"levelID" json:"levelID" binding:"required"`
}

func (ul *UserLevel) Validate() bool {
	if ul.UserUUID == "" || ul.LevelID == "" {
		return false
	}
	return true
}

type UserPoint struct {
	UserUUID string `bson:"userUUID,omitempty" json:"userUUID,omitempty"`
	Point    int    `bson:"point" json:"point"`
}

type UserPassword struct {
	UserUUID string `bson:"userUUID" json:"userUUID"`
	Secret   string `bson:"secret" json:"secret"`
}
