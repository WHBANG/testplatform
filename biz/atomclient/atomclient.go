package atomclient

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	apiserver "git.supremind.info/products/atom/apiserver/client"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	privateTokenKey = "private"
)

type AtomClientConfig struct {
	Doamin string `json:"domain"`
	// APIServer     string `json:"api_server"`
	// Hodor         string `json:"hodor"`
	PersonalToken string `json:"personal_token"`
	Insecure      bool   `json:"insecure"`
}

type AtomClient struct {
	config          *AtomClientConfig
	token           string
	apiserverClient apiserver.Interface
}

func NewAtomClient(config AtomClientConfig) (*AtomClient, error) {
	client := &AtomClient{
		config: &config,
	}
	err := client.init()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return client, nil
}

func (c *AtomClient) init() error {
	err := c.loginHodor()
	if err != nil {
		return err
	}
	err = c.loginAPIServer()
	if err != nil {
		return err
	}
	return nil
}

func (c *AtomClient) Close() error {
	return c.apiserverClient.Close()
}

func (c *AtomClient) loginHodor() error {
	var endpoint string
	if c.config.Insecure {
		endpoint = fmt.Sprintf("http://hodor.%s", c.config.Doamin)
	} else {
		endpoint = fmt.Sprintf("https://hodor.%s", c.config.Doamin)
	}
	log.Println(endpoint, c.config.PersonalToken)
	req, e := http.NewRequest("GET", fmt.Sprintf("%s/login/gitlab", endpoint), nil)
	if e != nil {
		return errors.Wrap(e, "init http request failed")
	}
	req.Header.Set("Authorization", privateTokenKey+" "+c.config.PersonalToken)
	// req = req.WithContext(context.Background())

	resp, e := http.DefaultClient.Do(req)
	if e != nil {
		return errors.Wrap(e, "http request failed")
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http request failed, status [%d] %s", resp.StatusCode, resp.Status)
	}

	idt, _ := ioutil.ReadAll(resp.Body)
	// user, e := identity.UnsafeParseIDToken(string(idt))
	// if e != nil {
	// 	return errors.Wrap(e, "parse id token failed")
	// }
	log.Infof("token : %s", string(idt))
	c.token = string(idt)
	return nil
}

func (c *AtomClient) GetAPIServer() apiserver.Interface {
	return c.apiserverClient
}

func (c *AtomClient) loginAPIServer() error {
	if c.apiserverClient != nil {
		c.apiserverClient.Close()
	}
	var endpoint string
	if c.config.Insecure {
		endpoint = fmt.Sprintf("apiserver.%s", c.config.Doamin)
	} else {
		endpoint = fmt.Sprintf("apiserver.%s:443", c.config.Doamin)
	}
	var apiserverClient apiserver.Interface
	opts := []apiserver.Option{
		// apiserver.WithLogger(app.zlog),
		apiserver.WithConstantIDToken(c.token),
		apiserver.WithRetry(time.Second),
		apiserver.SetMaxRecvMsgSize(1 << 30),
		apiserver.SetMaxSendMsgSize(1 << 30),
	}

	if c.config.Insecure {
		opts = append(opts, apiserver.WithInsecure())
	} else {
		opts = append(opts, apiserver.WithDefaultCertPool())
	}

	apiserverClient, err := apiserver.New(endpoint, opts...)
	if err != nil {
		log.Error(err)
		return errors.Wrap(err, "can not connect to apiserver")
	}

	c.apiserverClient = apiserverClient

	return nil
}
