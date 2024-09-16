package main

import (
	"fmt"
	"log"
	"os/exec"
)

func basicExec() {
	var out1, _ = exec.Command("ls").Output()
	fmt.Printf("-> ls \n%s\n", out1)

	var out2, _ = exec.Command("pwd").Output()
	fmt.Printf("-> pwd \n%s\n", out2)

	var out3, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Printf("-> git config user.name \n%s\n", out3)
}

func useLookPath() {
	path, err := exec.LookPath("go")
	if err != nil {
		log.Fatal("Golang is not defined. Please install it")
	}

	fmt.Println("Golang command located at", path)
}

func main() {
	//basicExec()
	useLookPath()
}

// full
// https://pkg.go.dev/os/exec
