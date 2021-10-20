package invertedtree

import (
	"fmt"
	"labdb/internal/core/domain/textprocessing"
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

type ItterationSlice []*nodeStringIntMapOfIntSliceTreeMap
