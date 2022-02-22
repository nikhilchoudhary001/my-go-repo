package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Result struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var url string = "https://jsonplaceholder.typicode.com/todos/%s"

func main() {
	//https://jsonplaceholder.typicode.com/todos/1

	res := make(chan *Result, 3)
	wg := &sync.WaitGroup{}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go Consume(i, res, wg)
	}

	wg.Wait()
	close(res)
	for response := range res {
		fmt.Printf("Response : %v", response)
	}

}

func Consume(id int, res chan *Result, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(url, strconv.Itoa(id)), nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	var responseObject *Result
	json.NewDecoder(resp.Body).Decode(&responseObject)
	// Can also be read using ioutil instead of NewDecoder

	// bodyBytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }

	// json.Unmarshal(bodyBytes, &responseObject)

	//v.Decode(&responseObject)

	res <- responseObject
}
