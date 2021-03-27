package main

import (
	"influxgo"
)

func main() {
	//make a new client
	c := influxgo.NewClient("192.168.1.1", 8086, "", "")
	//create data struct
	data := influxgo.WriteData{}
	//init map of tags and values
	data.Tags = make(map[string]interface{})
	data.Fields = make(map[string]interface{})
	//set the request db
	data.DB = "test1"
	//set measurement
	data.Measurement = "m3"
	//set tags
	data.Tags["myTag1"] = "t1"
	data.Tags["myTag2"] = "t2"
	//set fields
	data.Fields["myField1"] = "f1"
	//write the data into db
	err := c.Write(data)
	if err != nil {
		panic(err)
	}

}
