//go:build !windows || idea
// +build !windows idea

package logger

import "fmt"

type log struct {
	name string
}

func NewLog(name string) *log {
	return &log{
		name: name,
	}
}

func (l *log) Error(message ...interface{}) {
	fmt.Printf("\033[31mError %s -> %s: %v\033[0m\n", getTime(), l.name, message)
}

func (l *log) Info(message ...interface{}) {
	fmt.Printf("\033[32mInfo %s -> %s: %v\033[0m\n", getTime(), l.name, message)
}
