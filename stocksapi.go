package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	readTokenFile()
	MakeBookRequest("GE")
}

func MakeRequest() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func MakeBookRequest(symbol string) {
	resp, err := http.Get("https://cloud.iexapis.com/v1/stock/" + symbol + "/book?token=" + myToken)
	if err != nil {
		log.Fatalln(err)
	}

	//body, err := ioutil.ReadAll(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	log.Println(result["quote"])
	if err != nil {
		log.Fatalln(err)
	}

	//log.Println(string(body))
}

var myToken string

func readTokenFile() {
	dat, err := ioutil.ReadFile("dat")
	if err != nil {
		log.Fatalln(err)
	}
	myToken = string(dat)

}

type Message struct {
	symbol          string
	companyName     string
	primaryExchange string
	latestPrice     float64
	previousClose   float64
	change          float64
}

var M Message

// func JSONtoMap(input string) {

// }
