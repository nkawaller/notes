package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	compileBinary("./cmd/server", "./bin/runserver")
	compileBinary("./cmd/staticgen", "./bin/staticgen")
	compileBinary("./cmd/buildindex", "./bin/buildindex")

    executeBinary("./bin/buildindex")
    executeBinary("./bin/staticgen")
	copyStaticFile("./web/static/output.css", "./deploy/static/output.css")

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

func executeBinary(file string) {
	cmd := exec.Command(file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	fmt.Printf("%s executed successfully\n", file)
}

func copyStaticFile(source, output string) {
	cmd := exec.Command("cp", source, output)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	fmt.Printf("File copied: %s -> %s\n", source, output)
}
