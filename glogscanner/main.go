package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parse(s string) string {
	if strings.ContainsAny(s, ",.;-:") {
		replacer := strings.NewReplacer(
			",", "",
			".", "",
			";", "",
			":", "",
		)
		s = replacer.Replace(s)
	}
	return s
}

func parseRx(s string, r *regexp.Regexp) string {
	matches := r.FindStringSubmatch(s)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

func parseRx2(s string, r *regexp.Regexp) string {
	// fmt.Println("debug:", r.FindStringSubmatch(s))
	return r.ReplaceAllString(s, "")
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	words := map[string]int{}
	r, err := regexp.Compile(`[^a-z]+`)
	if err != nil {
		fmt.Println("failed to compile regex")
	}
	for in.Scan() {
		// w := parse(in.Text())
		w := parseRx(in.Text(), r)
		if w != "" {
			words[w] += 1
		}

	}
	fmt.Println(words)
	// pretty.Print(words)
	// fmt.Println("scanned:", in.Text())
	// file, err := os.Open("./sample")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// // optionally, resize scanner's capacity for lines over 64K, see next example
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
