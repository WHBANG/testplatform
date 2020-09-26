package proto

import (
	"time"

	"git.supremind.info/testplatform/biz/analyzerclient"
	"gopkg.in/mgo.v2/bson"
)

type ImageStatus string

const (
	Created  ImageStatus = "CREATED"
	Building ImageStatus = "BUILDING"
	Done     ImageStatus = "DONE"
	Failed   ImageStatus = "FAILED"
)

type ImageType string

const (
	FLOW ImageType = "analyzer-flow"
	IO   ImageType = "analyzer-io"
)

type ImageInfo struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Image       string        `json:"image" bson:"image"`
	UserID      int           `json:"user_id" bson:"user_id"`
	Status      ImageStatus   `json:"status" bson:"status"`
	Type        ImageType     `json:"type" bson:"type"`
	Description string        `json:"description" bson:"description"`
	Product     string        `json:"product" bson:"product"`
	Models      []Model       `json:"models" bson:"models"`
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
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}

const (
	EngineStatusNone    = "NONE"
	EngineStatusCreated = "CREATED"
	EngineStatusStarted = "STARTED"
	EngineStatusFailed  = "FAILED"
	EngineStatusStoped  = "STOPPED"
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
type ModelNames []string

var (
	BianJian = ModelNames{
		"ducha_det.tronmodel",
		"ducha_cls.tronmodel",
		"bk_fight.tronmodel",
	}
	Massiveflow = ModelNames{
		"crowd_count_model.tronmodel",
		"banner_detect_east_model.tronmodel",
		"banner_detect_od_model.tronmodel",
		"head_count_model.tronmodel",
		"fight_classify_local_model.tronmodel",
		"queue_count_local_model.tronmodel",
		"crowd_region_count_local_model.tronmodel",
	}
	DuCha = ModelNames{
		"ducha_det.tronmodel",
		"ducha_cls.tronmodel",
		"bk_fight.tronmodel",
	}
)

type AnalyzerTypeInfo struct {
	ID            bson.ObjectId `json:"_id" bson:"_id"`
	AnalyzerType  string        `json:"analyzer_type" bson:"analyzer_type"`
	ModelNameList ModelNames    `json:"model_name_list" bson:"model_name_list"`
	CreateTime    time.Time     `json:"create_time" bson:"create_time"`
}
