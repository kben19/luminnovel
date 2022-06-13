package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDBManager interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}

type MongoDBCursorProvider interface {
	Close(ctx context.Context) error
	Decode(val interface{}) error
	Next(ctx context.Context) bool
}

type repository struct {
	database mongoDBManager
}

func New(db mongoDBManager) *repository {
	return &repository{db}
}

func (repo *repository) Find(ctx context.Context, collectionName string, filter map[string]interface{}, findOptions ...FindOptions) (MongoDBCursorProvider, error) {
	filterDB := bson.D{}
	for key, value := range filter {
		filterDB = append(filterDB, bson.E{Key: key, Value: value})
	}

	optionsDB := make([]*options.FindOptions, len(findOptions))
	for index, option := range findOptions {
		optionsDB[index] = convertFindOptions(option)
	}

	cursor, err := repo.database.Collection(collectionName).Find(ctx, filterDB, optionsDB...)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

func convertFindOptions(option FindOptions) *options.FindOptions {
	return &options.FindOptions{
		AllowDiskUse:        option.AllowDiskUse,
		AllowPartialResults: option.AllowPartialResults,
		BatchSize:           option.BatchSize,
		Comment:             option.Comment,
		Hint:                option.Hint,
		Limit:               option.Limit,
		Max:                 option.Max,
		MaxAwaitTime:        option.MaxAwaitTime,
		MaxTime:             option.MaxTime,
		Min:                 option.Min,
		NoCursorTimeout:     option.NoCursorTimeout,
		Projection:          option.Projection,
		ReturnKey:           option.ReturnKey,
		ShowRecordID:        option.ShowRecordID,
		Skip:                option.Skip,
		Sort:                option.Sort,
		Let:                 option.Let,
	}
}
