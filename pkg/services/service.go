package service

import (
	"encoding/json"
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/util"
	"log"

	"github.com/globalsign/mgo/bson"
)

const (
	CollectionCounter      = "Counter"
	CollectionInit         = "Config"
	CollectionUser         = "User"
	CollectionUserPassword = "UserPassword"
	CollectionLevel        = "Level"
	CollectionForm         = "Form"
	CollectionFormStatus   = "FormStatus"
	CollectionPointHistory = "PointHistory"
)

const configName = "init-config"

type Config struct {
	Name  string `bson:"name"`
	Admin bool   `bson:"admin"`
	Level bool   `bson:"level"`
}

type DataAccess struct {
	db         *db.Mongo
	collection string

	// Access services
	Auth  *AuthService
	User  *UserService
	Level *LevelService
	Form  *FormService
	Point *PointService
}

func New(db *db.Mongo) *DataAccess {
	da := &DataAccess{db: db, collection: CollectionInit}
	da.Level = newLevelService(db)
	da.Point = newPointService(db)
	da.User = newUserService(db, da.Level, da.Point)
	da.Auth = newAuthService(db, da.User)
	da.Form = newFormService(db, da.User)
	return da
}

func (svc *DataAccess) initConfig(config *Config) error {
	config.Name = configName
	return svc.db.Insert(svc.collection, config)
}

func (svc *DataAccess) updateConfig(config *Config) error {
	return svc.db.Update(svc.collection, bson.M{"name": config.Name}, config)
}

func (svc *DataAccess) findConfig() (*Config, error) {
	result := &Config{}
	query := bson.M{"name": configName}
	if err := svc.db.FindOne(svc.collection, query, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *DataAccess) AlreadyInitAdmin() bool {
	if config, err := svc.findConfig(); err == nil && config != nil {
		return config.Admin
	}
	return false
}

func (svc *DataAccess) AlreadyInitLevel() bool {
	if config, err := svc.findConfig(); err == nil && config != nil {
		return config.Level
	}
	return false
}

func (svc *DataAccess) CreateConfig() error {
	_, err := svc.findConfig()
	if err != nil {
		conf := &Config{Admin: false, Level: false}
		if err := svc.initConfig(conf); err != nil {
			return err
		}
	}
	return nil
}

func (svc *DataAccess) InitAdmin(reg *model.User, pwd string) error {
	config, err := svc.findConfig()
	if err != nil {
		return err
	}

	secret := util.MD5Encode(pwd)
	if !svc.User.IsExistByEmail(reg.Email) {
		if err := svc.Auth.Register(reg, secret); err != nil {
			return err
		}

		user, err := svc.User.FindByEmail(reg.Email)
		if err != nil {
			return err
		}
		stat := &model.UserStatus{UserUUID: user.UUID, Block: false, Active: true}
		if err := svc.User.UpdateStatus(stat); err != nil {
			return err
		}

		role := &model.UserRole{UserUUID: user.UUID, Role: model.RoleAdmin}
		if err := svc.User.UpdateRole(role); err != nil {
			return err
		}

		log.Printf("Admin email: %s", reg.Email)
		log.Printf("Admin password: %s", pwd)

		config.Admin = true
		if err := svc.updateConfig(config); err != nil {
			return err
		}
	}
	return nil
}

func (svc *DataAccess) InitLevels(objs interface{}) error {
	config, err := svc.findConfig()
	if err != nil {
		return err
	}

	for _, obj := range objs.([]interface{}) {
		lv := map[string]interface{}{}
		for k, v := range obj.(map[interface{}]interface{}) {
			lv[k.(string)] = v
		}

		data, err := json.Marshal(lv)
		if err != nil {
			return err
		}

		level := &model.Level{}
		if err := json.Unmarshal(data, &level); err != nil {
			return err
		}

		if err := svc.Level.Insert(level); err != nil {
			return err
		}
		log.Printf("Default level created: %s.\n", level.Name)
	}

	config.Level = true
	if err := svc.updateConfig(config); err != nil {
		return err
	}
	return nil
}
