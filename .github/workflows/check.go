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
		bodyclose.Analyzer,
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
