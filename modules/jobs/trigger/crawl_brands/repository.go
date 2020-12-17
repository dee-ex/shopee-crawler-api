package crawl_brands

import (
  "gorm.io/gorm"

  "github.com/dee-ex/shopee_crawler_api/entities"
  "github.com/dee-ex/shopee_crawler_api/infrastructures"
)

type (
  Repository interface {
    Create(brand *entities.Brand) error
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
  return repo.db.Create(brand).Error
}
