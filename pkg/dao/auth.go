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

	pwd := &models.Password{UserUUID: user.UUID, Secret: password}
	if err := op.user.password.Insert(pwd); err != nil {
		return err
	}
	return nil
}

func (op *AuthOp) VerifyAccount(email, password string) bool {
	user, err := op.user.FindByEmail(email)
	if err != nil {
		return false
	}
	return op.user.password.IsExist(user.UUID, password)
}

func (op *AuthOp) Reset(email, password string) error {
	user, err := op.user.FindByEmail(email)
	if err != nil {
		return err
	}

	pwd := &models.Password{
		UserUUID: user.UUID,
		Secret:   password,
	}
	return op.user.password.Update(pwd)
}
