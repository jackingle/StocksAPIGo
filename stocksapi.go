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

	// MakeBookRequest("F")
	//MakeBookRequest("T")
	//JSONBookRequest("T")
	http.HandleFunc("/stocks/", book_api)
	http.HandleFunc("/json/", jsonbook_api)
	http.ListenAndServe(":8000", nil)
}

func book_api(w http.ResponseWriter, r *http.Request) {
	s := r.URL.RequestURI()
	s = strings.SplitAfter(s, "/")[2]

	MakeBookRequest(s, w, r)

}
func jsonbook_api(w http.ResponseWriter, r *http.Request) {
	s := r.URL.RequestURI()
	s = strings.SplitAfter(s, "/")[2]

	JSONBookRequest(s, w, r)

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

func MakeBookRequest(symbol string, w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://cloud.iexapis.com/v1/stock/" + symbol + "/book?token=" + myToken)
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	json.Unmarshal([]byte(body), &message)
	b := result["quote"].(map[string]interface{})
	//log.Println(message())
	log.Println("The Company Name is: " + b["companyName"].(string))
	log.Println("The ticker symbol is: " + symbol)
	log.Println("The Primary Exchange this company belongs to is: " + b["primaryExchange"].(string))
	log.Printf("The latest price is: %v\n", b["latestPrice"].(float64))
	log.Printf("The previous closing price was: %v\n", b["previousClose"].(float64))
	log.Println("")
	//
	fmt.Fprint(w, "The Company Name is: "+b["companyName"].(string)+"\n")
	fmt.Fprint(w, "The ticker symbol is: "+symbol+"\n")
	fmt.Fprint(w, "The Primary Exchange this company belongs to is: "+b["primaryExchange"].(string)+"\n")
	fmt.Fprintf(w, "The latest price is: %v\n", b["latestPrice"].(float64))
	fmt.Fprintf(w, "The previous closing price was: %v\n", b["previousClose"].(float64))
	fmt.Fprintln(w, "")
}
func JSONBookRequest(symbol string, w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://cloud.iexapis.com/v1/stock/" + symbol + "/book?token=" + myToken)
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(body), &message)

	b, err := json.Marshal(message)
	log.Println(string(b))
	fmt.Fprint(w, string(b))
}

func readTokenFile() {
	dat, err := ioutil.ReadFile("dat")
	if err != nil {
		log.Fatalln(err)
	}
	myToken = string(dat)

}

type Message struct {
	Quote       Quote
	Bids        Bids
	Asks        Asks
	Trades      Trades
	SystemEvent SystemEvent
}
type Quote struct {
	Symbol                string
	CompanyName           string
	PrimaryExchange       string
	CalculationPrice      string
	Open                  string
	OpenTime              string
	Close                 string
	CloseTime             string
	High                  string
	Low                   string
	LatestPrice           float64
	LatestSource          string
	LatestTime            string
	LatestUpdate          int64
	LatestVolume          string
	IEXRealtimePrice      float64
	IEXRealtimeSize       int64
	IEXLastUpdated        int64
	DelayedPrice          string
	DelayedPriceTime      string
	ExtendedPrice         string
	ExtendedChange        string
	ExtendedChangePercent string
	ExtendedPriceTime     string
	PreviousClose         int64
	PreviousVolume        int64
	Change                float64
	ChangePercent         float64
	Volume                string
	IEXMarketPercent      float64
	IEXVolume             int64
	AVGTotalVolume        int64
	IEXBidPrice           float64
	IEXBidSize            int64
	IEXAskPrice           float64
	IEXAskSize            int64
	MarketCap             int64
	PeRatio               float64
	Week52High            float64
	Week52Low             float64
	YTDChange             float64
	LastTradeTime         int64
	IsUSMarketOpen        bool
}
type Bids struct {
	Price     float64
	Size      int64
	Timestamp string
}
type Asks struct {
	Price     float64
	Size      int64
	Timestamp string
}
type Trades struct {
	Price                 float64
	Size                  int64
	TradeID               int64
	IsISO                 bool
	IsOddLot              bool
	IsOutsideRegularHours bool
	IsSinglePriceCross    bool
	IsTradeThroughExempt  bool
	Timestamp             string
}
type SystemEvent struct {
	SystemEvent string
	Timestamp   string
}
