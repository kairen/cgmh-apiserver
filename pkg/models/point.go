package model

type Point struct {
	UserUUID  string `bson:"userUUID" json:"userUUID" binding:"required"`
	AdminUUID string `bson:"adminUUID,omitempty" json:"adminUUID" binding:"required"`
	Value     int    `bson:"value" json:"value" binding:"required"`
	Time      string `bson:"time,omitempty" json:"time,omitempty"`
}

func (p *Point) Validate() bool {
	if p.UserUUID == "" || p.AdminUUID == "" || p.Value == 0 {
		return false
	}
	return true
}
