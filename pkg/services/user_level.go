package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type UserLevelService struct {
	db         *db.Mongo
	collection string
}

func newUserLevelService(db *db.Mongo) *UserLevelService {
	return &UserLevelService{db: db, collection: CollectionUserLevel}
}

func (svc *UserLevelService) Insert(level *models.UserLevel) error {
	return svc.db.Insert(svc.collection, level)
}

func (svc *UserLevelService) FindByUserUUID(uuid string) (*models.UserLevel, error) {
	result := &models.UserLevel{}
	query := bson.M{"userUUID": uuid}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *UserLevelService) FindAllByName(name string) ([]models.UserLevel, error) {
	result := []models.UserLevel{}
	query := bson.M{"name": name}
	if err := svc.db.FindAll(svc.collection, query, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *UserLevelService) Update(level *models.UserLevel) error {
	return svc.db.Update(svc.collection, bson.M{"userUUID": level.UserUUID}, level)
}

func (svc *UserLevelService) Remove(uuid string) error {
	return svc.db.Remove(svc.collection, bson.M{"userUUID": uuid})
}
