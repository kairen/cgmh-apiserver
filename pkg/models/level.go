package model

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

func (lv *Level) Validate() bool {
	if lv.Name == "" || lv.Day <= 0 || lv.NumberOfGPU <= 0 {
		return false
	}
	return true
}
