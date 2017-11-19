package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shoenig/mockery/libmockery"
)

type flags struct {
	version bool
	iface   string
	stdout  bool
	comment string
	pkgname string
}

func main() {
	config := parseFlags(os.Args)

	if config.version {
		fmt.Println("mockery " + libmockery.Version)
		return
	}

	if config.iface == "" {
		fmt.Println("-interface is required")
		os.Exit(1)
	}

	if config.pkgname == "" {
		fmt.Println("-package is required")
		os.Exit(1)
	}

	visitor := &libmockery.GeneratorVisitor{
		Comment:           config.comment,
		OutputProvider:    outputProvider(config),
		OutputPackageName: config.pkgname,
	}

	walker := libmockery.Walker{
		BaseDir:   ".",
		Interface: config.iface,
	}

	generated := walker.Walk(visitor)

	if !generated {
		fmt.Printf("Unable to find interface %q in any go files under this path\n", config.iface)
		os.Exit(1)
	}
}

func outputProvider(config flags) libmockery.OutputStreamProvider {
	if config.stdout {
		return &libmockery.StdoutStreamProvider{}
	}
	return &libmockery.FileOutputStreamProvider{
		BaseDir: config.pkgname,
	}
}

func parseFlags(args []string) flags {
	config := flags{}

	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.BoolVar(&config.version, "version", false, "print the version of this mockery executable")
	flagSet.StringVar(&config.iface, "interface", "", "name or matching regular expression of interface to generate mock for")
	flagSet.BoolVar(&config.stdout, "stdout", false, "print the generated mock to stdout instead of writing to disk")
	flagSet.StringVar(&config.comment, "comment", "", "comment to insert into prologue of each generated file")
	flagSet.StringVar(&config.pkgname, "package", "", "package name containing generated mocks")

	flagSet.Parse(args[1:])

	return config
}