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
    UpdateBrandName(brand_id int, brandname string) error
    UpdateLogo(brand_id int, logo string) error
    Delete(brand_id int) error
    GetAllProducts(shopid uint64) ([]entities.Product, error)
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

func (repo *MySQL) UpdateBrandName(brand_id int, brandname string) error {
  res := repo.db.Model(&entities.Brand{}).Where("id = ?", brand_id).Update("brand_name", brandname)
  return res.Error
}

func (repo *MySQL) UpdateLogo(brand_id int, logo string) error {
  res := repo.db.Model(&entities.Brand{}).Where("id = ?", brand_id).Update("logo", logo)
  return res.Error
}

func (repo *MySQL) Delete(brand_id int) error {
  res := repo.db.Where("id = ?", brand_id).Delete(&entities.Brand{})
  return res.Error
}

func (repo *MySQL) GetAllProducts(shopid uint64) ([]entities.Product, error) {
  var products []entities.Product
  res := repo.db.Where("shopid = ?", shopid).Find(&products)
  return products, res.Error
}
