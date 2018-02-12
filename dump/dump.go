package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/beevik/hexdump"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Printf("Usage: dump [file]\n")
		os.Exit(0)
	}

	b, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	hexdump.Dump(b, hexdump.FormatGo, os.Stdout)
}
