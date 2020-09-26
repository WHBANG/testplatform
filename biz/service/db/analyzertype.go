package db

import (
	"time"

	bproto "git.supremind.info/testplatform/biz/proto"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	analyzerTypeCollName = "analyzertype"
)

type AnalyzerMgnt interface {
	Insert(*bproto.AnalyzerTypeInfo) (*bproto.AnalyzerTypeInfo, error)
	Find(id bson.ObjectId) (*bproto.AnalyzerTypeInfo, error)
	FindAll() ([]bproto.AnalyzerTypeInfo, error)
	Delete(id bson.ObjectId) (*bproto.AnalyzerTypeInfo, error)
}

type MongoAnalyzerType struct {
	collName string
	DB       string
	session  *mgo.Session
}

func NewMongoAnalyzerType(s *mgo.Session, db string) (AnalyzerMgnt, error) {
	anzlyzerType := &MongoAnalyzerType{
		session:  s,
		DB:       db,
		collName: analyzerTypeCollName,
	}

	index := mgo.Index{
		Key:    []string{"analyzer_type"},
		Unique: true,
	}
	s.DB(db).C(analyzerTypeCollName).EnsureIndex(index)
	return anzlyzerType, nil
}

func (d *MongoAnalyzerType) Insert(an *bproto.AnalyzerTypeInfo) (*bproto.AnalyzerTypeInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(analyzerTypeCollName)

	if an.ID == "" {
		an.ID = bson.NewObjectId()
	}
	now := time.Now()
	an.CreateTime = now

	err := c.Insert(*an)
	if err != nil {
		log.Error("Insert Error: ", err)
		return nil, err
	}
	return an, err
}

func (d *MongoAnalyzerType) Delete(id bson.ObjectId) (*bproto.AnalyzerTypeInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(analyzerTypeCollName)
	log.Println(id)
	an := &bproto.AnalyzerTypeInfo{}
	c.Find(bson.M{"_id": id}).One(&an)
	err := c.Remove(bson.M{"_id": id})
	if err != nil {
		log.Error("Delete Error:", err)
		return nil, err
	}
	return an, err
}

func (d *MongoAnalyzerType) Find(id bson.ObjectId) (*bproto.AnalyzerTypeInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(analyzerTypeCollName)

	analyzerType := &bproto.AnalyzerTypeInfo{}
	err := c.Find(bson.M{"_id": id}).One(analyzerType)
	if err != nil {
		log.Error("Find Error:", err)
		return nil, err
	}
	return analyzerType, err
}

func (d *MongoAnalyzerType) FindAll() ([]bproto.AnalyzerTypeInfo, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(analyzerTypeCollName)

	analyzerTypes := []bproto.AnalyzerTypeInfo{}
	err := c.Find(nil).All(&analyzerTypes)
	if err != nil {
		log.Error("Find Error:", err)
		return nil, err
	}
	return analyzerTypes, err
}
