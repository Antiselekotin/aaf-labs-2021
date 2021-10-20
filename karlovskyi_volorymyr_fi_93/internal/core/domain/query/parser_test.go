package query

import "testing"

func TestParseOnCorrectInsertQueries(t *testing.T) {
	queries := map[string]Insert{
		"Insert schema \"something interesting\";": Insert{
			CollectionName: "schema",
			Content:        "something interesting",
		},
		"inSeRt sCheMa\n\"something INTeresting\";": Insert{
			CollectionName: "sCheMa",
			Content:        "something INTeresting",
		},
		"inSeRt mmasfgwe_wegeg \n\n\n \"INTeresting\"\n\n\n;": Insert{
			CollectionName: "mmasfgwe_wegeg",
			Content:        "INTeresting",
		},
	}

	for query, expect := range queries {
		result, err := Parse(query)
		if err != nil {
			t.Errorf("Parse func return error: '%v' on input '%v' but it must not", err, query)
		}
		insert, ok := result.(Insert)

		if !ok {
			t.Errorf("Parse func return not Insert type on input '%v'", query)
		}
		if insert.CollectionName != expect.CollectionName {
			t.Errorf("Insert.CollectionName = '%v', but expect '%v'", insert.CollectionName, expect.CollectionName)
		}
		if insert.Content != expect.Content {
			t.Errorf("insert.Content = '%v', but expect '%v'", insert.Content, expect.Content)
		}
	}
}

func TestParseOnIncorrectInsertQueries(t *testing.T) {
	queries := []string{
		"insert schema;",
		"insert schema \"something;",
		"insert schema \" something\" ooo;",
	}
	for _, query := range queries {
		_, err := Parse(query)
		if err == nil {
			t.Errorf("Parse func accept incorrect input '%v'", query)
		}
	}
}

func TestParseOnCorrectCreateQueries(t *testing.T) {
	queries := map[string]Create{
		"create\nt;": Create{
			Name: "t",
		},
		"create\nschema;": Create{
			Name: "schema",
		},
		"cReAte\n\nsCheMa\n\n;": Create{
			Name: "sCheMa",
		},
		"creATE mmasfgwe_wegeg;": Create{
			Name: "mmasfgwe_wegeg",
		},
	}

	for query, expect := range queries {
		result, err := Parse(query)
		if err != nil {
			t.Errorf("Parse func return error: '%v' on input '%v' but it must not", err, query)
		}
		create, ok := result.(Create)

		if !ok {
			t.Errorf("Parse func return not Create type on input '%v'", query)
		}
		if create.Name != expect.Name {
			t.Errorf("create.Name = '%v', but expect '%v'", create.Name, expect.Name)
		}

	}
}

func TestParseOnIncorrectCreateQueries(t *testing.T) {
	queries := []string{
		"create \"schema\";",
		"create schema lala;",
	}

	for _, query := range queries {
		_, err := Parse(query)
		if err == nil {
			t.Errorf("Parse func accept incorrect input '%v'", query)
		}
	}
}

func TestParseOnCorrectSearchQueries(t *testing.T) {
	queries := map[string]Search{
		"search \n\nschema\n\n;":           Search{CollectionName: "schema", Where: &WhereNone{}},
		"SearCH scHema Where \"word\";":    Search{CollectionName: "scHema", Where: &WhereWord{"word"}},
		"search schema where \"prefix\"*;": Search{CollectionName: "schema", Where: &WherePrefix{"prefix"}},
		"search schema where \"first\" <12> \"last\"": Search{CollectionName: "schema", Where: &WhereInterval{
			FirstWord: "first", LastWord: "last", Interval: 12,
		}},
	}
	for query, expect := range queries {
		result, err := Parse(query)
		if err != nil {
			t.Errorf("Parse func return error: '%v' on input '%v' but it must not", err, query)
		}
		search, ok := result.(Search)

		if !ok {
			t.Errorf("Parse func return not Create type on input '%v'", query)
		}
		if search.CollectionName != expect.CollectionName {
			t.Errorf("search.CollectionName = '%v', but expect '%v'", search.CollectionName, expect.CollectionName)
		}
		switch where := search.Where.(type) {
		case *WhereNone:
			_, ok = expect.Where.(*WhereNone)
			if !ok {
				t.Errorf("Where type is not WhereNone on input '%v'", query)
			}
		case *WhereWord:
			w, ok := expect.Where.(*WhereWord)
			if !ok {
				t.Errorf("Where type is not WhereWord on input '%v'", query)
				break
			}
			if w.Word != where.Word {
				t.Errorf("w.Word is %v on input '%v'", w.Word, query)
			}
		case *WherePrefix:
			w, ok := expect.Where.(*WherePrefix)
			if !ok {
				t.Errorf("Where type is not WherePrefix on input '%v'", query)
				break
			}
			if w.Prefix != where.Prefix {
				t.Errorf("w.Prefix is %v on input '%v'", w.Prefix, query)
			}
		case *WhereInterval:
			w, ok := expect.Where.(*WhereInterval)
			if !ok {
				t.Errorf("Where type is not WherePrefix on input '%v'", query)
				break
			}
			if w.FirstWord != where.FirstWord {
				t.Errorf("w.FirstWord is %v on input '%v'", w.FirstWord, query)
			}
			if w.LastWord != where.LastWord {
				t.Errorf("w.LastWord is %v on input '%v'", w.LastWord, query)
			}
			if w.Interval != where.Interval {
				t.Errorf("w.Interval is %v on input '%v'", w.Interval, query)
			}
		}
	}
}

func TestParseOnIncorrectSearchQueries(t *testing.T) {
	queries := []string{
		"search \n\nschema\n\n\";",
		"SearCH scHema Where \"word\" df;",
		"search schema where \"prefix\" *;",
		"search schema where \"first\" <1gg2> \"last\"",
	}
	for _, query := range queries {
		_, err := Parse(query)
		if err == nil {
			t.Errorf("Parse func accept incorrect input '%v'", query)
		}
	}
}

func TestParseOnCorrectPrintQueries(t *testing.T) {
	queries := map[string]Print{
		"pRint_index schema;": Print{CollectionName: "schema",},
		"print_index\nsCC": Print{CollectionName: "sCC",},
	}

	for query, expect := range queries {
		result, err := Parse(query)
		if err != nil {
			t.Errorf("Parse func return error: '%v' on input '%v' but it must not", err, query)
		}
		print, ok := result.(Print)
		if !ok {
			t.Error("Parse func return not Print but it must")
		}
		if print.CollectionName != expect.CollectionName {
			t.Errorf("print.CollectionName = '%v', but expect '%v'", print.CollectionName, expect.CollectionName)
		}
	}
}

func TestParseOnIncorrectPrintQueries(t *testing.T) {
	queries := []string{
		"print index;",
		"print_index word word;",
		"print_index \"schema\";",
	}
	for _, query := range queries {
		_, err := Parse(query)
		if err == nil {
			t.Errorf("Parse func accept incorrect input '%v'", query)
		}
	}
}
