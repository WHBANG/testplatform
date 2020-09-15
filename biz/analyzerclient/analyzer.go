package analyzerclient

import (
	"context"

	log "github.com/sirupsen/logrus"

	"git.supremind.info/products/atom/proto/go/api"
	"git.supremind.info/testplatform/biz/atomclient"
)

const (
	resourceReference string = "2080ti-1" //todo 可配置
)

type AnalyzerClientConfig struct {
	atomclient.AtomClientConfig
	ResourceRef        string `json:"resource_ref"`
	ResourceRefCreator string `json:"resource_ref_creator"`
}

type AnalyzerDeployInfo struct {
	JobName   string
	Image     string
	ConfigMap map[string]string
	Args      string
}

type AnalyzerClient struct {
	atomC  *atomclient.AtomClient
	config *AnalyzerClientConfig
}

type JobInfo struct {
	Name    string
	Creator string
}

type AnalyzerInterface interface {
	CreateAnalyzer(info *AnalyzerDeployInfo) (*JobInfo, error)
	StartAnalyzer(info *JobInfo) error
	StopAnalyzer(info *JobInfo) error
	RemoveAnalyzer(info *JobInfo) error
	Status(info *JobInfo) error
	Close() error
}

func NewAnalyzerClient(ctx context.Context, config *AnalyzerClientConfig) (AnalyzerInterface, error) {
	atomC, err := atomclient.NewAtomClient(config.AtomClientConfig)
	if err != nil {
		log.Errorf("NewAtomClient  err: %s ", err)
		return nil, err
	}

	c := &AnalyzerClient{
		atomC:  atomC,
		config: config,
	}
	return c, nil
}

func (c *AnalyzerClient) Close() error {
	return c.atomC.Close()
}

func (c *AnalyzerClient) CreateAnalyzer(info *AnalyzerDeployInfo) (*JobInfo, error) {
	if info.Args == "" {
		if _, ok := info.ConfigMap["start.sh"]; ok {
			info.Args = "sh /workspace/mnt/config/start.sh"
		} else {
			info.Args = "sh /workspace/start.sh"
		}
	}

	as := c.atomC.GetAPIServer()
	j, err := as.Job().CreateJob(context.Background(), &api.CreateJobReq{
		Job: &api.Job{
			Metadata: &api.Metadata{
				Id:   api.NewID(),
				Name: info.JobName,
				// Creator: "wuzz",
			},
			Spec: &api.JobSpec{
				Kind: api.TrainingJob,
				JobCommon: &api.JobCommon{
					Image:  info.Image,
					Args:   []string{info.Args},
					Config: info.ConfigMap,
					// Env: map[string]string{
					// 	"test-env": "value",
					// },
					Package:  &api.ResourceReference{Name: c.config.ResourceRef, Creator: c.config.ResourceRefCreator},
					Mounting: &api.JobMounting{
						// Volumes:  []*api.VolumeMounting{{Volume: &api.ResourceReference{Name: "test-vol", Creator: hackUsername}}},
						// Datasets: []*api.DatasetVersionRef{{Dataset: "test-ds", Creator: hackUsername, Version: "test-version"}},
						// Storages: []*api.StorageMounting{{Storage: &api.ResourceReference{Name: "test-storage", Creator: hackUsername}, ReadOnly: false}},
						// // secrets
					},
				},
				JobInstruction: &api.JobInstruction{
					Training: &api.TrainingSpec{
						EnableJupyter: false,
						EnableSSH:     false,
						EnableFinder:  false,
						EnableLogger:  true,
					},
				},
			},
		},
	})
	if err != nil {
		log.Errorf("CreateJob err: %s ", err)
		return nil, err
	}
	log.Info(*j)

	return &JobInfo{
		Name:    j.GetName(),
		Creator: j.GetCreator(),
	}, nil
}

func (c *AnalyzerClient) StartAnalyzer(info *JobInfo) error {
	as := c.atomC.GetAPIServer()

	j, err := as.Job().StartJob(context.Background(), &api.StartJobReq{
		Name:    info.Name,
		Creator: info.Creator,
	})
	if err != nil {
		log.Errorf("StartJob  err: %s ", err)
		return err
	}
	log.Infof("start job:%s", *j)
	return nil
}

func (c *AnalyzerClient) StopAnalyzer(info *JobInfo) error {
	as := c.atomC.GetAPIServer()

	j, err := as.Job().StopJob(context.Background(), &api.StopJobReq{
		Name:    info.Name,
		Creator: info.Creator,
	})
	if err != nil {
		log.Errorf("StopJobs  err: %s ", err)
		return err
	}
	log.Infof("stop job:%s", *j)

	return nil
}
func (c *AnalyzerClient) RemoveAnalyzer(info *JobInfo) error {
	as := c.atomC.GetAPIServer()

	j, err := as.Job().RemoveJob(context.Background(), &api.RemoveJobReq{
		Name:    info.Name,
		Creator: info.Creator,
	})
	if err != nil {
		log.Errorf("RemoveJob  err: %s ", err)
		return err
	}
	log.Infof("remove job:%s", *j)
	return nil
}

//todo
func (c *AnalyzerClient) Status(info *JobInfo) error {
	as := c.atomC.GetAPIServer()

	j, err := as.Job().GetJob(context.Background(), &api.GetJobReq{
		Name:    info.Name,
		Creator: info.Creator,
	})
	if err != nil {
		log.Errorf("GetJob  err: %s ", err)
		return err
	}
	log.Infof("get job:%s", *j)

	// j.GetStatus().Phase
	return nil
}
