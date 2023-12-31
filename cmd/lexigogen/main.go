package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/MysteriousPotato/lexigo/internal/generator"
)

var (
	helpFlag    = flag.Bool("h", false, "Help")
	outPathFlag = flag.String("o", "", "Output path")
	srcPathFlag = flag.String("p", "", "Source path (Directory containing locale files)")
	pkgNameFlag = flag.String("pkg", "", "Package name")
	varNameFlag = flag.String("var", "Locales", "Name of the variable containing the locales")
)

func main() {
	flag.Parse()

	if *helpFlag {
		flag.PrintDefaults()
		return
	}

	var errs []error
	if *outPathFlag == "" {
		errs = append(errs, fmt.Errorf("missing -o argument"))
	}
	if *pkgNameFlag == "" {
		errs = append(errs, fmt.Errorf("missing -pkg argument"))
	}
	if *varNameFlag == "" {
		errs = append(errs, fmt.Errorf("argument '-var' must not be empty"))
	}
	if *srcPathFlag == "" {
		errs = append(errs, fmt.Errorf("argument '-var' must not be empty"))
	}

	if errs != nil {
		log.Println(errors.Join(errs...))
		flag.PrintDefaults()
		os.Exit(1)
	}

	out, err := os.Create(*outPathFlag)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	defer func() {
		if reason := recover(); reason != nil {
			reasonErr, ok := reason.(error)
			if !ok {
				reasonErr = fmt.Errorf("%v", reasonErr)
			}

			if err := os.Remove(out.Name()); err != nil {
				log.Fatal(errors.Join(reasonErr, err))
			}
			if err := out.Close(); err != nil {
				log.Fatal(errors.Join(reasonErr, err))
			}
			log.Fatal(reason, string(debug.Stack()))
		}
	}()

	g, err := generator.NewGenerator(generator.Params{
		SrcPath: *srcPathFlag,
		PkgName: *pkgNameFlag,
	})
	if err != nil {
		panic(err)
	}

	b, err := g.Exec(*varNameFlag)
	if err != nil {
		panic(err)
	}

	if _, err := out.Write(b); err != nil {
		panic(err)
	}
}
