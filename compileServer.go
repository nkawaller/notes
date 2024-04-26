package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {

    compileServer("./cmd/server", "./bin/runserver")

}

func compileServer(source, output string) {
    cmd := exec.Command("go", "build", "-o", output, source)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        panic(err)
    }
    fmt.Printf("Binary created in %s\n", output)
}
