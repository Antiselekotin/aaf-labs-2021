package invertedtree

import (
	"fmt"
	"labdb/internal/core/domain/contentprocessing"
	"strings"
)

func (t *StringIntMapOfIntSliceTreeMap) RootNode() *nodeStringIntMapOfIntSliceTreeMap {
	root := t.beginNode
	for root.parent != nil {
		root = root.parent
	}
	return root
}

func (t *StringIntMapOfIntSliceTreeMap) Begin() *nodeStringIntMapOfIntSliceTreeMap {
	return t.beginNode
}

func (n *nodeStringIntMapOfIntSliceTreeMap) Parent() *nodeStringIntMapOfIntSliceTreeMap {
	return n.parent
}

func (t *StringIntMapOfIntSliceTreeMap) PrintTree() {
	tellAboutYourself(t.RootNode(), 0, "root")
}

func tellAboutYourself(n *nodeStringIntMapOfIntSliceTreeMap, depth int, prefix string) {
	if n == nil {
		return
	}
	fmt.Println(
		contentprocessing.ShiftString(depth, fmt.Sprintf("(%v) -> %v", prefix, n.Key())))
	tellAboutYourself(n.left, depth+1, "left")
	tellAboutYourself(n.right, depth+1, "right")
}

func (n *nodeStringIntMapOfIntSliceTreeMap) Left() *nodeStringIntMapOfIntSliceTreeMap {
	return n.left
}

func (n *nodeStringIntMapOfIntSliceTreeMap) Right() *nodeStringIntMapOfIntSliceTreeMap {
	return n.right
}

func (n *nodeStringIntMapOfIntSliceTreeMap) Key() string {
	if n == nil {
		return ""
	}
	return n.key
}

func (n *nodeStringIntMapOfIntSliceTreeMap) Value() map[int][]int {
	return n.value
}

func (t *StringIntMapOfIntSliceTreeMap) SearchByPrefix(prefix string) []int {
	root := t.RootNode().Left().goToNextPrefix(prefix, t.Less)

	if root == nil {
		return []int{}
	}
	idsMap := map[int]struct{}{}
	root.searchByPrefix(prefix, &idsMap, t.Less)
	ids := make([]int, len(idsMap))
	i := 0
	for id := range idsMap {
		ids[i] = id
		i++
	}
	return ids
}

func (n *nodeStringIntMapOfIntSliceTreeMap) searchByPrefix(prefix string, docIds *map[int]struct{}, less func(string, string) bool) {
	if !strings.HasPrefix(n.Key(), prefix) {
		n = n.goToNextPrefix(prefix, less)
	}
	if n == nil {
		return
	}
	for id := range n.Value() {
		(*docIds)[id] = struct{}{}
	}
	if l := n.Left(); l != nil {
		n.Left().searchByPrefix(prefix, docIds, less)
	}

	if r := n.Right(); r != nil {
		n.Right().searchByPrefix(prefix, docIds, less)
	}
}

func (n *nodeStringIntMapOfIntSliceTreeMap) goToNextPrefix(prefix string, less func(string, string) bool) *nodeStringIntMapOfIntSliceTreeMap {
	for n != nil && !strings.HasPrefix(n.Key(), prefix) {
		if less(n.Key(), prefix) {
			n = n.Right()
		} else {
			n = n.Left()
		}
	}
	return n
}

type ItterationSlice []*nodeStringIntMapOfIntSliceTreeMap
