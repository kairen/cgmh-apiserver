package service

import (
	"fmt"
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type LevelService struct {
	db         *db.Mongo
	collection string
	counter    *CounterService
}

func newLevelService(db *db.Mongo) *LevelService {
	level := &LevelService{db: db, collection: CollectionLevel}
	level.counter = newCounterService(db)
	return level
}

func (svc *LevelService) Insert(level *model.Level) error {
	id, err := svc.counter.Increase("level-serial-id")
	if err != nil {
		return err
	}
	level.ID = fmt.Sprintf("lv%05d", id)
	return svc.db.Insert(svc.collection, level)
}

func (svc *LevelService) IsExist(id string) bool {
	return svc.db.IsExist(svc.collection, bson.M{"_id": id})
}

func (svc *LevelService) FindByID(id string) (*model.Level, error) {
	result := &model.Level{}
	query := bson.M{"_id": id}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *LevelService) FindDefault() (*model.Level, error) {
	result := &model.Level{}
	query := bson.M{"default": true}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *LevelService) FindAll() ([]model.Level, error) {
	result := []model.Level{}
	if err := svc.db.FindAll(svc.collection, nil, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *LevelService) Update(level *model.Level) error {
	return svc.db.Update(svc.collection, bson.M{"_id": level.ID}, level)
}

func (svc *LevelService) Remove(id string) error {
	return svc.db.Remove(svc.collection, bson.M{"_id": id})
}
