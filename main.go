package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os"
	"os/exec"
)

//go:embed lib/qjs
var qjs []byte

func main() {
	tmpFile, err := os.CreateTemp("", "njs")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer os.Remove(tmpFile.Name())

	if _, err := io.Copy(tmpFile, bytes.NewReader(qjs)); err != nil {
		fmt.Println("Error writing to temp file:", err)
		return
	}

	tmpFile.Close()
	if err := os.Chmod(tmpFile.Name(), 0755); err != nil {
		fmt.Println("Error setting file permissions:", err)
		return
	}

	cmd := exec.Command(tmpFile.Name(), "-e", "console.log('Hello, world')")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing file:", err)
		return
	}
	fmt.Println(string(output))
}
