package brands

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

func (repo *MockMySQL) Create(brand *entities.Brand) error {
  res := repo.db.Create(brand)
  return res.Error
}

func (repo *MockMySQL) GetAllBrands() ([]entities.Brand, error) {
  var brands []entities.Brand
  res := repo.db.Find(&brands)
  return brands, res.Error
}

func (repo *MockMySQL) GetBrandByID(brand_id int) (*entities.Brand, error) {
  var brand entities.Brand
  res := repo.db.Where("id = ?", brand_id).Find(&brand)
  return &brand, res.Error
}

func (repo *MockMySQL) GetBrandByShopid(shopid uint64) (*entities.Brand, error) {
  var brand entities.Brand
  res := repo.db.Where("shopid = ?", shopid).Find(&brand)
  return &brand, res.Error
}

func (repo *MockMySQL) GetBrandByUsername(username string) (*entities.Brand, error) {
  var brand entities.Brand
  res := repo.db.Where("username = ?", username).Find(&brand)
  return &brand, res.Error
}

func (repo *MockMySQL) UpdateBrandName(brand_id int, brandname string) error {
  res := repo.db.Model(&entities.Brand{}).Where("id = ?", brand_id).Update("brand_name", brandname)
  return res.Error
}

func (repo *MockMySQL) UpdateLogo(brand_id int, logo string) error {
  res := repo.db.Model(&entities.Brand{}).Where("id = ?", brand_id).Update("logo", logo)
  return res.Error
}

func (repo *MockMySQL) Delete(brand_id int) error {
  res := repo.db.Where("id = ?", brand_id).Delete(&entities.Brand{})
  return res.Error
}

func (repo *MockMySQL) GetAllProducts(shopid uint64) ([]entities.Product, error) {
  var products []entities.Product
  res := repo.db.Where("shopid = ?", shopid).Find(&products)
  return products, res.Error
}
