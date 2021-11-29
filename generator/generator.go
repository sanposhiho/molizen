package generator

import (
	"bytes"
	"fmt"
	"log"

	toolsimports "golang.org/x/tools/imports"
)

type Generator struct {
	buf                       bytes.Buffer
	indent                    string
	Destination               string // may be empty
	srcPackage, srcInterfaces string // may be empty
	CopyrightHeader           string

	packageMap map[string]string // map from import path to package name
}

// Output returns the generator's output, formatted in the standard Go style.
func (g *Generator) Output() []byte {
	src, err := toolsimports.Process(g.Destination, g.buf.Bytes(), nil)
	if err != nil {
		log.Fatalf("Failed to format generated source code: %s\n%s", err, g.buf.String())
	}
	return src
}

func (g *Generator) p(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, g.indent+format+"\n", args...)
}

func (g *Generator) in() {
	g.indent += "\t"
}

func (g *Generator) out() {
	if len(g.indent) > 0 {
		g.indent = g.indent[0 : len(g.indent)-1]
	}
}
