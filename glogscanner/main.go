package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kylelemons/godebug/pretty"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	words := map[string]int{}
	for in.Scan() {
		w := in.Text()
		if strings.ContainsAny(w, ",.;-") {
			replacer := strings.NewReplacer(",", "", ".", "", ";", "")
			w = replacer.Replace(w)
		}
		words[w] += 1
	}
	fmt.Println(words)
	pretty.Print(words)
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
