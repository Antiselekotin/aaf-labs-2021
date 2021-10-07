package engine

import (
	"fmt"
	"labdb/internal/core/domain/query"
)

type Database interface {
	Create(q query.Create) (string, error)
}

type database struct {
	schemas         map[string]schema
	schemasRegistry map[string]bool
}

func New() *database {
	return &database{
		schemas:         make(map[string]schema),
		schemasRegistry: make(map[string]bool),
	}
}

type schema struct {
	content []string
}

func (db *database) Create(q query.Create) (success string, err error) {
	if db.schemasRegistry[q.Name] {
		return "", fmt.Errorf("schema %v already exists", q.Name)
	}

	db.schemas[q.Name] = schema{}
	db.schemasRegistry[q.Name] = true

	return fmt.Sprintf("schema %v has been created", q.Name), nil
}
