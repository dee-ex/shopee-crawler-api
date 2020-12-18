package brands

type (
  BrandCreation struct {
    Shopid uint64
    Username string
    BrandName string
    Logo string
  }

  BrandUpdate struct {
    BrandName *string
    Logo *string
  }
)
