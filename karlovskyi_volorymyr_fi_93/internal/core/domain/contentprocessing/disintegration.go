package contentprocessing

import (
	"strings"
	"bytes"
)

func SplitBytesWithPositions(p []byte) (m map[string][]int) {
	m = make(map[string][]int)
	split := bytes.Split(p, []byte{' '})
	for i := 0; i < len(split); i++ {
		m[string(split[i])] = append(m[string(split[i])], i)
	}
	return
}

func SplitStringWithPositions(str string) (m map[string][]int) {
	m = make(map[string][]int)
	split := strings.Split(str, " ")
	for i := 0; i < len(split); i++ {
		m[split[i]] = append(m[split[i]], i)
	}
	return
}
