package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name    string  `csv:name`
	Age     int     `csv:age`
	Savings float64 `csv:savings`
}

func (p *Person) Introduce() string {
	return fmt.Sprintf("My name is %s. I'm %d years old. I have %f on my bank account.", p.Name, p.Age, p.Savings)
}

func NewPerson(record []string) (Person, error) {
	if strings.Join(record, ",") == "name,age,savings" {
		return Person{}, errors.New("this is header, skipping")
	}
	age, err := strconv.Atoi(record[1])
	if err != nil {
		return Person{}, err
	}

	savings, err := strconv.ParseFloat(record[2], 64)
	if err != nil {
		return Person{}, err
	}

	return Person{
		record[0],
		age,
		savings,
	}, nil
}

func main() {
	f, _ := os.Open("./test.csv")
	defer f.Close()
	// r := bufio.NewScanner(f)
	// p := []byte{}
	// for r.Scan() {
	// 	csv.
	// }
	persons := []Person{}

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		person, err := NewPerson(record)
		if err != nil {
			log.Println("error parsing record: ", err)
			continue
		}
		persons = append(persons, person)
	}

	fmt.Println(persons)

}
