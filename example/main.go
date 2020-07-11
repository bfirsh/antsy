package main

import (
	"fmt"
	"time"

	"github.com/bfirsh/antsy"
	"github.com/fatih/color"
)

func main() {
	printer := antsy.NewPrinter()

	for i := 1; i <= 100; i++ {
		printer.Print(color.GreenString("0 tests passed ") + "\n" +
			color.RedString(fmt.Sprintf("%d tests failed", i)) + "\n")

		time.Sleep(100 * time.Millisecond)
	}
}
