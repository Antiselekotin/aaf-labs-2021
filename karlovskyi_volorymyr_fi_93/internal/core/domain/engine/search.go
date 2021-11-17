package engine

import (
	"labdb/internal/core/domain/query"
)

func (db *database) Search(q query.Search) ([]string, error) {
	if !db.collectionsRegistry[q.CollectionName] {
		return nil, ErrCollectionNotExists
	}
	collection := db.collections[q.CollectionName]

	ids := q.Where.Filter(collection.reversedIndex)
	content := make([]string, len(ids))

	for i, id := range ids {
		content[i] = collection.content[id]
	}

	return content, nil
}
