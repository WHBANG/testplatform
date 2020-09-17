package proto

type DeeperFacePeopleInfo struct {
	PTS         [4][2]int `json:"pts"`
	Score       float64   `json:"score"`
	Orientation string    `json:"orientation"`
	Quality     string    `json:"quality"`
	FaceUri     []byte    `json:"faceUri"`

	YituPeople []YituSearchResultItem `json:"yituPeople"` // 保存从依图搜索出来的结果
}

type DeeperFaceInfoArray struct {
	Infos      []DeeperFacePeopleInfo `json:"infos,omitempty"`
	FeatureURI string                 `json:"featureUri,omitempty"`
	FaceInfo   string                 `json:"faceInfo,omitempty"`
}

type DeeperFaceInfo struct {
	People    []DeeperFaceInfoArray `json:"people,omitempty"` // 顺序和Snapshot保持一致, snapshot与之一一对应
	ErrorInfo string                `json:"errorInfo,omitempty"`
}

type DeeperInfo struct {
	Face DeeperFaceInfo `json:"face,omitempty"`
}

type YituSearchResultItem struct {
	YituID           string  `json:"yituId" bson:"yituId"`
	YituFaceImageUri string  `json:"yituFaceImageUri" bson:"yituFaceImageUri"`
	Name             string  `json:"name" bson:"name"`
	Nation           string  `json:"nation" bson:"nation"`
	Sex              string  `json:"sex" bson:"sex"` //  1: Male, 0: Female
	IDCard           string  `json:"idCard" bson:"idCard"`
	Similarity       float64 `json:"similarity" bson:"similarity"`
	Offset           int     `json:"offset" bson:"offset"`

	FaceImageId  string `json:"faceImageId" bson:"faceImageId"`
	FaceImageUrl string `json:"faceImageUrl" bson:"faceImageUrl"`
	FaceRect     struct {
		W int `json:"w" bson:"w"`
		H int `json:"h" bson:"h"`
		X int `json:"x" bson:"x"`
		Y int `json:"y" bson:"y"`
	} `json:"faceRect" bson:"faceRect"`

	FaceLibId   int    `json:"faceLibId" bson:"faceLibId"`
	FaceLibName string `json:"faceLibName" bson:"faceLibName"`

	UpdateTime string `json:"updateTime" bson:"updateTime"`

	GeoCheck int `json:"geoCheck" bson:"geoCheck"` // 位置校验，1位置匹配，0位置不匹配，-1查询不到手机号
}
