// +build ca

package client

import (
	"context"
	"crypto/md5"
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"git.supremind.info/product/visionmind/com/flow"
	"git.supremind.info/product/visionmind/com/flow/lease"
	"git.supremind.info/product/visionmind/com/go_sdk"
	"git.supremind.info/product/visionmind/util"
	mgoutil "github.com/qiniu/db/mgoutil.v3"
	"gopkg.in/mgo.v2/bson"
)

var signPubKeyBytes = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3mRnrZ00onoBX2TvcF/L
/jXq8essh7PdSsqIPAa4IyLyDpvcjEzkr7mhrRelj5hcaJIS9DjjG1CFvyKCWf15
psvKJ4XrEEEnS2EGnL6GTwXqCNhYvKl3gCFmJwah2VArBUGHrUY1q6bYhCDJ2N1q
11egBw6Ka7ZCmLQTcjK1dU5w7RaXofXWmeuDsvUroRRPqfVcJ3OLO6rXKTZOUGwh
mL26WzPpzli4QhgvGNRMx5ILbsO4pHCqQbdoyyb9pgaNGURHILZkXJsXFax89/7D
6qCPiDrZSEq9z7T6xVBx/5hudkNGr3OZ8GNp+55SB8/CuHqeASrKyjV5Rhf62I6p
fwIDAQAB
-----END RSA PUBLIC KEY-----`)

const syncInterval time.Duration = 300

type CAClientConfig struct {
	CaServerHost string `json:"host"`
}

type CAClient struct {
	CAClientConfig

	pubKey *rsa.PublicKey

	group  string // client的group, 同一个group共享quota
	id     string // client的unique-id
	dbColl *mgoutil.Collection

	lock sync.Mutex
}

func NewCAClient(cfg *CAClientConfig, tasksColl *mgoutil.Collection) (*CAClient, error) {
	if cfg == nil {
		return nil, errors.New("ca config required")
	}
	pubKey, err := util.BytesToPublicKey(signPubKeyBytes)
	if err != nil {
		return nil, err
	}
	dbLiveServers := tasksColl.Database.Session.LiveServers()
	sort.Strings(dbLiveServers)
	return &CAClient{
		CAClientConfig: *cfg,
		pubKey:         pubKey,
		id:             fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%d", util.GatherDeivceInfo().ToHexString(), os.Getpid())))),
		group:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%d", strings.Join(dbLiveServers, ","), tasksColl.Database.Name)))),
		dbColl:         tasksColl,
	}, nil
}

func (c *CAClient) Start(ctx context.Context) {
	println("starting ca loop")
	c.do()
	go c.syncQuotaLoop(ctx)
}

func (c *CAClient) CheckQuota(name string) (enough bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	quotas, err := c.fetchQuota()
	if err != nil {
		// 心跳校验不通过, 则直接退出程序
		fmt.Printf("fetch quota failed: %v\n", err)
		return
	}
	runningTasks, err := c.fetchNonStoppedTasks()
	if err != nil {
		// db 不可访问时, 忽略这个错误
		return
	}
	_, quotaLeft := c.pickTaskShouldBeOff(runningTasks, quotas)
	return quotaLeft[name] > 0
}

func (c *CAClient) FetchQuotas() (quotas []go_sdk.Quota, err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	quotas, err = c.fetchQuota()
	if err != nil {
		fmt.Printf("fetch quota failed: %v\n", err)
		return
	}
	return
}

func (c *CAClient) do() {
	c.lock.Lock()
	defer c.lock.Unlock()
	quotas, err := c.fetchQuota()
	if err != nil {
		// 心跳校验不通过, 则直接退出程序
		fmt.Printf("fetch quota failed: %v\n", err)
		os.Exit(1)
	}
	runningTasks, err := c.fetchNonStoppedTasks()
	if err != nil {
		// db 不可访问时, 忽略这个错误
		fmt.Printf("fetch running tasks failed: %v\n", err)
		return
	}
	tasksShouldBeOff, _ := c.pickTaskShouldBeOff(runningTasks, quotas)
	c.turnOffTask(tasksShouldBeOff)
}

func (c *CAClient) syncQuotaLoop(ctx context.Context) {
	// NOTE 该函数理论上不应该退出, 上游应该关注这个context, 如果done了, 则需要重跑
	ticker := time.NewTicker(time.Second * syncInterval)
	for {
		select {
		case <-ticker.C:
			c.do()
		case <-ctx.Done():
			// 上游cancel掉当前循环, 直接退出
			return
		}
	}
}

func (c *CAClient) fetchQuota() (quotas []go_sdk.Quota, err error) {
	// TODO 区分出哪些错误是需要重试的, 比如证书验证失败的, 则直接退出比较好
	randomToken := strconv.Itoa(rand.Int()) // 需要判断当前的quota值是否合法, 防止拿着同样的resp回放
	request, _ := http.NewRequest("GET", fmt.Sprintf("http://%s/quota", c.CaServerHost), nil)
	request.Header.Set("X-SupreMind-Group", c.group)
	request.Header.Set("X-SupreMind-ID", c.id)
	request.Header.Set("X-SupreMind-Token", randomToken)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		err = fmt.Errorf("invalid ca server host")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("ca server resp failed")
		return
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to read ca resp")
		return
	}
	signature, err := hex.DecodeString(resp.Header.Get("X-SupreMind-Signature"))
	if err != nil {
		err = fmt.Errorf("signature is broken")
		return
	}

	verifyBs := []byte(strings.Join([]string{c.group, c.id, randomToken}, "."))
	err = util.VerySignWithSha1(verifyBs, signature, c.pubKey)
	if err != nil {
		err = fmt.Errorf("signature is invalid")
		return
	}
	err = json.Unmarshal(bs, &quotas)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal response")
		return
	}
	return
}

func (c *CAClient) fetchNonStoppedTasks() (tasks []*flow.Task, err error) {
	coll := c.dbColl.CopySession()
	defer func() {
		_ = coll.CloseSession()
	}()

	tasks = []*flow.Task{}
	findM := bson.M{
		"status": bson.M{"$ne": lease.STATUS_STOPPED},
		"$or":    []bson.M{{"is_deleted": bson.M{"$exists": false}}, {"is_deleted": false}},
	}
	err = coll.Find(findM).All(&tasks)
	return
}

func (c *CAClient) pickTaskShouldBeOff(tasks []*flow.Task, quotas []go_sdk.Quota) (tasksShouldBeOff []*flow.Task, quotaLeft map[string]int) {
	// quota转成map
	quotaLeft = make(map[string]int)
	for _, quota := range quotas {
		quotaLeft[quota.Name] = quota.Quota
	}

	// NOTE 选出需要关闭的task, 安装创建的时候, 先创建的继续保活
	sort.SliceStable(tasks, func(i, j int) bool { return tasks[i].CreatedAt.Before(tasks[j].CreatedAt) })
	for _, task := range tasks {
		c, exist := quotaLeft[task.Type]
		if !exist || c == 0 {
			tasksShouldBeOff = append(tasksShouldBeOff, task)
		} else {
			quotaLeft[task.Type]--
		}
	}

	return
}

func (c *CAClient) turnOffTask(tasks []*flow.Task) (err error) {
	coll := c.dbColl.CopySession()
	defer func() {
		_ = coll.CloseSession()
	}()
	for _, task := range tasks {
		if task.Status == lease.STATUS_FAILED {
			continue
		}
		err = c.dbColl.Update(bson.M{"id": task.ID}, bson.M{"$set": bson.M{"status": lease.STATUS_STOPPED, "ver": 0, "holder": "", "updated_at": time.Now(), "switched_at": time.Now()}})
		if err != nil {
			fmt.Println("shutdown task failed...")
			return
		}
		fmt.Println("shutdown task success: ", task.ID)
	}
	return

}
