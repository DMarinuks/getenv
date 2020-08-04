package getenv

import (
	"errors"
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
		c := color
		if fatal {
			c = red
		}
		if colorMode {
			logger.Printf(makeFormat(c, len(args)), args...)
		} else {
			logger.Println(args...)
		}
	}
	if fatal {
		os.Exit(1)
	}
}

// Int - takes ENV_KEY, out *int, fatal bool
// It will override "out" value if ENV_KEY is not empty
func Int(key string, out *int, fatal bool) {
	logMsg(false, "GetEnv Key:", key, "=", os.Getenv(key))
	if val := os.Getenv(key); len(val) > 0 {
		i, err := strconv.Atoi(val)
		if err != nil {
			if fatal {
				logMsg(fatal, "Could not load ENV", key)
			}
			return
		}
		*out = i
		return
	}
	if fatal {
		logMsg(fatal, "Could not load ENV", key)
	}
}

// Bool - takes ENV_KEY, out *string, fatal bool
// It will override "out" value if ENV_KEY is not empty
func Bool(key string, out *bool, fatal bool) {
	logMsg(false, "GetEnv Key:", key, "=", os.Getenv(key))
	if val := strings.ToLower(os.Getenv(key)); len(val) > 0 {
		if val == "false" || val == "true"  {
			*out = strings.ToLower(val) == "true"
			return
		}
	}
	if fatal {
		logMsg(fatal, "Could not load ENV", key)
	}
}

// Str - takes ENV_KEY, out *string, fatal bool
// It will override "out" value if ENV_KEY is not empty
func Str(key string, out *string, fatal bool)  {
	logMsg(false, "GetEnv Key:", key, "=", os.Getenv(key))
	if val := os.Getenv(key); len(val) > 0 {
		*out = val
		return
	}
	if fatal {
		logMsg(fatal, "Could not load ENV", key)
	}
}

// StrSlice takes ENV_KEY string, separator string, fatal bool.
// It will return slice of strings and error value.
func StrSlice(key string, separator string, fatal bool) []string {
	logMsg(false, "GetEnv Key:", key, "=", os.Getenv(key))
	var results []string
	if val := os.Getenv(key); len(val) > 0 {
		arr := strings.Split(val, separator)
		for _, v := range arr {
			results = append(results, strings.TrimSpace(v))
		}
		return results
	}
	if fatal {
		logMsg(fatal, "Could not load ENV", key)
	}
	return results
}

// IntSlice takes ENV_KEY string, separator string, fatal bool
// It will return slice of strings and error value.
func IntSlice(key string, separator string, fatal bool) ([]int, error){
	logMsg(false, "GetEnv Key:", key, "=", os.Getenv(key))
	var results []int
	if val := os.Getenv(key); len(val) > 0 {
		arr := strings.Split(val, separator)
		for _, v := range arr {
			i, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				if fatal {
					logMsg(fatal, "Could not load ENV", key)
				}
				return []int{}, err
			}
			results = append(results, i)
		}
		return results, nil
	}
	if fatal {
		logMsg(fatal, "Could not load ENV", key)
	}
	return results, errors.New("environment variable not found")
}

// BoolSlice takes ENV_KEY string, separator string, fatal bool
// It will return slice of strings and error value.
func BoolSlice(key string, separator string, fatal bool) ([]bool, error){
	logMsg(false, "GetEnv Key:", key, "=", os.Getenv(key))
	var results []bool
	if val := os.Getenv(key); len(val) > 0 {
		arr := strings.Split(val, separator)
		for _, v := range arr {
			v = strings.TrimSpace(strings.ToLower(v))
			if v == "false" || v == "true" {
				results = append(results, strings.ToLower(v) == "true")
			} else {
				if fatal {
					logMsg(fatal, "Could not load ENV", key)
				}
				return []bool{}, errors.New("not bool value")
			}

		}
		return results, nil
	}
	if fatal {
		logMsg(fatal, "Could not load ENV", key)
	}
	return results, errors.New("environment variable not found")
}