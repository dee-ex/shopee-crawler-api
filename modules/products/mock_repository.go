package products

import (
  "gorm.io/gorm"

  "github.com/dee-ex/shopee_crawler_api/entities"
  "github.com/dee-ex/shopee_crawler_api/infrastructures"
)

type (
  MockMySQL struct {
    db *gorm.DB
  }
)

func NewMockMySQLRepository() *MockMySQL {
  db, err := infrastructures.NewMockMySQLSession()
  if err != nil {
    return nil
  }
  return &MockMySQL{db: db}
}

func (repo *MockMySQL) Create(product *entities.Product) error {
  res := repo.db.Create(product)
  return res.Error
}

func (repo *MockMySQL) GetAllProducts(limit, offset int) ([]entities.Product, error) {
  var products []entities.Product
  res := repo.db.Limit(limit).Offset(offset).Find(&products)
  return products, res.Error
}

func (repo *MockMySQL) GetProductByID(product_id int) (*entities.Product, error) {
  var product entities.Product
  res := repo.db.Where("id = ?", product_id).Find(&product)
  return &product, res.Error
}

func (repo *MockMySQL) GetProductByItemid(itemid uint64) (*entities.Product, error) {
  var product entities.Product
  res := repo.db.Where("itemid = ?", itemid).Find(&product)
  return &product, res.Error
}

func (repo *MockMySQL) UpdateShopid(product_id int, shopid uint64) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("shopid", shopid)
  return res.Error
}

func (repo *MockMySQL) UpdatePriceMax(product_id int, pricemax uint64) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("price_max", pricemax)
  return res.Error
}

func (repo *MockMySQL) UpdatePriceMin(product_id int, pricemin uint64) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("price_min", pricemin)
  return res.Error
}

func (repo *MockMySQL) UpdateName(product_id int, name string) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("name", name)
  return res.Error
}

func (repo *MockMySQL) UpdateImages(product_id int, images string) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("images", images)
  return res.Error
}

func (repo *MockMySQL) UpdateHistoricalSold(product_id int, historicalsold uint32) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("historical_sold", historicalsold)
  return res.Error
}

func (repo *MockMySQL) UpdateRating(product_id int, rating string) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("rating", rating)
  return res.Error
}

func (repo *MockMySQL) Delete(product_id int) error {
  res := repo.db.Where("id = ?", product_id).Delete(&entities.Product{})
  return res.Error
}
