package db

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2"
)

type MongodbConfig struct {
	Host string `json:"host"`
	DB   string `json:"db"`
}

var dbConfig *MongodbConfig
var databaseSession *mgo.Session

func InitDB(config *MongodbConfig) error {
	session, err := mgo.Dial(config.Host)
	if err != nil {
		return fmt.Errorf("mongo dial error:%s ", err)
	}
	session.SetMode(mgo.Monotonic, true)
	databaseSession = session
	dbConfig = config
	return nil
}

func GetDBName() string {
	if dbConfig == nil {
		return ""
	}
	return dbConfig.DB
}

func GetMgoDBSession() (*mgo.Session, error) {
	if databaseSession == nil {
		return nil, errors.New("database connection is not init")
	}
	return databaseSession, nil
}
