# go-check-action

Parse [go/analysis](https://pkg.go.dev/golang.org/x/tools/go/analysis) reports and anotate diagnostics on the GitHub

[![Go Report Card](https://goreportcard.com/badge/github.com/kyoh86/go-check-action)](https://goreportcard.com/report/github.com/kyoh86/go-check-action)
[![Coverage Status](https://img.shields.io/codecov/c/github/kyoh86/go-check-action.svg)](https://codecov.io/gh/kyoh86/go-check-action)
[![Release](https://github.com/kyoh86/go-check-action/workflows/Release/badge.svg)](https://github.com/kyoh86/go-check-action/releases)

## Example

<img src="go-check-1.png" width="450" height="400">

<img src="go-check-2.png" width="450" height="500">

## Usage

```yaml
    steps:
      - name: go vet
        run: go vet -json ./... 2> diagnostics.json

      - name: annotate diagnostics
        uses: kyoh86/go-check-action/annotate@v1
        with:
          level: error
          exit-code: 1
```

Get more information: [USAGE.md](USAGE.md)

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
