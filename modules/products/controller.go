package products

import (
  "strconv"
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
  "github.com/dee-ex/shopee_crawler_api/utils"
)

func HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
  var data ProductCreation
  err := json.NewDecoder(r.Body).Decode(&data)
  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }
  if data.Shopid == 0 {
    http.Error(w, "Empty shopid of product", 400)
    return
  }
  if len(data.Name) == 0 {
    http.Error(w, "Empty name of product", 400)
    return
  }

  mysql := NewMySQLRepository()
  if mysql == nil {
    http.Error(w, "Cannot connect to database", 500)
    return
  }
  serv := NewService(mysql)

  dup, err := serv.IsDuplicateItemid(data.Itemid)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if dup {
    http.Error(w, "\"Itemid\" already exists", 409)
    return
  }

  product, err := serv.Create(data)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  utils.RespondJSON(w, http.StatusOK, product)
}

func HandleGetAllProducts(w http.ResponseWriter, r *http.Request) {
  limit := utils.GetNumericQuery(r, "limit", 100)
  offset := utils.GetNumericQuery(r, "offset", 0)

  mysql := NewMySQLRepository()
  if mysql == nil {
    http.Error(w, "Cannot connect to database", 500)
    return
  }
  serv := NewService(mysql)

  products, err := serv.GetAllProducts(limit, offset)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  utils.RespondJSON(w, http.StatusOK, products)
}

func HandleGetProductByID(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  product_id, err := strconv.Atoi(vars["product_id"])
  if err != nil {
    http.Error(w, err.Error(), 403)
    return
  }

  mysql := NewMySQLRepository()
  if mysql == nil {
    http.Error(w, "Cannot connect to database", 500)
    return
  }
  serv := NewService(mysql)

  product, err := serv.GetProductByID(product_id)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if product.ID == 0 {
    http.Error(w, "Not found", 404)
    return
  }
  utils.RespondJSON(w, http.StatusOK, product)
}

func HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  product_id, err := strconv.Atoi(vars["product_id"])
  if err != nil {
    http.Error(w, err.Error(), 403)
    return
  }

  mysql := NewMySQLRepository()
  if mysql == nil {
    http.Error(w, "Cannot connect to database", 500)
    return
  }
  serv := NewService(mysql)

  product, err := serv.GetProductByID(product_id)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if product.ID == 0 {
    http.Error(w, "Not found", 404)
    return
  }

  var data ProductUpdate
  err = json.NewDecoder(r.Body).Decode(&data)
  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }

  err =serv.Update(product.ID, data)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  utils.Respond(w, http.StatusOK, "Updated successfully")
}

func HandleDeleteByID(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  product_id, err := strconv.Atoi(vars["product_id"])
  if err != nil {
    http.Error(w, err.Error(), 403)
    return
  }

  mysql := NewMySQLRepository()
  if mysql == nil {
    http.Error(w, "Cannot connect to database", 500)
    return
  }
  serv := NewService(mysql)

  product, err := serv.GetProductByID(product_id)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if product.ID == 0 {
    http.Error(w, "Not found", 404)
    return
  }

  err = serv.Delete(product_id)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  utils.Respond(w, http.StatusOK, "Deleted successfully")
}
