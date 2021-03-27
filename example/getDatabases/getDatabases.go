package main

import (
	"influxgo"
	"log"
)

func main() {
	c := influxgo.NewClient("192.168.0.1", 8086, "", "")
	databases, err := c.GetDatabases()
	if err != nil {
		panic(err)
	} else {
		log.Printf("db:%v\n", databases)
	}
}
