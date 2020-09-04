package proto

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CommonRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func DefaultRet(c *gin.Context, data interface{}, err error) {
	var res CommonRes
	if err != nil {
		res.Code = DefaultErrorCode
		res.Msg = err.Error()
	} else {
		res.Code = 200
		res.Msg = "success"
	}
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func ParseMapFromStruct(stru interface{}, m *map[string]interface{}) (err error) {
	temp, err := json.Marshal(stru)
	if err != nil {
		return err
	}
	json.Unmarshal(temp, m)
	return nil
}

//用来解决时间序列化和反序列化问题
const TimeFormat = "2006-01-02 15:04:05"

type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) Value() ([]byte, error) {
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalTime(tTime)
	return nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}
