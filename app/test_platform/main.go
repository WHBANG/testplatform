package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"runtime"

	"git.supremind.info/testplatform/biz/analyzerclient"
	"git.supremind.info/testplatform/biz/atomclient"
	"git.supremind.info/testplatform/biz/service"
	"git.supremind.info/testplatform/biz/service/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Host        string                      `json:"host"`
	AtomClient  atomclient.AtomClientConfig `json:"atom_client"`
	ConfigFiles map[string]string           `json:"config_files"`
	Mongodb     db.MongodbConfig            `json:"mongodb"`
}

func main() {
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
	fileMap := make(map[string]string)
	for k, v := range conf.ConfigFiles {
		data, err := ioutil.ReadFile(v)
		if err != nil {
			log.Panicf("read file %s err: %s ", v, err)
		}
		fileMap[k] = string(data)
	}

	err = db.InitDB(&conf.Mongodb)
	if err != nil {
		log.Panic(err)
	}
	session, _ := db.GetMgoDBSession()
	defer session.Close()

	atomC, err := atomclient.NewAtomClient(conf.AtomClient)
	if err != nil {
		log.Panicf("NewAtomClient  err: %s ", err)
	}
	defer atomC.Close()

	analyzerC, err := analyzerclient.NewAnalyzerClient(context.Background(), atomC)
	if err != nil {
		log.Panicf("NewAtomClient  err: %s ", err)
	}
	log.Println(analyzerC)

	r := gin.Default()
	group := r.Group("v1")

	_, err = service.NewTestPlatformSvc(context.Background(), &service.Config{}, group)
	if err != nil {
		log.Panicf("NewTestPlatformSvc err: %s ", err)
	}

	_, err = service.NewTestSvc(context.Background(), group, analyzerC)
	if err != nil {
		log.Panicf("NewTestSvc err: %s ", err)
	}

	err = r.Run(conf.Host)
	log.Panic(err)
}
