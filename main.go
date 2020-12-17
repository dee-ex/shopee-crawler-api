package main

import (
  "os"
  "log"
  "net/http"

  "github.com/rs/cors"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "github.com/dee-ex/shopee_crawler_api/modules/jobs/trigger/crawl_brands"
  "github.com/dee-ex/shopee_crawler_api/modules/jobs/trigger/crawl_products"
)

func start_server() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/done", crawl_brands.HandleCrawl).Methods("GET")
  router.HandleFunc("/", crawl_products.HandleCrawl).Methods("GET")
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
  })
  enhanced_router := c.Handler(router)
  enhanced_router = handlers.LoggingHandler(os.Stdout, enhanced_router)
  log.Fatal(http.ListenAndServe(":8080", enhanced_router))
}

func main() {
  log.Println("Hosting: Local server: localhost:8080")
  start_server()
}
