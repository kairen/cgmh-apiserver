package dao

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type UserStatusOp struct {
	db         *db.Database
	collection string
}

func (op *UserStatusOp) Insert(status *models.UserStatus) error {
	return op.db.Insert(op.collection, status)
}

func (op *UserStatusOp) FindOne(uuid string) (*models.UserStatus, error) {
	result := &models.UserStatus{}
	query := bson.M{"userUUID": uuid}
	if err := op.db.FindOne(op.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (op *UserStatusOp) Update(stat *models.UserStatus) error {
	return op.db.Update(op.collection, bson.M{"userUUID": stat.UserUUID}, stat)
}

func (op *UserStatusOp) Remove(uuid string) error {
	return op.db.Remove(op.collection, bson.M{"userUUID": uuid})
}
