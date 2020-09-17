package util

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	log "qiniupkg.com/x/log.v7"
)

func Post(url string, data []byte, timeout ...time.Duration) (body []byte, err error) {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	client := http.DefaultClient
	if len(timeout) > 0 {
		client = &http.Client{
			Timeout: timeout[0],
		}
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if response.StatusCode >= 300 || response.StatusCode < 200 {
		log.Println("response status code error :", response.StatusCode, "; body:", string(body))
		return nil, errors.New(string(body))
	}
	// log.Println(string(body), response.StatusCode)
	return body, nil
}

func Get(url string, timeout ...time.Duration) (body []byte, err error) {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	client := http.DefaultClient
	if len(timeout) > 0 {
		client = &http.Client{
			Timeout: timeout[0],
		}
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if response.StatusCode >= 300 || response.StatusCode < 200 {
		log.Println("response status code error :", response.StatusCode, "; body:", string(body))
		return nil, errors.New(string(body))
	}
	// log.Println(string(body), response.StatusCode)
	return body, nil
}

func PostRaw(url string, data []byte, header map[string]string, timeout ...time.Duration) (body []byte, err error) {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	client := http.DefaultClient
	if len(timeout) > 0 {
		client = &http.Client{
			Timeout: timeout[0],
		}
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if response.StatusCode >= 300 || response.StatusCode < 200 {
		log.Println("response status code error :", response.StatusCode, "; body:", string(body))
		return nil, errors.New(string(body))
	}
	// log.Println(string(body), response.StatusCode)
	return body, nil
}

func BetterResponseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var result struct {
			Code    int         `json:"code"`
			Message interface{} `json:"message,omitempty"`
		}
		rawWriter := c.Response().Writer
		recorder := httptest.NewRecorder()
		c.Response().Writer = recorder
		err = next(c)
		c.Response().Writer = rawWriter
		if err == nil {
			for k, vs := range recorder.HeaderMap {
				for _, v := range vs {
					rawWriter.Header().Add(k, v)
				}
			}
			if !c.Response().Committed {
				return c.JSON(http.StatusOK, result)
			}
			if !strings.Contains(rawWriter.Header().Get(echo.HeaderContentType), echo.MIMEApplicationJSON) {
				_, err = io.Copy(rawWriter, recorder.Body)
				return
			}
			rawWriter.Write([]byte(`{"code":0,"data":`))
			io.Copy(rawWriter, recorder.Body)
			_, err = io.Copy(rawWriter, recorder.Body)
			rawWriter.Write([]byte(`}`))
			return
		}
		if herr, ok := err.(*echo.HTTPError); ok {
			result.Code = herr.Code
			result.Message = herr.Message
		} else {
			result.Code = http.StatusInternalServerError
			result.Message = err.Error()
		}
		return c.JSON(http.StatusOK, result)
	}
}
