package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"strings"

	"github.com/spf13/pflag"
)

// the selected output Format
var outputFormat string

const (
	outputFormatPlain = "plain"
	outputFormatJSON  = "json"
	outputFormatXML   = "xml"
	outputFormatYAAML = "yaml"
)

var outputFormats = []string{outputFormatPlain, outputFormatJSON, outputFormatXML, outputFormatYAAML}

func addFormatFlag(flagSet *pflag.FlagSet) {
	flagSet.StringVarP(&outputFormat, "format", "f", outputFormatPlain, "The output format ("+strings.Join(outputFormats, "/")+")")
}

func checkOutputFormat() error {
	for i := range outputFormats {
		if outputFormats[i] == outputFormat {
			return nil
		}
	}

	return errInvalidFormat(outputFormat)
}

type outputWriter func(w io.Writer)

func writeOutput(w io.Writer, v interface{}, plainWriter outputWriter) error {
	switch outputFormat {
	case outputFormatPlain:
		plainWriter(w)
		return nil
	case outputFormatJSON:
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		encoder.SetEscapeHTML(false)
		return encoder.Encode(v)
	case outputFormatXML:
		return xml.NewEncoder(w).Encode(v)
	case outputFormatYAAML:
		return yaml.NewEncoder(w).Encode(v)
	default:
		return errInvalidFormat(outputFormat)
	}
}

type errInvalidFormat string

func (err errInvalidFormat) Error() string {
	return fmt.Sprintf("invalid output format: %s", outputFormat)
}
