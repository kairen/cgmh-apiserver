package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type LevelService struct {
	db         *db.Mongo
	collection string
}

func newLevelService(db *db.Mongo) *LevelService {
	return &LevelService{db: db, collection: CollectionLevel}
}

func (svc *LevelService) Insert(level *model.Level) error {
	level.ID = bson.NewObjectId()
	return svc.db.Insert(svc.collection, level)
}

func (svc *LevelService) IsExistByName(name string) bool {
	return svc.db.IsExist(svc.collection, bson.M{"name": name})
}

func (svc *LevelService) FindAll() ([]model.Level, error) {
	result := []model.Level{}
	if err := svc.db.FindAll(svc.collection, nil, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *LevelService) FindByID(id string) (*model.Level, error) {
	result := &model.Level{}
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *LevelService) Update(level *model.Level) error {
	return svc.db.Update(svc.collection, bson.M{"_id": level.ID}, level)
}

func (svc *LevelService) Remove(id string) error {
	return svc.db.Remove(svc.collection, bson.M{"_id": bson.ObjectIdHex(id)})
}
