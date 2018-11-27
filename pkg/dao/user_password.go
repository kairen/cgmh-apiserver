package dao

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type UserPasswordOp struct {
	db         *db.Database
	collection string
}

func (op *UserPasswordOp) Insert(pwd *models.UserPassword) error {
	return op.db.Insert(op.collection, pwd)
}

func (op *UserPasswordOp) IsExist(uuid, passwd string) bool {
	return op.db.IsExist(op.collection, bson.M{"userUUID": uuid, "secret": passwd})
}

func (op *UserPasswordOp) Update(pwd *models.UserPassword) error {
	return op.db.Update(op.collection, bson.M{"userUUID": pwd.UserUUID}, pwd)
}

func (op *UserPasswordOp) Remove(uuid string) error {
	return op.db.Remove(op.collection, bson.M{"userUUID": uuid})
}
