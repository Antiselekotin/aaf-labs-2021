package main

import (
	"fmt"

	"github.com/fatih/color"
)

type cliResponseAdapter struct {
}

func (ra *cliResponseAdapter) OnSuccess(str string) {
	color.Set(color.FgGreen)
	fmt.Printf("%v\n", str)
	color.Unset()
}

func (ra *cliResponseAdapter) OnError(err error) {
	color.Set(color.FgRed)
	fmt.Printf("%v\n", err)
	color.Unset()
}

func (ra *cliResponseAdapter) OnCreateSuccess(str string) {
	ra.OnSuccess(str)
}

func (ra *cliResponseAdapter) OnCreateFailure(err error) {
	ra.OnError(err)
}

func (ra *cliResponseAdapter) OnInsertSuccess(str string) {
	ra.OnSuccess(str)
}

func (ra *cliResponseAdapter) OnInsertFailure(err error) {
	ra.OnError(err)
}
