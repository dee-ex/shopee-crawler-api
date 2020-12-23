package products

import (
  "testing"
)

var (
  testshopid uint64 = 999999
  testitemid uint64 = 999998
  testpricemax uint64 = 999997
  testpricemin uint64 = 999996
  testname string = "testname"
  testimages string = "testimage"
  testhistoricalsold uint32 = 999995
  testrating string = "testrating"


  updateshopid uint64 = 999994
  updatepricemax uint64 = 999993
  updatepricemin uint64 = 999992
  updatename string = "updatename"
  updateimages string = "updateimage"
  updatehistoricalsold uint32 = 999991
  updaterating string = "updaterating"
)

func TestCreate(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  data := ProductCreation{testshopid, testitemid, testpricemax, testpricemin, testname, testimages, testhistoricalsold, testrating}
  _, err := serv.Create(data)
  if err != nil {
    t.Fatal("Server error")
  }
}

func TestGetAllProducts(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  _, err := serv.repo.GetAllProducts(-1, 0)
  if err != nil {
    t.Fatal("Server error")
  }
}

func TestGetProductByID(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  product, err := serv.GetProductByID(-1)
  if err != nil {
    t.Fatal("Server error")
  }
  if product.ID != 0 {
    t.Error("Select negative id")
  }

  product, err = serv.GetProductByID(999999)
  if err != nil {
    t.Fatal("Server error")
  }
  if product.ID != 0 {
    t.Error("Select non-exists id")
  }
}

func TestGetProductByItemid(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  product, err := serv.GetProductByItemid(0)
  if err != nil {
    t.Fatal("Server error")
  }
  if product.ID != 0 {
    t.Error("Select non-exists id")
  }

  product, err = serv.GetProductByItemid(testitemid)
  if err != nil {
    t.Fatal("Server error")
  }
  if product.Itemid != testitemid && product.ID != 0 {
    t.Error("Select wrong object")
  }
}

func TestIsDuplicateItemid(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  dup, err := serv.IsDuplicateItemid(0)
  if err != nil {
    t.Fatal("Server error")
  }
  if dup {
    t.Error("Fail to verify duplicate Itemid. FT")
  }


  dup, err = serv.IsDuplicateItemid(testitemid)
  if err != nil {
    t.Fatal("Server error")
  }
  if !dup {
    t.Error("Fail to verify duplicate Itemid. TF")
  }
}

func TestUpdate(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  product, err := serv.GetProductByItemid(testitemid)
  if err != nil {
    t.Fatal("Server error")
  }

  data := ProductUpdate{&updateshopid, &updatepricemax, &updatepricemin, &updatename, &updateimages, &updatehistoricalsold, &updaterating}
  err = serv.Update(product.ID, data)
  product, err = serv.GetProductByID(product.ID)
  if err != nil {
    t.Fatal("Server error")
  }
  if product.Shopid != updateshopid {
    t.Error("Fail to update Shopid")
  }
  if product.PriceMax != updatepricemax {
    t.Error("Fail to update Pricemax")
  }
  if product.PriceMin != updatepricemin {
    t.Error("Fail to update Pricemin")
  }
  if product.Name != updatename {
    t.Error("Fail to update Name")
  }
  if product.Images != updateimages {
    t.Error("Fail to update Images")
  }
  if product.HistoricalSold != updatehistoricalsold {
    t.Error("Fail to update HistoricalSold")
  }
  if product.Rating != updaterating {
    t.Error("Fail to update Rating")
  }
}

func TestDelete(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)


  product, err := serv.GetProductByItemid(testitemid)
  if err != nil {
    t.Fatal("Server error")
  }

  err = serv.Delete(product.ID)
  if err != nil {
    t.Fatal("Server error")
  }
  product, err = serv.GetProductByID(product.ID)
  if err != nil {
    t.Fatal("Server error")
  }
  if product.ID != 0 {
    t.Error("Fail to delete product")
  }
}
