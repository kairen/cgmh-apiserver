package models

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var (
	host     = os.Getenv("MONGODB_HOST")
	source   = os.Getenv("MONGODB_SOURCE")
	user     = os.Getenv("MONGODB_USER")
	password = os.Getenv("MONGODB_PASSWORD")
	db       = os.Getenv("MONGODB_DB")
)

const (
	CollectionCounter = iota
	CollectionUser
	CollectionPassword
	CollectionForm
)

var collections = map[int]string{
	CollectionCounter:  "counter",
	CollectionUser:     "user",
	CollectionPassword: "password",
	CollectionForm:     "form",
}

var globalSession *mgo.Session

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{host},
		Source:   source,
		Username: user,
		Password: password,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln("create session error ", err)
	}
	globalSession = s
}

func connect(collection string) (*mgo.Session, *mgo.Collection) {
	s := globalSession.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

func getCollection(c int) string {
	return collections[c]
}

func Insert(collctionType int, docs ...interface{}) error {
	collection := getCollection(collctionType)
	ms, c := connect(collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func IsExist(collctionType int, query interface{}) bool {
	collection := getCollection(collctionType)
	ms, c := connect(collection)
	defer ms.Close()
	count, _ := c.Find(query).Count()
	return count > 0
}

func FindOne(collctionType int, query, selector, result interface{}) error {
	collection := getCollection(collctionType)
	ms, c := connect(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(collctionType int, query, selector, result interface{}) error {
	collection := getCollection(collctionType)
	ms, c := connect(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func Pipe(collctionType int, pipeline, result interface{}) error {
	collection := getCollection(collctionType)
	ms, c := connect(collection)
	defer ms.Close()
	return c.Pipe(pipeline).All(result)
}

func Update(collctionType int, query, update interface{}) error {
	collection := getCollection(collctionType)
	ms, c := connect(collection)
	defer ms.Close()
	return c.Update(query, update)
}

func Remove(collctionType int, query interface{}) error {
	collection := getCollection(collctionType)
	ms, c := connect(collection)
	defer ms.Close()
	return c.Remove(query)
}

type Query struct {
	StartDate string `form:"startDate,omitempty"`
	EndDate   string `form:"endDate,omitempty"`
	UserUUID  string `form:"userUUID,omitempty"`
}

func (q *Query) toBSON() bson.M {
	b := bson.M{}
	if q.UserUUID != "" {
		b["ownerUserUUID"] = q.UserUUID
	}

	if q.StartDate != "" {
		b["startDate"] = q.StartDate
	}

	if q.EndDate != "" {
		b["endDate"] = q.EndDate
	}
	return b
}
