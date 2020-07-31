# go-env-loader
Loads ENV into int, bool or string

## Description
getenv has 3 load functions Int(), Bool() or Str()  
each function accepts 3 parameters:  
getenv.Int(ENV_KEY, variable_pointer, fatal)

## Example

```
import "github.com/DMarinuks/go-env-loader/getenv"

var serverPort int
var debug bool
var uploadDir string

func main() {
  getenv.DebugMode(true)
  getenv.ColorMode(true)
  // choose color from:
  // blue
  // lightBlue
  // yellow
  // red
  // teal
  getenv.Color("yellow")
  
  getenv("SERVER_PORT", &serverPort, false)
  getenv("DEBUG", &serverPort, false)
  getenv("UPLOAD_DIR", &serverPort, true)
}
```
