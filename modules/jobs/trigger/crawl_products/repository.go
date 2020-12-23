package crawl_products

import (
  "gorm.io/gorm"

  "github.com/dee-ex/shopee_crawler_api/entities"
  "github.com/dee-ex/shopee_crawler_api/infrastructures"
)

type (
  Repository interface {
    GetShopidsAndUsernames(limit int) ([]QueryParameters, error)
    GetShopidByUsername(username string) (uint64, error)
    Create(product *entities.Product) error
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

func (repo *MySQL) GetShopidsAndUsernames(limit int) ([]QueryParameters, error) {
  var qrpmses []QueryParameters
  res := repo.db.Model(&entities.Brand{}).Select("shopid", "username").Limit(limit).Find(&qrpmses)
  return qrpmses, res.Error
}

func (repo *MySQL) GetShopidByUsername(username string) (uint64, error) {
  var shopid uint64
  res := repo.db.Model(&entities.Brand{}).Where("username = ?", username).Select("shopid").Find(&shopid)
  return shopid, res.Error
}

func (repo *MySQL) Create(product *entities.Product) error {
  res := repo.db.Create(product)
  return res.Error
}
