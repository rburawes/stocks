package main

import (
	"fmt"
	"github.com/rburawes/stocks/controllers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Tests the application if it will be accessible.
// Tells the application is running.
func TestAppHealth(t *testing.T) {
	executeTest("/", controllers.Index, t)
}

// Tests the Stocks API with multiple symbol and stock exchange.
func TestMultipleSymbolsAndStockExchange(t *testing.T) {
	executeTest("/api/v1/stocks/NOG,AAPL?stock_exchange=AMEX,NASDAQ", controllers.Stocks, t)
}

// Tests the Stocks API with single symbol and a stock exchange.
func TestWithSingleSymbolAndStockExchange(t *testing.T) {
	executeTest("/api/v1/stocks/AAPL?stock_exchange=NASDAQ", controllers.Stocks, t)
}

// Executes the test based on the given conditions.
func executeTest(url string, f func(w http.ResponseWriter, req *http.Request), t *testing.T) {
	// Create a request to pass to a handler that manages stocks.
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Records the response.
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(f)

	handler.ServeHTTP(rec, req)

	// The status code must be '200'
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("request returns unexpected result with status code: RECEIVED: %v EXPECTED %v",
			status, http.StatusOK)
	}

	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Status Code: ", resp.StatusCode)
	fmt.Println("Content Type: ", resp.Header.Get("Content-Type"))
	fmt.Println("Response Body: ", string(body))

}
