package brands

import (
  "strconv"
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
  "github.com/dee-ex/shopee_crawler_api/utils"
)

func HandleCreateBrand(w http.ResponseWriter, r *http.Request) {
  var data BrandCreation
  err := json.NewDecoder(r.Body).Decode(&data)
  if err != nil {
    http.Error(w, err.Error(), 400)
    return
  }
  if len(data.Username)*len(data.BrandName) == 0 {
    http.Error(w, "Empty username or brandname", 400)
    return
  }

  mysql := NewMySQLRepository()
  if mysql == nil {
    http.Error(w, "Cannot connect to database", 500)
    return
  }
  serv := NewService(mysql)

  dup, err := serv.IsDuplicateShopid(data.Shopid)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if dup {
    http.Error(w, "\"shopid\" already exists", 409)
    return
  }

  dup, err = serv.IsDuplicateUsername(data.Username)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if dup {
    http.Error(w, "\"username\" already exists", 409)
    return
  }

  brand, err := serv.Create(data)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  utils.RespondJSON(w, http.StatusOK, brand)
}

func HandleGetAllBrands(w http.ResponseWriter, r *http.Request) {
  mysql := NewMySQLRepository()
  if mysql == nil {
    http.Error(w, "Cannot connect to database", 500)
    return
  }
  serv := NewService(mysql)

  brands, err := serv.GetAllBrands()
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  utils.RespondJSON(w, http.StatusOK, brands)
}

func HandleGetBrandByID(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  brand_id, err := strconv.Atoi(vars["brand_id"])
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

  brand, err := serv.GetBrandByID(brand_id)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if brand.ID == 0 {
    http.Error(w, "Not found", 404)
    return
  }
  utils.RespondJSON(w, http.StatusOK, brand)
}

func HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  brand_id, err := strconv.Atoi(vars["brand_id"])
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

  brand, err := serv.GetBrandByID(brand_id)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if brand.ID == 0 {
    http.Error(w, "Not found", 404)
    return
  }

  var data BrandUpdate
  err = json.NewDecoder(r.Body).Decode(&data)
  if err != nil {
    http.Error(w, err.Error(), 403)
    return
  }

  err = serv.Update(brand.ID, data)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  utils.Respond(w, http.StatusOK, "Updated successfully")
}

func HandleDeleteByID(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  brand_id, err := strconv.Atoi(vars["brand_id"])
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

  brand, err := serv.GetBrandByID(brand_id)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if brand.ID == 0 {
    http.Error(w, "Not found", 404)
    return
  }

  err = serv.Delete(brand.ID)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  utils.Respond(w, http.StatusOK, "Deleted successfully")
}
