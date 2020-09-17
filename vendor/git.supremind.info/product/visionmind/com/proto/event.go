package proto

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	StatusInit     = "init"
	StatusFinished = "finished"
)

// 打标操作状态
const (
	MarkingInit    = "init"
	MarkingIllegal = "illegal"
	MarkingDiscard = "discard"
	MarkingDone    = "done"
)

type IndexData struct {
	Hour int    `json:"hour" bson:"hour"` //发生时间，24小时制 ，用于统计
	Date string `json:"date" bson:"date"` //发生时间日期 2006-15-04，用于统计
}

//================================================================
// 某个标签及置信度
type Detail struct {
	ID     string   `json:"id,omitempty" bson:"id,omitempty"`         // 标签的ID，通过“/”拼接，可标示标签之间的归属关系
	Pts    [][2]int `json:"pts,omitempty" bson:"pts,omitempty"`       // 检测的区域，分类则为空
	Score  float64  `json:"score" bson:"score"`                       // 置信度
	Scene  string   `json:"scene" bson:"scene"`                       // 场景，如“车牌识别”
	Label  string   `json:"label" bson:"label"`                       // 标签值，如“AB123”
	Group  string   `json:"group,omitempty" bson:"group,omitempty"`   // 标签分组，如“特种车辆”
	Sample *Sample  `json:"sample,omitempty" bson:"sample,omitempty"` // 比对场景下的示例，如人脸比对场景下的底图信息
}

// 比对场景下的，底图示例
type Sample struct {
	URI string   `json:"uri"`
	Pts [][2]int `json:"pts,omitempty"`
}

// constant for Scene：标签场景枚举值
const (
	// JDC
	SCENE_DETECT = "detect" // 目标检测
	SCENE_CLASS  = "class"  // 目标分类
	SCENE_PLATE  = "plate"  // 目标号牌
)

type Snapshot struct {
	FeatureURI       string `json:"featureUri" bson:"featureUri"`   // 特写图片
	SnapshotURI      string `json:"snapshotUri" bson:"snapshotUri"` // 截帧图片
	SnapshotURIRaw   string `json:"snapshotUriRaw" bson:"snapshotUriRaw"`
	SnapshotURIThumb string `json:"snapshotUriThumb" bson:"snapshotUriThumb"` // 缩略图，对 SnapshotURI 的画框图进行缩小处理

	// 老数据结构，暂时不动
	Pts        [][2]int `json:"pts" bson:"pts"`             // 检测目标坐标
	SizeRatio  float64  `json:"sizeRatio" bson:"sizeRatio"` // 缩放比例
	Score      float64  `json:"score" bson:"score"`
	Class      int      `json:"class" bson:"class"`           //类别（非机动车-外卖类型，饿了么/美团等）
	Label      string   `json:"label" bson:"label"`           // 标牌（非机动车-车牌）
	LabelScore float64  `json:"labelScore" bson:"labelScore"` // 标牌置信度

	// 新数据结构，结构化表示一张图中的标签信息
	Main    bool     `json:"main,omitempty" bson:"main,omitempty"`
	Details []Detail `json:"details,omitempty" bson:"details,omitempty"`

	// 原始结构化信息
	Origin interface{} `json:"origin,omitempty" bson:"origin,omitempty"` // 原始信息
}

const (
	ReadStatusUnread   = 0
	ReadStatusHaveRead = 1

	ProcessStatusUnprocessed = 0
	ProcessStatusProcessed   = 1
)

// 打标相关字段
type Mark struct {
	Marking            string          `json:"marking" bson:"marking"`                       // 打标，init/illegal/discard
	IsClassEdit        bool            `json:"isClassEdit" bson:"isClassEdit"`               // 类别是否被编辑
	IsLabelEdit        bool            `json:"isLabelEdit" bson:"isLabelEdit"`               // 标牌是否被编辑
	IsNonmotorTypeEdit bool            `json:"isNonmotorTypeEdit" bson:"isNonmotorTypeEdit"` // 非机动车车辆类型是否被编辑
	DiscardReason      int             `json:"discardReason" bson:"discardReason"`           // 作废原因
	CommonEditMap      map[string]bool `json:"commonEditMap" bson:"commonEditMap"`           // 通用数组，是否被编辑
}

type EventExtra struct {
	ReadStatus    int `json:"readStatus" bson:"readStatus"`       // 读取状态，0未读/1已读
	ProcessStatus int `json:"processStatus" bson:"processStatus"` // 处理状态，0未处理/1已处理
	ProcessReason int `json:"processReason" bson:"processReason"` //处理原因
}

//视频事件检测消息
type EventMsg struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	EventID   string        `json:"eventId" bson:"eventId"`      // 事件ID ，全局唯一
	EventType int           `json:"eventType" bson:"eventType" ` // 事件类型
	TaskID    string        `json:"taskId" bson:"taskId"`        // 任务ID
	TaskType  string        `json:"taskType" bson:"taskType"`    // 任务类型
	Region    string        `json:"region" bson:"region"`        // 事件发生区域

	CameraID string     `json:"cameraId" bson:"cameraId"`  // 设备ID
	Snapshot []Snapshot `json:"snapshot" bson:"snapshot"`  // 事件截图
	VideoURI string     `json:"videoUri"  bson:"videoUri"` // 事件视频

	Lane        *string          `json:"lane,omitempty"  bson:"lane,omitempty"`                   // 车道
	Direction   *string          `json:"direction,omitempty"  bson:"direction,omitempty"`         // 方向
	Count       *int32           `json:"count,omitempty"  bson:"count,omitempty"`                 // 数量
	CountByType map[string]int32 `json:"count_by_type,omitempty"  bson:"count_by_type,omitempty"` // 数量

	AvgSpeed       *float32 `json:"avg_speed,omitempty"  bson:"avg_speed,omitempty"`               // 平均速度
	AvgSpeedCount  *int32   `json:"avg_speed_count,omitempty"  bson:"avg_speed_count,omitempty"`   // 平均速度的总车数
	RoadCoverRatio *float32 `json:"road_cover_ratio,omitempty"  bson:"road_cover_ratio,omitempty"` // 路面占用比

	// 事件总结，定性，利于查询
	// scene: label，某个场景下，打上某个标签
	Summary map[string]struct {
		Score float64 `json:"score" bson:"score"`
		Label string  `json:"label" bson:"label"`
	} `json:"summary,omitempty" bson:"summary,omitempty"`

	// 原始配置信息
	OriginalConfig         interface{} `json:"originalConfig,omitempty" bson:"originalConfig,omitempty"`
	OriginalViolationIndex int         `json:"originalViolationIndex" bson:"originalViolationIndex"`

	StartTime time.Time `json:"startTime" bson:"startTime"` // 事件起始时间
	EndTime   time.Time `json:"endTime" bson:"endTime"`     // 事件结束时间
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"` // 数据产生时间
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"` // 数据更新时间

	Mark      Mark      `json:"mark" bson:"mark"`
	IndexData IndexData `json:"indexData" bson:"indexData"` // 索引

	Extra interface{} `json:"extra" bson:"extra"` // 额外扩展信息，比如人脸、档案等其它信息

	// 具体项目相关信息
	Zone       interface{} `json:"zone" bson:"zone"` //划线或监测区域配置
	DeeperInfo interface{} `json:"deeperInfo" bson:"deeperInfo"`
	Status     string      `json:"status" bson:"status"`                       // 状态，init/finished
	IsWhite    bool        `json:"isWhite,omitempty" bson:"isWhite,omitempty"` // “非违规”的正常事件

	// 事件相关
	EventExtra     EventExtra                `json:"eventExtra" bson:"eventExtra"`         // 事件扩展信息
	ComponentExtra map[string]ComponentExtra `json:"componentExtra" bson:"componentExtra"` //各个系统处理字段
}

type ComponentExtra struct {
	Status string      `json:"status" bson:"status"`
	Result interface{} `json:"result" bson:"result"`
}
