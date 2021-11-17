package engine

import (
	"fmt"
	"labdb/internal/core/domain/query"
	"labdb/internal/core/domain/contentprocessing"
)

func (db *database) Print(q query.Print) (str string, err error) {
	if !db.collectionsRegistry[q.CollectionName] {
		return str, ErrCollectionNotExists
	}

	collection := db.collections[q.CollectionName]
	index := collection.reversedIndex

	for iterator := index.Iterator(); iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		insideMap := iterator.Value()

		str += contentprocessing.ShiftAndNewLineString(0, fmt.Sprintf("\"%v\":", key))
		for document, positions := range insideMap {
			str += contentprocessing.ShiftAndNewLineString(1, fmt.Sprintf("document #%v -> %v", document, positions))
		}
	}

	return str, nil
}
