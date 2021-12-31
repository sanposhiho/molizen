package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/sanposhiho/mock/mockgen/model"
	"github.com/sanposhiho/mock/mockgen/parse"
	"github.com/sanposhiho/mock/mockgen/util"
	"github.com/sanposhiho/molizen/generator"
)

func main() {
	var (
		source        = flag.String("source", "", "Input Go source file.")
		destination   = flag.String("destination", "", "Output file; defaults to stdout.")
		packageOut    = flag.String("package", "", "Package of the generated code; defaults to the package of the input with a 'actor_' prefix.")
		copyrightFile = flag.String("copyright_file", "", "Copyright file used to add copyright header")
	)

	flag.Parse()

	var pkg *model.Package
	var err error
	pkg, err = parse.SourceMode(*source)
	if err != nil {
		log.Fatalf("Loading input failed: %v", err)
	}

	dst := os.Stdout
	if len(*destination) > 0 {
		if err := os.MkdirAll(filepath.Dir(*destination), os.ModePerm); err != nil {
			log.Fatalf("Unable to create directory: %v", err)
		}
		f, err := os.Create(*destination)
		if err != nil {
			log.Fatalf("Failed opening destination file: %v", err)
		}
		defer f.Close()
		dst = f
	}

	outputPackageName := *packageOut
	if outputPackageName == "" {
		// pkg.Name in reflect mode is the base name of the import path,
		// which might have characters that are illegal to have in package names.
		outputPackageName = "actor_" + generator.Sanitize(pkg.Name)
	}

	outputPackagePath := ""
	if *destination != "" {
		dstPath, err := filepath.Abs(filepath.Dir(*destination))
		if err == nil {
			pkgPath, err := util.ParsePackageImport(dstPath)
			if err == nil {
				outputPackagePath = pkgPath
			} else {
				log.Println("Unable to infer -self_package from destination file path:", err)
			}
		} else {
			log.Println("Unable to determine destination file path:", err)
		}
	}

	g := new(generator.Generator)
	g.Destination = *destination

	if *copyrightFile != "" {
		header, err := ioutil.ReadFile(*copyrightFile)
		if err != nil {
			log.Fatalf("Failed reading copyright file: %v", err)
		}

		g.CopyrightHeader = string(header)
	}

	if err := g.Generate(pkg, outputPackageName, outputPackagePath); err != nil {
		log.Fatalf("Failed generating actor: %v", err)
	}

	if _, err := dst.Write(g.Output()); err != nil {
		log.Fatalf("Failed writing to destination: %v", err)
	}
}
