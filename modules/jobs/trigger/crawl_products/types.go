package crawl_products

type (
  JSONProduct struct {
    Nomore bool
    Items []map[string]interface{}
  }
  QueryParameters struct {
    Shopid uint64
    Username string
  }
)
