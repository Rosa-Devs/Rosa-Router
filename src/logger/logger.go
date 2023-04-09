package logger

import (
	"log"

	"github.com/fatih/color"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarn
	LogLevelInfo
)

func LogD(level LogLevel, component string, msg string) {
	var prefix string
	switch level {
	case LogLevelError:
		prefix = color.RedString("[ERROR]")
	case LogLevelWarn:
		prefix = color.YellowString("[WARN]")
	case LogLevelInfo:
		prefix = color.GreenString("[INFO]")
	}

	colors := map[string]*color.Color{
		"DB":  color.New(color.FgHiYellow),
		"HS":  color.New(color.FgHiBlue),
		"CLI": color.New(color.FgGreen),
		"WEB": color.New(color.FgCyan),
	}

	colorFunc, ok := colors[component]
	if !ok {
		colorFunc = color.New(color.Reset)
	}
	componentText := colorFunc.Sprint(component)

	log.Printf("%s %s %s\n", prefix, componentText+":", msg)

	//logWithComponent(LogLevelInfo, "CLI", "String")
	//logWithComponent(LogLevelInfo, "DB", "String")
	//logWithComponent(LogLevelInfo, "MAIN", "String")
	//logWithComponent(LogLevelError, "HS", "Guff")
	//logWithComponent(LogLevelWarn, "WEB", "guffffff")
}
