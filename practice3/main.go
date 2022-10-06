package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type ResultPerIp struct {
	Total int
	OK    int
	Bad   int
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
		line := sc.Text()
		// log.Println("line: ", line)
		fields := strings.Fields(line)
		// result.IpRequests[fields[0]] = ResultPerIp{Total: 1}
		if v, ok := result.IpRequests[fields[0]]; ok {
			v.Total++
		} else {
			result.IpRequests[fields[0]] = &ResultPerIp{
				Total: 1,
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
