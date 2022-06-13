package rightstufanime

import (
	"context"

	"luminnovel/internal/entity"
)

const (
	rightStufAnimeSite = "RightStufAnime"
)

func (rsc *resource) FindSeriesByTitleFromDB(ctx context.Context, collectionName string, title string) ([]entity.Source, error) {
	cursor, err := rsc.dbRepo.Find(ctx, collectionName, map[string]interface{}{
		"title": title,
		"site":  rightStufAnimeSite,
	})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = cursor.Close(ctx)
	}()

	var sources []entity.Source
	for cursor.Next(ctx) {
		var source entity.Source
		err = cursor.Decode(&source)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}
	return sources, nil
}
