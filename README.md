# StocksAPIGo

## Introduction
This program is designed to serve the IEX Cloud API and provide up to date stock information 
on a variety of stocks.  As it currently exists, the latest price is returned for a given stock
at 15 second intervals.  The output is printed to terminal in JSON format.
## Setup
In order to use this program on your own, you will need to sign up with the IEX Group [here](https://iextrading.com/developers/).
After cloning the repository from Github, create an empty file called "dat" without quotes.
Log into your IEX Cloud console and copy your API token to the dat file.
Now the program can be run!
## Running the program
`go run stocksapi.go`

This will initialize the webserver at http://localhost:8000

Access the endpoint at http://localhost:8000/stocks/ and enter a stock ticker symbol at the end

of the URI in order to obtain the latest stock price.
