# go-check-action

Parse [go/analysis](https://pkg.go.dev/golang.org/x/tools/go/analysis) reports and anotate diagnostics on the GitHub

[![Go Report Card](https://goreportcard.com/badge/github.com/kyoh86/go-check-action)](https://goreportcard.com/report/github.com/kyoh86/go-check-action)
[![Coverage Status](https://img.shields.io/codecov/c/github/kyoh86/go-check-action.svg)](https://codecov.io/gh/kyoh86/go-check-action)
[![Release](https://github.com/kyoh86/go-check-action/workflows/Release/badge.svg)](https://github.com/kyoh86/go-check-action/releases)

## Usage

### In the GitHub Action

```yaml
step:
  - name: go vet
    run: go vet -json ./... > diagnostics.json

  - name: annotate diagnostics
    uses: docker://docker.pkg.github.com/kyoh86/go-check-action/go-check:latest
    with:
      level: error
```

If you want to use other custom go/analysis checkers:

```yaml
step:
  - name: other custom checker
    run: some-custom-checker -json ./... > diagnostics.json

  - name: annotate diagnostics
    uses: docker://docker.pkg.github.com/kyoh86/go-check-action/go-check:latest
    with:
      level: error
```

NOTE: `unitchecker`, `singlechecker` and `multichecker` in the go/analysis support `-json` flag.

### By the docker

```console
$ docker pull docker.pkg.github.com/kyoh86/go-check-action/go-check:latest
$ docker run docker.pkg.github.com/kyoh86/go-check-action/go-check:latest --help
```

### As a go program

```console
$ go get github.com/kyoh86/go-check-action
$ go-check-action --help
```

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
