package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var message Message
var myToken string

func main() {
	readTokenFile()
	http.HandleFunc("/stocks/", bookAPI)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func bookAPI(w http.ResponseWriter, r *http.Request) {
	s := r.URL.RequestURI()
	s = strings.SplitAfter(s, "/")[2]
	jsonBookRequest(s, w, r)

}

func jsonBookRequest(symbol string, w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://cloud.iexapis.com/v1/stock/" + symbol + "/book?token=" + myToken)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &message)
	b, err := json.Marshal(message)
	defer resp.Body.Close()
	log.Println(string(b))
	fmt.Fprint(w, string(b))
	w.Write(b)
}

func readTokenFile() {
	dat, err := ioutil.ReadFile("dat")
	if err != nil {
		log.Fatalln(err)
	}
	myToken = string(dat)
}

type Message struct {
	Quote Quote
}
type Quote struct {
	LatestPrice float64
}
