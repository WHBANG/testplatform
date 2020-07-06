package proto

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	EngineStatusStarting = "STARTING" //启动中
	EngineStatusWaiting  = "WAITING"
	EngineStatusRunning  = "RUNNING"
	EngineStatusFailed   = "FAILED"
	EngineStatusFinished = "FINISHED"
)

type EngineRuntimeInfo struct {
	Status string `json:"status" bson:"status"` //enginestatus
	// Region      string        `json:"region" bson:"region"`     //根据不同的region来获取任务，区别测试
	AnalyzerConfig interface{} `json:"analyzer_config" bson:"analyzer_config"` //配置文件，覆盖原有的
	ErrorInfo      string      `json:"error_info" bson:"error_info"`
	StartTime      time.Time   `json:"start_time" bson:"start_time"`
	StopTime       time.Time   `json:"stop_time" bson:"stop_time"`
}

type ImageInfo struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	ImageType   string        `json:"image_type" bson:"image_type"`
	Image       string        `json:"image" bson:"image"`
	Models      string        `json:"models" bson:"models"`
	UserID      int           `json:"user_id" bson:"user_id"`
	Description string        `json:"description" bson:"description"`
	Product     string        `json:"product" bson:"product"`
	// Status      string            `json:"status" bson:"status"`
	Runtime   EngineRuntimeInfo `json:"runtime" bson:"runtime"`
	CreatedAt time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time         `json:"updated_at" bson:"updated_at"`
}

type TestCaseInfo struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	TestData    interface{}   `json:"test_data" bson:"test_data"`
	UserID      int           `json:"user_id" bson:"user_id"`
	Description string        `json:"description" bson:"description"`
	Product     string        `json:"product" bson:"product"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}

const (
// EngineStatusStarting = "STARTING" //启动中
// EngineStatusWaiting  = "WAITING"
// EngineStatusRunning  = "RUNNING"
// EngineStatusFailed   = "FAILED"
// EngineStatusFinished = "FINISHED"
)

//测试实例
type ImageTestInfo struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	ImageID     bson.ObjectId `json:"image_id" bson:"image_id"` //镜像表中id
	UserID      int           `json:"user_id" bson:"user_id"`
	Description string        `json:"description" bson:"description"`
	Product     string        `json:"product" bson:"product"`

	TestStatus string       `json:"test_status" bson:"test_status"`
	TestCase   TestCaseInfo `json:"test_case" bson:"test_case"`
	TestResult string       `json:"test_result" bson:"test_result"`
	StartTime  time.Time    `json:"start_time" bson:"start_time"`
	StopTime   time.Time    `json:"stop_time" bson:"stop_time"`
	Duration   int          `json:"duration" bson:"duration"`
	ErrorInfo  string       `json:"error_info" bson:"error_info"`
	Report     string       `json:"report" bson:"report"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
