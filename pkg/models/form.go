package models

import (
	"github.com/globalsign/mgo/bson"
)

type Date struct {
	Start string `bson:"startDate" json:"startDate"`
	End   string `bson:"endDate" json:"endDate"`
}

type Form struct {
	ID                    bson.ObjectId `bson:"_id" json:"id"`
	OwnerUserUUID         string        `bson:"ownerUserUUID" json:"ownerUserUUID" binding:"required"`
	AccessTime            string        `bson:"accessTime" json:"accessTime" binding:"required"`
	Active                bool          `bson:"active" json:"active"`
	Resource              string        `bson:"resource" json:"resource"`
	ModelArchitecture     string        `bson:"modelArchitecture" json:"modelArchitecture"`
	ModelLayer            int           `bson:"modelLayer" json:"modelLayer"`
	OS                    string        `bson:"os" json:"os" binding:"required"`
	GPUModel              string        `bson:"gpuModel" json:"gpuModel"`
	EpochTime             string        `bson:"epochTime" json:"epochTime"`
	EpochSize             int           `bson:"epochSize" json:"epochSize"`
	BatchSize             int           `bson:"batchSize" json:"batchSize"`
	TrainingModelSource   string        `bson:"trainingModelSource" json:"trainingModelSource"`
	TrainingModelDiskSize int           `bson:"trainingModelDiskSize" json:"trainingModelDiskSize"`
	TrainingModelData     int           `bson:"trainingModelData" json:"trainingModelData"`
	TrainingEnvironment   string        `bson:"trainingEnvironment" json:"trainingEnvironment" binding:"required"`
	TrainingType          string        `bson:"trainingType" json:"trainingType"`
	DataType              string        `bson:"dataType" json:"dataType"`
	NumGPUs               int           `bson:"numGPUs" json:"numGPUs"`
	Language              string        `bson:"language" json:"language"`
	ProjectName           string        `bson:"projectName" json:"projectName"`
	FundingSource         string        `bson:"fundingSource" json:"fundingSource"`
	ProjectSchedule       Date          `bson:"projectSchedule" json:"projectSchedule"`
	ExpectTime            Date          `bson:"expectTime" json:"expectTime"`
}

func (o *Form) Insert(f *Form) error {
	userDAO := &User{}
	_, err := userDAO.FindByUUID(f.OwnerUserUUID)
	if err != nil {
		return err
	}
	f.ID = bson.NewObjectId()
	f.Active = false
	return Insert(CollectionForm, f)
}

func (o *Form) FindAll(q *Query) ([]Form, error) {
	result := []Form{}
	if err := FindAll(CollectionForm, q.toBSON(), nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (o *Form) FindByID(id string) (*Form, error) {
	result := &Form{}
	err := FindOne(CollectionForm, bson.M{"_id": bson.ObjectIdHex(id)}, nil, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *Form) Update(f *Form) error {
	return Update(CollectionForm, bson.M{"_id": f.ID}, f)
}

func (o *Form) RemoveByID(id string) error {
	return Remove(CollectionForm, bson.M{"_id": bson.ObjectIdHex(id)})
}
