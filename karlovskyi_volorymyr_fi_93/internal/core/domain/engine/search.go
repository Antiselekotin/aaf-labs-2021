package engine

import (
	"fmt"
	"labdb/internal/core/domain/query"
)

func (db *database) Search(q query.Search) ([][]byte, error) {
	if !db.collectionsRegistry[q.CollectionName] {
		return nil, fmt.Errorf("Collection %v does not already exists", q.CollectionName)
	}
	collection := db.collections[q.CollectionName]

	ids := q.Where.Filter(collection.reversedIndex)
	content := make([][]byte, len(ids))

	for i, id := range ids {
		content[i] = collection.content[id]
	}

	return content, nil
}
