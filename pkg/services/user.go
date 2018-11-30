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
	point    *UserPointService
	counter  *CounterService
}

func newUserService(db *db.Mongo) *UserService {
	user := &UserService{db: db, collection: CollectionUser}
	user.counter = newCounterService(db)
	user.password = newUserPasswordService(db)
	user.role = newUserRoleService(db)
	user.status = newUserStatusService(db)
	user.level = newUserLevelService(db)
	user.point = newUserPointService(db)
	return user
}

func (svc *UserService) Insert(user *model.User) error {
	id, err := svc.counter.Increase("user-serial-id")
	if err != nil {
		return err
	}

	user.UUID = fmt.Sprintf("u%05d", id)
	if err := svc.db.Insert(svc.collection, user); err != nil {
		return err
	}

	stat := &model.UserStatus{UserUUID: user.UUID, Block: false, Active: false}
	if err := svc.status.Insert(stat); err != nil {
		return err
	}

	role := &model.UserRole{UserUUID: user.UUID, Name: model.RoleUser}
	if err := svc.role.Insert(role); err != nil {
		return err
	}

	level := &model.UserLevel{UserUUID: user.UUID, Name: model.LevelNone}
	if err := svc.level.Insert(level); err != nil {
		return err
	}

	point := &model.Point{UserUUID: user.UUID, Value: 0}
	if err := svc.point.Insert(point); err != nil {
		return err
	}
	return nil
}

func (svc *UserService) getRelationalObjects(user *model.User) error {
	role, err := svc.role.FindByUserUUID(user.UUID)
	if err != nil {
		return err
	}
	user.Role = role.Name

	stat, err := svc.status.FindByUserUUID(user.UUID)
	if err != nil {
		return err
	}
	user.Active = stat.Active
	user.Block = stat.Block

	level, err := svc.level.FindByUserUUID(user.UUID)
	if err != nil {
		return err
	}
	user.Level = level.Name

	point, err := svc.point.FindByUserUUID(user.UUID)
	if err != nil {
		return err
	}
	user.Point = point.Value
	return nil
}

func (svc *UserService) IsExistByEmail(email string) bool {
	return svc.db.IsExist(svc.collection, bson.M{"email": email})
}

func (svc *UserService) FindAll() ([]model.User, error) {
	result := []model.User{}
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

func (svc *UserService) FindByEmail(email string) (*model.User, error) {
	result := &model.User{}
	query := bson.M{"email": email}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}

	if err := svc.getRelationalObjects(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *UserService) FindByUUID(uuid string) (*model.User, error) {
	result := &model.User{}
	query := bson.M{"uuid": uuid}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}

	if err := svc.getRelationalObjects(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *UserService) FindUserLevels(levelName string) ([]model.UserLevel, error) {
	return svc.level.FindAllByName(levelName)
}

func (svc *UserService) Update(user *model.User) error {
	if err := svc.db.Update(svc.collection, bson.M{"uuid": user.UUID}, user); err != nil {
		return err
	}
	return svc.getRelationalObjects(user)
}

func (svc *UserService) UpdateRole(role *model.UserRole) error {
	return svc.role.Update(role)
}

func (svc *UserService) UpdateStatus(stat *model.UserStatus) error {
	return svc.status.Update(stat)
}

func (svc *UserService) UpdateLevel(level *model.UserLevel) error {
	return svc.level.Update(level)
}

func (svc *UserService) UpdatePoint(point *model.Point) error {
	return svc.point.Update(point)
}

func (svc *UserService) UpdateLevelsByName(oldLevel, newLevel string) error {
	levels, err := svc.level.FindAllByName(oldLevel)
	if err != nil {
		return err
	}

	for _, level := range levels {
		level.Name = newLevel
		if err := svc.level.Update(&level); err != nil {
			return err
		}
	}
	return nil
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

	if err := svc.point.Remove(uuid); err != nil {
		return err
	}
	return svc.password.Remove(uuid)
}
