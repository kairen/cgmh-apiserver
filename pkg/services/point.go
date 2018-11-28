package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"
)

type PointService struct {
	db         *db.Mongo
	collection string
}

func newPointService(db *db.Mongo) *PointService {
	return &PointService{db: db, collection: CollectionPointHistory}
}

func (svc *PointService) Insert(point *model.Point) error {
	return svc.db.Insert(svc.collection, point)
}

func (svc *PointService) FindAll(query *model.Query) ([]model.Point, error) {
	result := []model.Point{}
	if err := svc.db.FindAll(svc.collection, query.ToBSON(), nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
