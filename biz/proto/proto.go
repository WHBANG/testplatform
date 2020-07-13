package proto

import (
	"time"

	"git.supremind.info/testplatform/biz/analyzerclient"
	"gopkg.in/mgo.v2/bson"
)

type ImageInfo struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Image       string        `json:"image" bson:"image"`
	UserID      int           `json:"user_id" bson:"user_id"`
	Description string        `json:"description" bson:"description"`
	Product     string        `json:"product" bson:"product"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}

type AnalyzerFlowImageInfo struct {
	ImageInfo
	Models string `json:"models" bson:"models"`
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
	EngineStatusNone    = "NONE"
	EngineStatusCreated = "CREATED"
	EngineStatusStarted = "STARTED"
	EngineStatusFailed  = "FAILED"
	EngineStatusStoped  = "STOPED"
)

type EngineDeployInfo struct {
	ID          bson.ObjectId          `json:"id" bson:"_id"`
	Image       string                 `json:"image" bson:"image"`
	UserID      int                    `json:"user_id" bson:"user_id"`
	Description string                 `json:"description" bson:"description"`
	Product     string                 `json:"product" bson:"product"`
	Status      string                 `json:"status" bson:"status"` //enginestatus
	JobInfo     analyzerclient.JobInfo `json:"job_info" bson:"job_info"`
	// Region      string        `json:"region" bson:"region"`     //根据不同的region来获取任务，区别测试
	AnalyzerConfig string    `json:"analyzer_config,omitempty" bson:"analyzer_config"` //配置文件，覆盖原有的
	ErrorInfo      string    `json:"error_info" bson:"error_info"`
	StartTime      time.Time `json:"start_time" bson:"start_time"`
	StopTime       time.Time `json:"stop_time" bson:"stop_time"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

//测试实例
type ImageTestInfo struct {
	ID bson.ObjectId `json:"_id" bson:"_id"`

	EngineID bson.ObjectId `json:"engine_id"`

	UserID      int    `json:"user_id" bson:"user_id"`
	Description string `json:"description" bson:"description"`
	Product     string `json:"product" bson:"product"`

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
