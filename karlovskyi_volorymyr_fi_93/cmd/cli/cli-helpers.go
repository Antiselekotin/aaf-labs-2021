package main

func searchLastQuoteIndex(buf *[]byte) int {
	for i := len(*buf) - 1; i >= 0; i-- {
		if (*buf)[i] == '"' {
			return i
		}
	}
	return -1
}

func lineEnded(buf *[]byte) (bool, int) {
	for i := searchLastQuoteIndex(buf) + 1; i < len(*buf); i++ {
		if (*buf)[i] == ';' {
			return true, i
		}
	}
	return false, -1
}
