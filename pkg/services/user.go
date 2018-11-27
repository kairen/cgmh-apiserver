package service

import (
	"fmt"
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type UserService struct {
	db         *db.Mongo
	collection string

	// Refers objects
	role     *UserRoleService
	status   *UserStatusService
	level    *UserLevelService
	password *UserPasswordService
	counter  *CounterService
}

func newUserService(db *db.Mongo) *UserService {
	user := &UserService{db: db, collection: CollectionUser}
	user.counter = newCounterService(db)
	user.password = newUserPasswordService(db)
	user.role = newUserRoleService(db)
	user.status = newUserStatusService(db)
	user.level = newUserLevelService(db)
	return user
}

func (svc *UserService) Insert(user *models.User) error {
	id, err := svc.counter.Increase("user-serial-id")
	if err != nil {
		return err
	}

	user.UUID = fmt.Sprintf("u%05d", id)
	if err := svc.db.Insert(svc.collection, user); err != nil {
		return err
	}

	stat := &models.UserStatus{UserUUID: user.UUID, Block: false, Active: false}
	if err := svc.status.Insert(stat); err != nil {
		return err
	}

	role := &models.UserRole{UserUUID: user.UUID, Name: models.RoleUser}
	if err := svc.role.Insert(role); err != nil {
		return err
	}

	level := &models.UserLevel{UserUUID: user.UUID, Name: models.LevelNone}
	if err := svc.level.Insert(level); err != nil {
		return err
	}
	return nil
}

func (svc *UserService) getRelationalObjects(user *models.User) error {
	role, err := svc.role.FindOne(user.UUID)
	if err != nil {
		return err
	}
	user.Role = role.Name

	stat, err := svc.status.FindOne(user.UUID)
	if err != nil {
		return err
	}
	user.Active = stat.Active
	user.Block = stat.Block

	level, err := svc.level.FindOne(user.UUID)
	if err != nil {
		return err
	}
	user.Level = level.Name
	return nil
}

func (svc *UserService) IsExistByEmail(email string) bool {
	return svc.db.IsExist(svc.collection, bson.M{"email": email})
}

func (svc *UserService) FindAll() ([]models.User, error) {
	result := []models.User{}
	if err := svc.db.FindAll(svc.collection, nil, nil, &result); err != nil {
		return nil, err
	}

	for index := range result {
		if err := svc.getRelationalObjects(&result[index]); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (svc *UserService) FindByEmail(email string) (*models.User, error) {
	result := &models.User{}
	query := bson.M{"email": email}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}

	if err := svc.getRelationalObjects(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *UserService) FindByUUID(uuid string) (*models.User, error) {
	result := &models.User{}
	query := bson.M{"uuid": uuid}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}

	if err := svc.getRelationalObjects(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *UserService) Update(user *models.User) error {
	if err := svc.db.Update(svc.collection, bson.M{"uuid": user.UUID}, user); err != nil {
		return err
	}
	return svc.getRelationalObjects(user)
}

func (svc *UserService) UpdateRole(role *models.UserRole) error {
	return svc.role.Update(role)
}

func (svc *UserService) UpdateStatus(stat *models.UserStatus) error {
	return svc.status.Update(stat)
}

func (svc *UserService) UpdateLevel(level *models.UserLevel) error {
	return svc.level.Update(level)
}

func (svc *UserService) RemoveByUUID(uuid string) error {
	if err := svc.db.Remove(svc.collection, bson.M{"uuid": uuid}); err != nil {
		return err
	}

	if err := svc.role.Remove(uuid); err != nil {
		return err
	}

	if err := svc.status.Remove(uuid); err != nil {
		return err
	}
	return svc.password.Remove(uuid)
}
