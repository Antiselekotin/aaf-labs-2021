package main

import (
	"bufio"
	"fmt"
	"labdb/internal/core/services/search"
	"log"
	"os"

	"github.com/fatih/color"
)

var stdin = os.Stdin
var searchService = search.NewSearch()

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
				res, err := searchService.Execute(str)
				if err != nil {
					color.Set(color.FgRed)
					fmt.Printf("%v\n", err)
					color.Unset()
				}
				color.Set(color.FgGreen)
				fmt.Printf("%v\n", res)
				color.Unset()
				buf = []byte{}
			}
		}
	}
}
