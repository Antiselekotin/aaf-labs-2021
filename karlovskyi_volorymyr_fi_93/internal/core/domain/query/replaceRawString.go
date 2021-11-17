package query

import (
	"fmt"
	"labdb/internal/core/domain/contentprocessing"
)

func replaceRawString(str string) (res string, memMap map[string]string) {
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
				res += contentprocessing.ReplaceNotAllowedChars(string(buf[:mem])) + memStr
				memMap[memStr] = string(buf[mem+1 : i])
				buf = buf[i+1:]
				i = 0
				mem = -1
			}
		}
	}

	if len(res) == 0 {
		res = string(buf)
	} else {
		res += string(buf)
	}
	return
}
