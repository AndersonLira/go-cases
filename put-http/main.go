package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	for i := 0; i < 200; i++ {
		go doPut(i)
	}
	time.Sleep(30 * time.Second)
}
func doPut(i int) {
	client := &http.Client{}
	url := fmt.Sprintf("http://localhost:9200/test/matematica/%d", i)
	json := `{"title":"valor","text":"Just trying this out...","date":"2014/01/01"}`
	request, err := http.NewRequest("PUT", url, strings.NewReader(json))
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		_, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
		fmt.Println("   ", response.StatusCode)
		// hdr := response.Header
		// for key, value := range hdr {
		// 	fmt.Println("   ", key, ":", value)
		// }
		// fmt.Println(contents)
	}
}
