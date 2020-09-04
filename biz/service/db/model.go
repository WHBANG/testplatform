package db

import (
	bproto "git.supremind.info/testplatform/biz/proto"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	modelsCollName = "models"
)

type ModelMgnt interface {
	Init(models []bproto.Model) error
	Reset([]bproto.Model) ([]bproto.Model, error)
	FindByModelID(int) (*bproto.Model, error)
	FindByModelType(string) ([]bproto.Model, error)

	Insert(req *bproto.Model) (*bproto.Model, error)
	Update(id int, updateInfo map[string]interface{}) (*bproto.Model, error)
	Delete(id int) (*bproto.Model, error)
	Likefind(query bson.M, page, size int) ([]bproto.Model, int, error)
	BatchDelete(id []int) ([]bproto.Model, error)
}

type MongoModel struct {
	collName string
	DB       string
	session  *mgo.Session
}

func NewMongoModel(s *mgo.Session, db string) (ModelMgnt, error) {
	mgnt := &MongoModel{
		session:  s,
		DB:       db,
		collName: modelsCollName,
	}
	index := mgo.Index{
		Key:    []string{"model_id", "model_name", "model_url"},
		Unique: true,
	}
	s.DB(db).C(modelsCollName).EnsureIndex(index)
	return mgnt, nil
}

func (d *MongoModel) Update(id int, updateInfo map[string]interface{}) (*bproto.Model, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(modelsCollName)

	selector := bson.M{"model_id": id}
	//data := bson.M{"$set": updateInfo}

	err := c.Update(selector, map[string]interface{}{
		"$set": updateInfo,
	})
	if err != nil {
		log.Error("Update Error: ", err)
		return nil, err
	}
	model := &bproto.Model{}
	err = c.Find(bson.M{"model_id": id}).One(model)
	return model, err
}

func (d *MongoModel) Delete(id int) (*bproto.Model, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(modelsCollName)

	model := &bproto.Model{}
	c.Find(bson.M{"model_id": id}).One(model)
	err := c.Remove(bson.M{"model_id": id})
	if err != nil {
		log.Error("Delete Error: ", err)
		return nil, err
	}
	return model, err
}

func (d *MongoModel) BatchDelete(ids []int) ([]bproto.Model, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(modelsCollName)

	models := []bproto.Model{}
	c.Find(bson.M{"model_id": bson.M{"$in": ids}}).All(&models)
	_, err := c.RemoveAll(bson.M{"model_id": bson.M{"$in": ids}})
	if err != nil {
		log.Error("Delete Error!: ", err)
		return nil, err
	}
	return models, err
}

func (d *MongoModel) Insert(req *bproto.Model) (*bproto.Model, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(modelsCollName)

	err := c.Insert(*req)
	if err != nil {
		log.Error("Insert Error: ", err)
	}
	return req, err
}

func (d *MongoModel) Likefind(query bson.M, page, size int) ([]bproto.Model, int, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(modelsCollName)

	mlist := make([]bproto.Model, size)
	err := c.Find(query).Skip((page - 1) * size).Limit(size).Sort("-_id").All(&mlist)
	if err != nil {
		log.Error("Model query Error: ", err)
		return nil, 0, err
	}
	total, err := c.Find(query).Count()
	if err != nil {
		log.Error("Model Query Count Error:", err)
		return nil, 0, err
	}
	return mlist, total, err
}

func (d *MongoModel) Init(models []bproto.Model) error {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(modelsCollName)

	var list_id []int
	var tmp bproto.Model
	for i := 0; i < len(models); i++ {
		list_id = append(list_id, models[i].ModelID)
		err := c.Find(bson.M{"model_id": models[i].ModelID}).One(&tmp)
		if err != nil {
			err = c.Insert(models[i])
			if err != nil {
				log.Printf("Update Model Data Error: ", err)
				return err
			}
		}
		//else {
		//	err = c.Update(bson.M{"model_id": models[i].ModelID}, models[i])
		//}
	}
	_, err := c.RemoveAll(bson.M{"model_id": bson.M{"$nin": list_id}})
	if err != nil {
		log.Printf("Update Model Data Error: ", err)
		return err
	}
	//保持心跳，每10分钟进行一次更新
	//统计数量是不是一样多
	//如果是一次根据ID进行更新
	//如果不是全部删除后，重新插入（数量可能一样但是由于id不同必然后导致新库中数据的增加，一旦数据不一致，肯定是旧库增加或者删除导致的

	return nil
}

func (d *MongoModel) Reset(req []bproto.Model) ([]bproto.Model, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(modelsCollName)
	var err error
	_, err = c.RemoveAll(nil)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(req); i++ {
		err = c.Insert(req[i])
		if err != nil {
			log.Error("Insert Error: ", err)
			return nil, err
		}
	}
	return req, err
}

func (d *MongoModel) FindByModelID(id int) (*bproto.Model, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(modelsCollName)

	model := &bproto.Model{}
	err := c.Find(bson.M{"model_id": id}).One(model)
	if err != nil {
		log.Error("Find Error: ", err)
		return nil, err
	}
	return model, err
}

func (d *MongoModel) FindByModelType(model_type string) ([]bproto.Model, error) {
	session := d.session.Clone()
	defer session.Close()
	c := session.DB(d.DB).C(modelsCollName)

	models := []bproto.Model{}
	err := c.Find(bson.M{"model_type": model_type}).All(&models)
	if err != nil {
		log.Error("Find Error: ", err)
		return nil, err
	}
	return models, err
}
