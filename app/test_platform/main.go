package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"runtime"

	"git.supremind.info/testplatform/biz/analyzerclient"
	"git.supremind.info/testplatform/biz/atomclient"
	"git.supremind.info/testplatform/biz/service"
	"git.supremind.info/testplatform/biz/service/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

type Config struct {
	Host        string                      `json:"host"`
	AtomClient  atomclient.AtomClientConfig `json:"atom_client"`
	ConfigFiles map[string]string           `json:"config_files"`
	Mongodb     struct {
		Host string `json:"host"`
		DB   string `json:"db"`
	} `json:"mongodb"`
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
		os.Exit(-1)
	}
	log.Infof("config load, %s", string(configData))

	var conf Config
	err = json.Unmarshal(configData, &conf)
	if err != nil {
		log.Panicf("config unmarshal err: %s ", err)

		os.Exit(-1)
	}
	fileMap := make(map[string]string)
	for k, v := range conf.ConfigFiles {
		data, err := ioutil.ReadFile(v)
		if err != nil {
			log.Panicf("read file %s err: %s ", v, err)
		}
		fileMap[k] = string(data)
	}

	session, err := mgo.Dial(conf.Mongodb.Host)
	if err != nil {
		log.Panicf("mongo dial error:%s ", err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	imageMgnt, err := db.NewMongoImage(session.DB(conf.Mongodb.DB))
	if err != nil {
		log.Panicf("NewMongoImage error:%s ", err)
	}

	atomC, err := atomclient.NewAtomClient(conf.AtomClient)
	if err != nil {
		log.Panicf("NewAtomClient  err: %s ", err)
	}
	defer atomC.Close()

	analyzerPlat, err := analyzerclient.NewAnalyzerClient(context.Background(), atomC)
	if err != nil {
		log.Panicf("NewAtomClient  err: %s ", err)
	}
	log.Println(analyzerPlat)

	r := gin.Default()
	group := r.Group("v1")

	_, err = service.NewTestPlatformSvc(context.Background(), &service.Config{}, group, imageMgnt)
	if err != nil {
		log.Panicf("NewTestPlatformSvc err: %s ", err)
	}

	err = r.Run(conf.Host)
	log.Panic(err)
}
