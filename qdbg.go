package main

import (
    "fmt"
    "os"
    "os/exec"
    "github.com/pkg/browser"

    "qdbg/internal"
)


func main() {

    // Extract command to run
    cmd, args := internal.ParseArgs(os.Args)

    // Execute command
    output, err := exec.Command(cmd, args...).CombinedOutput()

    // On error, search for the user
    if err != nil {
        if browser.OpenURL(internal.GetSearchUrl(output)) != nil {
            fmt.Println("Unable to open brower. Is one installed?")
        }
    }

    // Flush command output
    fmt.Printf("%s", output)
}
