package crawl_products

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/dee-ex/shopee_crawler_api/utils"
  "github.com/dee-ex/shopee_crawler_api/entities"
)

func GetFrom(r *http.Request, _max int) int {
  from :=  utils.GetNumericQuery(r, "from", 0)
  if from >= _max {
    return 0
  }
  return from
}

func GetTo(r *http.Request, _max, from int) int {
  to :=  utils.GetNumericQuery(r, "to", _max)
  if to > _max || to <= from {
    return _max
  }
  return to
}

func HandleCrawlProducts(w http.ResponseWriter, r *http.Request) {
  // limit -1 if we want to select without limit
  limit := utils.GetNumericQuery(r, "limit", 5)

  mysql := NewMySQLRepository()
  if mysql == nil {
    http.Error(w, "Cannot connect to database.", 500)
    return
  }
  serv := NewService(mysql)

  qrpmses, err := serv.GetShopidsAndUsernames(limit)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }

  from := GetFrom(r, len(qrpmses))
  to := GetTo(r, len(qrpmses), from)

  var all_products []entities.Product
  for _, qrpms := range qrpmses[from:to] {
    products, err := Crawl(qrpms.Shopid, qrpms.Username)
    if err != nil {
      log.Println(err.Error())
    }
    log.Println(qrpms.Shopid, qrpms.Username, len(products))

    mysql := NewMySQLRepository()
    if mysql == nil {
      http.Error(w, "Cannot connect to database.", 500)
      continue
    }
    serv := NewService(mysql)

    for _, product := range products {
      err = serv.Create(&product)
      if err != nil {
        log.Println(err.Error())
      }
    }

    if len(all_products) + len(products) <= 1000 {
      all_products = append(all_products, products...)
    }
  }
  utils.RespondJSON(w, http.StatusOK, all_products)
}

func HandleCrawlByUsername(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  mysql := NewMySQLRepository()
  if mysql == nil {
    http.Error(w, "Cannot connect to database.", 500)
    return
  }
  serv := NewService(mysql)

  shopid, err := serv.GetShopidByUsername(vars["brand_username"])
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if shopid == 0 {
    http.Error(w, "Not found", 404)
    return
  }

  products, err := Crawl(shopid, vars["brand_username"])
  if err != nil {
    log.Fatal(err.Error())
  }

  for _, product := range products {
    err = serv.Create(&product)
    if err != nil {
      log.Println(err.Error())
    }
  }
  utils.RespondJSON(w, http.StatusOK, products)
}
