package util

import (
	"bytes"
	"flag"
	"github.com/ghodss/yaml"
	xlog "github.com/qiniu/xlog.v1"
	"io/ioutil"
	"qbox.us/cc"
)

var (
	confName *string
	NL       = []byte{'\n'}
	ANT      = []byte{'#'}
)

func Init(cflag, app, default_conf string) {

	confDir, _ := cc.GetConfigDir(app)
	confName = flag.String(cflag, confDir+"/"+default_conf, "the yml config")
}

func Load() (conf map[string]interface{}, err error) {
	if !flag.Parsed() {
		flag.Parse()
	}

	xlog.Info("Use the config file of ", *confName)
	return LoadEx(*confName)
}

func LoadEx(confName string) (conf map[string]interface{}, err error) {
	data, err := ioutil.ReadFile(confName)
	if err != nil {
		xlog.Error("Load conf failed:", err)
		return
	}
	data = trimComments(data)

	err = yaml.Unmarshal(data, &conf)

	if err != nil {
		xlog.Error("Parse conf failed:", err)
	}
	return
}

func trimComments(data []byte) (data1 []byte) {

	conflines := bytes.Split(data, NL)
	for k, line := range conflines {
		conflines[k] = trimCommentsLine(line)
	}
	return bytes.Join(conflines, NL)
}

func trimCommentsLine(line []byte) []byte {

	var newLine []byte
	var i, quoteCount int
	lastIdx := len(line) - 1
	for i = 0; i <= lastIdx; i++ {
		if line[i] == '\\' {
			if i != lastIdx && (line[i+1] == '\\' || line[i+1] == '"') {
				newLine = append(newLine, line[i], line[i+1])
				i++
				continue
			}
		}
		if line[i] == '"' {
			quoteCount++
		}
		if line[i] == '#' {
			if quoteCount%2 == 0 {
				break
			}
		}
		newLine = append(newLine, line[i])
	}
	return newLine
}
