package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

const (
	// SortOrder used to request stocks in descending order.
	// use 'asc' for ascending order.
	SortOrder = "desc"
	// SortBy is a field to use in sorting the result.
	// symbol, name or list_order can be used.
	SortBy = "symbol"
	// Output is the expected result format.
	// It can be in 'json' or 'csv'
	Output = "json"
)

// Property holds information like URL, database connection
// and others that should not be hardcoded.
var Property Properties

// Properties holds the connection properties for t
type Properties struct {
	URL                     string `json:"url"`
	URLWithSorting          string `json:"urlWithSorting"`
	URLWithSortingAndOutput string `json:"urlWithSortingAndOutput"`
	APIToken                string `json:"apiToken"`
}

// Initializes database connection.
func init() {
	var err error

	Property, err = loadConnectionProperties()

	if err != nil {
		panic(err)
		fmt.Println("unable to load property details: " + err.Error())
	}

	if len(Property.URL) == 0 {
		panic(errors.New("invalid property values"))
		fmt.Println("invalid property values")
	}

	fmt.Println("property details successfully loaded")
}

// Parses the config file and load to Config struct.
func loadConnectionProperties() (Properties, error) {

	var property Properties
	data, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return Properties{}, err
	}

	err = json.Unmarshal(data, &property)
	if err != nil {
		return Properties{}, err
	}

	return property, nil

}

// GetTradingURL returns the complete trading url acceptable on the trading site.
func (config Properties) GetTradingURL(symbol, sortOrder, sortBy, output string) string {
	switch {
	case len(sortOrder) > 0:
		return fmt.Sprintf(config.URLWithSorting, symbol, config.APIToken, SortOrder)
	case len(sortOrder) > 0 && len(sortBy) > 0:
		return fmt.Sprintf(config.URLWithSorting, symbol, config.APIToken, sortBy)
	case len(sortOrder) > 0 && len(sortBy) > 0 && len(output) > 0:
		return fmt.Sprintf(config.URLWithSortingAndOutput, symbol, config.APIToken, sortBy, output)
	default:
		return fmt.Sprintf(config.URL, symbol)
	}
}
