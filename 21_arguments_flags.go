package main

import (
	"flag"
	"fmt"
	"os"
)

// Example use
// go run file.go hallo iam verdi
// go build file.go
// file.exe hallo iam verdi

func useArgs() {
	var args = os.Args
	fmt.Printf("-> %#v\n", args)
	// -> []string{"<tmp exe path.exe>", "hello", "iam", "verdi"}

	args = args[1:]
	fmt.Printf("-> %#v\n", args)
	// ->	[]string{"hello", "iam", "verdi"}

	//	Exe output
	// -> []string{"<exe path.exe>", "hello", "iam", "verdi"}
	// -> []string{"hello", "iam", "verdi"}
}

func useFlags() {
	var name = flag.String("name", "anonymous", "Your name")
	var age int
	flag.IntVar(&age, "age", 18, "Your age")

	flag.Parse()
	fmt.Printf("Name	: %s\n", *name)
	fmt.Printf("Age	: %d", age)
}

func main() {
	//useArgs()
	useFlags()
}
