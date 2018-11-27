package models

import (
	"github.com/globalsign/mgo/bson"
)

type Level struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	NumberOfGPU int           `bson:"numberOfGPU" json:"numberOfGPU"`
	Day         int           `bson:"day" json:"day"`
	Description string        `bson:"description" json:"description"`
	Default     bool          `bson:"default" json:"default"`
}
