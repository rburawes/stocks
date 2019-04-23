package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/rburawes/stocks/models"
	"net/http"
	"strings"
)

const (
	pathStocks             = "/api/v1/stocks/"
	headerContentTypeKey   = "Content-Type"
	jsonType               = "application/json"
	jsonIndexPrefix        = ""
	jsonIndentValue        = "\t"
	errorMsg               = "Unable to retrieve data: "
	stockExchangeParamName = "stock_exchange"
	defaultStockExchange   = "AMEX"
)

// Index is the default page of the application.
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "I am running...\n")
}

// Stocks return prices for the given stock(s)
func Stocks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	stockExchangeValue := make([]string, 0)
	symbol := r.URL.Path[len(pathStocks):]

	queryParam := r.URL.Query()

	var stockExchangeParamValue string

	if len(queryParam) > 0 {
		stockExchangeParamValue = queryParam.Get(stockExchangeParamName)
	}

	if len(stockExchangeParamValue) > 0 {
		stockExchangeValue = append(stockExchangeValue, strings.Split(stockExchangeParamValue, ",")...)
	} else {
		stockExchangeValue = append(stockExchangeValue, defaultStockExchange)
	}

	if len(symbol) == 0 {
		http.Error(w, errorMsg+"invalid parameter", http.StatusBadRequest)
		fmt.Println(errorMsg + "invalid parameter")
		return
	}

	result, err := models.GetData(symbol, stockExchangeValue)

	if err != nil {
		http.Error(w, errorMsg+err.Error(), http.StatusInternalServerError)
		fmt.Println(errorMsg + err.Error())
		return
	}

	if len(result) == 0 {
		http.Error(w, errorMsg+"no result", http.StatusBadRequest)
		fmt.Println(errorMsg + "no result")
		return
	}

	ConvertToJSON(w, result)

}

// ConvertToJSON converts the target struct to json object
func ConvertToJSON(w http.ResponseWriter, s interface{}) {

	uj, err := json.MarshalIndent(s, jsonIndexPrefix, jsonIndentValue)
	if err != nil {
		http.Error(w, errorMsg+err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set(headerContentTypeKey, jsonType)
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}
