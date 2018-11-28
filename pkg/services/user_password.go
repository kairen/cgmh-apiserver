package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type UserPasswordService struct {
	db         *db.Mongo
	collection string
}

func newUserPasswordService(db *db.Mongo) *UserPasswordService {
	return &UserPasswordService{db: db, collection: CollectionUserPassword}
}

func (svc *UserPasswordService) Insert(pwd *model.UserPassword) error {
	return svc.db.Insert(svc.collection, pwd)
}

func (svc *UserPasswordService) IsExist(uuid, passwd string) bool {
	return svc.db.IsExist(svc.collection, bson.M{"userUUID": uuid, "secret": passwd})
}

func (svc *UserPasswordService) Update(pwd *model.UserPassword) error {
	return svc.db.Update(svc.collection, bson.M{"userUUID": pwd.UserUUID}, pwd)
}

func (svc *UserPasswordService) Remove(uuid string) error {
	return svc.db.Remove(svc.collection, bson.M{"userUUID": uuid})
}
