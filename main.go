package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/alecthomas/kingpin"
)

// nolint
var (
	version = "snapshot"
	commit  = "snapshot"
	date    = "snapshot"
)

var (
	flag struct {
		level    string
		file     string
		exitCode int
	}
)

func main() {
	app := kingpin.New("go-check", "Parse go/analysis reports and anotate diagnostics on the GitHub").Version(version).Author("kyoh86")
	app.Flag("level", "Annotation level").Default("warning").EnumVar(&flag.level, "warning", "error")
	app.Flag("exit-code", "Exit code when any diagnostics found").Default("1").IntVar(&flag.exitCode)
	app.Arg("go-vet-JSON", "JSON files which `go vet -json` reported.").ExistingFileVar(&flag.file)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	log.SetFlags(0)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	diagnostics, err := parse()
	if err != nil {
		return err
	}
	var formatFilePath func(string) string = func(f string) string { return f }
	for _, d := range diagnostics {
		fmt.Printf(
			"::%s file=%s,line=%d,col=%d::%s\n",
			flag.level,
			formatFilePath(d.File),
			d.Line,
			d.Col,
			d.Message,
		)
	}
	if len(diagnostics) > 0 {
		log.Printf("%d diagnostics found", len(diagnostics))
		os.Exit(flag.exitCode)
	}
	return nil
}

func parse() ([]Diagnostic, error) {
	var r io.Reader
	if flag.file == "-" {
		r = os.Stdin
	} else {
		file, err := os.Open(flag.file)
		if err != nil {
			return nil, fmt.Errorf("opening file: %w", err)
		}
		defer file.Close()
		r = file
	}
	return ParseDiagnostics(r)
}
