package proto

import "time"

type GetJTEventRes struct {
	Content    []JTEventInfo `json:"content"`
	Page       int           `json:"page"`
	PerPage    int           `json:"per_page"`
	TotalPage  int           `json:"total_page"`
	TotalCount int           `json:"total_count"`
}

type JTEventInfo struct {
	ID                     string         `json:"id"`
	EventID                string         `json:"eventId"`
	EventType              string         `json:"eventType"`
	TaskID                 string         `json:"taskId"`
	TaskType               string         `json:"task_type"`
	Region                 string         `json:"region"`
	CameraID               string         `json:"cameraId"`
	Snapshots              []Snapshot     `json:"snapshot"`
	VideoURI               string         `json:"videoUri"`
	Summ                   Summary        `json:"summary"`
	OriginalConfi          OriginalConfig `json:"originalConfig"`
	OriginalViolationIndex int            `json:"originalViolationIndex"`
	StartTime              string         `json:"startTime"`
	EndTime                string         `json:"endTime"`
	CreatedAt              string         `json:"createdAt"`
	UpdatedAt              string         `json:"updatedAt"`
	Mark                   Mark           `json:"mark"`
	IndexData              IndexData      `json:"indexData"`
	Extra                  interface{}    `json:"extra"`
	Zone                   interface{}    `json:"zone"`
	DeeperInfo             interface{}    `json:"deeperInfo"`
	Status                 string         `json:"status"`
	EventExtra             EventExtra     `json:"eventExtra"`
	ComponentExtra         interface{}    `json:"componentExtra"`
	EventTypeStr           string         `json:"eventTypeStr"`
	ClassStr               []string       `json:"classStr"`
}

type Snapshot struct {
	FeatureURI       string   `json:"featureUri"`
	SnapshotURI      string   `json:"snapshotUri"`
	SnapshotURIRaw   string   `json:"snapshotUriRaw"`
	SnapshotUUIThumb string   `json:"snapshotUriThumb"`
	Pts              PTS      `json:"pts"`
	SizeRatio        int64    `json:"sizeRatio"`
	Score            int64    `json:"score"`
	Class            int64    `json:"class"`
	Label            string   `json:"label"`
	LabelScore       int64    `json:"labelScore"`
	Details          []Detail `json:"details"`
	Origin           Origi    `json:"origin"`
}

type PTS [2][2]int64

type Detail struct {
	ID    string `json:"id"`
	Pts   PTS    `json:"pts"`
	Score int    `json:"score"`
	Scene string `json:"scene"`
	Label string `json:"label"`
}

type Origi struct {
	Now     time.Time `json:"now"`
	Objects []Object  `json:"objects"`
	PTS     time.Time `json:"pts"`
}

type Object struct {
	Box                 [4]int `json:"box"`
	Direction           string `json:"direction"`
	DirectionScore      int    `json:"direction_score"`
	Score               int    `json:"score"`
	SpecialCarType      string `json:"special_car_type"`
	SpecialCarTypeScore int    `json:"special_car_type_score"`
	SubType             string `json:"sub_type"`
	SubTypeScore        int    `json:"sub_type_score"`
	Type                string `json:"type"`
}

type Summary struct {
	Class  ClassDetect `json:"class"`
	Detect ClassDetect `json:"detect"`
}

type ClassDetect struct {
	Score int    `json:"score"`
	Label string `json:"label"`
}

type Rois [][]int

type OriginalConfig struct {
	Rois              Rois        `json:"rois"`
	StreamPushAddress string      `json:"stream_push_address"`
	Violations        []Violation `json:"violations"`
}

type Violation struct {
	Code            string      `json:"code"`
	Conditions      []Condition `json:"conditions"`
	EnableRoi       bool        `json:"enable_roi"`
	Name            string      `json:"name"`
	On              bool        `json:"on"`
	Roi             []int       `json:"roi"`
	Threshold       int         `json:"threshold"`
	TrafficLightBox []int       `json:"traffic_light_box"`
}

type Condition struct {
	Data []int  `json:"data"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type Mark struct {
	Marking            string      `json:"marking"`
	IsClassEdit        bool        `json:"isClassEdit"`
	IsLabelEdit        bool        `json:"isLabelEdit"`
	IsNonmotorTypeEdit bool        `json:"isNonmotorTypeEdit"`
	DiscardReason      int         `json:"discardReason"`
	CommonEditMap      interface{} `json:"commonEditMap"`
}

type IndexData struct {
	Hour int    `json:"hour"`
	Date string `json:"date"`
}

type EventExtra struct {
	ReadStatus    int `json:"readStatus"`
	ProcessStatus int `json:"processStatus"`
	ProcessReason int `json:"processReason"`
}
