package search

import (
	"fmt"
	"labdb/internal/core/domain/engine"
	"labdb/internal/core/domain/query"
)

var hiddenDB = engine.New()

type service struct {
	db engine.Database
	responseAdapter ResponseAdapter
}

type Service interface {
	Execute(string)
}

type ResponseAdapter interface {
	OnError(error)
	OnCreateSuccess(string)
	OnCreateFailure(error)
}

func NewSearch(r ResponseAdapter) Service {
	return &service{
		db: hiddenDB,
		responseAdapter: r,
	}
}

func (s *service) Execute(str string) {
	q, err := query.Parse(str)
	if err != nil {
		s.responseAdapter.OnError(err)
		return
	}

	switch typed := q.(type) {
	case query.Create:
		str, err := s.db.Create(typed)
		if err != nil {
			s.responseAdapter.OnCreateFailure(err)
			break
		}
		s.responseAdapter.OnCreateSuccess(str)
	case query.Insert:
		fmt.Printf("%[1]T: %[1]v", typed)
	case query.Search:
		fmt.Printf("%[1]T: %[1]v", typed)
	case query.Print:
		fmt.Printf("%[1]T: %[1]v", typed)
	default:
		panic("unknown query")
	}
}


