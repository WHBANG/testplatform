package jenkinsclient

import (
	"context"
	"strconv"

	"git.supremind.info/testplatform/biz/service/proto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var Jc JenkinsClient

// @Summary show build result
// @Description show build result by build number
// @Accept json
// @Param number query string true "number"
// @Success 200 {object}  proto.CommonRes{data=map[string]string}
// @Router /jenkins/show/{number} [GET]
func JenkinsClientShowBuildResultHandler(c *gin.Context) {

	job, err := Jc.Cli.GetJob(Jc.Name)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		log.Println("get job ", Jc.Name, "error:", err)
		return
	}
	jobIDStr := c.Param("number")
	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		log.Println("paremter number error:", jobIDStr)
		return
	}
	build, err := job.GetBuild(jobID)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		log.Println("this build is wrong or not exist! ", err)
	} else if "" == build.GetResult() {
		proto.DefaultRet(c, map[string]string{"Result": "building, please wait a moment"}, nil)
	} else {
		proto.DefaultRet(c, map[string]string{"Result": build.GetResult()}, nil)
		log.Println("the build result is :", build.GetResult())
	}
}

// @Summary get the last successful build
// @Description show the last successful build number
// @Accept json
// @Success 200 {object} proto.CommonRes{data=map[string]int64}
// @Router /jenkins/last [GET]
func JenkinsClientGetLastBuildResultHandler(c *gin.Context) {

	job, err := Jc.Cli.GetJob(Jc.Name)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		log.Println("get job ", Jc.Name, "error:", err)
		return
	}
	lastSuccessBuild, err := job.GetLastSuccessfulBuild()
	if err != nil {
		proto.DefaultRet(c, nil, err)
		log.Println("Last sueecssBuild not exist")
		return
	}
	if "SUCCESS" == lastSuccessBuild.GetResult() {
		proto.DefaultRet(c, map[string]int64{"buildNumber": lastSuccessBuild.GetBuildNumber()}, nil)
		log.Println("the last build succeeded number is:", lastSuccessBuild.GetBuildNumber())
	}
}

// @Summary build
// @Description do build
// @Accept json
// @Param example body jenkinsclient.Parmeter true "Parmeter"
// @Success 200 {object} proto.CommonRes{data=map[string]int64}
// @Router /jenkins/build [POST]
func JenkinsClientBuildHandler(c *gin.Context) {

	var reqJSON Parmeter
	err := c.ShouldBindJSON(&reqJSON)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		log.Println("read post body error! ", err)
		return
	}
	log.Printf("%+v", reqJSON)

	id, err := Jc.BuildJob(&reqJSON, Jc.config)
	if err != nil {
		proto.DefaultRet(c, nil, err)
		log.Println("build job error! ", err)
		return
	}
	proto.DefaultRet(c, map[string]int64{"buildNumber": id}, nil)
	log.Println("this time build job id: ", id)

}

func JenkinsClientSvc(ctx context.Context, group *gin.RouterGroup, conf Config) {

	err := JenkinsInit(conf, &Jc)
	if err != nil {
		log.Println("init error:", err)
		return
	}
	group.GET("/jenkins/show/:number", JenkinsClientShowBuildResultHandler)
	group.POST("/jenkins/build", JenkinsClientBuildHandler)
	group.GET("/jenkins/last", JenkinsClientGetLastBuildResultHandler)
}
