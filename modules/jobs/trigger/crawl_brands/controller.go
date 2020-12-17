package crawl_brands

import (
  "log"
  "net/http"

  "github.com/dee-ex/shopee_crawler_api/utils"
)

func HandleCrawl(w http.ResponseWriter, r *http.Request) {
  brands, err := Crawl()
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  mysql := NewMySQLRepository()
  if mysql == nil {
    http.Error(w, "Cannot connect to database", 500)
    return
  }
  serv := NewService(mysql)
  for _, brand := range brands {
    err = serv.Create(&brand)
    if err != nil {
      log.Println(err.Error())
    }
  }
  utils.RespondJSON(w, http.StatusOK, brands)
}
