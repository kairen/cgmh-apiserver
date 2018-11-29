package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type FormService struct {
	db         *db.Mongo
	collection string
	user       *UserService
	status     *FormStatusService
}

func newFormService(db *db.Mongo, user *UserService) *FormService {
	fs := &FormService{db: db, collection: CollectionForm}
	fs.user = user
	fs.status = newFormStatusService(db)
	return fs
}

func (svc *FormService) getRelationalObjects(form *model.Form) error {
	status, err := svc.status.FindByID(form.ID.Hex())
	if err != nil {
		return err
	}
	form.State = status.State
	return nil
}

func (svc *FormService) Insert(form *model.Form) error {
	if _, err := svc.user.FindByUUID(form.UserUUID); err != nil {
		return err
	}

	form.ID = bson.NewObjectId()
	if err := svc.db.Insert(svc.collection, form); err != nil {
		return err
	}

	status := &model.FormStatus{FormID: form.ID, State: model.FormInactiveState}
	if err := svc.status.Insert(status); err != nil {
		return err
	}
	return nil
}

func (svc *FormService) FindAll(query *model.Query) ([]model.Form, error) {
	result := []model.Form{}
	if err := svc.db.FindAll(svc.collection, query.ToBSON(), nil, &result); err != nil {
		return nil, err
	}

	for index := range result {
		if err := svc.getRelationalObjects(&result[index]); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (svc *FormService) FindByID(id string) (*model.Form, error) {
	result := &model.Form{}
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}

	if err := svc.getRelationalObjects(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *FormService) Update(form *model.Form) error {
	if err := svc.db.Update(svc.collection, bson.M{"_id": form.ID}, form); err != nil {
		return err
	}
	return svc.getRelationalObjects(form)
}

func (svc *FormService) UpdateStatus(status *model.FormStatus) error {
	return svc.status.Update(status)
}

func (svc *FormService) Remove(id string) error {
	data := bson.M{"_id": bson.ObjectIdHex(id)}
	if err := svc.db.Remove(svc.collection, data); err != nil {
		return err
	}
	return svc.status.Remove(id)
}
