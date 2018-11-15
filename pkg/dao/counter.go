package dao

import (
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"

	"github.com/globalsign/mgo/bson"
)

type CounterOp struct {
	db         *db.Database
	collection string
}

func (op *CounterOp) init(c *models.Counter) error {
	c.Init()
	return op.db.Insert(op.collection, c)
}

func (op *CounterOp) getValues(c *models.Counter) error {
	return op.db.FindOne(op.collection, bson.M{"name": c.Name}, nil, &c)

}

func (op *CounterOp) inc(c *models.Counter) error {
	c.Inc()
	return op.db.Update(op.collection, bson.M{"name": c.Name}, c)
}

func (op *CounterOp) Increase(name string) (int, error) {
	c := &models.Counter{Name: name}
	if err := op.getValues(c); err != nil {
		if err := op.init(c); err != nil {
			return -1, err
		}
	}

	if err := op.inc(c); err != nil {
		return -1, err
	}
	return c.Seq, nil
}
