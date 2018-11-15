package dao

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type FormOp struct {
	db         *db.Database
	collection string

	// Refers objects
	user *UserOp
}

func (op *FormOp) Insert(form *models.Form) error {
	if _, err := op.user.FindByUUID(form.OwnerUserUUID); err != nil {
		return err
	}

	form.ID = bson.NewObjectId()
	form.Active = false
	return op.db.Insert(op.collection, form)
}

func (op *FormOp) FindAll(query *models.Query) ([]models.Form, error) {
	result := []models.Form{}
	if err := op.db.FindAll(op.collection, query.ToBSON(), nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (op *FormOp) FindByID(id string) (*models.Form, error) {
	result := &models.Form{}
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	if err := op.db.FindOne(op.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (op *FormOp) Update(form *models.Form) error {
	return op.db.Update(op.collection, bson.M{"_id": form.ID}, form)
}

func (op *FormOp) RemoveByID(id string) error {
	return op.db.Remove(op.collection, bson.M{"_id": bson.ObjectIdHex(id)})
}
