package api

import (
	"strings"

	"github.com/globalsign/mgo/bson"
)

func NewID() *ID {
	return &ID{Hex: bson.NewObjectId().Hex()}
}

func (id *ID) ObjectID() bson.ObjectId {
	return bson.ObjectIdHex(id.Hex)
}

// json.Marshaler
func (id *ID) MarshalJSON() ([]byte, error) {
	return id.ObjectID().MarshalJSON()
}

// json.Unmarshaler
func (id *ID) UnmarshalJSON(data []byte) error {
	id.Hex = removeQuotes(string(data))
	return nil
}

// text.Marshaler
func (id *ID) MarshalText() ([]byte, error) {
	return id.ObjectID().MarshalText()
}

// text.Unmarshaler
func (id *ID) UnmarshalText(data []byte) error {
	id.Hex = removeQuotes(string(data))
	return nil
}

// bson.Marshaler
func (id *ID) GetBSON() (interface{}, error) {
	return id.ObjectID(), nil
}

// bson.Unmarshaler
func (id *ID) SetBSON(raw bson.Raw) error {
	var objectID bson.ObjectId
	if e := raw.Unmarshal(&objectID); e != nil {
		return e
	}

	id.Hex = objectID.Hex()
	return nil
}

func removeQuotes(s string) string {
	s = strings.TrimPrefix(s, `"`)
	s = strings.TrimSuffix(s, `"`)
	return s
}
