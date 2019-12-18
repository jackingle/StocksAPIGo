package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	readTokenFile()
	MakeBookRequest("TM")
	MakeBookRequest("F")
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

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}

	json.Unmarshal([]byte(body), &result)

	b := result["quote"].(map[string]interface{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("The Company Name is: " + b["companyName"].(string))
	log.Println("The ticker symbol is: " + symbol)
	log.Println("The Primary Exchange this company belongs to is: " + b["primaryExchange"].(string))
	log.Printf("The latest price is: %v\n", b["latestPrice"].(float64))
	log.Printf("The previous closing price was: %v\n", b["previousClose"].(float64))
	log.Println("")
}

var myToken string

func readTokenFile() {
	dat, err := ioutil.ReadFile("dat")
	if err != nil {
		log.Fatalln(err)
	}
	myToken = string(dat)

}
