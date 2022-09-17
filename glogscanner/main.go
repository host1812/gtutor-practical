package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
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
	return r.ReplaceAllString(s, "")
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	words := []string{}
	wordsStorage := map[string]int{}
	// r, err := regexp.Compile(`[^a-z]+`)
	r, err := regexp.Compile(`\b\w+\b`)
	if err != nil {
		fmt.Println("failed to compile regex")
	}
	for in.Scan() {
		// w := parse(in.Text())
		w := parseRx(in.Text(), r)
		if w != "" {
			if _, ok := wordsStorage[w]; !ok {
				words = append(words, w)
			}
			wordsStorage[w] += 1
		}
	}
	sort.Strings(words)
	for _, w := range words {
		fmt.Printf("%s: %d\n", w, wordsStorage[w])
	}
	// fmt.Println(words)
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
