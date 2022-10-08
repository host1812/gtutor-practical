package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Issue struct {
	Url           string `json:"url"`
	RepositoryUrl string `json:"repository_url"`
	Id            int    `json:"id"`
	User          *User  `json:"user"`
}

type User struct {
	Login string `json:"login"`
}

func main() {
	url := "https://api.github.com/repositories/19438/issues"
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Microsecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Panicln(err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Panicln(err)
		}
	}()

	issues := []Issue{}

	json.Unmarshal(bytes, &issues)

	for _, issue := range issues {
		log.Printf("issue id: %d (opened by %s)", issue.Id, issue.User.Login)
	}

	// log.Println("res:", string(bytes))
}
