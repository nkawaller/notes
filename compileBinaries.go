package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	compileBinary("./cmd/server", "./bin/runserver")
	compileBinary("./cmd/staticgen", "./bin/staticgen")
}

func compileBinary(source, output string) {
	cmd := exec.Command("go", "build", "-o", output, source)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	fmt.Printf("Binary created: %s\n", output)
}
