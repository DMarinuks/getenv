# go-env-loader
Loads environment variables into an int, bool or string  
Has debug-mode to log what values environment variables have.  
Can log messages with color.  
You can change color and set custom logger.

## Description
`getenv` has 3 load functions `Int()`, `Bool()` or `Str()`  
Each function accepts 3 parameters:  
getenv.Int("ENV_KEY", &variable_pointer, log_fatal_bool)


## Example

```go
import "github.com/DMarinuks/go-env-loader"

var serverPort int
var debug bool
var uploadDir string

func main() {
  // set debug mode
  getenv.DebugMode(true)
  // set color mode
  getenv.ColorMode(true)

  // choose color from:
  // blue
  // lightBlue
  // yellow - default color
  // red
  // teal
  getenv.Color("yellow")

  // set custom logger
  myLogger = log.New(os.Stdout, "", 0)
  getenv.Logger(myLogger)
  
  // get environment variables
  getenv("SERVER_PORT", &serverPort, false)
  getenv("DEBUG", &serverPort, false)
  getenv("UPLOAD_DIR", &serverPort, true)
}
```
