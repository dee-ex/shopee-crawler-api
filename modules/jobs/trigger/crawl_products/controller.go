package crawl_products

import (
  "log"
  "strconv"
  "net/http"

  "github.com/dee-ex/shopee_crawler_api/utils"
  "github.com/dee-ex/shopee_crawler_api/entities"
)

func HandleCrawl(w http.ResponseWriter, r *http.Request) {
  // limit -1 if we want to select without limit
  var limit, from, to int
  var err error
  limit_query, ok := r.URL.Query()["limit"]
  if !ok {
    limit = 5
  } else {
    limit, err = strconv.Atoi(limit_query[0])
    if err != nil {
      limit = 5
    }
  }
  from_query, ok := r.URL.Query()["from"]
  if !ok {
    from = 0
  } else {
    from, err = strconv.Atoi(from_query[0])
    if err != nil {
      from = 0
    }
  }
  to_query, ok := r.URL.Query()["to"]
  if !ok {
    to = -1
  } else {
    to, err = strconv.Atoi(to_query[0])
    if err != nil {
      to = -1
    }
  }
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
  if from >= limit {
    from = 0
  }
  if to < 0 || to > len(qrpmses)  || to <= from {
    to = len(qrpmses)
  }
  var all_products []entities.Product
  for _, qrpms := range qrpmses[from:to] {
    products, err := Crawl(qrpms.Shopid, qrpms.Username)
    if err != nil {
      log.Println(err.Error())
    }
    log.Println(len(products))
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
