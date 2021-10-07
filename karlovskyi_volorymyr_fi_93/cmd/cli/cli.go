package main

import (
	"bufio"
	"labdb/internal/core/services/search"
	"log"
	"os"
)

var stdin = os.Stdin
var ra = cliResponseAdapter{}
var searchService = search.NewSearch(&ra)

func main() {

	reader := bufio.NewReader(stdin)
	buf := []byte{}
	isOpenQuote := false
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		for _, c := range line {
			if c == '"' {
				isOpenQuote = !isOpenQuote
			}
		}
		if len(buf) != 0 {
			buf = append(buf, '\n')
		}
		buf = append(buf, line...)

		if !isOpenQuote {
			isLineEnded, endedFrom := lineEnded(&buf)
			if isLineEnded {
				str := string(buf[:endedFrom+1])
				if str == ".EXIT;" {
					os.Exit(0)
				}
				searchService.Execute(str)
				buf = []byte{}
			}
		}
	}
}
