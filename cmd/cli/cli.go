package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
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
		buf = append(buf, line...)
		if !isOpenQuote {
			isLineEnded, endedFrom := lineEnded(&buf)
			if isLineEnded {
				fmt.Printf("\n%s\n", buf[:endedFrom+1])
				buf = []byte{}
			}
		}
	}
}
