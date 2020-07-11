package antsy

import (
	"fmt"
	"strings"

	"github.com/k0kubun/go-ansi"
)

type Printer struct {
	previousPrint string
}

func NewPrinter() *Printer {
	return &Printer{previousPrint: ""}
}

func (p *Printer) Print(s string) {
	if p.previousPrint != "" {
		// Erase everything we printed last time
		lineCount := strings.Count(p.previousPrint, "\n")
		// Erase current line
		ansi.EraseInLine(2)
		fmt.Print("\r")
		for i := 0; i < lineCount; i++ {
			// Go up to previous line and erase it
			ansi.CursorPreviousLine(1)
			ansi.EraseInLine(2)
		}
	}
	fmt.Print(s)
	p.previousPrint = s
}
