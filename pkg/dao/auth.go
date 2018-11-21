package dao

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"
)

type AuthOp struct {
	db *db.Database
	// Refers objects
	user *UserOp
}

func (op *AuthOp) Register(user *models.User, password string) error {
	if err := op.user.Insert(user); err != nil {
		return err
	}
	pwd := &models.Password{
		Email:    user.Email,
		UserUUID: user.UUID,
		Secret:   password,
	}
	if err := op.user.password.Insert(pwd); err != nil {
		return err
	}
	return nil
}

func (op *AuthOp) VerifyAccount(email, password string) bool {
	return op.user.password.IsExist(email, password)
}

func (op *AuthOp) Reset(email, password string) error {
	pwd := &models.Password{
		Email:  email,
		Secret: password,
	}
	return op.user.password.UpdateByEmail(pwd)
}
