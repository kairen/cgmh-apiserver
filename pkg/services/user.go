package service

import (
	"fmt"
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/util"

	"github.com/globalsign/mgo/bson"
)

type UserService struct {
	db         *db.Mongo
	collection string
	password   *UserPasswordService
	counter    *CounterService
	level      *LevelService
	point      *PointService
}

func newUserService(db *db.Mongo, level *LevelService, point *PointService) *UserService {
	user := &UserService{db: db, collection: CollectionUser}
	user.counter = newCounterService(db)
	user.password = newUserPasswordService(db)
	user.level = level
	user.point = point
	return user
}

func (svc *UserService) getRelationalObjects(user *model.User) error {
	if user.LevelID == "" {
		return nil
	}

	level, err := svc.level.FindByID(user.LevelID)
	if err != nil {
		return err
	}
	user.Level = level.Name
	return nil
}

func (svc *UserService) Insert(user *model.User) error {
	level, err := svc.level.FindDefault()
	if err != nil {
		return err
	}

	id, err := svc.counter.Increase("user-serial-id")
	if err != nil {
		return err
	}

	user.UUID = fmt.Sprintf("u%05d", id)
	user.Default()
	user.LevelID = level.ID
	if err := svc.db.Insert(svc.collection, user); err != nil {
		return err
	}
	return nil
}

func (svc *UserService) IsExistByEmail(email string) bool {
	return svc.db.IsExist(svc.collection, bson.M{"email": email})
}

func (svc *UserService) FindAll() ([]model.User, error) {
	result := []model.User{}
	if err := svc.db.FindAll(svc.collection, nil, nil, &result); err != nil {
		return nil, err
	}

	for index := range result {
		if err := svc.getRelationalObjects(&result[index]); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (svc *UserService) FindByEmail(email string) (*model.User, error) {
	result := &model.User{}
	query := bson.M{"email": email}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}

	if err := svc.getRelationalObjects(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *UserService) FindByUUID(uuid string) (*model.User, error) {
	result := &model.User{}
	query := bson.M{"uuid": uuid}
	if err := svc.db.FindOne(svc.collection, query, nil, result); err != nil {
		return nil, err
	}

	if err := svc.getRelationalObjects(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *UserService) FindUsersByLevelID(levelID string) ([]model.User, error) {
	result := []model.User{}
	query := bson.M{"levelID": levelID}
	if err := svc.db.FindAll(svc.collection, query, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *UserService) Update(user *model.UserPost) error {
	return svc.db.Update(svc.collection, bson.M{"uuid": user.UUID}, user)
}

func (svc *UserService) UpdateRole(role *model.UserRole) error {
	d := &model.UserRole{Role: role.Role}
	return svc.db.Update(svc.collection, bson.M{"uuid": role.UserUUID}, d)
}

func (svc *UserService) UpdateStatus(stat *model.UserStatus) error {
	d := &model.UserStatus{Active: stat.Active, Block: stat.Block}
	return svc.db.Update(svc.collection, bson.M{"uuid": stat.UserUUID}, d)
}

func (svc *UserService) UpdateLevel(level *model.UserLevel) error {
	d := &model.UserLevel{LevelID: level.LevelID}
	return svc.db.Update(svc.collection, bson.M{"uuid": level.UserUUID}, d)
}

func (svc *UserService) UpdatePoint(point *model.Point, deposit int) error {
	d := &model.UserPoint{Point: deposit}
	query := bson.M{"uuid": point.UserUUID}
	if err := svc.db.Update(svc.collection, query, d); err != nil {
		return err
	}

	point.Time = util.NowTime()
	return svc.point.Insert(point)
}

func (svc *UserService) Remove(uuid string) error {
	if err := svc.db.Remove(svc.collection, bson.M{"uuid": uuid}); err != nil {
		return err
	}
	return svc.password.Remove(uuid)
}
