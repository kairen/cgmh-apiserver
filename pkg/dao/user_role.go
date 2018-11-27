package dao

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type UserRoleOp struct {
	db         *db.Database
	collection string
}

func (op *UserRoleOp) Insert(role *models.UserRole) error {
	role.Validate()
	return op.db.Insert(op.collection, role)
}

func (op *UserRoleOp) FindOne(uuid string) (*models.UserRole, error) {
	result := &models.UserRole{}
	query := bson.M{"userUUID": uuid}
	if err := op.db.FindOne(op.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (op *UserRoleOp) Update(role *models.UserRole) error {
	role.Validate()
	return op.db.Update(op.collection, bson.M{"userUUID": role.UserUUID}, role)
}

func (op *UserRoleOp) Remove(uuid string) error {
	return op.db.Remove(op.collection, bson.M{"userUUID": uuid})
}
