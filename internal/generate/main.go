package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/MOZGIII/sdg-go"
)

func perform() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("must pass 2 arguments")
	}

	source := os.Args[1]
	destination := os.Args[2]

	rows, err := readFile(source)
	if err != nil {
		return err
	}
	return sdg.WriteTempalte(sdg.WriteTempalteArgs{
		DstPath: destination,
		Rows:    rows,
		Params:  params,
	})
}

func main() {
	if err := perform(); err != nil {
		log.Fatal(err)
	}
}

func readFile(path string) ([][]string, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer f.Close() // nolint: errcheck
	r := csv.NewReader(f)
	r.Comma = '\t'
	return r.ReadAll()
}

type templateParams struct {
	Preface string
	Var     string
	Type    string
	ValFn   func(row []string) string
}

var params = &templateParams{
	Preface: `type ZIPCode struct {
            PostalCode string
            PlaceName string
            State string
            StateUSPS string
            County string
            CountyFIPS string
        }`,
	Var:  "ZIPCodes",
	Type: "[]ZIPCode",
	ValFn: func(row []string) string {
		return fmt.Sprintf(`ZIPCode{%q, %q, %q, %q, %q, %q}`,
			row[1], row[2], row[3], row[4], row[5], row[6])
	},
}
