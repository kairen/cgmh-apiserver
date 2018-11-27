package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type UserStatusService struct {
	db         *db.Mongo
	collection string
}

func newUserStatusService(db *db.Mongo) *UserStatusService {
	return &UserStatusService{db: db, collection: CollectionUserStatus}
}

func (svc *UserStatusService) Insert(status *models.UserStatus) error {
	return svc.db.Insert(svc.collection, status)
}

func (svc *UserStatusService) FindOne(uuid string) (*models.UserStatus, error) {
	result := &models.UserStatus{}
	query := bson.M{"userUUID": uuid}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *UserStatusService) Update(stat *models.UserStatus) error {
	return svc.db.Update(svc.collection, bson.M{"userUUID": stat.UserUUID}, stat)
}

func (svc *UserStatusService) Remove(uuid string) error {
	return svc.db.Remove(svc.collection, bson.M{"userUUID": uuid})
}
