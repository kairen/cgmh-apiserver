package dao

import (
	"inwinstack/cgmh/apiserver/pkg/db"
)

const (
	CollectionCounter    = "Counter"
	CollectionUser       = "User"
	CollectionUserRole   = "UserRole"
	CollectionUserStatus = "UserStatus"
	CollectionPassword   = "Password"
	CollectionForm       = "Form"
)

type DataAccess struct {
	db *db.Database

	// Data access objects
	Auth *AuthOp
	User *UserOp
	Form *FormOp
}

func New(db *db.Database) *DataAccess {
	da := &DataAccess{db: db}
	// Init data access objects
	counter := &CounterOp{db: db, collection: CollectionCounter}
	pwd := &PasswordOp{db: db, collection: CollectionPassword}
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
	auth := &AuthOp{db: db, user: user}
	form := &FormOp{db: db, user: user, collection: CollectionForm}

	// Assign data access objects
	da.User = user
	da.Form = form
	da.Auth = auth
	return da
}
