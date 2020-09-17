package util

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"regexp"
	"sort"
	"strings"
)

func IsContainer() bool {
	f, err := os.Open("/proc/1/cgroup")
	if os.IsNotExist(err) {
		// if file not exsit
		return true
	}
	reader := bufio.NewReader(f)
	l, isPrefix, err := reader.ReadLine()
	if err != nil {
		// read the file error
		return true
	}
	if isPrefix {
		// line is too long, not a general linux
		return true
	}
	if bytes.Contains(l, []byte(":/docker/")) {
		// contains the `docker`
		// TODO but it may run in VM
		return true
	}
	return false
}

// DeviceInfo collects the
type DeviceInfo struct {
	Cpu          string `json:"cpu"`
	RootDiskUUID string `json:"root_disk_uuid"`
	Net          string `json:"network"`
}

func GatherDeivceInfo() *DeviceInfo {
	d := &DeviceInfo{}
	d.GetCpuInfo()
	d.GetRootDiskUUID()
	d.GetNetMacIp(false)
	return d
}

func NewDeviceInfoFromHexString(s string) (*DeviceInfo, error) {
	bs, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}

	d := &DeviceInfo{}
	err = json.Unmarshal(bs, d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (d *DeviceInfo) Equal(d2 *DeviceInfo) bool {
	bs, _ := json.Marshal(d)
	bs2, _ := json.Marshal(d2)
	return bytes.Equal(bs, bs2)
}

func (d *DeviceInfo) ToHexString() string {
	bs, _ := json.Marshal(d)
	return hex.EncodeToString(bs)
}

func (d *DeviceInfo) GetCpuInfo() (err error) {
	bs, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return errors.New("failed to read cpu info")
	}

	//使用确定不变的信息 CPU制造商|CPU产品系列代号|CPU属于其系列中的哪一代号|CPU属于的名字、编号、主频|每个CPU的逻辑核数|CPU的物理核数
	re := regexp.MustCompile(`vendor_id|cpu family|model|model name|siblings|cpu cores`)
	fixedContents := [][]byte{}
	for _, line := range bytes.Split(bs, []byte("\n")) {
		if re.Match(line) {
			fixedContents = append(fixedContents, line)
		}
	}
	bs = bytes.Join(fixedContents, []byte("\n"))

	d.Cpu = fmt.Sprintf("%x", md5.Sum(bs))
	return

}

func (d *DeviceInfo) GetNetMacIp(needIp bool) (err error) {
	result := ""
	for _, n := range []string{"eth0", "em1", "enp4s0"} {
		inter, err := net.InterfaceByName(n)
		if err != nil {
			continue
		}

		//mac地址
		mac := inter.HardwareAddr.String()
		result = fmt.Sprintf("%s,%s", n, mac)
		if needIp {
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}

			ips := []string{}
			for _, addr := range addrs {
				ipNet, isValidIpNet := addr.(*net.IPNet)
				if isValidIpNet && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						ips = append(ips, ipNet.IP.To4().String())
					}
				}
			}
			if len(ips) > 0 {
				sort.Strings(ips)
				result = fmt.Sprintf("%s,%s", result, strings.Join(ips, ","))
			}
		}

		d.Net = result
		return nil
	}
	return errors.New("root deivce uuid not found")
}

func (d *DeviceInfo) GetRootDiskUUID() (err error) {
	bs, err := ioutil.ReadFile("/etc/fstab")
	if err != nil {
		return
	}
	re := regexp.MustCompile(`^(?P<prefix>UUID=)(?P<uuid>[a-f0-9\-]{36})\ +(?P<mount>/)\ +`)
	/*
		必须严格匹配根目录上挂载的磁盘
		fmt.Println(re.FindAllStringSubmatch(`UUID=4cf77db1-60b7-42c1-98fc-ff11e8af32c4 /               ext4    errors=remount-ro 0       1`, 2))
		fmt.Println(re.FindAllStringSubmatch(`#UUID=4cf77db1-60b7-42c1-98fc-ff11e8af32c4 /               ext4    errors=remount-ro 0       1`, 2))
		fmt.Println(re.FindAllStringSubmatch(`  UUID=4cf77db1-60b7-42c1-98fc-ff11e8af32c4 /               ext4    errors=remount-ro 0       1`, 2))
		fmt.Println(re.FindAllStringSubmatch(`UUID=4cf77db1-60b7-42c1-98fc-ff11e8af32c4 /mnt             ext4    errors=remount-ro 0       1`, 2))
	*/
	for _, s := range strings.Split(string(bs), "\n") {
		result := re.FindAllStringSubmatch(s, 3)
		if len(result) > 0 && len(result[0]) == 4 {
			d.RootDiskUUID = result[0][2]
			return
		}
	}

	return errors.New("root deivce uuid not found")
}
