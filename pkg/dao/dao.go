package dao

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

	// Data access objects
	Auth  *AuthOp
	User  *UserOp
	Level *LevelOp
	Form  *FormOp
}

func New(db *db.Database) *DataAccess {
	da := &DataAccess{db: db}
	// Init data access objects
	counter := &CounterOp{db: db, collection: CollectionCounter}
	userPassword := &UserPasswordOp{db: db, collection: CollectionUserPassword}
	level := &LevelOp{db: db, collection: CollectionLevel}
	userRole := &UserRoleOp{db: db, collection: CollectionUserRole}
	userStatus := &UserStatusOp{db: db, collection: CollectionUserStatus}
	userLevel := &UserLevelOp{db: db, collection: CollectionUserLevel}
	user := &UserOp{
		db:         db,
		role:       userRole,
		status:     userStatus,
		level:      userLevel,
		password:   userPassword,
		counter:    counter,
		collection: CollectionUser,
	}

	// Assign data access objects
	da.User = user
	da.Form = &FormOp{db: db, user: user, collection: CollectionForm}
	da.Auth = &AuthOp{db: db, user: user}
	da.Level = level
	return da
}
