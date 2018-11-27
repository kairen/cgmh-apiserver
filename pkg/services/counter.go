package service

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type CounterService struct {
	db         *db.Database
	collection string
}

func newCounterService(db *db.Database) *CounterService {
	return &CounterService{db: db, collection: CollectionCounter}
}

func (svc *CounterService) init(c *models.Counter) error {
	c.Init()
	return svc.db.Insert(svc.collection, c)
}

func (svc *CounterService) getValues(c *models.Counter) error {
	return svc.db.FindOne(svc.collection, bson.M{"name": c.Name}, nil, &c)

}

func (svc *CounterService) inc(c *models.Counter) error {
	c.Inc()
	return svc.db.Update(svc.collection, bson.M{"name": c.Name}, c)
}

func (svc *CounterService) Increase(name string) (int, error) {
	c := &models.Counter{Name: name}
	if err := svc.getValues(c); err != nil {
		if err := svc.init(c); err != nil {
			return -1, err
		}
	}

	if err := svc.inc(c); err != nil {
		return -1, err
	}
	return c.Seq, nil
}
