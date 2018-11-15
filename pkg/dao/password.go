package dao

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type PasswordOp struct {
	db         *db.Database
	collection string
}

func (op *PasswordOp) Insert(pwd *models.Password) error {
	return op.db.Insert(op.collection, pwd)
}

func (op *PasswordOp) IsExist(email, passwd string) bool {
	return op.db.IsExist(op.collection, bson.M{"email": email, "secret": passwd})
}

func (op *PasswordOp) Remove(uuid string) error {
	return op.db.Remove(op.collection, bson.M{"userUUID": uuid})
}
