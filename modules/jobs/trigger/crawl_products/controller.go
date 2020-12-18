package crawl_products

import (
  "log"
  "strconv"
  "net/http"

  "github.com/dee-ex/shopee_crawler_api/utils"
  "github.com/dee-ex/shopee_crawler_api/entities"
)

func GetLimit(r *http.Request) int {
  limit_query, ok := r.URL.Query()["limit"]
  if !ok {
    return 5
  }
  limit, err := strconv.Atoi(limit_query[0])
  if err != nil {
    return 5
  }
  return limit
}

func GetFrom(r *http.Request, _max int) int {
  from_query, ok := r.URL.Query()["from"]
  if !ok {
    return 0
  }
  from, err := strconv.Atoi(from_query[0])
  if err != nil {
    return 0
  }
  if from >= _max {
    return 0
  }
  return from
}

func GetTo(r *http.Request, _max, from int) int {
  to_query, ok := r.URL.Query()["to"]
  if !ok {
    return _max
  }
  to, err := strconv.Atoi(to_query[0])
  if err != nil {
    return _max
  }
  if to > _max || to <= from {
    return _max
  }
  return to
}

func HandleCrawl(w http.ResponseWriter, r *http.Request) {
  // limit -1 if we want to select without limit
  limit := GetLimit(r)

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
