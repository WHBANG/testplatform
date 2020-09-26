package jenkinsclient

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bndr/gojenkins"
	log "github.com/sirupsen/logrus"
)

/*
	"Url":"jenkins.supremind.io"
	"JobName":"analyzer-io-build"
	"Username":"wanghao"
	"Token":"110fab2a9658a151ed8d7df4a6c2fba866"
*/
type Config struct {
	Host     string `json:"host"`
	JobName  string `json:"job_name"`
	Username string `json:"user_name"`
	Token    string `json:"token"`
}

type Model struct {
	CrowdCountModel            string `bson:"crowd_count_model.tronmodel" json:"crowd_count_model.tronmodel"`
	BannerDetectEastModel      string `bson:"banner_detect_east_model.tronmodel" json:"banner_detect_east_model.tronmodel"`
	BannerDetectOdModel        string `bson:"banner_detect_od_model.tronmodel" json:"banner_detect_od_model.tronmodel"`
	HeadCountModel             string `bson:"head_count_model.tronmodel" json:"head_count_model.tronmodel"`
	FightClassifyLocalModel    string `bson:"fight_classify_local_model.tronmodel" json:"fight_classify_local_model.tronmodel"`
	QueueCountLocalModel       string `bson:"queue_count_local_model.tronmodel" json:"queue_count_local_model.tronmodel"`
	CrowdRegionCountLocalModel string `bson:"crowd_region_count_local_model.tronmodel" json:"crowd_region_count_local_model.tronmodel"`
}

type Models struct {
	Models Model `bson:"models" json:"models"`
}

type Parmeter struct {
	AnalyzerIoBaseImage string `json:"analyzer_io_base_image" bson:"analyzer_io_base_image"`
	Branch              string `json:"branch" bson:"branch"`
	ModelsConfig        string `json:"models_config" bson:"models_config"`
	AnalyzerType        string `json:"analyzer_type" bson:"analyzer_type"`
	ImageName           string `json:"image_name"`
}

type JenkinsClientImp interface {
	BuildJob(string) error
}

type JenkinsClient struct {
	Cli    *gojenkins.Jenkins
	Name   string
	config Config
}

func JenkinsInit(config Config, jc *JenkinsClient) error {

	var build strings.Builder
	fmt.Fprintf(&build, "http://%s/", config.Host)
	/*
	*  用户名为userID，password为token
	 */
	cli := gojenkins.CreateJenkins(nil, build.String(), config.Username, config.Token)
	c, err := cli.Init()
	if err != nil {
		return err
	}
	log.Println("jenkins init success")
	jc.Cli = c
	jc.Name = config.JobName
	jc.config = config
	return nil
}

func (j *JenkinsClient) BuildJob(p *Parmeter, config Config) (int64, error) {

	log.Println(j.Name, ":building...")

	job, err := j.Cli.GetJob(j.Name)
	if err != nil {
		log.Println("GetJob error:", err)
		return 0, err
	}
	b, err := job.GetLastBuild()
	var number int64
	if err != nil {
		number = 1
	} else {
		number = b.GetBuildNumber() + 1
	}

	reqMap := make(map[string]interface{})
	reqMap["ANALYZER_IO_BASE_IMAGE"] = p.AnalyzerIoBaseImage
	reqMap["BRANCH"] = p.Branch
	data, _ := json.Marshal(p.ModelsConfig)
	reqMap["MODELS_CONFIG"] = string(data)
	reqMap["ANALYZER_TYPE"] = p.AnalyzerType
	reqMap["IMAGE_NAME"] = p.ImageName
	num, err := j.Cli.BuildJob(j.Name, reqMap)
	fmt.Println(num)
	return number, nil
}

/*
func JenkinsClientShowBuildResultHandler(w http.ResponseWriter, r *http.Request) {

	jc, err := JenkinsInit()
	if err != nil {
		log.Println("failed to init jenkins", err)
	}
	job, err := jc.Cli.GetJob(jc.Name)
	if err != nil {
		log.Println("get job ", jc.Name, "error:", err)
		return
	}
	jobIDStr := r.URL.Query().Get("id")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		log.Println("paremter id error:", jobIDStr)
		return
	}
	build, err := job.GetBuild(jobID)
	if err != nil {
		log.Println("this build is wrong or not exist! ", err)
	} else {
		log.Println("the build result is :", build.GetResult())
	}
}

//get the last successful build
func JenkinsClientGetLastBuildResultHandler(w http.ResponseWriter, r *http.Request) {

	jc, err := JenkinsInit()
	if err != nil {
		log.Println("failed to init jenkins", err)
	}
	job, err := jc.Cli.GetJob(jc.Name)
	if err != nil {
		log.Println("get job ", jc.Name, "error:", err)
		return
	}
	lastSuccessBuild, err := job.GetLastSuccessfulBuild()
	if err != nil {
		log.Println("Last sueecssBuild not exist")
	}
	if "SUCCESS" == lastSuccessBuild.GetResult() {
		log.Println("the last build succeeded id is:", lastSuccessBuild.GetBuildNumber())
	}
}

// build
func JenkinsClientBuildHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	paramter, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("read post body error! ", err)
		return
	}
	var p Parmeter
	if err = json.Unmarshal(paramter, &p); nil != err {
		log.Println("Unmarshal Error ", err)
		return
	}
	id, err := jc.BuildJob(&p)
	if err != nil {
		log.Println("build job error! ", err)
		return
	}
	log.Println("this time build job id: ", id)
}

func JenkinsClientSvc(group *gin.RouterGroup) {

	http.HandleFunc("/show", JenkinsClientShowBuildResultHandler)
	http.HandleFunc("/build", JenkinsClientBuildHandler)
	http.HandleFunc("/last", JenkinsClientGetLastBuildResultHandler)
	err := http.ListenAndServe(":9899", nil)
	if err != nil {
		log.Println("failed to start jenkins server")
		return
	}
	log.Println("Listening and serving HTTP on :9899")
}
*/
