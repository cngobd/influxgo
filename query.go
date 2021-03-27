package influxgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type QueryResp struct {
	Columns []string
	Values [][]interface{}
}
func (f *IfxCli) QueryMeasure(start, end time.Time) {

}

//query the latest time data by set the tag,
//n: the number of time duration
//timeType support:
//"s": second, "m": minute, "h": hour, "d": day, "w": week
func (f *IfxCli) QueryLastNTimeByTag(db, measure, tagName, tagValue string, n int, timeType string) (QueryResp, error) {
	t, err := parseTimeType(n, timeType)
	if err != nil {
		return QueryResp{}, err
	}
	q := fmt.Sprintf("select%%20*%%20from%%20\"%v\"%%20where%%20time%%20>%%20%v%%20and%%20\"%v\"%%20=%%20'%v'",
		measure, t, tagName, tagValue)
	url := fmt.Sprintf("%v/query?db=%v&u=%v&p=%v&q=%v",
		f.baseReqUrl, db, f.User, f.PassWord, q)
	log.Printf("url:%v", url)
	body, err := sendQueryRequest(url)
	if err != nil {
		return QueryResp{}, err
	}
	return umsQueryResp(body)
}

//query the latest time data;
//n: the number of time duration;
//timeType support:
//"s": second, "m": minute, "h": hour, "d": day, "w": week
func (f *IfxCli) QueryLastNTime(db, measure string, n int, tp string) (QueryResp, error) {
	t, err := parseTimeType(n, tp)
	if err != nil {
		return QueryResp{}, err
	}
	q := fmt.Sprintf("select%%20*%%20from%%20\"%v\"%%20where%%20time%%20>%%20%v", measure, t)
	url := fmt.Sprintf("%v/query?db=%v&u=%v&p=%v&q=%v",
		f.baseReqUrl, db, f.User, f.PassWord, q)
	log.Printf("url:%v", url)
	body, err := sendQueryRequest(url)
	if err != nil {
		return QueryResp{}, err
	}
	return umsQueryResp(body)

}

func parseTimeType(n int, tp string) (string, error){
	switch tp {
	case "w":
		return fmt.Sprintf("%%20now()%%20-%%20%dw", n), nil
	case "d":
		return fmt.Sprintf("%%20now()%%20-%%20%dd", n), nil
	case "h":
		return fmt.Sprintf("%%20now()%%20-%%20%dh", n), nil
	case "m":
		return fmt.Sprintf("%%20now()%%20-%%20%dm", n), nil
	case "s":
		return fmt.Sprintf("%%20now()%%20-%%20%ds", n), nil
	default:
		return "", errors.New("unknown time type")

	}
}
func sendQueryRequest(url string) ([]byte, error){
	get, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}

type query struct {
	Results []result `json:"results"`
}

type result struct {
	Series []seri
}
type seri struct {
	MeasureName string `json:"name"`
	Columns []string
	Values [][]interface{}
}

func umsQueryResp(data []byte) (QueryResp, error) {
	var u query
	err := json.Unmarshal(data, &u)
	if err != nil {
		return QueryResp{}, err
	}
	if len(u.Results) <= 0 ||
		len(u.Results[0].Series) <= 0 ||
		len(u.Results[0].Series[0].Columns) <= 0 ||
		len(u.Results[0].Series[0].Values) <= 0{
		return QueryResp{}, err
	} else {
		return QueryResp{
			Columns: u.Results[0].Series[0].Columns,
		Values: u.Results[0].Series[0].Values}, err
	}
}