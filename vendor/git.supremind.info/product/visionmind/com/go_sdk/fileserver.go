package go_sdk

type IFileServerClient interface {
	Init() error
	Save(key string, data []byte) (url string, err error)
	SaveKeepForever(key string, data []byte) (url string, err error)
	// GetURI(key string) (string, error)
}
