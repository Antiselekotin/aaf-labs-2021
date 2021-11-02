package invertedtree

import (
	"fmt"
	"labdb/internal/core/domain/textprocessing"
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
		textprocessing.ShiftString(depth, fmt.Sprintf("(%v) -> %v", prefix, n.Key())))
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
	root := t.RootNode().Left()
	for root != nil && !strings.HasPrefix(root.Key(), prefix) {
		if t.Less(root.Key(), prefix) {
			root = root.Right()
		} else {
			root = root.Left()
		}
	}

	if root == nil {
		return []int{}
	}
	idsMap := map[int]struct{}{}
	root.searchByPrefix(prefix, &idsMap)
	ids := make([]int, len(idsMap))
	i := 0
	for id := range idsMap {
		ids[i] = id
		i++
	}
	return ids
}

func (n *nodeStringIntMapOfIntSliceTreeMap) searchByPrefix(prefix string, docIds *map[int]struct{}) {
	if strings.HasPrefix(n.Key(), prefix) {
		for id := range n.Value() {
			(*docIds)[id] = struct{}{}
		}
		if l := n.Left(); l != nil {
			n.Left().searchByPrefix(prefix, docIds)
		}

		if r := n.Right(); r != nil {
			n.Right().searchByPrefix(prefix, docIds)
		}
	}
}

type ItterationSlice []*nodeStringIntMapOfIntSliceTreeMap
