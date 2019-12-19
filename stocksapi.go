package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var message Message
var myToken string

func main() {
	readTokenFile()

	// MakeBookRequest("F")
	//MakeBookRequest("T")

	http.HandleFunc("/stocks/", book_api)
	http.ListenAndServe(":8000", nil)
}

func book_api(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	key := keys[0]
	// u, err := url.Parse("http://localhost:8000/stocks/?key=")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// s := strings.Split(u.RequestURI(), "/")[1]
	log.Println(r.URL.RequestURI())
	MakeBookRequest(string(key))
	log.Println("Url Param 'key' is: " + string(key))

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
	json.Unmarshal([]byte(body), &message)
	b := result["quote"].(map[string]interface{})
	//log.Println(message())
	log.Println("The Company Name is: " + b["companyName"].(string))
	log.Println("The ticker symbol is: " + symbol)
	log.Println("The Primary Exchange this company belongs to is: " + b["primaryExchange"].(string))
	log.Printf("The latest price is: %v\n", b["latestPrice"].(float64))
	log.Printf("The previous closing price was: %v\n", b["previousClose"].(float64))
	log.Println("")
}
func JSONBookRequest(symbol string) {
	resp, err := http.Get("https://cloud.iexapis.com/v1/stock/" + symbol + "/book?token=" + myToken)
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(body), &message)

	//log.Println(message())

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
