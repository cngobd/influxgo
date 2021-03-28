package influxgo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//get all the name of databases
func (f *IfxCli) GetDatabases() ([]string, error) {
	u := f.baseReqUrl + "/query?" + f.upParam + "&q=show%20databases"
	q, err := queryWork(u)
	if err != nil {
		return nil, err
	}
	resp, err := umsResp(q)
	if err != nil {
		return nil, err
	}
	var arr []string
	for _, v := range resp.Values {
		if len(v) <= 0 {
			continue
		}
		arr = append(arr, v[0])

	}
	return arr, nil
}

func queryWork(url string) ([]byte, error){
	get, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	} else {
		return all, nil
	}
}
func (f *IfxCli) Ping() error {
	url := fmt.Sprintf("%v/ping", f.baseReqUrl)
	get, err := http.Get(url)
	if err != nil {
		return errors.New(fmt.Sprintf("ping fail:%v", err.Error()))
	}
	if get.StatusCode == 204 {
		return nil
	} else {
		return errors.New(fmt.Sprintf("ping fail, status code:%v", get.StatusCode))
	}
}