package model

type Level struct {
	ID          string `bson:"_id" json:"id"`
	Name        string `bson:"name" json:"name"`
	GPUPrice    int    `bson:"gpuPrice" json:"gpuPrice"`
	DayPrice    int    `bson:"dayPrice" json:"dayPrice"`
	Description string `bson:"description" json:"description"`
	Default     bool   `bson:"default" json:"default"`
}

func (lv *Level) Validate() bool {
	if lv.Name == "" || lv.DayPrice <= 0 {
		return false
	}
	return true
}
