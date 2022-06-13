package rightstufanime

import (
	"context"
)

const (
	collectionSource = "source"
)

func (svc *service) FetchProductDataByTitle(ctx context.Context, title string) ([]Product, error) {
	sources, err := svc.resource.FindSeriesByTitleFromDB(ctx, collectionSource, title)
	if err != nil {
		return nil, err
	}

	products := make([]Product, len(sources))
	for index, source := range sources {
		product, err := svc.resource.GetRequestFromHTTP(ctx, source)
		if err != nil {
			return nil, err
		}
		products[index] = Product{
			ProductRightStufAnime: product,
			Volume:                source.Volume,
		}
	}
	return products, nil
}
