package query

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile(`[^\w\s'.";<>]+`)
var whitespacePattern = regexp.MustCompile(`\s+`)

func Parse(str string) (Query, error) {
	str = pattern.ReplaceAllString(str, "")
	str, memMap := replaceRawStrings(str)
	str = whitespacePattern.ReplaceAllString(str, " ")
	str = strings.TrimLeft(str, " ")
	str = strings.TrimRight(str, " ")
	if strings.HasPrefix(strings.ToLower(str), "create ") {
		return parseCreateQuery(str, memMap)
	}
	if strings.HasPrefix(strings.ToLower(str), "insert ") {
		return parseInsertQuery(str, memMap)
	}
	if strings.HasPrefix(strings.ToLower(str), "search ") {
		return parseSearchQuery(str, memMap)
	}

	return nil, fmt.Errorf("unknown statement")
}

func replaceRawStrings(str string) (res string, memMap map[string]string) {
	memMap = make(map[string]string)
	buf := []byte(str)
	counter, mem := 1, -1
	for i := 0; i < len(buf); i++ {
		if buf[i] == '"' {
			if mem == -1 {
				mem = i
			} else {
				memStr := fmt.Sprintf("$%v", counter)
				counter++
				res += string(buf[:mem]) + memStr
				memMap[memStr] = string(buf[mem+1 : i])
				buf = buf[i:]
				i = 0
			}
		}
	}
	if len(res) == 0 {
		res = string(buf)
	}
	return
}

func parseCreateQuery(str string, memMap map[string]string) (Create, error) {
	split := strings.Split(str, " ")
	if len(split) != 2 {
		return Create{}, fmt.Errorf("create statement must have 2 words")
	}
	if len(memMap) != 0 {
		return Create{}, fmt.Errorf("create statement must have no quotes")
	}
	return Create{
		Name: split[1],
	}, nil
}

func parseInsertQuery(str string, memMap map[string]string) (Insert, error) {
	split := strings.Split(str, " ")
	if len(split) != 3 {
		return Insert{}, fmt.Errorf("insert statement must have 3 words")
	}
	if len(memMap) != 1 {
		return Insert{}, fmt.Errorf("create statement must have one quotes pair")
	}
	return Insert{
		CollectionName: split[1],
		Content:        memMap[split[2]],
	}, nil
}

func parseSearchQuery(str string, memMap map[string]string) (Search, error) {
	split := strings.Split(str, " ")
	if len(split) != 2 && len(split) != 4 && len(split) != 6 {
		return Search{}, fmt.Errorf("search statement must have 2 or 4 or 6 words")
	}

	if len(split) == 2 {
		return Search{
			CollectionName: split[1],
			Where:          WhereNone{},
		}, nil
	}

	if strings.ToLower(split[2]) != "where" {
		return Search{}, fmt.Errorf("There are must be where statement")
	}

	if len(memMap) != 1 && len(split) != 2 {
		return Search{}, fmt.Errorf("search query must have 1 or 2 search words in quotes")
	}

	search := Search{
		CollectionName: split[1],
	}
	mapIndex := split[3]

	if strings.HasSuffix(split[3], "*") && len(split) == 4 {
		if strings.HasSuffix(split[3], "**") {
			return Search{}, fmt.Errorf("prefix statement must have only one '*' symbol")
		}

		if len(mapIndex) != 0 {
			mapIndex = mapIndex[:len(mapIndex)-2]
		}
		if memMap[mapIndex] == "" {
			return search, fmt.Errorf("bad search parameter")
		}
		search.Where = WherePrefix{
			Prefix: memMap[mapIndex],
		}
		return search, nil
	}

	if memMap[mapIndex] == "" {
		return search, fmt.Errorf("bad search parameter")
	}

	if len(split) == 4 {
		search.Where = WhereWord{
			Word: memMap[mapIndex],
		}
		return search, nil
	}

	if len(split[4]) < 3 {
		return search, fmt.Errorf("bad internal parameter")
	}
	intervalStr := split[4]
	internal, err := strconv.ParseInt(intervalStr[1:len(intervalStr)-1], 10, 64)

	if err != nil {
		return search, fmt.Errorf("can not parse interval")
	}
	if memMap[split[len(split)-1]] == "" {
		return search, fmt.Errorf("bad search parameter")
	}

	where := WhereInterval{
		FirstWord: memMap[mapIndex],
		LastWord:  memMap[split[len(split)-1]],
		Interval:  int(internal),
	}

	search.Where = where

	return search, nil
}
