package common

import (
	"github.com/fatih/color"
)

type Logger struct {
	Verbose bool
}

func (r *Logger) Success(format string, a ...interface{}) {
	c := color.New(color.FgGreen).Add(color.Bold)
	c.Printf(format+"\n", a...)
}

func (r *Logger) Failed(format string, a ...interface{}) {
	if r.Verbose {
		c := color.New(color.FgRed).Add(color.Bold)
		c.Printf(format+"\n", a...)
	}
}

func (r *Logger) Info(format string, a ...interface{}) {
	c := color.New(color.FgWhite).Add(color.Bold)
	c.Printf(format+"\n", a...)
}

func (r *Logger) Debug(format string, a ...interface{}) {
	if r.Verbose {
		c := color.New(color.FgYellow).Add(color.Bold)
		c.Printf(format+"\n", a...)
	}
}
