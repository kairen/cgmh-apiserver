package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
)

const (
	CollectionCounter      = "Counter"
	CollectionUser         = "User"
	CollectionUserRole     = "UserRole"
	CollectionUserStatus   = "UserStatus"
	CollectionUserPassword = "UserPassword"
	CollectionUserLevel    = "UserLevel"
	CollectionLevel        = "Level"
	CollectionForm         = "Form"
)

type DataAccess struct {
	db *db.Database

	// Access services
	Auth  *AuthService
	User  *UserService
	Level *LevelService
	Form  *FormService
}

func New(db *db.Database) *DataAccess {
	da := &DataAccess{db: db}
	user := newUserService(db)
	da.User = user
	da.Auth = newAuthService(db, user)
	da.Form = newFormService(db, user)
	da.Level = newLevelService(db)
	return da
}
