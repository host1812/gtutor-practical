package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

const depth = 2
const maxObjectsDir = 10
const wordsFile = "/usr/share/dict/words"
const rootDir = "./generated"
const dirFactor = 20

func readWords() ([]string, error) {
	var words []string
	f, err := os.Open(wordsFile)
	if err != nil {
		return []string{}, err
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		words = append(words, sc.Text())
	}
	return words, nil
}

func getWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func createDir() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) < 20
}

func getRandomContent(words []string) string {
	sb := strings.Builder{}

	for i := 0; i < rand.Intn(1000); i++ {
		sb.WriteString(getWord(words))
		sb.WriteString(" ")
	}
	sb.WriteString("\n")
	return sb.String()
}

func createFile(dir string, words []string) error {
	file := path.Join(dir, getWord(words))
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	log.Printf("file created: %s\n", file)

	wr := bufio.NewWriter(f)
	wr.WriteString(getRandomContent(words))
	wr.Flush()
	log.Printf("file content updated: %s\n", file)
	return nil
}

func generateFiles(words []string) error {
	_, err := os.Stat(rootDir)
	if os.IsNotExist(err) {
		log.Printf("root dir '%s' doesn't exists, will create it\n", rootDir)
		err := os.Mkdir(rootDir, 0755)
		if err != nil {
			return nil
		}
		log.Printf("'%s' successfully created\n", rootDir)
	}
	for i := 0; i < 12; i++ {
		if createDir() {
			dir := path.Join(rootDir, getWord(words))
			err = os.Mkdir(dir, 0755)

			if err != nil {
				return err
			}
			log.Printf("dir created: %s\n", dir)
		} else {
			err := createFile(rootDir, words)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("started")
	words, err := readWords()
	if err != nil {
		log.Fatalln(err)
	}
	generateFiles(words)
	log.Println("finished")
}
