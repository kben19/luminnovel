package bookdepository

import (
	"context"
	"log"
	"regexp"
	"strconv"
	"strings"

	"luminnovel/internal/entity"
)

func (rsc *resource) GetRequestFromHTTP(ctx context.Context, params entity.Source) (entity.ProductBookDepository, error) {
	urlParam := params.Domain + params.Path
	paramMap := map[string]string{}
	for key, value := range params.Params {
		paramMap[key] = value
	}
	resp, err := rsc.httpRepo.Get(ctx, urlParam, paramMap, nil, nil)
	if err != nil {
		return entity.ProductBookDepository{}, err
	}

	return parseHTMLBookDepository(string(resp), params.Path), nil
}

func parseHTMLBookDepository(resp string, path string) entity.ProductBookDepository {
	var title string
	paths := strings.Split(path, "/")
	if len(paths) > 0 {
		title = paths[1]
	}

	priceRE := regexp.MustCompile(`dimension3','.*'`)
	stockRE := regexp.MustCompile(`Currently unavailable`)

	priceRaw := priceRE.FindString(resp)
	stockRaw := stockRE.FindString(resp)

	priceStr := strings.TrimSuffix(strings.TrimPrefix(priceRaw, "dimension3','"), "'")
	priceFloat, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		log.Println(err)
	}

	return entity.ProductBookDepository{
		Title:   title,
		Price:   priceFloat,
		InStock: stockRaw == "",
	}
}
