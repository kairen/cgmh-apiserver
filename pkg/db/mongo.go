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

type Mongo struct {
	session *mgo.Session
	flag    *Flag
}

func New(flag *Flag) (*Mongo, error) {
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

	database := &Mongo{session: s, flag: flag}
	return database, nil
}

func (d *Mongo) connect(collection string) (*mgo.Session, *mgo.Collection) {
	s := d.session.Copy()
	c := s.DB(d.flag.DB).C(collection)
	return s, c
}

func (d *Mongo) Insert(collection string, docs ...interface{}) error {
	ms, c := d.connect(collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func (d *Mongo) IsExist(collection string, query interface{}) bool {
	ms, c := d.connect(collection)
	defer ms.Close()
	count, _ := c.Find(query).Count()
	return count > 0
}

func (d *Mongo) FindOne(collection string, query, selector, result interface{}) error {
	ms, c := d.connect(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func (d *Mongo) FindAll(collection string, query, selector, result interface{}) error {
	ms, c := d.connect(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func (d *Mongo) Pipe(collection string, pipeline, result interface{}) error {
	ms, c := d.connect(collection)
	defer ms.Close()
	return c.Pipe(pipeline).All(result)
}

func (d *Mongo) Update(collection string, query, update interface{}) error {
	ms, c := d.connect(collection)
	defer ms.Close()
	return c.Update(query, update)
}

func (d *Mongo) Remove(collection string, query interface{}) error {
	ms, c := d.connect(collection)
	defer ms.Close()
	return c.Remove(query)
}
