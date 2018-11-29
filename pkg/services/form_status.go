package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type FormStatusService struct {
	db         *db.Mongo
	collection string
}

func newFormStatusService(db *db.Mongo) *FormStatusService {
	return &FormStatusService{db: db, collection: CollectionFormStatus}
}

func (svc *FormStatusService) Insert(status *model.FormStatus) error {
	status.ValidateState()
	return svc.db.Insert(svc.collection, status)
}

func (svc *FormStatusService) FindByID(id string) (*model.FormStatus, error) {
	result := &model.FormStatus{}
	query := bson.M{"formID": bson.ObjectIdHex(id)}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *FormStatusService) Update(status *model.FormStatus) error {
	status.ValidateState()
	return svc.db.Update(svc.collection, bson.M{"formID": status.FormID}, status)
}

func (svc *FormStatusService) Remove(id string) error {
	return svc.db.Remove(svc.collection, bson.M{"formID": bson.ObjectIdHex(id)})
}
