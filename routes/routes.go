package routes

import (
	"github.com/rburawes/stocks/controllers"
	"net/http"
)

// LoadRoutes handles routes to pages of the application.
func LoadRoutes() {

	// Index or main page.
	http.HandleFunc("/", controllers.Index)
	// Stocks endpoint
	http.HandleFunc("/api/v1/stocks/", controllers.Stocks)
	// Handles application icon if not available.
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// Listens and serve requests.
	http.ListenAndServe(":8080", nil)

}
