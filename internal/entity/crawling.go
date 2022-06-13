package entity

type CrawlingItem struct {
	Title string
	Price CrawlingPrice
	Stock CrawlingStock
	Weight float64
	ActualWeight float64
}

type CrawlingPrice struct {
	Amazon string
	RightStufAnime string
	InStockTrades string
	BookDepository string
}

type CrawlingStock struct {
	Amazon bool
	RightStufAnime bool
	InStockTrades bool
	BookDepository bool
}
