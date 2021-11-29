package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"unicode"

	"github.com/sanposhiho/mock/mockgen/model"
)

// Sanitize cleans up a string to make a suitable package name.
func Sanitize(s string) string {
	t := ""
	for _, r := range s {
		if t == "" {
			if unicode.IsLetter(r) || r == '_' {
				t += string(r)
				continue
			}
		} else {
			if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
				t += string(r)
				continue
			}
		}
		t += "_"
	}
	if t == "_" {
		t = "x"
	}
	return t
}

func (g *Generator) getArgNames(m *model.Method) []string {
	argNames := make([]string, len(m.In))
	for i, p := range m.In {
		name := p.Name
		if name == "" || name == "_" {
			name = fmt.Sprintf("arg%d", i)
		}
		argNames[i] = name
	}
	if m.Variadic != nil {
		name := m.Variadic.Name
		if name == "" {
			name = fmt.Sprintf("arg%d", len(m.In))
		}
		argNames = append(argNames, name)
	}
	return argNames
}

func (g *Generator) getArgTypes(m *model.Method, pkgOverride string) []string {
	argTypes := make([]string, len(m.In))
	for i, p := range m.In {
		argTypes[i] = p.Type.String(g.packageMap, pkgOverride)
	}
	if m.Variadic != nil {
		argTypes = append(argTypes, "..."+m.Variadic.Type.String(g.packageMap, pkgOverride))
	}
	return argTypes
}

func makeArgString(argNames, argTypes []string) string {
	args := make([]string, len(argNames))
	for i, name := range argNames {
		// specify the type only once for consecutive args of the same type
		if i+1 < len(argTypes) && argTypes[i] == argTypes[i+1] {
			args[i] = name
		} else {
			args[i] = name + " " + argTypes[i]
		}
	}
	return strings.Join(args, ", ")
}

// createPackageMap returns a map of import path to package name
// for specified importPaths.
func createPackageMap(importPaths []string) (map[string]string, error) {
	var pkg struct {
		Name       string
		ImportPath string
	}
	pkgMap := make(map[string]string)
	b := bytes.NewBuffer(nil)
	args := []string{"list", "-json"}
	args = append(args, importPaths...)
	cmd := exec.Command("go", args...)
	cmd.Stdout = b
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("run `go list -json`: %w", err)
	}
	dec := json.NewDecoder(b)
	for dec.More() {
		err := dec.Decode(&pkg)
		if err != nil {
			return nil, fmt.Errorf("decode `go list -json` output: %w", err)
		}
		pkgMap[pkg.ImportPath] = pkg.Name
	}
	return pkgMap, nil
}
