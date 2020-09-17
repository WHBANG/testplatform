package proto

type Model struct {
	ModelID      int    `json:"model_id" bson:"model_id"`
	ModelName    string `json:"model_name" bson:"model_name"`
	ModelType    string `json:"model_type" bson:"model_type"`
	ModelURL     string `json:"model_url" bson:"model_url"`
	ModelVersion string `json:"model_version" bson:"model_version"`
	ModelUser    string `json:"model_user" bson:"model_user"`
	Description  string `json:"description" bson:"description"`
}
