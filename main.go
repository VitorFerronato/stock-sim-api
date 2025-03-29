package main

import (
	"log"
	"net/http"
	"stocksim-api/config"
	"stocksim-api/handlers"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.NewConfig()
	defer cfg.DB.Close()

	h := handlers.NewHandler(cfg)

	http.HandleFunc("/api/search_stock", h.SearchStock)
	http.HandleFunc("/api/stock_price", h.StockPrice)
	http.HandleFunc("/api/stocks", h.SaveStock)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
