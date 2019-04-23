# Stocks

This application utilises the World Trading Data API to display the stocks prices based on the provided symbol and stock exchange.

1. Reads JSON response file from `https://www.worldtradingdata.com/api/v1/stock?symbol={symbol}&api_token={token}`.
2. Select the stocks based on the stock exchange query param and return the filtered result.
3. The data can be accessed through `http://<server>/api/v1/stocks/{symbol}?stock_exchange={stock exchange values}`
   
   e.g. 
   
   - http://<server>/api/v1/stocks/AAPL?stock_exchange=NASDAQ
   - http://<server>/api/v1/stocks/NOG,AAPL?stock_exchange=AMEX,NASDAQ
   - http://<server>/api/v1/stocks/NOG,AAPL

   
# The application flow

1. Application starts with the file 'main.go'
2. The first time the application runs, it loads the content of the config.json (contains property details like url, token etc).
3. 'main.go' calls 'LoadRoutes' from 'routes' package to prepare the application for handling request and response.
4. 'routes/routes.go' defines all the paths that the application can serve and accept.
5. 'controllers/web.go' handles the processing of the web request and response. 

# Running the app

The binary file is available in the 'stock_app' directory.

1. Go to 'stock_app' directory and execute './stocks'
2. Open 'http://localhost:8080', a message indicates that the application is running will be displayed.
3. Hit 'http://<server>/api/v1/stocks/{symbol}?stock_exchange={stock exchange values}' via browser or curl.

# Deploying the app
1. `stock_app` can be compressed and deployed anywhere.

# Tests

Tests are in the 'stocks_test.go' file.

```
$ go test -v
property details successfully loaded
=== RUN   TestAppHealth
Status Code:  200
Content Type:  text/plain; charset=utf-8
Response Body:  I am running...

--- PASS: TestAppHealth (0.00s)
=== RUN   TestMultipleSymbolsAndStockExchange
Status Code:  200
Content Type:  application/json
Response Body:  [
        {
                "stock_exchange": "AMEX",
                "stocks": [
                        {
                                "symbol": "NOG",
                                "name": "Northern Oil \u0026 Gas, Inc.",
                                "price": "2.73",
                                "close_yesterday": "2.61",
                                "currency": "USD",
                                "market_cap": "1029163597",
                                "volume": "4792239",
                                "timezone": "EDT",
                                "timezone_name": "America/New_York",
                                "gmt_offset": "-14400",
                                "last_trade_time": "2019-04-22 16:00:00"
                        }
                ]
        },
        {
                "stock_exchange": "NASDAQ",
                "stocks": [
                        {
                                "symbol": "AAPL",
                                "name": "Apple Inc.",
                                "price": "204.53",
                                "close_yesterday": "203.86",
                                "currency": "USD",
                                "market_cap": "964416212644",
                                "volume": "19439545",
                                "timezone": "EDT",
                                "timezone_name": "America/New_York",
                                "gmt_offset": "-14400",
                                "last_trade_time": "2019-04-22 16:00:01"
                        }
                ]
        }
]

--- PASS: TestMultipleSymbolsAndStockExchange (1.02s)
=== RUN   TestWithSingleSymbolAndStockExchange
Status Code:  200
Content Type:  application/json
Response Body:  [
        {
                "stock_exchange": "NASDAQ",
                "stocks": [
                        {
                                "symbol": "AAPL",
                                "name": "Apple Inc.",
                                "price": "204.53",
                                "close_yesterday": "203.86",
                                "currency": "USD",
                                "market_cap": "964416212644",
                                "volume": "19439545",
                                "timezone": "EDT",
                                "timezone_name": "America/New_York",
                                "gmt_offset": "-14400",
                                "last_trade_time": "2019-04-22 16:00:01"
                        }
                ]
        }
]

--- PASS: TestWithSingleSymbolAndStockExchange (0.23s)
PASS

```
