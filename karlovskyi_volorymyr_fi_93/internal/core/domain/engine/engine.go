package engine

import (
	"fmt"
	"labdb/internal/core/domain/query"
	"labdb/internal/core/domain/textprocessing"
)

//go:generate gotemplate "github.com/igrmk/treemap" "stringIntMapOfIntSliceTreeMap(string, map[int][]int)"

type Database interface {
	Create(q query.Create) (string, error)
	Insert(q query.Insert) (string, error)
	Print(q query.Print) (string, error)
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
	content       []string
	reversedIndex stringIntMapOfIntSliceTreeMap
}

func stringIntMapOfIntSliceTreeMapLess(a, b string) bool { return a < b }

func (db *database) Create(q query.Create) (success string, err error) {
	if db.collectionsRegistry[q.Name] {
		return "", fmt.Errorf("Collection %v already exists", q.Name)
	}
	reversedIndex := newStringIntMapOfIntSliceTreeMap(stringIntMapOfIntSliceTreeMapLess)
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
	contentNoPunc := textprocessing.RemovePunctuation(originalContent)
	content := textprocessing.RemoveIndent(contentNoPunc)
	insertIndex := len(collection.content)
	splitMap := textprocessing.SplitStringWithPositions(content)

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
