package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

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
		level string
		file  string
		rel   bool
		wd    string
	}
)

func main() {
	app := kingpin.New("go-check", "Parse go/analysis reports and anotate diagnostics on the GitHub").Version(version).Author("kyoh86")
	app.Flag("level", "Annotation level").Default("warning").EnumVar(&flag.level, "warning", "error")
	app.Flag("wd", "Working directory").ExistingDirVar(&flag.wd)
	app.Flag("rel", "Show relative paths").BoolVar(&flag.rel)
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
	if flag.rel {
		wd, err := getWd()
		if err != nil {
			return fmt.Errorf("getting working directory: %w", err)
		}
		formatFilePath = func(f string) string {
			r, err := filepath.Rel(wd, f)
			if err != nil {
				return f
			}
			return r
		}
	}
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

func getWd() (string, error) {
	if flag.wd == "" {
		return os.Getwd()
	}
	return filepath.Abs(flag.wd)
}
