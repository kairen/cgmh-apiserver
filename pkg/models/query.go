package model

import (
	"github.com/globalsign/mgo/bson"
)

type Query struct {
	StartDate string `form:"startDate,omitempty"`
	EndDate   string `form:"endDate,omitempty"`
	UserUUID  string `form:"userUUID,omitempty"`
}

func (q *Query) ToBSON() bson.M {
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
