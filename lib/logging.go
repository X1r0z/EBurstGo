package lib

import (
	"github.com/fatih/color"
)

type Logging struct {
	Verbose bool
	IsDebug bool
}

func (logging *Logging) Success(format string, a ...interface{}) {
	c := color.New(color.FgGreen).Add(color.Bold)
	c.Printf(format, a...)
	c.Print("\n")
}

func (logging *Logging) Failed(format string, a ...interface{}) {
	if logging.Verbose {
		c := color.New(color.FgRed).Add(color.Bold)
		c.Printf(format, a...)
		c.Print("\n")
	}
}

func (logging *Logging) Info(format string, a ...interface{}) {
	c := color.New(color.FgWhite).Add(color.Bold)
	c.Printf(format, a...)
	c.Print("\n")
}

func (logging *Logging) Debug(format string, a ...interface{}) {
	if logging.IsDebug {
		c := color.New(color.FgYellow).Add(color.Bold)
		c.Printf(format, a...)
		c.Print("\n")
	}
}
