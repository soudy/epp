package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/flosch/pongo2"
)

var (
	Version   string
	GitCommit string

	output  = flag.String("o", "", "output file")
	version = flag.Bool("version", false, "print epp version")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stderr, "epp %s (%s)\n", Version, GitCommit)
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "error: an input file is required")
		os.Exit(1)
	}

	pongo2.RegisterFilter("b64enc", filterBase64Encode)

	fileContents, err := readInput(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %s\n", err)
		os.Exit(1)
	}

	out, err := parse(fileContents)
	if err != nil {
		fmt.Fprintf(os.Stderr, "templating error: %s\n", err)
		os.Exit(1)
	}

	if *output == "" {
		fmt.Printf(string(out))
		return
	}

	err = ioutil.WriteFile(*output, out, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "IO error: %s\n", err)
		os.Exit(1)
	}
}

func parse(input []byte) ([]byte, error) {
	tpl, err := pongo2.FromString(string(input))
	if err != nil {
		return nil, err
	}

	context := environToContext()
	return tpl.ExecuteBytes(context)
}

func readInput(input string) ([]byte, error) {
	if inputFile := flag.Arg(0); inputFile == "-" {
		return ioutil.ReadAll(os.Stdin)
	}

	return ioutil.ReadFile(input)
}

func environToContext() pongo2.Context {
	ctx := pongo2.Context{}

	for _, env := range os.Environ() {
		variable := strings.SplitN(env, "=", 2)
		key, value := variable[0], variable[1]

		ctx[key] = value
	}

	return ctx
}

func filterBase64Encode(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(base64.StdEncoding.EncodeToString([]byte(in.String()))), nil
}
