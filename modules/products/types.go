package products

type (
  ProductCreation struct {
    Shopid uint64
    Itemid uint64
    PriceMax uint64
    PriceMin uint64
    Name string
    Images string
    HistoricalSold uint32
    Rating string
  }

  ProductUpdate struct {
    Shopid *uint64
    PriceMax *uint64
    PriceMin *uint64
    Name *string
    Images *string
    HistoricalSold *uint32
    Rating *string
  }
)
