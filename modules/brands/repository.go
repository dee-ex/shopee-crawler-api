package brands

import (
  "gorm.io/gorm"

  "github.com/dee-ex/shopee_crawler_api/entities"
  "github.com/dee-ex/shopee_crawler_api/infrastructures"
)

type (
  Repository interface {
    Create(brand *entities.Brand) error
    GetAllBrands() ([]entities.Brand, error)
    GetBrandByID(brand_id int) (*entities.Brand, error)
    GetBrandByShopid(shopid uint64) (*entities.Brand, error)
    GetBrandByUsername(username string) (*entities.Brand, error)
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

func (repo *MySQL) Create(brand *entities.Brand) error {
  res := repo.db.Create(brand)
  return res.Error
}

func (repo *MySQL) GetAllBrands() ([]entities.Brand, error) {
  var brands []entities.Brand
  res := repo.db.Find(&brands)
  return brands, res.Error
}

func (repo *MySQL) GetBrandByID(brand_id int) (*entities.Brand, error) {
  var brand entities.Brand
  res := repo.db.Where("id = ?", brand_id).Find(&brand)
  return &brand, res.Error
}

func (repo *MySQL) GetBrandByShopid(shopid uint64) (*entities.Brand, error) {
  var brand entities.Brand
  res := repo.db.Where("shopid = ?", shopid).Find(&brand)
  return &brand, res.Error
}

func (repo *MySQL) GetBrandByUsername(username string) (*entities.Brand, error) {
  var brand entities.Brand
  res := repo.db.Where("username = ?", username).Find(&brand)
  return &brand, res.Error
}
