package main

import (
	"os"

	"github.com/beevik/hexdump"
)

func main() {
	hexdump.Dump(make([]byte, 108), hexdump.FormatGo, os.Stdout)
}
