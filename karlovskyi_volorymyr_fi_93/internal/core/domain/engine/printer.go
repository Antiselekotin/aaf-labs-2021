package engine

import (
	"fmt"
	"labdb/internal/core/domain/query"
	"labdb/internal/core/domain/textprocessing"
)

func (db *database) Print(q query.Print) (str string, err error) {
	if !db.collectionsRegistry[q.CollectionName] {
		return str, fmt.Errorf("Collection %v does not already exists", q.CollectionName)
	}

	collection := db.collections[q.CollectionName]
	index := collection.reversedIndex

	for iterator := index.Iterator(); iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		insideMap := iterator.Value()

		str += textprocessing.ShiftAndNewLineString(0, fmt.Sprintf("\"%v\":", key))
		for document, positions := range insideMap {
			str += textprocessing.ShiftAndNewLineString(1, fmt.Sprintf("document #%v -> %v", document, positions))
		}
	}

	return str, nil
}
