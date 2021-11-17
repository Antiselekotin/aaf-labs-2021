package contentprocessing

import (
	"strings"
)

func SplitStringWithPositions(str string) (m map[string][]int) {
	m = make(map[string][]int)
	split := strings.Split(str, " ")
	for i := 0; i < len(split); i++ {
		m[split[i]] = append(m[split[i]], i)
	}
	return
}
