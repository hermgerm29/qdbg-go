package internal

import (
    "fmt"
    "net/url"
    "os"
    "strings"
)


// Parse program arguments into a command and arguments.
// If no command is provided, exit with error.
func ParseArgs(progArgs []string) (cmd string, args []string) {
    if len(progArgs) < 2 {
        fmt.Println("[qdbg-go error] Please provide a command to run.")
        os.Exit(1)
    }
    return progArgs[1], progArgs[2:]
}

// Parse a trace and return last non-empty line.
func ParseTrace(trace string) string {
    var traceLines []string = strings.Split(trace, "\n")
    for i := len(traceLines)-1; i >= 0; i-- {
        var line string = traceLines[i]
        if line != "" {
            return line
        }
    }
    return ""
}

// Parse a trace and return last non-empty line. 
// If a non-empty line does not exist, exit with error.
func GetSearchUrl(traceBytes []byte) string {
    var traceStr string = string(traceBytes[:])
    var errorMsg string = ParseTrace(traceStr)
    return fmt.Sprintf("https://you.com/search?q=%s", url.QueryEscape(errorMsg))
}
