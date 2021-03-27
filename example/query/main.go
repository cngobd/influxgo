package main

import (
	"influxgo"
	"log"
)

func main() {
	c := influxgo.NewClient("192.168.0.1", 8086, "", "")
	q, err := c.QueryLastNTime("test1", "m3", 30, "d")
	if err != nil {
		panic(err)
	}
	log.Printf("query result:%v\n", q)
	q, err = c.QueryLastNTimeByTag("test1", "m3", "tg1", "4", 5, "h")
	if err != nil {
		panic(err)
	}
	log.Printf("query result:%v\n", q)
}
