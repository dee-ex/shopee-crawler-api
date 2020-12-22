package main

import (
  "os"
  "log"
  "net/http"

  "github.com/rs/cors"
  "github.com/joho/godotenv"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "github.com/dee-ex/shopee_crawler_api/modules/brands"
  "github.com/dee-ex/shopee_crawler_api/modules/products"
  "github.com/dee-ex/shopee_crawler_api/modules/jobs/trigger/crawl_brands"
  "github.com/dee-ex/shopee_crawler_api/modules/jobs/trigger/crawl_products"
)

func start_server() {
  router := mux.NewRouter().StrictSlash(true)


  router.HandleFunc("/brands", brands.HandleCreateBrand).Methods("POST")
  router.HandleFunc("/brands", brands.HandleGetAllBrands).Methods("GET")
  router.HandleFunc("/brands/{brand_id}", brands.HandleGetBrandByID).Methods("GET")
  router.HandleFunc("/brands/{brand_id}", brands.HandleUpdateByID).Methods("PUT")
  router.HandleFunc("/brands/{brand_id}", brands.HandleDeleteByID).Methods("DELETE")
  router.HandleFunc("/brands/{brand_id}/products", brands.HandleGetAllProducts).Methods("GET")

  router.HandleFunc("/products", products.HandleCreateProduct).Methods("POST")
  router.HandleFunc("/products", products.HandleGetAllProducts).Methods("GET")
  router.HandleFunc("/products/{product_id}", products.HandleGetProductByID).Methods("GET")
  router.HandleFunc("/products/{product_id}", products.HandleUpdateByID).Methods("PUT")
  router.HandleFunc("/products/{product_id}", products.HandleDeleteByID).Methods("DELETE")

  router.HandleFunc("/jobs/trigger/crawl_brands", crawl_brands.HandleCrawl).Methods("GET")
  router.HandleFunc("/jobs/trigger/crawl_products", crawl_products.HandleCrawl).Methods("GET")


  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
  })
  enhanced_router := c.Handler(router)
  enhanced_router = handlers.LoggingHandler(os.Stdout, enhanced_router)

  log.Fatal(http.ListenAndServe(":" + os.Getenv("SERVER_PORT"), enhanced_router))
}

func main() {
  err := godotenv.Load("env/database.env")
  if err != nil {
    log.Fatal("Error loading database.env file")
  }
  if os.Getenv("DATABASE_CONFIGURED") == "NO" {
    log.Fatal("Database is not configured yet")
  }

  err = godotenv.Load("env/server.env")
  if err != nil {
    log.Fatal("Error loading server.env file")
  }

  log.Println("Hosting: Local server: localhost:" + os.Getenv("SERVER_PORT"))
  start_server()
}
