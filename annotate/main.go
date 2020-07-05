package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	flags struct {
		level    string
		files    string
		exitCode int
	}
)

func main() {
	log.SetFlags(0)

	flag.StringVar(&flags.level, "level", "warning", "Annotation level, 'warning' or 'error'")
	flag.IntVar(&flags.exitCode, "exit-code", 0, "Exit code when any diagnostics found")
	flag.Parse()
	flags.files = flag.Arg(0)

	if flags.level != "warning" && flags.level != "error" {
		flag.Usage()
		os.Exit(2)
	}

	count, err := run()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%d diagnostics found", count)
	if count > 0 {
		os.Exit(flags.exitCode)
	}
}

func run() (int, error) {
	total := 0
	for _, filename := range filepath.SplitList(flags.files) {
		diagnostics, err := parse(filename)
		if err != nil {
			return 0, err
		}
		var formatFilePath func(string) string = func(f string) string { return f }
		for _, d := range diagnostics {
			fmt.Printf(
				"::%s file=%s,line=%d,col=%d::%s\n",
				flags.level,
				formatFilePath(d.File),
				d.Line,
				d.Col,
				d.Message,
			)
		}
		total += len(diagnostics)
	}

	return total, nil
}

func parse(filename string) ([]Diagnostic, error) {
	var r io.Reader
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()
	r = file
	return ParseDiagnostics(r)
}
