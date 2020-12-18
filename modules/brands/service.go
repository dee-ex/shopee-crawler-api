package brands

import (

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

func (serv *Service) Create(data BrandCreation) (*entities.Brand, error) {
  brand := entities.Brand{0, data.Shopid, data.Username, data.BrandName, data.Logo}
  err := serv.repo.Create(&brand)
  if err != nil {
    return nil, err
  }
  return &brand, nil
}

func (serv *Service) GetAllBrands() ([]entities.Brand, error) {
  brands, err := serv.repo.GetAllBrands()
  return brands, err
}

func (serv *Service) GetBrandByID(brand_id int) (*entities.Brand, error) {
  brand, err := serv.repo.GetBrandByID(brand_id)
  return brand, err
}

func (serv *Service) GetBrandByShopid(shopid uint64) (*entities.Brand, error) {
  brand, err := serv.repo.GetBrandByShopid(shopid)
  return brand, err
}

func (serv *Service) GetBrandByUsername(username string) (*entities.Brand, error) {
  brand, err := serv.repo.GetBrandByUsername(username)
  return brand, err
}

func (serv *Service) IsDuplicateShopid(shopid uint64) (bool, error) {
  brand, err := serv.GetBrandByShopid(shopid)
  if err != nil {
    return true, err
  }
  if brand.ID != 0 {
    return true, nil
  }
  return false, nil
}

func (serv *Service) IsDuplicateUsername(username string) (bool, error) {
  brand, err := serv.GetBrandByUsername(username)
  if err != nil {
    return true, err
  }
  if brand.ID != 0 {
    return true, nil
  }
  return false, nil
}