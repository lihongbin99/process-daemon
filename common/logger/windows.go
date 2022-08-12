//go:build windows && !idea
// +build windows,!idea

package logger

import (
	"fmt"
	"syscall"
)

var (
	kernel32    = syscall.NewLazyDLL(`kernel32.dll`)
	proc        = kernel32.NewProc(`SetConsoleTextAttribute`)
	CloseHandle = kernel32.NewProc(`CloseHandle`)
)

type log struct {
	name string
}

func NewLog(name string) *log {
	return &log{
		name: name,
	}
}

func (l *log) Error(message ...interface{}) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(12))
	fmt.Printf("Error %s -> %s: %v\n", getTime(), l.name, message)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	_, _, _ = CloseHandle.Call(handle)
}

func (l *log) Info(message ...interface{}) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(10))
	fmt.Printf("Info %s -> %s: %v\n", getTime(), l.name, message)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	_, _, _ = CloseHandle.Call(handle)
}
