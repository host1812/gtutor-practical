package main

import (
	"fmt"
	"os"
	"strconv"
)

func usage() {
	fmt.Println("usage:")
	fmt.Println("./authn <user> <password>")
}

func main() {
	uDb := map[string]string{
		"alex": "pass",
		"bob":  "1234",
	}
	// fmt.Printf("%v, %T", uDb, uDb)

	if len(os.Args) < 3 {
		fmt.Println("error: no username and/or password provided")
		usage()
		return
	}

	user := os.Args[1]
	pass := os.Args[2]

	if p, ok := uDb[user]; ok {
		if pass == p {
			fmt.Println("access granted")
		} else {
			fmt.Println("access denied")
		}
	} else {
		fmt.Println("access denied")
	}

	n, err := strconv.Atoi("test")
	if err != nil {
		fmt.Printf("type: %T, value: %v", err, err)
	}
	fmt.Println(n)

}
