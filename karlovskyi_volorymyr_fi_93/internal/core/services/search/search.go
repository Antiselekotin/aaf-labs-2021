package search

import (
	"fmt"
	"labdb/internal/core/domain/query"
)

type service struct {
}

type Service interface {
	Execute(string) (string, error)
}

func (s *service) Execute(str string) (string, error) {
	q, err := query.Parse(str)
	if err != nil {
		return "", err
	}

	switch typed := q.(type) {
	case query.Create:
		fmt.Printf("%[1]T: %[1]v", typed)
	case query.Insert:
		fmt.Printf("%[1]T: %[1]v", typed)
	case query.Search:
		fmt.Printf("%[1]T: %[1]v", typed)
	default:
		panic("unknown query")
	}
	return str, nil
}

func NewSearch() Service {
	return &service{}
}
