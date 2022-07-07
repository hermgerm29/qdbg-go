package main

import (
    "fmt"
    "os"
    "os/exec"
    "github.com/pkg/browser"
    "github.com/hermgerm29/qdbg/internal"
)


func main() {

    // Extract command to run
    cmd, args := internal.ParseArgs(os.Args)

    // Execute command
    output, err := exec.Command(cmd, args...).CombinedOutput()

    if err != nil {
        if len(output) == 0 {
            fmt.Printf("Command not found %s\n", cmd)
            os.Exit(1)
        }
        if browser.OpenURL(internal.GetSearchUrl(output)) != nil {
            fmt.Println("Unable to open brower. Is one installed?")
        }
    }

    // Flush command output
    fmt.Printf("%s", output)
}
