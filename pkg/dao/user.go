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
	password *PasswordOp
	counter  *CounterOp
}

func (op *UserOp) Register(user *models.User, password string) error {
	id, err := op.counter.Increase("user-serial-id")
	if err != nil {
		return err
	}

	user.UUID = fmt.Sprintf("u%05d", id)
	if err := op.db.Insert(op.collection, user); err != nil {
		return err
	}

	pwd := &models.Password{
		Email:    user.Email,
		UserUUID: user.UUID,
		Secret:   password,
	}
	if err := op.password.Insert(pwd); err != nil {
		return err
	}
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
	return result, nil
}

func (op *UserOp) FindByEmail(email string) (*models.User, error) {
	result := &models.User{}
	query := bson.M{"email": email}
	if err := op.db.FindOne(op.collection, query, nil, result); err != nil {
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
	return result, nil
}

func (op *UserOp) VerifyAccount(email, password string) bool {
	return op.password.IsExist(email, password)
}

func (op *UserOp) Update(user *models.User) error {
	return op.db.Update(op.collection, bson.M{"uuid": user.UUID}, user)
}

func (op *UserOp) RemoveByUUID(uuid string) error {
	if err := op.db.Remove(op.collection, bson.M{"uuid": uuid}); err != nil {
		return err
	}
	return op.password.Remove(uuid)
}
