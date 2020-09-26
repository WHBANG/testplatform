package service

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"git.supremind.info/product/visionmind/com/flow"
	client "git.supremind.info/product/visionmind/com/go_sdk"
	"git.supremind.info/testplatform/biz/service/proto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TargetHost struct {
	Host    string `json:"host"`
	IsHTTPS bool   `json:"is_https"`
	CAPath  string `json:"ca_path"`
}

func Forward(c *gin.Context, targetHost *TargetHost) {
	HostReverseProxy(c.Writer, c.Request, targetHost)
}

func MiddleWare(cl *Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		username, passwd, isOK := c.Request.BasicAuth()
		if (!isOK) || username == "" || passwd == "" {
			// c.JSON(http.StatusOK, Response{
			// 	Code:    400002,
			// 	Message: "用户未登录",
			// })
			// c.Abort()
			// return
			c.Request.SetBasicAuth(cl.Username, cl.Password)
		}
		c.Next()
	}

}

func Forward01(host, way, param string) (data interface{}, err error) {

	var (
		urlstr string
		req    *http.Request
	)
	urlstr = "http://" + host + "/" + param
	if way == "get" {
		req, err = http.NewRequest("GET", urlstr, nil)
	} else {
		req, err = http.NewRequest("POST", urlstr, nil)
	}
	if err != nil {
		log.Println("NewRequest Error: ", err)
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Client.Do Error: ", err)
		return nil, err
	}
	var resp proto.GetTaskRes
	if res != nil {
		defer res.Body.Close()

		buffer := bytes.NewBuffer(make([]byte, 1024*8))
		_, err := io.Copy(buffer, res.Body)
		if err != nil {
			log.Println("Buffer Capacity Error: ", err)
			return nil, err
		}
		respData := buffer.Bytes()
		str := strings.Trim(string(respData), "\x00 \n")
		log.Println(str)
		err = json.Unmarshal([]byte(str[:len(str)]), &resp)
		if err != nil {
			log.Println("Json Unmarshal Error: ", err)
			return nil, err
		}
	}
	return resp.Data, nil
}

func Forward02(host, way, param string) (data []proto.JTEventInfo, err error) {

	var (
		urlstr string
		req    *http.Request
	)
	urlstr = "http://" + host + "/" + param
	if way == "get" {
		req, err = http.NewRequest("GET", urlstr, nil)
	} else {
		req, err = http.NewRequest("POST", urlstr, nil)
	}
	if err != nil {
		log.Println("NewRequest Error: ", err)
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Client.Do Error: ", err)
		return nil, err
	}
	var resp proto.GetJTEventRes
	if res != nil {
		defer res.Body.Close()

		buffer := bytes.NewBuffer(make([]byte, 1024*8))
		_, err := io.Copy(buffer, res.Body)
		if err != nil {
			log.Println("Buffer Capacity Error: ", err)
			return nil, err
		}
		respData := buffer.Bytes()
		str := strings.Trim(string(respData), "\x00 \n")
		log.Println(str)
		err = json.Unmarshal([]byte(str[:len(str)]), &resp)
		if err != nil {
			log.Println("Json Unmarshal Error: ", err)
			return nil, err
		}
	}
	return resp.Content, nil
}

func HostReverseProxy(w http.ResponseWriter, req *http.Request, targetHost *TargetHost) {
	host := ""
	if targetHost.IsHTTPS {
		host = host + "https://"
	} else {
		host = host + "http://"
	}
	remote, err := url.Parse(host + targetHost.Host)
	if err != nil {
		log.Errorf("err:%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)

	if targetHost.IsHTTPS {
		tls, err := GetVerTLSConfig(targetHost.CAPath)
		if err != nil {
			log.Errorf("https crt error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var pTransport http.RoundTripper = &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(10))
				if err != nil {
					return nil, err
				}
				return c, nil
			},
			ResponseHeaderTimeout: time.Second * time.Duration(10),
			TLSClientConfig:       tls,
		}
		proxy.Transport = pTransport
	}

	proxy.ServeHTTP(w, req)
}

var TlsConfig *tls.Config

func GetVerTLSConfig(CAPath string) (*tls.Config, error) {
	caData, err := ioutil.ReadFile(CAPath)
	if err != nil {
		log.Errorf("read wechat ca fail", err)
		return nil, err
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	TlsConfig = &tls.Config{
		RootCAs: pool,
	}
	return TlsConfig, nil
}

type CreateSubDeviceParams struct {
	Type           int                `json:"type"`
	OrganizationID string             `json:"organization_id"`
	DeviceID       string             `json:"device_id"`
	Channel        int                `json:"channel"`
	Attribute      SubDeviceAttribute `json:"attribute"`
}

type SubDeviceAttribute struct {
	Name              string `json:"name"`
	DiscoveryProtocol int    `json:"discovery_protocol"`
	UpstreamURL       string `json:"upstream_url"`
	Vendor            int    `json:"vendor"`
}

func CreateSubDevice(flowCli client.IFlowClient, deviceName, url string, maxChan int, globalDevId, namePrefix string) (device *flow.Device, err error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(maxChan-1)))
	if err != nil {
		log.Errorf("rand.Int(rand.Reader, big.NewInt(TestArg.MaxChannel)): %+v", err)
		return
	}
	channel := int(n.Int64()) + 1
	subDevice := CreateSubDeviceParams{
		Type:           1,
		OrganizationID: "000000000000000000000000",
		DeviceID:       globalDevId,
		Channel:        channel,
		Attribute: SubDeviceAttribute{
			Name:              namePrefix + "_" + deviceName,
			DiscoveryProtocol: 2,
			UpstreamURL:       url,
			Vendor:            1,
		},
	}

	device, err = flowCli.CreateDevice(context.Background(), subDevice)
	if err != nil {
		log.Errorf("t.FlowClient.CreateDevice(%+v):%+v", subDevice, err)
	}
	return
}
