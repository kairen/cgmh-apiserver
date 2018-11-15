package db

import (
	"time"

	"github.com/globalsign/mgo"
)

type Flag struct {
	Host     string
	Source   string
	User     string
	Password string
	DB       string
}

type Database struct {
	session *mgo.Session
	flag    *Flag
}

func New(flag *Flag) (*Database, error) {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{flag.Host},
		Source:   flag.Source,
		Username: flag.User,
		Password: flag.Password,
		Timeout:  3 * time.Second,
	}

	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}

	database := &Database{session: s, flag: flag}
	return database, nil
}

func (d *Database) connect(collection string) (*mgo.Session, *mgo.Collection) {
	s := d.session.Copy()
	c := s.DB(d.flag.DB).C(collection)
	return s, c
}

func (d *Database) Insert(collection string, docs ...interface{}) error {
	ms, c := d.connect(collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func (d *Database) IsExist(collection string, query interface{}) bool {
	ms, c := d.connect(collection)
	defer ms.Close()
	count, _ := c.Find(query).Count()
	return count > 0
}

func (d *Database) FindOne(collection string, query, selector, result interface{}) error {
	ms, c := d.connect(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func (d *Database) FindAll(collection string, query, selector, result interface{}) error {
	ms, c := d.connect(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func (d *Database) Pipe(collection string, pipeline, result interface{}) error {
	ms, c := d.connect(collection)
	defer ms.Close()
	return c.Pipe(pipeline).All(result)
}

func (d *Database) Update(collection string, query, update interface{}) error {
	ms, c := d.connect(collection)
	defer ms.Close()
	return c.Update(query, update)
}

func (d *Database) Remove(collection string, query interface{}) error {
	ms, c := d.connect(collection)
	defer ms.Close()
	return c.Remove(query)
}
