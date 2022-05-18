# Usage of go-check action

## Linters

We can use `go vet` or the many custom checkers providing `*go/analysis.Analyzer`.

### With `go vet`

Call `go vet` and pass the output JSON file to go-check/annotate action.

```yaml
on: [push]

jobs:
  check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
        with:
          go-version: '^1.18'

      - name: go vet
        run: go vet -json ./... 2> diagnostics.json

      - name: annotate diagnostics
        uses: kyoh86/go-check-action/annotate@v1
        with:
          level: error
          exit-code: 1
```

### With custom linters

If you want to use custom go/analysis checkers

- Prepare `check.go` calling `go/analysis/mutichecker.Main`
    - Recommend: put a build tag like `//+build lint`
- Call it from action with `-json`.
    - Recommend: with a build tag like `-tags lint`
- Pass the output JSON file to go-check/annotate action.

#### check.go sample

.github/workflows/check.go:

```go
//+build lint

package main

import (
	err113 "github.com/Djarvur/go-err113"
	"github.com/jingyugao/rowserrcheck/passes/rowserr"
	printfunc "github.com/jirfag/go-printf-func-name/pkg/analyzer"
	"github.com/kyoh86/exportloopref"
	"github.com/kyoh86/looppointer"
	"github.com/maratori/testpackage/pkg/testpackage"
	"github.com/nishanths/exhaustive"
	"github.com/sonatard/noctx"
	"github.com/tdakkota/asciicheck"
	"github.com/timakin/bodyclose/passes/bodyclose"
	magicnumbers "github.com/tommy-muehle/go-mnd"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/asmdecl"
	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/atomic"
	"golang.org/x/tools/go/analysis/passes/bools"
	"golang.org/x/tools/go/analysis/passes/buildtag"
	"golang.org/x/tools/go/analysis/passes/cgocall"
	"golang.org/x/tools/go/analysis/passes/composite"
	"golang.org/x/tools/go/analysis/passes/copylock"
	"golang.org/x/tools/go/analysis/passes/errorsas"
	"golang.org/x/tools/go/analysis/passes/httpresponse"
	"golang.org/x/tools/go/analysis/passes/ifaceassert"
	"golang.org/x/tools/go/analysis/passes/loopclosure"
	"golang.org/x/tools/go/analysis/passes/lostcancel"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shift"
	"golang.org/x/tools/go/analysis/passes/stdmethods"
	"golang.org/x/tools/go/analysis/passes/stringintconv"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"golang.org/x/tools/go/analysis/passes/tests"
	"golang.org/x/tools/go/analysis/passes/unmarshal"
	"golang.org/x/tools/go/analysis/passes/unreachable"
	"golang.org/x/tools/go/analysis/passes/unsafeptr"
	"golang.org/x/tools/go/analysis/passes/unusedresult"
	"honnef.co/go/tools/simple"
	"honnef.co/go/tools/staticcheck"
	"honnef.co/go/tools/stylecheck"
	"honnef.co/go/tools/unused"
)

func main() {
	var analyzers []*analysis.Analyzer
	for _, v := range simple.Analyzers {
		analyzers = append(analyzers, v)
	}
	for _, v := range staticcheck.Analyzers {
		analyzers = append(analyzers, v)
	}
	for _, v := range stylecheck.Analyzers {
		analyzers = append(analyzers, v)
	}

	analyzers = append(analyzers,
		asciicheck.NewAnalyzer(),
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		bools.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		err113.NewAnalyzer(),
		errorsas.Analyzer,
		exhaustive.Analyzer,
		exportloopref.Analyzer,
		httpresponse.Analyzer,
		ifaceassert.Analyzer,
		loopclosure.Analyzer,
		looppointer.Analyzer,
		lostcancel.Analyzer,
		magicnumbers.Analyzer,
		nilfunc.Analyzer,
		noctx.Analyzer,
		printf.Analyzer,
		printfunc.Analyzer,
		rowserr.NewAnalyzer(),
		shift.Analyzer,
		stdmethods.Analyzer,
		stringintconv.Analyzer,
		structtag.Analyzer,
		testpackage.NewAnalyzer(),
		tests.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		unsafeptr.Analyzer,
		unused.Analyzer,
		unusedresult.Analyzer,
	)

	multichecker.Main(analyzers...)
}
```

#### Sample of workflow

.github/workflows/check.yml:

```yaml
on: [push]

jobs:
  check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
        with:
          go-version: '^1.14'

      - name: check
        run: go run -tags lint ./.github/workflows/check.go -json ./... > diagnostics.json

      - name: annotate diagnostics
        uses: kyoh86/go-check-action/annotate@v1
        with:
          level: error
          exit-code: 1
```

## Action Parameters

| Name        | Default          | Description                                                   |
| ---         | ---              | ---                                                           |
| level       | warning          | Which level to annotate, `warning` or `error`                 |
| exit-code   | 0                | Exit code when any diagnostics found                          |
| go-vet-json | diagnostics.json | A JSON file that a go/analysis (e.g. `go vet -json`) reported |
