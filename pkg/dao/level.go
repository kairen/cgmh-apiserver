package dao

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type LevelOp struct {
	db         *db.Database
	collection string
}

func (op *LevelOp) Insert(level *models.Level) error {
	level.ID = bson.NewObjectId()
	return op.db.Insert(op.collection, level)
}

func (op *LevelOp) IsExistByName(name string) bool {
	return op.db.IsExist(op.collection, bson.M{"name": name})
}

func (op *LevelOp) FindAll() ([]models.Level, error) {
	result := []models.Level{}
	if err := op.db.FindAll(op.collection, nil, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (op *LevelOp) Update(level *models.Level) error {
	return op.db.Update(op.collection, bson.M{"_id": level.ID}, level)
}

func (op *LevelOp) Remove(id string) error {
	return op.db.Remove(op.collection, bson.M{"_id": bson.ObjectIdHex(id)})
}
