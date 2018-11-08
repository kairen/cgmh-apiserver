package models

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
)

type User struct {
	UUID    string   `bson:"uuid" json:"uuid"`
	Email   string   `bson:"email" json:"email"`
	IsAdmin bool     `bson:"isAdmin" json:"isAdmin"`
	Active  bool     `bson:"active" json:"active"`
	Roles   []string `bson:"roles,omitempty" json:"roles,omitempty"`
}

type Password struct {
	UserUUID string `bson:"userUUID" json:"userUUID"`
	Email    string `bson:"email" json:"email"`
	Secret   string `bson:"secret" json:"secret"`
}

func (o *User) Insert(u *User, pwd *Password) error {
	counter := Counter{Name: "user-serial-id"}
	if err := counter.Increase(); err != nil {
		return err
	}

	u.UUID = fmt.Sprintf("u%05d", counter.Seq)
	if err := Insert(CollectionUser, u); err != nil {
		return err
	}

	pwd.UserUUID = u.UUID
	pwd.Email = u.Email
	if err := Insert(CollectionPassword, pwd); err != nil {
		return err
	}
	return nil
}

func (o *User) IsExistByEmail(email string) bool {
	return IsExist(CollectionUser, bson.M{"email": email})
}

func (o *User) FindAll() ([]User, error) {
	result := []User{}
	if err := FindAll(CollectionUser, nil, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (o *User) FindByEmail(email string) (*User, error) {
	result := &User{}
	err := FindOne(CollectionUser, bson.M{"email": email}, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *User) FindByUUID(uuid string) (*User, error) {
	result := &User{}
	err := FindOne(CollectionUser, bson.M{"uuid": uuid}, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *User) VerifyAccount(email, passwd string) bool {
	return IsExist(CollectionPassword, bson.M{"email": email, "secret": passwd})
}

func (o *User) Update(u *User) error {
	return Update(CollectionUser, bson.M{"uuid": u.UUID}, u)
}

func (o *User) RemoveByUUID(uuid string) error {
	if err := Remove(CollectionUser, bson.M{"uuid": uuid}); err != nil {
		return err
	}

	if err := Remove(CollectionPassword, bson.M{"userUUID": uuid}); err != nil {
		return err
	}
	return nil
}
