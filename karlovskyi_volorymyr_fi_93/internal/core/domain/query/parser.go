package query

import (
	"strconv"
	"strings"

	"labdb/internal/core/domain/contentprocessing"
)

func Parse(str string) (Query, error) {
	if str[len(str)-1] == ';' {
		str = str[:len(str)-1]
	}
	str, memMap := replaceRawString(str)
	str = contentprocessing.RemoveIndent(str)
	str = contentprocessing.Trim(str)
	if strings.HasPrefix(strings.ToLower(str), "create ") {
		return parseCreateQuery(str, memMap)
	}
	if strings.HasPrefix(strings.ToLower(str), "insert ") {
		return parseInsertQuery(str, memMap)
	}
	if strings.HasPrefix(strings.ToLower(str), "search ") {
		return parseSearchQuery(str, memMap)
	}
	if strings.HasPrefix(strings.ToLower(str), "print_index ") {
		return parsePrintQuery(str, memMap)
	}

	return nil, ErrUnknown
}

func parseCreateQuery(str string, memMap map[string]string) (Create, error) {
	split := strings.Split(str, " ")
	if len(split) != 2 {
		return Create{}, ErrCreateQuantity
	}
	if len(memMap) != 0 {
		return Create{}, ErrCreateQuotes
	}
	return Create{
		Name: split[1],
	}, nil
}

func parseInsertQuery(str string, memMap map[string]string) (Insert, error) {
	split := strings.Split(str, " ")
	if len(split) != 3 {
		return Insert{}, ErrInsertQuantity
	}
	if len(memMap) != 1 {
		return Insert{}, ErrInsertQuotes
	}
	return Insert{
		CollectionName: split[1],
		Content:        memMap[split[2]],
	}, nil
}

func parseSearchQuery(str string, memMap map[string]string) (Search, error) {
	split := strings.Split(str, " ")
	if len(split) != 2 && len(split) != 4 && len(split) != 6 {
		return Search{}, ErrSearchQuantity
	}

	if len(split) == 2 {
		return Search{
			CollectionName: split[1],
			Where:          &WhereNone{},
		}, nil
	}

	if strings.ToLower(split[2]) != "where" {
		return Search{}, ErrSearchNoWhere
	}

	if len(memMap) != 1 && len(memMap) != 2 {
		return Search{}, ErrWhereQuantity
	}

	search := Search{
		CollectionName: split[1],
	}
	mapIndex := split[3]

	if strings.HasSuffix(split[3], "*") && len(split) == 4 {
		if strings.HasSuffix(split[3], "**") {
			return Search{}, ErrPrefixSyntax
		}
		if len(mapIndex) != 0 {
			mapIndex = mapIndex[:len(mapIndex)-1]
		}
		if memMap[mapIndex] == "" {
			return search, ErrSearchParams
		}
		search.Where = &WherePrefix{
			Prefix: memMap[mapIndex],
		}
		return search, nil
	}

	if memMap[mapIndex] == "" {
		return search, ErrSearchParams
	}

	if len(split) == 4 {
		search.Where = &WhereWord{
			Word: memMap[mapIndex],
		}
		return search, nil
	}

	if len(split[4]) < 3 {
		return search, ErrIntervalParams
	}
	intervalStr := split[4]
	internal, err := strconv.ParseInt(intervalStr[1:len(intervalStr)-1], 10, 64)

	if err != nil {
		return search, ErrIntervalParsing
	}
	if memMap[split[len(split)-1]] == "" {
		return search, ErrSearchParams
	}

	where := WhereInterval{
		FirstWord: memMap[mapIndex],
		LastWord:  memMap[split[len(split)-1]],
		Interval:  int(internal),
	}

	search.Where = &where

	return search, nil
}

func parsePrintQuery(str string, memMap map[string]string) (Print, error) {
	split := strings.Split(str, " ")
	if len(split) != 2 {
		return Print{}, ErrPrintQuantity
	}
	if len(memMap) != 0 {
		return Print{}, ErrPrintQuotes
	}
	return Print{
		CollectionName: split[1],
	}, nil
}
