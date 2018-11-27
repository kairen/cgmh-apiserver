package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"
)

type AuthService struct {
	db *db.Database

	// Refers service
	user *UserService
}

func newAuthService(db *db.Database, user *UserService) *AuthService {
	return &AuthService{db: db, user: user}
}

func (svc *AuthService) Register(user *models.User, password string) error {
	if err := svc.user.Insert(user); err != nil {
		return err
	}

	pwd := &models.UserPassword{UserUUID: user.UUID, Secret: password}
	if err := svc.user.password.Insert(pwd); err != nil {
		return err
	}
	return nil
}

func (svc *AuthService) VerifyAccount(email, password string) bool {
	user, err := svc.user.FindByEmail(email)
	if err != nil {
		return false
	}
	return svc.user.password.IsExist(user.UUID, password)
}

func (svc *AuthService) Reset(email, password string) error {
	user, err := svc.user.FindByEmail(email)
	if err != nil {
		return err
	}

	pwd := &models.UserPassword{
		UserUUID: user.UUID,
		Secret:   password,
	}
	return svc.user.password.Update(pwd)
}
