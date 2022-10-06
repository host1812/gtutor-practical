package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ResultPerIp struct {
	Total   int
	Good    int
	Bad     int
	UserErr int
	URLs    map[string]int
}

type Result struct {
	TotalLines              int
	MeanLatencyMilliseconds float64
	IpRequests              map[string]*ResultPerIp
}

func GetLogStats(fileName string) Result {
	// open and read file line-by-line to optimize memory
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	result := Result{
		TotalLines:              0,
		MeanLatencyMilliseconds: 0.0,
		IpRequests:              map[string]*ResultPerIp{},
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var good, bad, userErr int
		line := sc.Text()
		// log.Println("line: ", line)
		fields := strings.Fields(line)
		responseCode, err := strconv.Atoi(fields[8])
		if err != nil {
			log.Panic("failed to get response code, err:", err)
		}

		if responseCode >= 200 && responseCode < 400 {
			good = 1
		} else if responseCode >= 400 && responseCode < 500 {
			userErr = 1
		} else {
			bad = 1
		}

		if v, ok := result.IpRequests[fields[0]]; ok {
			v.Total++
			v.Good += good
			v.Bad += bad
			v.UserErr += userErr
		} else {
			result.IpRequests[fields[0]] = &ResultPerIp{
				Total:   1,
				Good:    good,
				Bad:     bad,
				UserErr: userErr,
			}
		}
		result.TotalLines++
	}
	return result
}

func main() {
	result := GetLogStats("./access.log")
	fmt.Printf("Total Lines: %d\n", result.TotalLines)
	for k, v := range result.IpRequests {
		fmt.Printf("\tTotal Requests for '%s': %d\n", k, v)
	}

}
