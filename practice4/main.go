package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var ips = []string{
	"127.0.0.1",
	"127.0.0.2",
	"127.0.0.3",
	"127.0.0.4",
	"127.0.0.5",
	"192.168.0.1",
	"192.168.0.2",
	"192.168.0.3",
	"192.168.0.4",
	"192.168.0.5",
	"10.10.0.1",
	"10.10.0.2",
	"10.10.0.3",
	"10.10.0.4",
	"10.10.0.5",
}

var resources = []Resource{
	{"/", 2326},
	{"/admin", 4040},
	{"/sales", 5050},
	{"/sales/product/1", 12954},
	{"/sales/product/12", 9754},
	{"/sales/product/102", 16545},
	{"/sales/product/1001", 10556},
}

type Resource struct {
	URI       string
	SizeBytes int
}

var clientErrRate, serverErrRate = 1, 2

func generateLine() string {
	rand.Seed(time.Now().UnixNano())
	responseCode := 200
	ip := ips[rand.Intn(len(ips))]
	resource := resources[rand.Intn(len(resources))]

	if rand.Intn(100) <= clientErrRate {
		responseCode = 404
	}
	if rand.Intn(100) <= serverErrRate {
		responseCode = 503
	}
	time := time.Now().Format("10/Dec/2006:15:04:05 -0700")
	return fmt.Sprintf("%s - Scott [%s] \"GET %s HTTP/1.1\" %d 2326 %d\n", ip, time, resource.URI, responseCode, resource.SizeBytes)
}

func generateLog(fileName string, lines int) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("failed to open file to write to, err:", err)
	}
	defer f.Close()

	// 127.0.0.1 - Scott [10/Dec/2019:13:55:36 -0700] "GET /server-status HTTP/1.1" 200 2326
	for i := 0; i <= lines; i++ {
		f.WriteString(generateLine())
	}
	return nil
}

func main() {
	_ = generateLog("./access.log", 10000)
}
