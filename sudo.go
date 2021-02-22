package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"strings"
	"syscall"
)

func main() {
	var exe string
	var args string

	argLen := len(os.Args)
	if argLen == 1 {
		fmt.Printf(
			`Run command as Administrator
Example: sudo <command line>
`)
		os.Exit(0)
	}
	exe = os.Args[1]
	args = strings.Join(os.Args[2:], " ")
	cwd, _ := os.Getwd()

	err := runAsAdministrator(exe, args, cwd)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func runAsAdministrator(exe, args, cwd string) error {
	verb := "runas"

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = windows.SW_NORMAL //SW_NORMAL or SW_HIDE

	return windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
}
