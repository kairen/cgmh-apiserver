package models

import "github.com/globalsign/mgo/bson"

type Counter struct {
	Name string `bson:"name" json:"name" binding:"required"`
	Seq  int    `bson:"seq" json:"seq" binding:"required"`
}

func (c *Counter) init() error {
	c.Seq = 0
	return Insert(CollectionCounter, c)
}

func (c *Counter) getValues() error {
	return FindOne(CollectionCounter, bson.M{"name": c.Name}, nil, &c)
}

func (c *Counter) increase() error {
	c.Seq++
	return Update(CollectionCounter, bson.M{"name": c.Name}, c)
}

func (c *Counter) Increase() error {
	if err := c.getValues(); err != nil {
		if err := c.init(); err != nil {
			return err
		}
	}

	if err := c.increase(); err != nil {
		return err
	}
	return nil
}
