package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rburawes/stocks/config"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Stock holds information about specific stock e.g. Apple, Facebook etc.
type Stock struct {
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Price          string `json:"price"`
	CloseYesterday string `json:"close_yesterday"`
	Currency       string `json:"currency"`
	MarketCap      string `json:"market_cap"`
	Volume         string `json:"volume"`
	Timezone       string `json:"timezone"`
	TimezoneName   string `json:"timezone_name"`
	GmtOffset      string `json:"gmt_offset"`
	LastTradeTime  string `json:"last_trade_time"`
}

// StockData holds data for the stock price in a stock exchange.
type StockData struct {
	StockExchange string  `json:"stock_exchange"`
	Stocks        []Stock `json:"stocks"`
}

// TradingData holds information about the stock found wrapped in
// the response body.
type TradingData struct {
	Symbol             string `json:"symbol"`
	Name               string `json:"name"`
	Currency           string `json:"currency`
	Price              string `json:"price"`
	PriceOpen          string `json:"price_open"`
	DayHigh            string `json:"day_high"`
	DayLow             string `json:"day_low"`
	FiftyTwoWeekHigh   string `json:"52_week_high"`
	FiftyTwoWeekLow    string `json:"52_week_low"`
	DayChange          string `json:"day_change"`
	ChangePct          string `json:"change_pct"`
	CloseYesterday     string `json:"close_yesterday"`
	MarketCap          string `json:"market_cap"`
	Volume             string `json:"volume"`
	VolumeAvg          string `json:"volume_avg"`
	Shares             string `json:"shares"`
	StockExchangeLong  string `json:"stock_exchange_long"`
	StockExchangeShort string `json:"stock_exchange_short"`
	Timezone           string `json:"timezone"`
	TimezoneName       string `json:"timezone_name"`
	GmtOffset          string `json:"gmt_offset"`
	LastTradeTime      string `json:"last_trade_time"`
}

// TradingResponse is the response body receive from
// trading website.
type TradingResponse struct {
	SymbolsRequested int32         `json:"symbols_requested"`
	SymbolsReturned  int32         `json:"symbols_returned"`
	Data             []TradingData `json:"data"`
}

// GetData reads json data from the 'www.worldtrading.com'.
func GetData(symbol string, stockExchange []string) ([]StockData, error) {

	url := config.Property.GetTradingURL(strings.ToUpper(symbol), config.SortOrder, config.SortBy, config.Output)

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)

	if err != nil {
		return []StockData{}, errors.New("failed to access trading API")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []StockData{}, fmt.Errorf("unable to retrieve data from the given url %d", resp.StatusCode)
	}

	data, err := readBody(resp.Body)

	if err != nil {
		return []StockData{}, fmt.Errorf("unable to process stock(s) data %s", err)
	}

	stocksData := make([]StockData, 0)

	for _, se := range stockExchange {
		upperSE := strings.ToUpper(se)
		sd := StockData{
			StockExchange: upperSE,
			Stocks:        make([]Stock, 0),
		}
		for _, d := range data.Data {
			if upperSE == strings.ToUpper(d.StockExchangeShort) {
				stock := Stock{
					Symbol:         d.Symbol,
					Name:           d.Name,
					Price:          d.Price,
					CloseYesterday: d.CloseYesterday,
					Currency:       d.Currency,
					MarketCap:      d.MarketCap,
					Volume:         d.Volume,
					Timezone:       d.Timezone,
					TimezoneName:   d.TimezoneName,
					GmtOffset:      d.GmtOffset,
					LastTradeTime:  d.LastTradeTime,
				}
				sd.Stocks = append(sd.Stocks, stock)
			}
		}
		if len(sd.Stocks) > 0 {
			stocksData = append(stocksData, sd)
		}
	}
	return stocksData, nil
}

// Reads the content of the http response body to get the json data.
func readBody(reader io.Reader) (TradingResponse, error) {

	var tradingResponse TradingResponse

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return TradingResponse{}, fmt.Errorf("something went wrong while reading trading data: %v", err)
	}

	err = json.Unmarshal(data, &tradingResponse)
	if err != nil {
		return TradingResponse{}, fmt.Errorf("something went wrong while processing trading data: %v", err)
	}

	return tradingResponse, nil
}
