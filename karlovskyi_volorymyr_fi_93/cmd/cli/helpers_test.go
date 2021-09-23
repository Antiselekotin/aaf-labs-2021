package main

import "testing"

func TestSearchLastQuoteIndex(t *testing.T) {
	testMap := map[int][]byte{
		11: []byte("hello world\""),
		0:  []byte("\"Hi, Mr. Jacobs"),
		-1: []byte("There are no quote"),
	}
	for expect, str := range testMap {
		i := searchLastQuoteIndex(&str)
		if i != expect {
			t.Errorf("index must be 10 but it %v", i)
		}
	}
}

func TestLineEnded(t *testing.T) {
	type answer struct {
		IsEnded bool
		Index   int
	}
	testMap := map[string]answer{
		"hello world\";\"": answer{
			false, -1,
		},
		"create ; something": answer{
			true, 7,
		},
	}

	for str, ans := range testMap {
		b := []byte(str)
		isEnded, index := lineEnded(&b)
		if isEnded != ans.IsEnded || index != ans.Index {
			t.Errorf("lineEnded(%v) != (%v, %v)", str, ans.IsEnded, ans.Index)
		}
	}
}
