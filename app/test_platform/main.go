package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"runtime"

	"net/http"
	_ "net/http/pprof"

	_ "git.supremind.info/testplatform/app/test_platform/docs"
	"git.supremind.info/testplatform/biz/analyzerclient"
	"git.supremind.info/testplatform/biz/jenkinsclient"
	"git.supremind.info/testplatform/biz/service"
	"git.supremind.info/testplatform/biz/service/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
 
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

type Config struct {
	Host                 string                              `json:"host"`
	Mode                 string                              `json:"mode"`
	AnalyzerClientConfig analyzerclient.AnalyzerClientConfig `json:"analyzer_client"`
	ConfigFiles          map[string]string                   `json:"config_files"`
	Mongodb              db.MongodbConfig                    `json:"mongodb"`
	Jenkins              jenkinsclient.Config                `json:jenkins`
	VMRClient            service.VMRClient                   `json:"vmr_client"`
}

func checkConfigFiles(configM map[string]string) error {
	if analyzerConfig, ok := configM["analyzer.conf"]; ok {
		var m = make(map[string]interface{})
		err := json.Unmarshal([]byte(analyzerConfig), &m)
		if err != nil {
			err = fmt.Errorf("unmarshal analyzer.conf error:%s", err)
			return err
		}
	}
	return nil
}

func main() {

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	runtime.GOMAXPROCS(runtime.NumCPU())

	var configFile string
	flag.StringVar(&configFile, "f", "service.conf", "config file path")
	flag.Parse()

	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Panicf("config file read err: %s ", err)
	}
	log.Infof("config load, %s", string(configData))

	var conf Config
	err = json.Unmarshal(configData, &conf)
	if err != nil {
		log.Panicf("config unmarshal err: %s ", err)
	}
	if conf.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	fileMap := make(map[string]string)
	for k, v := range conf.ConfigFiles {
		data, err := ioutil.ReadFile(v)
		if err != nil {
			log.Panicf("read file %s err: %s ", v, err)
		}
		fileMap[k] = string(data)
	}
	err = checkConfigFiles(fileMap)
	if err != nil {
		log.Panicf("check config files error: %s ", err)
	}

	err = db.InitDB(&conf.Mongodb)
	if err != nil {
		log.Panic(err)
	}
	session, _ := db.GetMgoDBSession()
	defer session.Close()

	analyzerC, err := analyzerclient.NewAnalyzerClient(context.Background(), &conf.AnalyzerClientConfig)
	if err != nil {
		log.Panicf("NewAtomClient  err: %s ", err)
	}
	log.Println(analyzerC)

	r := gin.Default()
	group := r.Group("v1")

	_, err = service.NewTestPlatformSvc(context.Background(), &service.Config{VMRClient: conf.VMRClient}, group)
	if err != nil {
		log.Panicf("NewTestPlatformSvc err: %s ", err)
	}

	_, err = service.NewEngineTestSvc(context.Background(), group, analyzerC, fileMap)
	if err != nil {
		log.Panicf("NewEngineTestSvc err: %s ", err)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	jenkinsclient.JenkinsClientSvc(context.Background(), group, conf.Jenkins)

	err = r.Run(conf.Host)
	log.Panic(err)
}
