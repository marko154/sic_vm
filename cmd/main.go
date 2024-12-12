package main

import (
	"flag"
	"os"
	"sic_vm/simulator"
	"sic_vm/tui"

	log "github.com/sirupsen/logrus"
)

func main() {
	debug := flag.Bool("debug", false, "enable debug logging")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	if flag.NArg() < 1 {
		log.Error("missing filename")
		log.Fatal("usage: sic_vm <filename>")
	}

	filename := flag.Arg(0)
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("error: Failed to open '%s': %v\n", filename, err)
	}

	defer file.Close()

	// vm := vm.NewVM(file)
	sim := simulator.NewSimulator(file)
	tui.Setup(sim)

	// if err := vm.Load(); err != nil {
	// 	log.Fatalf("error: Failed to load '%s': %v\n", filename, err)
	// }

	// if err = vm.Run(); err != nil {
	// 	log.Fatalf("failed to run program %v\n", err)
	// }
}
