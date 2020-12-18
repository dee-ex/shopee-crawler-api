package crawl_brands

import (
  "fmt"
  "errors"
  "net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/dee-ex/shopee_crawler_api/entities"
)

type (
  Service struct {
    repo Repository
  }
)

func NewService(repo Repository) *Service {
  return &Service{repo: repo}
}

func Crawl() ([]entities.Brand, error) {
  res, err := http.Get("https://shopee.vn/api/v2/brand_lists/get")
  if err != nil {
    return nil, err
  }
  if res.StatusCode != 200 {
    return nil, errors.New(fmt.Sprintf("Status code error: %d %s", res.StatusCode, res.Status))
  }

  body, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    return nil, err
  }

  var raw_json JSONBrand
  var brands []entities.Brand
  json.Unmarshal(body, &raw_json)
  for _, v := range raw_json.Data {
    shopid := uint64(v.(map[string]interface{})["shopid"].(float64))
    username := v.(map[string]interface{})["username"].(string)
    brand_name := v.(map[string]interface{})["brand_name"].(string)
    logo := v.(map[string]interface{})["logo"].(string)

    brands = append(brands, entities.Brand{0, shopid, username, brand_name, logo})
  }
  return brands, nil
}

func (serv *Service) Create(brand *entities.Brand) error {
  err := serv.repo.Create(brand)
  return err
}
