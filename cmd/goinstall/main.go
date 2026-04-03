package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tinywasm/goinstall"
)

func main() {
	version := flag.String("version", goinstall.DefaultVersion, "Go version to install")
	verbose := flag.Bool("v", false, "Enable verbose output")
	help := flag.Bool("h", false, "Show help")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	opts := []goinstall.Option{
		goinstall.WithVersion(*version),
	}

	if *verbose {
		opts = append(opts, goinstall.WithLogger(func(msg string) {
			fmt.Println(msg)
		}))
	}

	goPath, err := goinstall.EnsureInstalled(opts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Go installed at: %s\n", goPath)
	fmt.Println("Note: If you are using a new shell, you might need to run 'hash -r' to update the binary path cache.")
}
