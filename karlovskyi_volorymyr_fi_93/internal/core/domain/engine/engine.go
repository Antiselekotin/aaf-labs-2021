package engine

import (
	"bytes"
	"fmt"
	"labdb/internal/core/domain/contentprocessing"
	tree "labdb/internal/core/domain/invertedtree"
	"labdb/internal/core/domain/query"
)

type Database interface {
	Create(q query.Create) (string, error)
	Insert(q query.Insert) (string, error)
	Print(q query.Print) (string, error)
	Search(q query.Search) ([]string, error)
}

type database struct {
	collections         map[string]*Collection
	collectionsRegistry map[string]bool
}

func New() *database {
	return &database{
		collections:         make(map[string]*Collection),
		collectionsRegistry: make(map[string]bool),
	}
}

type Collection struct {
	content       [][]byte
	reversedIndex tree.StringIntMapOfIntSliceTreeMap
}

func stringIntMapOfIntSliceTreeMapLess(a, b string) bool { return a < b }

func (db *database) Create(q query.Create) (success string, err error) {
	if db.collectionsRegistry[q.Name] {
		return "", fmt.Errorf("Collection %v already exists", q.Name)
	}
	reversedIndex := tree.NewStringIntMapOfIntSliceTreeMap(stringIntMapOfIntSliceTreeMapLess)
	db.collections[q.Name] = &Collection{reversedIndex: *reversedIndex}
	db.collectionsRegistry[q.Name] = true

	return fmt.Sprintf("Collection %v has been successfully created", q.Name), nil
}

func (db *database) Insert(q query.Insert) (success string, err error) {
	if !db.collectionsRegistry[q.CollectionName] {
		return "", fmt.Errorf("Collection %v does not already exists", q.CollectionName)
	}
	collection := db.collections[q.CollectionName]
	originalContent := q.Content
	contentNoPunc := contentprocessing.RemovePunctuationBytes(originalContent)
	content := bytes.ToLower(contentNoPunc)
	insertIndex := len(collection.content)
	splitMap := contentprocessing.SplitBytesWithPositions(content)

	for word, positions := range splitMap {
		oldMap, ok := collection.reversedIndex.Get(word)
		if !ok {
			oldMap = make(map[int][]int)
		}
		oldMap[insertIndex] = positions
		collection.reversedIndex.Set(word, oldMap)
	}
	collection.content = append(collection.content, originalContent)
	return "Content has been added", nil
}
