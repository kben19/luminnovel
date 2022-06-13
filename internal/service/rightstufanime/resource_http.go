package rightstufanime

import (
	"context"
	"encoding/json"
	"luminnovel/internal/entity"
)

func (rsc *resource) GetRequestFromHTTP(ctx context.Context, params entity.Source) (entity.ProductRightStufAnime, error) {
	urlParam := params.Domain + params.Path
	paramMap := map[string]string{}
	for key, value := range params.Params {
		paramMap[key] = value
	}
	resp, err := rsc.httpRepo.Get(ctx, urlParam, paramMap, nil, nil)
	if err != nil {
		return entity.ProductRightStufAnime{}, err
	}

	var productResp HTTPResponse
	err = json.Unmarshal(resp, &productResp)
	if err != nil {
		return entity.ProductRightStufAnime{}, err
	}

	if len(productResp.Items) == 0 {
		return entity.ProductRightStufAnime{}, nil
	}
	return entity.ProductRightStufAnime{
		Title:    productResp.Items[0].Title,
		Price:    productResp.Items[0].Price.Price,
		PriceFmt: productResp.Items[0].Price.PriceFmt,
		InStock:  productResp.Items[0].InStock,
		PreOrder: productResp.Items[0].PreOrder != "",
	}, nil
}
