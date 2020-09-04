package proto

type CommonTaskRes struct {
	Code int         `json:"code"`
	Data interface{} `bson:"data"`
}
