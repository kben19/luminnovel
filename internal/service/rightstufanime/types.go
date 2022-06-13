package rightstufanime

import "luminnovel/internal/entity"

type HTTPResponse struct {
	Items []ItemsResponse `json:"items"`
}

type ItemsResponse struct {
	Title    string        `json:"storedisplayname2"`
	Price    PriceResponse `json:"onlinecustomerprice_detail"`
	InStock  bool          `json:"isinstock"`
	PreOrder string        `json:"custitem_rs_new_releases_preorders"`
}

type PriceResponse struct {
	Price    float64 `json:"onlinecustomerprice"`
	PriceFmt string  `json:"onlinecustomerprice_formatted"`
}

type Product struct {
	entity.ProductRightStufAnime
	Volume string
}
