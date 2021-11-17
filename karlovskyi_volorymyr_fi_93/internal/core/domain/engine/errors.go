package engine

import "errors"

var (
	ErrCollectionNotExists = errors.New("collection is not exists")
	ErrCollectionAlreadyExists = errors.New("collection already exists")
)
