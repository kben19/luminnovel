package crawling

import "luminnovel/internal/entity"

type CrawlingConfig struct {
	Series entity.ProductTitle
	Source CrawlingSource
}

type CrawlingSeries struct {
	Bofuri              bool
	DeathMarch          bool
	DevilPartTimer      bool
	EightySix           bool
	ImSpiderSoWhat      bool
	KillingSlime        bool
	Konosuba            bool
	ReZero              bool
	ReincarnatedAsSlime bool
	Smartphone          bool
}

type CrawlingSource struct {
	Amazon         bool
	RightStufAnime bool
	InStockTrades  bool
	BookDepository bool
}

type CrawlingPayload struct {
	Volume   int
	Price    string
	InStock  bool
	Weight   float64
	PreOrder bool
}
