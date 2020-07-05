package main

import (
	"bufio"
	"bytes"
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type Diagnostic struct {
	Position
	Message string
	VetTool string
}

type RawDiagnostic struct {
	Position Position `json:"posn"`
	Message  string   `json:"message"`
}

type Position struct {
	File string
	Line int64
	Col  int64
}

func (p *Position) UnmarshalText(t []byte) error {
	terms := strings.Split(string(t), ":")
	if len(terms) < 3 {
		return errors.New("position has no lines and columns")
	}
	col, err := strconv.ParseInt(terms[len(terms)-1], 10, 64)
	if err != nil {
		return fmt.Errorf("parsing col position: %w", err)
	}
	line, err := strconv.ParseInt(terms[len(terms)-2], 10, 64)
	if err != nil {
		return fmt.Errorf("parsing line position: %w", err)
	}
	p.Col = col
	p.Line = line
	p.File = strings.Join(terms[:len(terms)-2], ":")
	return nil
}

var _ encoding.TextUnmarshaler = (*Position)(nil)

func ParseDiagnostics(r io.Reader) ([]Diagnostic, error) {
	raw, err := readAllWithoutComment(r)
	if err != nil {
		return nil, fmt.Errorf("reading input: %w", err)
	}
	var rawMap map[string]map[string][]RawDiagnostic
	if err := json.Unmarshal(raw, &rawMap); err != nil {
		return nil, err
	}

	var diagnostics []Diagnostic
	// rawMap is [package path]:[toolMap]
	for _, toolMap := range rawMap {
		// toolMap is [tool name]:[diagnostics]
		for vetTool, diags := range toolMap {
			for _, d := range diags {
				diagnostics = append(diagnostics, Diagnostic{
					Position: d.Position,
					Message:  d.Message,
					VetTool:  vetTool,
				})
			}
		}
	}
	return diagnostics, nil
}

func readAllWithoutComment(r io.Reader) ([]byte, error) {
	// ignore first line if it begins with '#'
	// (go vet prints # {package name} on first line)
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	skip := 0

	scanner := bufio.NewScanner(bytes.NewBuffer(raw))
	if scanner.Scan() {
		line := scanner.Bytes()
		if line[0] == '#' {
			skip = len(line)
		}
	}
	return raw[skip:], nil
}
