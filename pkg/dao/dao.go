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
	pwd := &UserPasswordOp{db: db, collection: CollectionUserPassword}
	level := &LevelOp{db: db, collection: CollectionLevel}
	userRole := &UserRoleOp{db: db, collection: CollectionUserRole}
	userStatus := &UserStatusOp{db: db, collection: CollectionUserStatus}
	user := &UserOp{
		db:         db,
		role:       userRole,
		status:     userStatus,
		password:   pwd,
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
