package models

type Counter struct {
	Name string `bson:"name" json:"name" binding:"required"`
	Seq  int    `bson:"seq" json:"seq" binding:"required"`
}

func (c *Counter) Init() {
	c.Seq = 0
}

func (c *Counter) Inc() {
	c.Seq++
}
