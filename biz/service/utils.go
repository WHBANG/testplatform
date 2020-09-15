package service

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

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

/*
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
*/
