package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	executeBinary("./bin/staticgen")
	copyStaticFile("./web/static/output.css", "./deploy/static/output.css")

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
