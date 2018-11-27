package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type FormService struct {
	db         *db.Mongo
	collection string

	// Refers service
	user *UserService
}

func newFormService(db *db.Mongo, user *UserService) *FormService {
	return &FormService{db: db, user: user, collection: CollectionForm}
}

func (svc *FormService) Insert(form *models.Form) error {
	if _, err := svc.user.FindByUUID(form.OwnerUserUUID); err != nil {
		return err
	}

	form.ID = bson.NewObjectId()
	form.Active = false
	return svc.db.Insert(svc.collection, form)
}

func (svc *FormService) FindAll(query *models.Query) ([]models.Form, error) {
	result := []models.Form{}
	if err := svc.db.FindAll(svc.collection, query.ToBSON(), nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *FormService) FindByID(id string) (*models.Form, error) {
	result := &models.Form{}
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *FormService) Update(form *models.Form) error {
	return svc.db.Update(svc.collection, bson.M{"_id": form.ID}, form)
}

func (svc *FormService) RemoveByID(id string) error {
	return svc.db.Remove(svc.collection, bson.M{"_id": bson.ObjectIdHex(id)})
}
