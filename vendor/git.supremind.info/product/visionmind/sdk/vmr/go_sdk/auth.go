package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.supremind.info/product/visionmind/com/go_sdk"
	"qiniupkg.com/x/log.v7"
)

// auth version: reg.supremind.info/application/java_base/micro/traffic_auth:0.0.1

type AuthClient struct {
	Host string
}

type CommonResp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

var _ go_sdk.IAuthClient = &AuthClient{}

func NewAuthClient(host string) go_sdk.IAuthClient {
	return &AuthClient{Host: host}
}

func (auth *AuthClient) CheckToken(token string) (err error) {
	_, err = auth.getReq("v1/auth/login/check-token", token)
	return
}

func (auth *AuthClient) GetSysOrgTree(token string) (root *go_sdk.SysOrgNode, err error) {
	retBz, err := auth.getReq("v1/auth/sys/organ/get-organ-tree", token)
	if err != nil {
		return root, err
	}

	root = &go_sdk.SysOrgNode{}
	err = json.Unmarshal(retBz, root)
	if err != nil {
		return root, err
	}

	return
}

func (auth *AuthClient) getReq(path, token string) (retData []byte, err error) {
	url := fmt.Sprintf("http://%s/%s", auth.Host, path)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return retData, err
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		err = fmt.Errorf("err, StatusCode = %d", resp.StatusCode)
		return retData, err
	}

	bz, err := ioutil.ReadAll(resp.Body)
	log.Debugf("body: %s", string(bz))

	apiResp := &CommonResp{}
	err = json.Unmarshal(bz, apiResp)
	if err != nil {
		return
	}
	if apiResp.Code != 0 && apiResp.Code != http.StatusOK {
		err = fmt.Errorf("respCode: %d, msg: %s", apiResp.Code, apiResp.Msg)
		return
	}
	retData, err = json.Marshal(apiResp.Data)

	return retData, err
}
