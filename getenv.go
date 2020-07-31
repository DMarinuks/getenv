package getenv

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	blue      = "\033[1;34m"
	lightBlue = "\033[1;36m"
	yellow    = "\033[1;33m"
	red       = "\033[1;31m"
	teal      = "\033[0;36m"
)

var logger = log.New(os.Stdout, "", 0)
var debugMode = false
var colorMode = false
var color = yellow

func Logger(newLogger *log.Logger) {
	logger = newLogger
}

func DebugMode(v bool) {
	debugMode = v
}

func ColorMode(v bool) {
	colorMode = v
}

func Color(v string) {
	switch v {
	case "blue":
		color = blue
	case "lightBlue":
		color = lightBlue
	case "yellow":
		color = yellow
	case "red":
		color = red
	case "teal":
		color = teal
	default:
		log.Fatal("setColor color", v, "not found")

	}
}

func makeFormat(color string, argsCount int) string {
	format := ""
	for i := 0; i < argsCount; i++ {
		format += "%v "
	}
	return color + format + "\033[0m"
}

func logMsg(fatal bool, args ...interface{}) {
	if debugMode || fatal {
		if colorMode {
			logger.Printf(makeFormat(color, len(args)), args...)
		} else {
			logger.Println(args...)
		}
	}
	if fatal {
		os.Exit(1)
	}
}

// Int loads env key and returns it as an int value
func Int(key string, variable *int, fatal bool) {
	logMsg(false, "GetEnv Key:", key, "=", os.Getenv(key))
	if val := os.Getenv(key); len(val) > 0 {
		i, err := strconv.Atoi(val)
		if err != nil {
			if fatal {
				logMsg(fatal, "Could not load ENV", key)
			}
		}
		*variable = i
	} else {
		if fatal {
			logMsg(fatal, "Could not load ENV", key)
		}
	}
}

// Bool loads env key and returns it as a bool value
func Bool(key string, variable *bool, fatal bool) {
	logMsg(false, "GetEnv Key:", key, "=", os.Getenv(key))
	if val := os.Getenv(key); len(val) > 0 {
		*variable = strings.ToLower(val) != "false"
	} else {
		if fatal {
			logMsg(fatal, "Could not load ENV", key)
		}
	}
}

// Str loads env key and returns it as a string value
func Str(key string, variable *string, fatal bool) {
	logMsg(false, "GetEnv Key:", key, "=", os.Getenv(key))
	if val := os.Getenv(key); len(val) > 0 {
		*variable = val
	} else {
		if fatal {
			logMsg(fatal, "Could not load ENV", key)
		}
	}
}
