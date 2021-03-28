package influxgo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type WriteData struct {
	DB          string
	Measurement string
	Tags        map[string]interface{}
	Fields      map[string]interface{}
}

func (f *IfxCli) Write(data WriteData) error {
	if err := checkReq(data); err != nil {
		return err
	}
	postData := data.Measurement
	for k, v := range data.Tags {
		postData = fmt.Sprintf("%v,%v=%v", postData, k, v)
	}
	postData = fmt.Sprintf("%v ", postData)
	for k, v := range data.Fields {
		postData = fmt.Sprintf("%v%v=%v,", postData, k, v)
	}
	postData = postData[:len(postData) - 1]
	url := fmt.Sprintf(
		"http://152.32.173.198:3306/write?db=%v&u=%v&p=%v",
		data.DB, f.User, f.PassWord,
	)
	log.Printf("write url:%v\n", url)
	log.Printf("post data:%v\n", postData)
	return sendWriteRequest(postData, url)
}

func sendWriteRequest(postData string, url string) error {
	c := http.Client{}
	r, err := http.NewRequest("POST", url, strings.NewReader(postData))
	if err != nil {
		return err
	}
	resp, err := c.Do(r)
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(all) > 0 {
		return errors.New(string(all))
	}
	return nil
}

func checkReq(data WriteData) error {
	if data.DB == "" {
		return errors.New("db is nil")
	}
	if data.Measurement == "" {
		return errors.New("measurement is nil")
	}
	var FieldCount int
	for range data.Fields {
		FieldCount++
	}
	if FieldCount <= 0 {
		return errors.New("field is nil")
	}
	return nil
}