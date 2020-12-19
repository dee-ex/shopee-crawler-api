package products

import (
  "errors"

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

func (serv *Service) Create(data ProductCreation) (*entities.Product, error) {
  product := entities.Product{0, data.Shopid, data.Itemid, data.PriceMax, data.PriceMin, data.Name, data.Images, data.HistoricalSold, data.Rating}
  err := serv.repo.Create(&product)
  if err != nil {
    return nil, err
  }
  return &product, nil
}

func (serv *Service) GetAllProducts(limit, offset int) ([]entities.Product, error) {
  products, err := serv.repo.GetAllProducts(limit, offset)
  return products, err
}

func (serv *Service) GetProductByID(product_id int) (*entities.Product, error) {
  product, err := serv.repo.GetProductByID(product_id)
  return product, err
}

func (serv *Service) GetProductByItemid(itemid uint64) (*entities.Product, error) {
  product, err := serv.repo.GetProductByItemid(itemid)
  return product, err
}

func (serv *Service) IsDuplicateItemid(itemid uint64) (bool, error) {
  product, err := serv.GetProductByItemid(itemid)
  if err != nil {
    return true, err
  }
  if product.ID != 0 {
    return true, nil
  }
  return false, nil
}

func (serv *Service) Update(product_id int, data ProductUpdate) error {
  if data.Shopid != nil {
    if *data.Shopid == 0 {
      return errors.New("\"Shopid\" cannot be 0")
    }
    err := serv.repo.UpdateShopid(product_id, *data.Shopid)
    if err != nil {
      return err
    }
  }


  if data.PriceMax != nil {
    err := serv.repo.UpdatePriceMax(product_id, *data.PriceMax)
    if err != nil {
      return err
    }
  }

  if data.PriceMin != nil {
    err := serv.repo.UpdatePriceMin(product_id, *data.PriceMin)
    if err != nil {
      return err
    }
  }


  if data.Name != nil {
    if *data.Name == "" {
      return errors.New("\"Name\" cannot be empty")
    }
    err := serv.repo.UpdateName(product_id, *data.Name)
    if err != nil {
      return err
    }
  }

  if data.Images != nil {
    err := serv.repo.UpdateImages(product_id, *data.Images)
    if err != nil {
      return err
    }
  }

  if data.HistoricalSold != nil {
    err := serv.repo.UpdateHistoricalSold(product_id, *data.HistoricalSold)
    if err != nil {
      return err
    }
  }

  if data.Rating != nil {
    err := serv.repo.UpdateRating(product_id, *data.Rating)
    if err != nil {
      return err
    }
  }

  return nil
}

func (serv *Service) Delete(product_id int) error {
  err := serv.repo.Delete(product_id)
  return err
}
