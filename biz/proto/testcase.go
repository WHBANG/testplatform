package proto

import (
	"time"

	"git.supremind.info/product/visionmind/com/flow"
	"gopkg.in/mgo.v2/bson"
)

type EventData map[string]interface{}

type MetaData struct {
	Case  *MetaCase   `json:"case"`
	Task  *flow.Task  `json:"task"`
	Event []EventData `json:"event"`
	Files *FileCase   `json:"files"`
}

type FileCase struct {
	Videos []string `json:"videos"`
	Images []string `json:"images"`
}

type MetaCase struct {
	Name           string            `json:"name"`
	Product        string            `json:"product"`
	ProductVersion string            `json:"product_version"`
	Label          map[string]string `json:"label"`
}

type TestCase struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	TestData    *MetaData     `json:"test_data" bson:"test_data"`
	UserID      int           `json:"user_id" bson:"user_id"`
	Description string        `json:"description" bson:"description"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}
