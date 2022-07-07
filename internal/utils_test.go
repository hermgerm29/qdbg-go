package internal

import (
    "fmt"
    "testing"
)


// Test that parse args correctly returns the command and args
func TestParseArgsCorrectly(t *testing.T) {
    progArgs := []string{"qdbg", "cmd", "arg1", "arg2"}
    cmd, args := ParseArgs(progArgs)
    fmt.Println(args)
    if cmd != "cmd" {
        t.Fatalf("Error in ParseArgs. Got command %s but expected `cmd`", cmd)
    }
}
