package rightstufanime

import (
	"context"
	"luminnovel/internal/repository/mongodb"
)

type httpRepoProvider interface {
	Get(ctx context.Context, urlParam string, queryParam map[string]string, header map[string]string, body []byte) ([]byte, error)
}

type dbRepoProvider interface {
	Find(ctx context.Context, collectionName string, filter map[string]interface{}, findOptions ...mongodb.FindOptions) (mongodb.MongoDBCursorProvider, error)
}

type resource struct {
	httpRepo httpRepoProvider
	dbRepo   dbRepoProvider
}

func NewResource(repo httpRepoProvider, db dbRepoProvider) *resource {
	return &resource{repo, db}
}
