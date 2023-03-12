package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"os"
	"strings"
)

const fileName = "users.csv"

func main() {

	var data = make(chan string)
	var completed = make(chan int)
	for i := 0; i < 8; i++ {
		go generateData(data, completed)
	}
	go writeData(data)
	counter := 0
	for elem := range completed {
		fmt.Println(elem)
		counter += 1
		if counter == 8 {
			close(data)
		}
	}
}

func writeData(data chan string) {
	file, err2 := os.Create(fileName)
	if err2 != nil {
		panic("error creatung csv")
	}
	for row := range data {
		_, err := file.WriteString(row)
		if err != nil {
			panic(err)
		}
	}
	err := file.Close()
	if err != nil {
		panic(err)
	}
	os.Exit(0)

}

func generateData(data chan string, completed chan int) {
	for i := 0; i < 5_000; i++ {
		buffer := strings.Builder{}
		for i := 0; i < 1_000; i++ {
			address := faker.GetRealAddress()
			buffer.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s\n", faker.Name(), faker.Date(), faker.Email(), address.State, address.City))
		}
		data <- buffer.String()
	}
	completed <- 1
}
