//go:build windows
// +build windows

package misskey

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func terminalWidth() int {
	width, _, err := term.GetSize(int(syscall.Stdout))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %+v", err)
		os.Exit(1)
	}
	return width
}
