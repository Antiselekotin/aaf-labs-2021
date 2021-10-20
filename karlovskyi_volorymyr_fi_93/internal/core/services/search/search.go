package search

import (
	"labdb/internal/core/domain/engine"
	"labdb/internal/core/domain/query"
)

var hiddenDB = engine.New()

type service struct {
	db              engine.Database
	responseAdapter ResponseAdapter
}

type Service interface {
	Execute(string)
}

type ResponseAdapter interface {
	OnError(error)
	OnSuccess(string)
	OnCreateSuccess(string)
	OnCreateFailure(error)
	OnInsertSuccess(string)
	OnInsertFailure(error)
	OnPrintSuccess(string)
	OnPrintFailure(error)
	OnSearchSuccess([]string)
	OnSearchFailure(error)
}

func NewSearch(r ResponseAdapter) Service {
	return &service{
		db:              hiddenDB,
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
		str, err := s.db.Insert(typed)
		if err != nil {
			s.responseAdapter.OnInsertFailure(err)
			break
		}
		s.responseAdapter.OnInsertSuccess(str)
	case query.Search:
		strs, err := s.db.Search(typed)
		if err != nil {
			s.responseAdapter.OnSearchFailure(err)
		}
		s.responseAdapter.OnSearchSuccess(strs)
	case query.Print:
		str, err := s.db.Print(typed)
		if err != nil {
			s.responseAdapter.OnPrintFailure(err)
			break
		}
		s.responseAdapter.OnPrintSuccess(str)
	default:
		panic("unknown query")
	}
}
