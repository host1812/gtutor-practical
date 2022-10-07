package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func main() {
	log.Println("started")

	f, err := os.Open("./access.log")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(f)
	r, _ := regexp.Compile(`^(?P<IP>\d*\.\d*\.\d*\.\d*\s)`)

	for sc.Scan() {
		log.Println(sc.Text())
		log.Printf("%#v\n", r.FindStringSubmatch(sc.Text()))
		log.Printf("%#v\n", r.SubexpNames())
		// r := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)
		// fmt.Printf("%#v\n", r.FindStringSubmatch(`2015-05-27`))
		// fmt.Printf("%#v\n", r.SubexpNames())
		break
	}

	log.Println("finished")
}
