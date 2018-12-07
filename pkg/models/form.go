package model

import (
	"github.com/globalsign/mgo/bson"
)

type FormDate struct {
	Start string `bson:"startDate" json:"startDate"`
	End   string `bson:"endDate" json:"endDate"`
}

type FormCharge struct {
	Expect int `bson:"expect" json:"expect"`
	Actual int `bson:"actual" json:"actual"`
}

type FormState string

// These are the valid state of a Form.
const (
	FormRejectState   FormState = "Reject"
	FormInactiveState FormState = "Inactive"
	FormActiveState   FormState = "Active"
)

type Form struct {
	ID                    bson.ObjectId `bson:"_id" json:"id"`
	UserUUID              string        `bson:"userUUID" json:"userUUID" binding:"required"`
	ContactName           string        `bson:"contactName" json:"contactName" binding:"required"`
	ContactEmail          string        `bson:"contactEmail" json:"contactEmail" binding:"required"`
	ContactPhone          string        `bson:"contactPhone" json:"contactPhone" binding:"required"`
	ProjectName           string        `bson:"projectName" json:"projectName" binding:"required"`
	IrbID                 string        `bson:"irbID" json:"irbID" binding:"required"`
	FundingSource         string        `bson:"fundingSource" json:"fundingSource" binding:"required"`
	ModelArchitecture     string        `bson:"modelArchitecture" json:"modelArchitecture"`
	ModelLayer            int           `bson:"modelLayer" json:"modelLayer"`
	OS                    string        `bson:"os" json:"os"`
	GPUModel              string        `bson:"gpuModel" json:"gpuModel"`
	EpochTime             string        `bson:"epochTime" json:"epochTime"`
	EpochSize             int           `bson:"epochSize" json:"epochSize"`
	BatchSize             int           `bson:"batchSize" json:"batchSize"`
	TrainingModelSource   string        `bson:"trainingModelSource" json:"trainingModelSource" binding:"required"`
	TrainingModelDiskSize int           `bson:"trainingModelDiskSize" json:"trainingModelDiskSize" binding:"required"`
	TrainingEnvironment   string        `bson:"trainingEnvironment" json:"trainingEnvironment" binding:"required"`
	TrainingModelData     int           `bson:"trainingModelData" json:"trainingModelData"`
	TrainingType          string        `bson:"trainingType" json:"trainingType"`
	DataType              string        `bson:"dataType" json:"dataType"`
	NumberOfGPU           int           `bson:"numberOfGPU" json:"numberOfGPU"`
	Language              string        `bson:"language" json:"language" binding:"required"`
	ProjectSchedule       FormDate      `bson:"projectSchedule" json:"projectSchedule" binding:"required"`
	ExpectTime            FormDate      `bson:"expectTime" json:"expectTime" binding:"required"`
	Charge                FormCharge    `bson:"charge" json:"charge" binding:"required"`
	Reason                string        `bson:"reason" json:"reason" binding:"required"`
	State                 FormState     `bson:"state" json:"state"`
	CreationTime          string        `bson:"creationTime,omitempty" json:"creationTime,omitempty"`
	LastUpdateTime        string        `bson:"lastUpdateTime,omitempty" json:"lastUpdateTime,omitempty"`
}

func (f *Form) Validate() bool {
	if f.ContactName == "" || f.ContactEmail == "" || f.ContactPhone == "" || f.NumberOfGPU < 0 {
		return false
	}
	return true
}

func (f *Form) CanUpdate() bool {
	return f.State == FormInactiveState
}

type FormStatus struct {
	FormID bson.ObjectId `bson:"formID" json:"formID" binding:"required"`
	State  FormState     `bson:"state" json:"state" binding:"required"`
}

func (fs *FormStatus) Validate() bool {
	if fs.State == "" {
		return false
	}
	return true
}

func (fs *FormStatus) ValidateState() {
	if fs.State != FormRejectState &&
		fs.State != FormInactiveState &&
		fs.State != FormActiveState {
		fs.State = FormInactiveState
	}
}
