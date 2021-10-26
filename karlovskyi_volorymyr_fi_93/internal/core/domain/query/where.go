package query

import (
	tree "labdb/internal/core/domain/invertedtree"
	"strings"
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
	ids := []int{}
	m, ok := t.Get(a.Word)
	if !ok {
		return ids
	}
	for key, _ := range m {
		ids = append(ids, key)
	}
	return ids
}

type WherePrefix struct {
	Prefix string
}

func (a *WherePrefix) Filter(t tree.StringIntMapOfIntSliceTreeMap) []int {
	ids := []int{}
	prefix := a.Prefix

	root := t.RootNode().Left()
	for root != nil && !strings.HasPrefix(root.Key(), prefix) {
		if t.Less(root.Key(), prefix) {
			root = root.Right()
		} else {
			root = root.Left()
		}
	}

	if root == nil {
		return ids
	}
	nodes := tree.ItterationSlice{root}
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		if strings.HasPrefix(node.Key(), prefix) {
			for id := range node.Value() {
				ids = append(ids, id)
			}
			if l := node.Left(); l != nil {
				nodes = append(nodes, l)
			}

			if r := node.Right(); r != nil {
				nodes = append(nodes, r)
			}
		}
		nodes = nodes[1:]
		i = 0
	}
	return ids
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
			if smallMap[b+a.Interval] == true || smallMap[b-a.Interval] == true {
				ids = append(ids, docFId)
				break
			}
		}
	}
	return ids
}
