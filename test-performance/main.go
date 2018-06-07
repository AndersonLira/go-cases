package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var ok int
var vu = 1000

func main() {
	for i := 0; i < vu; i++ {
		wg.Add(1)
		go request(i)
	}
	wg.Wait()
	fmt.Printf("Ok = %d\n", ok)
}

func request(i int) {
	defer wg.Done()
	fmt.Printf("Request: %d\n", i)
	URL := fmt.Sprintf("http://escoladigitaldes.vanzolini-gte.org.br/api/1/search?q=%d", i)
	response, err := http.Get(URL)
	if err != nil {
		fmt.Printf("Error %v\n", err)
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error parsin response")
		}
		s := search{}
		if response.StatusCode == http.StatusOK {
			json.Unmarshal(body, &s)
			fmt.Println(s)
			if s.Hits.Total > 0 {
				ok = ok + 1

			}
		}
	}
}

type search struct {
	Hits hits `json:"hits"`
}
type hits struct {
	Total int `json:"total"`
}
