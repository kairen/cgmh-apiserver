package dao

import (
	"fmt"
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type UserOp struct {
	db         *db.Database
	collection string

	// Refers objects
	role     *UserRoleOp
	status   *UserStatusOp
	level    *UserLevelOp
	password *UserPasswordOp
	counter  *CounterOp
}

func (op *UserOp) Insert(user *models.User) error {
	id, err := op.counter.Increase("user-serial-id")
	if err != nil {
		return err
	}

	user.UUID = fmt.Sprintf("u%05d", id)
	if err := op.db.Insert(op.collection, user); err != nil {
		return err
	}

	stat := &models.UserStatus{UserUUID: user.UUID, Block: false, Active: false}
	if err := op.status.Insert(stat); err != nil {
		return err
	}

	role := &models.UserRole{UserUUID: user.UUID, Name: models.RoleUser}
	if err := op.role.Insert(role); err != nil {
		return err
	}

	level := &models.UserLevel{UserUUID: user.UUID, Name: models.LevelNone}
	if err := op.level.Insert(level); err != nil {
		return err
	}
	return nil
}

func (op *UserOp) getRelationalObjects(user *models.User) error {
	role, err := op.role.FindOne(user.UUID)
	if err != nil {
		return err
	}
	user.Role = role.Name

	stat, err := op.status.FindOne(user.UUID)
	if err != nil {
		return err
	}
	user.Active = stat.Active
	user.Block = stat.Block

	level, err := op.level.FindOne(user.UUID)
	if err != nil {
		return err
	}
	user.Level = level.Name
	return nil
}

func (op *UserOp) IsExistByEmail(email string) bool {
	return op.db.IsExist(op.collection, bson.M{"email": email})
}

func (op *UserOp) FindAll() ([]models.User, error) {
	result := []models.User{}
	if err := op.db.FindAll(op.collection, nil, nil, &result); err != nil {
		return nil, err
	}

	for index := range result {
		if err := op.getRelationalObjects(&result[index]); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (op *UserOp) FindByEmail(email string) (*models.User, error) {
	result := &models.User{}
	query := bson.M{"email": email}
	if err := op.db.FindOne(op.collection, query, nil, result); err != nil {
		return nil, err
	}

	if err := op.getRelationalObjects(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (op *UserOp) FindByUUID(uuid string) (*models.User, error) {
	result := &models.User{}
	query := bson.M{"uuid": uuid}
	if err := op.db.FindOne(op.collection, query, nil, result); err != nil {
		return nil, err
	}

	if err := op.getRelationalObjects(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (op *UserOp) Update(user *models.User) error {
	if err := op.db.Update(op.collection, bson.M{"uuid": user.UUID}, user); err != nil {
		return err
	}
	return op.getRelationalObjects(user)
}

func (op *UserOp) UpdateRole(role *models.UserRole) error {
	return op.role.Update(role)
}

func (op *UserOp) UpdateStatus(stat *models.UserStatus) error {
	return op.status.Update(stat)
}

func (op *UserOp) UpdateLevel(level *models.UserLevel) error {
	return op.level.Update(level)
}

func (op *UserOp) RemoveByUUID(uuid string) error {
	if err := op.db.Remove(op.collection, bson.M{"uuid": uuid}); err != nil {
		return err
	}

	if err := op.role.Remove(uuid); err != nil {
		return err
	}

	if err := op.status.Remove(uuid); err != nil {
		return err
	}
	return op.password.Remove(uuid)
}
