package main

import (
	"fmt"

	"github.com/fatih/color"
)

type cliResponseAdapter struct {
}

func (ra *cliResponseAdapter) OnError(err error) {
	color.Set(color.FgRed)
	fmt.Printf("%v\n", err)
	color.Unset()
}

func (ra *cliResponseAdapter) OnCreateSuccess(str string) {
	color.Set(color.FgGreen)
	fmt.Printf("%v\n", str)
	color.Unset()
}

func (ra *cliResponseAdapter) OnCreateFailure(err error) {
	ra.OnError(err)
}
