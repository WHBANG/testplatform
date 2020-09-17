package util

import (
	"context"
	"errors"

	bb "git.supremind.info/product/visionmind/com/black_box"
	gclient "git.supremind.info/product/visionmind/com/go_sdk"
	log "qiniupkg.com/x/log.v7"
)

var VersionKey = "__visionmind_version__"

type VisionMindVersion struct {
	Software struct {
		App struct {
			Surveillance VisionMindProduct `json:"surveillance"`
		} `json:"app"`
		Platform struct {
			CA  VisionMindProduct `json:"ca"`
			VMR VisionMindProduct `json:"vmr"`
		} `json:"platform"`
	} `json:"software"`
}

type VisionMindProduct struct {
	Version string `json:version`
}

func GetVisionmindVersion(client gclient.IFlowClient) (version *VisionMindVersion, err error) {
	if client == nil {
		err = errors.New("flowClient is nil")
		return
	}

	config, err := client.ConfigCheck(context.Background(), VersionKey)
	if err != nil {
		log.Errorf(" client.ConfigCheck(%s), err: %v", VersionKey, err)
		return
	}

	var vmVersion VisionMindVersion
	err = ConvByJson(config.Info, &vmVersion)
	if err != nil {
		log.Errorf("convert VisionMindVersion failed, err: %v", err)
		return
	}

	version = &vmVersion
	return
}

func Collection(flowClient gclient.IFlowClient, blackBoxClient gclient.IBlackBoxClient,
	event bb.Event, name, info string, reason ...string) (err error) {
	if blackBoxClient == nil {
		log.Info("blackBoxClient is nil")
		return
	}

	version := &VisionMindVersion{}

	if flowClient != nil {
		version, err = GetVisionmindVersion(flowClient)
		if err != nil {
			return
		}
	}

	switch event.Product {
	case bb.ProductJiaotong:
		event.ProductVersion = version.Software.App.Surveillance.Version
	case bb.ProductVMR:
		event.ProductVersion = version.Software.Platform.VMR.Version
	case bb.ProductCA:
		event.ProductVersion = version.Software.Platform.CA.Version
	default:
	}

	if info == "success" || info == "fail" {
		info = name + " " + info
	}
	msg := bb.Message{
		Info: info,
	}
	if len(reason) > 0 {
		msg.Reason = reason[0]
	}

	event.Name = name
	event.Message = msg

	err = blackBoxClient.SetEvent(&event)
	if err != nil {
		log.Errorf("add collection(%+v): %+v", event, err)
	} else {
		log.Infof("add collection(%+v)", event)
	}

	return
}
