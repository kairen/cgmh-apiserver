package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	model "inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type UserPointService struct {
	db         *db.Mongo
	collection string
}

func newUserPointService(db *db.Mongo) *UserPointService {
	return &UserPointService{db: db, collection: CollectionUserPoint}
}

func (svc *UserPointService) Insert(point *model.Point) error {
	return svc.db.Insert(svc.collection, point)
}

func (svc *UserPointService) FindByUserUUID(uuid string) (*model.Point, error) {
	result := &model.Point{}
	query := bson.M{"userUUID": uuid}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *UserPointService) Update(point *model.Point) error {
	return svc.db.Update(svc.collection, bson.M{"userUUID": point.UserUUID}, point)
}

func (svc *UserPointService) Remove(uuid string) error {
	return svc.db.Remove(svc.collection, bson.M{"userUUID": uuid})
}
