package brands

import (
  "testing"
)

var (
  testshopid uint64 = 999999
  testusername string = "testusername"
  testbrandname string = "testbrandname"
  testlogo string = "testlogo"

  updatebrandname string = "updatebrandname"
  updatelogo string = "updatelogo"
)


func TestCreate(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  data := BrandCreation{testshopid, testusername, testbrandname, testlogo}
  _, err := serv.Create(data)
  if err != nil {
    t.Fatal("Server error")
  }
}

func TestGetAllBrands(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  _, err := serv.GetAllBrands()
  if err != nil {
    t.Fatal("Server error")
  }
}

func TestGetBrandByID(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  brand, err := serv.GetBrandByID(-1)
  if err != nil {
    t.Fatal("Server error")
  }
  if brand.ID != 0 {
    t.Error("Select negative id")
  }

  brand, err = serv.GetBrandByID(999999)
  if err != nil {
    t.Fatal("Server error")
  }
  if brand.ID != 0 {
    t.Error("Select non-exists id")
  }
}

func TestGetBrandByShopid(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  brand, err := serv.GetBrandByShopid(10)
  if err != nil {
    t.Fatal("Server error")
  }
  if brand.ID != 0 {
    t.Error("Select non-exists id")
  }

  brand, err = serv.GetBrandByShopid(testshopid)
  if err != nil {
    t.Fatal("Server error")
  }
  if brand.Shopid != testshopid && brand.ID != 0 {
    t.Error("Select wrong object")
  }
}

func TestGetBrandByUsername(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  brand, err := serv.GetBrandByUsername("non-exists_username")
  if err != nil {
    t.Fatal("Server error")
  }
  if brand.ID != 0 {
    t.Error("Select non-exists username")
  }


  brand, err = serv.GetBrandByUsername(testusername)
  if err != nil {
    t.Fatal("Server error")
  }
  if brand.Username != testusername && brand.ID != 0 {
    t.Error("Select wrong object")
  }
}

func TestIsDuplicateShopid(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  dup, err := serv.IsDuplicateShopid(10)
  if err != nil {
    t.Fatal("Server error")
  }
  if dup {
    t.Error("Fail to verify duplicate Shopid. FT")
  }

  dup, err = serv.IsDuplicateShopid(testshopid)
  if err != nil {
    t.Fatal("Server error")
  }
  if !dup {
    t.Error("Fail to verify duplicate Shopid. TF")
  }
}

func TestIsDuplicateUsername(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  dup, err := serv.IsDuplicateUsername("non-exists_username")
  if err != nil {
    t.Fatal("Server error")
  }
  if dup {
    t.Error("Fail to verify duplicate Username. FT")
  }


  dup, err = serv.IsDuplicateUsername(testusername)
  if err != nil {
    t.Fatal("Server error")
  }
  if !dup {
    t.Error("Fail to verify duplicate Username. TF")
  }
}

func TestUpdate(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  brand, err := serv.GetBrandByUsername(testusername)
  if err != nil {
    t.Fatal("Server error")
  }

  data := BrandUpdate{&updatebrandname, &updatelogo}
  err = serv.Update(brand.ID, data)
  brand, err = serv.GetBrandByID(brand.ID)
  if err != nil {
    t.Fatal("Server error")
  }
  if brand.BrandName != updatebrandname {
    t.Error("Fail to update BrandName")
  }
  if brand.Logo != updatelogo {
    t.Error("Fail to update Logo")
  }
}

func TestDelete(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)


  brand, err := serv.GetBrandByUsername(testusername)
  if err != nil {
    t.Fatal("Server error")
  }

  err = serv.Delete(brand.ID)
  if err != nil {
    t.Fatal("Server error")
  }
  brand, err = serv.GetBrandByID(brand.ID)
  if err != nil {
    t.Fatal("Server error")
  }
  if brand.ID != 0 {
    t.Error("Fail to delete brand")
  }
}

func TestGetAllProducts(t *testing.T) {
  mock := NewMockMySQLRepository()
  if mock == nil {
    t.Fatal("Cannot connect to database")
  }
  serv := NewService(mock)

  brands, err := serv.GetAllBrands()
  if err != nil {
    t.Fatal("Server error")
  }

  testbrand := brands[0]
  products, err := serv.GetAllProducts(testbrand.Shopid)
  if err != nil {
    t.Fatal("Server error")
  }
  if len(products) == 0 {
    t.Errorf("Fail to get all product of Shopid: %d", testbrand.Shopid)
  }
}
