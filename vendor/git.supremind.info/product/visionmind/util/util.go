package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"gopkg.in/mgo.v2/bson"
	"qiniupkg.com/x/log.v7"
)

// JsonStr
func JsonStr(obj interface{}) string {
	raw, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(raw)
}

// ConvByJson
func ConvByJson(src interface{}, dest interface{}) error {
	tmpbs, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(tmpbs, dest)
}

// ConvByBson
func ConvByBson(src interface{}, dest interface{}) error {
	tmpbs, err := bson.Marshal(src)
	if err != nil {
		return err
	}
	return bson.Unmarshal(tmpbs, dest)
}

// GetLocalIP
func GetLocalIP() string {
	addrArr, err := net.InterfaceAddrs()
	if nil != err {
		return ""
	}
	for _, addr := range addrArr {
		ipnet, ok := addr.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() &&
			ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	return ""
}

func Md5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func ReadFile(path string) (bytes []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Errorf("os.Open(%s): %+v", path, err)
		return
	}
	defer file.Close()

	bytes, err = ioutil.ReadAll(file)
	if err != nil {
		log.Errorf("ioutil.ReadAll(%s): %+v", path, err)
		return
	}

	return
}

func GenUuid() string {
	return uuid.NewV4().String()
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//convert UTF-8 to GBK
func EncodeGBK(s string) (string, error) {
	I := strings.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

func Utf8ToHZGB2312(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func CopyFile(from, to string) error {
	fromInfo, err := os.Stat(from)
	if err != nil {
		return err
	}

	switch fromInfo.Mode() & os.ModeType {
	case os.ModeSymlink:
		link, err := os.Readlink(from)
		if err != nil {
			return err
		}
		return os.Symlink(link, to)
	default:
		if !fromInfo.Mode().IsRegular() {
			return errors.New(from + " is not a regular file")
		}

		fromFile, err := os.Open(from)
		if err != nil {
			return err
		}
		defer fromFile.Close()

		toFile, err := os.OpenFile(to, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		defer toFile.Close()
		_, err = io.Copy(toFile, fromFile)
		return err
	}
}

func CopyRecursively(from, to string) error {
	fromInfo, err := os.Stat(from)
	if err != nil {
		return err
	}
	if !fromInfo.IsDir() {
		return CopyFile(from, to)
	}

	infos, err := ioutil.ReadDir(from)
	if err != nil {
		return err
	}

	for _, info := range infos {
		srcPath := filepath.Join(from, info.Name())
		destPath := filepath.Join(to, info.Name())
		if info.IsDir() {
			err = os.MkdirAll(destPath, 0755)
			if err != nil && !os.IsExist(err) {
				return err
			}
			err = CopyRecursively(srcPath, destPath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcPath, destPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
