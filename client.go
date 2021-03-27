package influxgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
)

type IfxCli struct {
	Ip         string
	Port       int
	User       string
	PassWord   string
	//db         string
	baseReqUrl string
	upParam    string
}
type IfxResp struct {
	Name    string     `json:"name"`
	Columns []string   `json:"columns"`
	Values  [][]string `json:"values"`
}

//create a new influxdb client; if the user and password was not be set
//in your influxdb server, set it as ""
func NewClient(ip string, port int, user, passWord string) *IfxCli {
	var userPassWord string
	if user != "" {
		userPassWord = fmt.Sprintf("u=%v&p=%v", user, passWord)
	}
	return &IfxCli{
		Ip:         ip,
		Port:       port,
		User:       user,
		PassWord:   passWord,
		baseReqUrl: fmt.Sprintf("http://%v:%v", ip, port),
		upParam:    userPassWord,
	}
}

func umsResp(data []byte) (*IfxResp, error) {
	resultsArr := gjson.GetBytes(data, "results").Array()
	if len(resultsArr) <= 0 {
		return nil, errors.New("results arr is nil")
	}
	seriesArr := resultsArr[0].Get("series").Array()
	if len(seriesArr) <= 0 {
		return nil, errors.New("results.series arr is nil")
	}
	var r IfxResp
	//log.Printf("se0:%v\n", seriesArr[0].String())
	err := json.Unmarshal([]byte(seriesArr[0].String()), &r)
	if err != nil {
		return nil, err
	} else {
		return &r, nil
	}
}
