package dao

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type UserLevelOp struct {
	db         *db.Database
	collection string
}

func (op *UserLevelOp) Insert(level *models.UserLevel) error {
	return op.db.Insert(op.collection, level)
}

func (op *UserLevelOp) FindOne(uuid string) (*models.UserLevel, error) {
	result := &models.UserLevel{}
	query := bson.M{"userUUID": uuid}
	if err := op.db.FindOne(op.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (op *UserLevelOp) Update(level *models.UserLevel) error {
	return op.db.Update(op.collection, bson.M{"userUUID": level.UserUUID}, level)
}

func (op *UserLevelOp) Remove(uuid string) error {
	return op.db.Remove(op.collection, bson.M{"userUUID": uuid})
}
