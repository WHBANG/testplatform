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

var databaseConn *mgo.Database
var databaseSession *mgo.Session

func InitDB(config *MongodbConfig) error {
	session, err := mgo.Dial(config.Host)
	if err != nil {
		return fmt.Errorf("mongo dial error:%s ", err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	databaseSession = session
	databaseConn = databaseSession.DB(config.DB)
	return nil
}

func GetMgoDB() (*mgo.Database, error) {
	if databaseConn == nil {
		return nil, errors.New("database connection is not init")
	}
	return databaseConn, nil
}

func GetMgoDBSession() (*mgo.Session, error) {
	if databaseSession == nil {
		return nil, errors.New("database connection is not init")
	}
	return databaseSession, nil
}
