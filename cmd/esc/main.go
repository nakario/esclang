package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	esc "github.com/nakario/esclang"
)

var (
	debug = flag.Bool("debug", false, "show debug information")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [-debug] <file>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *debug {
		esc.Logger = log.New(os.Stderr, "[DEBUG]", log.LstdFlags | log.Lshortfile)
	} else {
		esc.Logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	file := flag.Arg(0)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read: %s\n", file)
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
	p, err := esc.Load(bytes.NewReader(b))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse: %s\n", file)
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
	if err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
