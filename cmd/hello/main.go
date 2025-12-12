package main

import (
	"flag"
	"fmt"
	"os"
)

// These will be overridden via -ldflags by GoReleaser
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

func main() {
	showVersion := flag.Bool("version", false, "Print version information and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("goreleaser-starter\n")
		fmt.Printf("  Version: %s\n", Version)
		fmt.Printf("  Commit:  %s\n", Commit)
		fmt.Printf("  Date:    %s\n", Date)
		return
	}

	fmt.Println("Hello, GoReleaser & GitHub Releases! 👋")
	os.Exit(0)
}
