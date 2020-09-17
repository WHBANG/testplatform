package client

import (
	"log"

	"git.supremind.info/product/visionmind/com/go_sdk"
	"git.supremind.info/product/visionmind/util"
	simplejson "github.com/bitly/go-simplejson"
)

var _ go_sdk.IFileServerClient = &FileServerClient{}

type FileServerClient struct {
	host string
}

func NewFileServerClient(host string) go_sdk.IFileServerClient {
	return &FileServerClient{
		host: host,
	}
}

func (cli *FileServerClient) Init() error {
	return nil
}

func (cli *FileServerClient) Save(filename string, body []byte) (string, error) {
	return cli.uploadStream(body, filename, false)
}

func (cli *FileServerClient) SaveKeepForever(filename string, body []byte) (string, error) {
	return cli.uploadStream(body, filename, true)
}

// //TODO
// func (s *FileServerClient) GetURI(key string) (string, error) {
// 	return "", fmt.Errorf("not implemented")
// }

func (cli *FileServerClient) uploadStream(body []byte, filename string, fileKeepForever bool) (string, error) {
	header := map[string]string{
		"filename": filename,
	}
	if fileKeepForever {
		header["X-File-Keep-Forever"] = "true"
	}
	resp, err := util.PostRaw(cli.host+"/v1/upload/stream", body, header)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(string(resp))
	sj, err := simplejson.NewJson(resp)
	if err != nil {
		log.Println(err)
		return "", err
	}

	url, err := sj.Get("extra").Get("uri").String()
	if err != nil {
		log.Println(err)
		return "", err
	}

	return url, nil
}

type MockFileServerClient struct{}

var _ go_sdk.IFileServerClient = &MockFileServerClient{}

func NewMockFileServerClient() go_sdk.IFileServerClient {
	return &MockFileServerClient{}
}

func (cli *MockFileServerClient) Init() error {
	return nil
}

func (cli *MockFileServerClient) Save(key string, data []byte) (url string, err error) {
	return "http://mock/url", nil
}

func (cli *MockFileServerClient) SaveKeepForever(filename string, body []byte) (string, error) {
	return "http://mock/url", nil
}
