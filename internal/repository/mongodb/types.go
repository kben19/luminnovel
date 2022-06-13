package mongodb

import "time"

type FindOptions struct {
	AllowDiskUse        *bool
	AllowPartialResults *bool
	BatchSize           *int32
	Comment             *string
	Hint                interface{}
	Limit               *int64
	Max                 interface{}
	MaxAwaitTime        *time.Duration
	MaxTime             *time.Duration
	Min                 interface{}
	NoCursorTimeout     *bool
	OplogReplay         *bool
	Projection          interface{}
	ReturnKey           *bool
	ShowRecordID        *bool
	Skip                *int64
	Snapshot            *bool
	Sort                interface{}
	Let                 interface{}
}
