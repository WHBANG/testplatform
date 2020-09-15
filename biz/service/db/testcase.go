package db

import (
	"errors"
	"time"

	"git.supremind.info/testplatform/biz/proto"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	testCaseCollName = "testcase"
)

type TestCaseMgnt interface {
	Insert(*proto.TestCaseInfo) (*proto.TestCaseInfo, error)
	Update(id bson.ObjectId, updateInfo map[string]interface{}) (*proto.TestCaseInfo, error)
	Delete(id bson.ObjectId) (*proto.TestCaseInfo, error)
	Find(id bson.ObjectId) (*proto.TestCase, error)
	LikeFind(query bson.M, pageint, size int) (data []proto.TestCaseInfo, total int, err error)
	BatchDelete(list []bson.ObjectId) ([]proto.TestCaseInfo, error)
}

type TestCase struct {
	collName string
	DB       string
	session  *mgo.Session
}

func NewMongoCase(c *mgo.Session, db string) (TestCaseMgnt, error) {
	img := &TestCase{
		session:  c,
		DB:       db,
		collName: testCaseCollName,
	}
	c.DB(db).C(img.collName).EnsureIndexKey("user_id", "description", "product", "created_at")
	return img, nil
}

func (m *TestCase) Insert(tc *proto.TestCaseInfo) (*proto.TestCaseInfo, error) {
	session := m.session.Clone()
	defer session.Close()
	c := session.DB(m.DB).C(m.collName)

	if tc.ID == "" {
		tc.ID = bson.NewObjectId()
	}
	tc.CreatedAt = time.Now()
	tc.UpdatedAt = time.Now()
	err := c.Insert(*tc)
	if err != nil {
		log.Error("Insert Error: ", err)
		return nil, err
	}
	return tc, err
}

func (m *TestCase) Update(id bson.ObjectId, updateInfo map[string]interface{}) (*proto.TestCaseInfo, error) {
	session := m.session.Clone()
	defer session.Close()
	tc := session.DB(m.DB).C(m.collName)

	updateInfo["updated_at"] = time.Now()
	err := tc.UpdateId(id, updateInfo)
	if err != nil {
		log.Error("Update Error: ", err)
		return nil, err
	}
	testcase := &proto.TestCaseInfo{}
	tc.FindId(id).One(testcase)
	return testcase, nil
}

func (m *TestCase) Delete(id bson.ObjectId) (*proto.TestCaseInfo, error) {
	session := m.session.Clone()
	defer session.Close()
	tc := session.DB(m.DB).C(m.collName)

	testcase := &proto.TestCaseInfo{}
	tc.Find(bson.M{"_id": id}).One(testcase)
	err := tc.Remove(bson.M{"_id": id})
	if err != nil {
		log.Error("Delete Error: ")
		return nil, err
	}
	return testcase, nil
}

func (m *TestCase) BatchDelete(list []bson.ObjectId) ([]proto.TestCaseInfo, error) {
	session := m.session.Clone()
	defer session.Close()
	tc := session.DB(m.DB).C(m.collName)

	testcase := []proto.TestCaseInfo{}
	tc.Find(bson.M{"_id": bson.M{"$in": list}}).All(&testcase)
	_, err := tc.RemoveAll(bson.M{"_id": bson.M{"$in": list}})
	if err != nil {
		log.Error("Delete Error: ")
		return nil, err
	}
	return testcase, nil
}

func (m *TestCase) Find(id bson.ObjectId) (*proto.TestCase, error) {
	session := m.session.Clone()
	defer session.Close()
	tc := session.DB(m.DB).C(m.collName)

	testcase := &proto.TestCase{}
	err := tc.Find(bson.M{"_id": id}).One(testcase)
	if err != nil {
		log.Error("Find Error: ", err)
		return nil, err
	}
	return testcase, nil
}

func (m *TestCase) LikeFind(query bson.M, page, size int) ([]proto.TestCaseInfo, int, error) {
	session := m.session.Clone()
	defer session.Close()
	tc := session.DB(m.DB).C(m.collName)

	tlist := make([]proto.TestCaseInfo, size)
	err := tc.Find(query).Skip((page - 1) * size).Limit(size).Sort("-_id").All(&tlist)
	if err != nil {
		log.Error("Testase Query Error: ", err)
		return nil, 0, err
	}
	total, err := tc.Find(query).Count()
	if err != nil {
		log.Error("Testcase Count Error: ", err)
		return nil, 0, err
	}
	if total > 0 && (total/size+1) < page {
		err := errors.New("Testcase Page Error: ")
		return nil, total, err
	}
	return tlist, total, nil

}
