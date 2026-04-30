// Package main is the entry point for the rompatch application.
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/vetlekise/rompatch/patcher"
)

var base string
var patch string
var out string

func init() {
	flag.StringVar(&base, "base", "", "Path to the base ROM.")
	flag.StringVar(&patch, "patch", "", "Path to the ROM patch.")
	flag.StringVar(&out, "out", "", "Location to save the patched ROM.")
	flag.Parse()

	if base == "" || patch == "" || out == "" {
		fmt.Println("Error: Missing required arguments.")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Patching...\nBase: %s\nPatch: %s\nOut: %s\n", base, patch, out)
}

func main() {
	var p patcher.Patcher = patcher.IPS{}

	err := p.Apply(base, patch, out)
	if err != nil {
		slog.Error("failed to patch", "err", err)
		os.Exit(1)
	}

	fmt.Println("patch complete")
}
