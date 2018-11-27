package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type UserRoleService struct {
	db         *db.Database
	collection string
}

func newUserRoleService(db *db.Database) *UserRoleService {
	return &UserRoleService{db: db, collection: CollectionUserRole}
}

func (svc *UserRoleService) Insert(role *models.UserRole) error {
	role.Validate()
	return svc.db.Insert(svc.collection, role)
}

func (svc *UserRoleService) FindOne(uuid string) (*models.UserRole, error) {
	result := &models.UserRole{}
	query := bson.M{"userUUID": uuid}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *UserRoleService) Update(role *models.UserRole) error {
	role.Validate()
	return svc.db.Update(svc.collection, bson.M{"userUUID": role.UserUUID}, role)
}

func (svc *UserRoleService) Remove(uuid string) error {
	return svc.db.Remove(svc.collection, bson.M{"userUUID": uuid})
}
