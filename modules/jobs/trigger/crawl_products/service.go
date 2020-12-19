package crawl_products

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "crypto/md5"
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

func GenerateIfNoneMatch_(query string) string {
  return fmt.Sprintf("%x", md5.Sum([]byte("55b03" + fmt.Sprintf("%x", md5.Sum([]byte(query)))  +"55b03")))
}

func GenerateRequest(limit, offset int, shopid uint64, username string) (*http.Request, error) {
  query := fmt.Sprintf("by=pop&limit=%d&match_id=%d&newest=%d&order=desc&page_type=shop&version=2", limit, shopid, offset)
  req, err := http.NewRequest("GET", "https://shopee.vn/api/v2/search_items/?" + query, nil)
  if err != nil {
    return nil, err
  }
  req.Header.Add("if-none-match-", "55b03-" + GenerateIfNoneMatch_(query))
  req.Header.Add("authority", "shopee.vn")
  req.Header.Add("referer", "https://shopee.vn/" + username)
  req.Header.Add("x-api-source", "pc")
  return req, nil
}

func Crawl(shopid uint64, username string) ([]entities.Product, error) {
  var products []entities.Product
  var raw_json JSONProduct
  limit, offset := 100, 0
  for ; !raw_json.Nomore; {
    req, err := GenerateRequest(limit, offset, shopid, username)
    if err != nil {
      return nil, err
    }

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
      return nil, err
    }

    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
      return nil, err
    }

    json.Unmarshal(body, &raw_json)
    for _, item := range raw_json.Items {
      shopid := uint64(item["shopid"].(float64))
      itemid := uint64(item["itemid"].(float64))

      pricemax := uint64(item["price_max_before_discount"].(float64) / 100000)
      if pricemax == 0 {
        pricemax = uint64(item["price_max"].(float64) / 100000)
      }
      pricemin := uint64(item["price_min_before_discount"].(float64) / 100000)
      if pricemin == 0 {
        pricemin = uint64(item["price_min"].(float64) / 100000)
      }

      name := item["name"].(string)

      images := ""
      for _, img := range item["images"].([]interface{}) {
        images += img.(string) + "-"
      }
      images = images[:len(images) - 1]

      historical_sold := int(item["historical_sold"].(float64))

      rating := ""
      for _, r := range item["item_rating"].(map[string]interface{})["rating_count"].([]interface{}) {
        rating += fmt.Sprintf("%0.f", r.(float64)) + "-"
      }
      rating = rating[:len(rating) - 1]

      products = append(products, entities.Product{0, shopid, itemid, pricemax, pricemin, name, images, historical_sold, rating})
    }
    offset += limit
  }
  return products, nil
}

func (serv *Service) GetShopidsAndUsernames(limit int) ([]QueryParameters, error) {
  qrpmses, err := serv.repo.GetShopidsAndUsernames(limit)
  return qrpmses, err
}

func (serv *Service) Create(product *entities.Product) error {
  err := serv.repo.Create(product)
  return err
}
