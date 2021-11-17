package query

import (
	"errors"
)

var (
	ErrCreateQuantity = errors.New("create statement must have 2 words")
	ErrCreateQuotes = errors.New("create statement must have no quotes")

	ErrInsertQuantity = errors.New("insert statement must have 3 words")
	ErrInsertQuotes = errors.New("create statement must have one quotes pair")

	ErrSearchQuantity = errors.New("search statement must have 2 or 4 or 6 words")
	ErrSearchNoWhere = errors.New("there are must be where statement")
	ErrWhereQuantity = errors.New("search query must have 1 or 2 search words in quotes")

	ErrPrefixSyntax = errors.New("prefix statement must have only one '*' symbol")
	ErrSearchParams = errors.New("bad search parameter")
	ErrIntervalParams = errors.New("bad interval parameter")
	ErrIntervalParsing = errors.New("can not parse interval")

	ErrPrintQuantity = errors.New("print_index statement must have 2 words")
	ErrPrintQuotes = errors.New("print_index statement must have no quotes")

	ErrUnknown = errors.New("unknown statement")
)