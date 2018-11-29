package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/util"
	"log"
)

const (
	CollectionCounter      = "Counter"
	CollectionUser         = "User"
	CollectionUserRole     = "UserRole"
	CollectionUserStatus   = "UserStatus"
	CollectionUserPassword = "UserPassword"
	CollectionUserLevel    = "UserLevel"
	CollectionUserPoint    = "UserPoint"
	CollectionLevel        = "Level"
	CollectionForm         = "Form"
	CollectionFormStatus   = "FormStatus"
	CollectionPointHistory = "PointHistory"
)

type DataAccess struct {
	db *db.Mongo

	// Access services
	Auth  *AuthService
	User  *UserService
	Level *LevelService
	Form  *FormService
	Point *PointService
}

func New(db *db.Mongo) *DataAccess {
	da := &DataAccess{db: db}
	user := newUserService(db)
	da.User = user
	da.Auth = newAuthService(db, user)
	da.Form = newFormService(db, user)
	da.Level = newLevelService(db)
	da.Point = newPointService(db)
	return da
}

func (svc *DataAccess) InitAdminUser() error {
	hex, err := util.RandomHex(8)
	if err != nil {
		return err
	}

	pwd := util.GetEnv("INIT_ADMIN_PASSWORD", hex)
	secret := util.MD5Encode(pwd)
	reg := &model.User{
		Email: util.GetEnv("INIT_ADMIN_EMAIL", "admin@inwinstack.com"),
		Name:  "administrator",
	}

	if !svc.User.IsExistByEmail(reg.Email) {
		if err := svc.Auth.Register(reg, secret); err != nil {
			return err
		}

		user, err := svc.User.FindByEmail(reg.Email)
		if err != nil {
			return err
		}
		stat := &model.UserStatus{UserUUID: user.UUID, Block: false, Active: true}
		if err := svc.User.UpdateStatus(stat); err != nil {
			return err
		}

		role := &model.UserRole{UserUUID: user.UUID, Name: model.RoleAdmin}
		if err := svc.User.UpdateRole(role); err != nil {
			return err
		}

		log.Printf("Admin email: %s", reg.Email)
		log.Printf("Admin password: %s", pwd)
	}
	return nil
}
