package entities

type (
  Brand struct {
    Shopid uint64 `json: "shopid"`
    Username string `json: "username"`
    BrandName string `json: "brand_name"`
    Logo string `json: "logo"`
  }
)
