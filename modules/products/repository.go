package products

import (
  "gorm.io/gorm"

  "github.com/dee-ex/shopee_crawler_api/entities"
  "github.com/dee-ex/shopee_crawler_api/infrastructures"
)

type (
  Repository interface {
    Create(product *entities.Product) error
    GetAllProducts(limit, offset int) ([]entities.Product, error)
    GetProductByID(product_id int) (*entities.Product, error)
    GetProductByItemid(itemid uint64) (*entities.Product, error)
    UpdateShopid(product_id int, shopid uint64) error
    UpdatePriceMax(product_id int, pricemax uint64) error
    UpdatePriceMin(product_id int, pricemin uint64) error
    UpdateName(product_id int, name string) error
    UpdateImages(product_id int, images string) error
    UpdateHistoricalSold(product_id int, historicalsold uint32) error
    UpdateRating(product_id int, rating string) error
    Delete(product_id int) error
  }

  MySQL struct {
    db *gorm.DB
  }
)

func NewMySQLRepository() *MySQL {
  db, err := infrastructures.NewMySQLSession()
  if err != nil {
    return nil
  }
  return &MySQL{db: db}
}

func (repo *MySQL) Create(product *entities.Product) error {
  res := repo.db.Create(product)
  return res.Error
}

func (repo *MySQL) GetAllProducts(limit, offset int) ([]entities.Product, error) {
  var products []entities.Product
  res := repo.db.Limit(limit).Offset(offset).Find(&products)
  return products, res.Error
}

func (repo *MySQL) GetProductByID(product_id int) (*entities.Product, error) {
  var product entities.Product
  res := repo.db.Where("id = ?", product_id).Find(&product)
  return &product, res.Error
}

func (repo *MySQL) GetProductByItemid(itemid uint64) (*entities.Product, error) {
  var product entities.Product
  res := repo.db.Where("itemid = ?", itemid).Find(&product)
  return &product, res.Error
}

func (repo *MySQL) UpdateShopid(product_id int, shopid uint64) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("shopid", shopid)
  return res.Error
}

func (repo *MySQL) UpdatePriceMax(product_id int, pricemax uint64) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("price_max", pricemax)
  return res.Error
}

func (repo *MySQL) UpdatePriceMin(product_id int, pricemin uint64) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("price_min", pricemin)
  return res.Error
}

func (repo *MySQL) UpdateName(product_id int, name string) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("name", name)
  return res.Error
}

func (repo *MySQL) UpdateImages(product_id int, images string) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("images", images)
  return res.Error
}

func (repo *MySQL) UpdateHistoricalSold(product_id int, historicalsold uint32) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("historical_sold", historicalsold)
  return res.Error
}

func (repo *MySQL) UpdateRating(product_id int, rating string) error {
  res := repo.db.Model(&entities.Product{}).Where("id = ?", product_id).Update("rating", rating)
  return res.Error
}

func (repo *MySQL) Delete(product_id int) error {
  res := repo.db.Where("id = ?", product_id).Delete(&entities.Product{})
  return res.Error
}
