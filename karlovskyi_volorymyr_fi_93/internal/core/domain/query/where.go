package query

import (
	tree "labdb/internal/core/domain/invertedtree"
)

// Visitor

type Where interface {
	Filter(tree.StringIntMapOfIntSliceTreeMap) []int
}

type WhereNone struct {
}

func (a *WhereNone) Filter(t tree.StringIntMapOfIntSliceTreeMap) []int {
	idsSet := map[int]struct{}{}
	for iterator := t.Iterator(); iterator.Valid(); iterator.Next() {
		for id := range iterator.Value() {
			idsSet[id] = struct{}{}
		}
	}
	ids := make([]int, len(idsSet))
	i := 0
	for id := range idsSet {
		ids[i] = id
		i++
	}
	return ids
}

type WhereWord struct {
	Word string
}

func (a *WhereWord) Filter(t tree.StringIntMapOfIntSliceTreeMap) []int {
	m, ok := t.Get(a.Word)
	if !ok {
		return []int{}
	}
	ids := make([]int, 0, len(m))
	for key := range m {
		ids = append(ids, key)
	}
	return ids
}

type WherePrefix struct {
	Prefix string
}

func (a *WherePrefix) Filter(t tree.StringIntMapOfIntSliceTreeMap) []int {
	prefix := a.Prefix
	return t.SearchByPrefix(prefix)
}

type WhereInterval struct {
	FirstWord, LastWord string
	Interval            int
}

func (a *WhereInterval) Filter(t tree.StringIntMapOfIntSliceTreeMap) []int {
	f, ok1 := t.Get(a.FirstWord)
	l, ok2 := t.Get(a.LastWord)
	ids := []int{}
	if !ok1 || !ok2 {
		return ids
	}
	var smaller, bigger []int = nil, nil

	for docFId, positionsF := range f {
		positionsL := l[docFId]
		if len(positionsF) > len(positionsL) {
			smaller = positionsL
			bigger = positionsF
		} else {
			smaller = positionsF
			bigger = positionsL
		}
		smallMap := make(map[int]bool, len(smaller))
		for _, pos := range smaller {
			smallMap[pos] = true
		}
		for _, b := range bigger {
			if smallMap[b+a.Interval] || smallMap[b-a.Interval] {
				ids = append(ids, docFId)
				break
			}
		}
	}
	return ids
}
