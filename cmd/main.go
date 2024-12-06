package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	vm "sic_vm/internal"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("missing filename")
		fmt.Println("usage: sic_vm <filename>")
		os.Exit(1)
	}

	filename := flag.Arg(0)
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("error: Failed to open '%s': %v\n", filename, err)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	vm := vm.NewVM(reader)

	if err := vm.Load(); err != nil {
		fmt.Printf("error: Failed to load '%s': %v\n", filename, err)
		os.Exit(1)
	}

	vm.Run()
}
