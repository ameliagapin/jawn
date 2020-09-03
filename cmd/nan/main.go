package main

import (
	"fmt"
	"os"
)

func main() {
	var cmd string

	if len(os.Args) > 1 {
		cmd = os.Args[1]
		os.Args = append([]string{os.Args[0]}, os.Args[2:]...)
	}

	switch cmd {
	case "gen":
		generateDafault()
	case "extra":
		generateExtra()
	default:
		fmt.Println(`
		nan gen - generate implementations for the specified marshalers
		nan extra - generate marshalers implementations for extra types, defined by user
		`)
	}
}
