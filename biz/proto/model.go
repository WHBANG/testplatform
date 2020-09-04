package proto

type Model struct {
	ModelID   int    `json:"model_id" bson:"model_id"`
	ModelName string `json:"model_name" bson:"model_name"`
	ModelType string `json:"model_type" bson:"model_type"`
	ModelURL  string `json:"model_url" bson:"model_url"`
}
 