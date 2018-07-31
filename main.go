package main

import (
	"flag"
	"github.com/miyo/p4_16_compiler_go/repl"
	"log"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	repl.Start(file, os.Stdout)

}
