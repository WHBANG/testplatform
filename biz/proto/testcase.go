package proto

type MetaData struct {
	Case  *MetaCase   `json:"case"`
	Task  *Task       `json:"task"`
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
