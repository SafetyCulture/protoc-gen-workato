// protoc-gen-workato is used to generate supporting files for https://github.com/envoyproxy/workato.
//
// It is a protoc plugin, and can be invoked by passing `--workato_out` and `--workato_opt` arguments to protoc.
//
// Example: generate workato configuration files for Envoy:
//
//     protoc --workato_out=. --doc_opt=workato/config.yaml protos/*.proto
//
// For more details, check out the README at https://github.com/pseudomuto/protoc-gen-doc
package main

import (
	"github.com/pseudomuto/protokit"

	"log"
	"os"

	genworkato "github.com/SafetyCulture/protoc-gen-workato"

	_ "github.com/SafetyCulture/protoc-gen-workato/extensions/protoc_gen_openapiv2" // imported for side effects
	_ "github.com/SafetyCulture/protoc-gen-workato/extensions/protoc_gen_workato"   // imported for side effects
	_ "github.com/pseudomuto/protoc-gen-doc/extensions/google_api_http"             // imported for side effects
)

func main() {
	if flags := ParseFlags(os.Stdout, os.Args); HandleFlags(flags) {
		os.Exit(flags.Code())
	}

	if err := protokit.RunPlugin(new(genworkato.Plugin)); err != nil {
		log.Fatal(err)
	}
}

// HandleFlags checks if there's a match and returns true if it was "handled"
func HandleFlags(f *Flags) bool {
	if !f.HasMatch() {
		return false
	}

	if f.ShowHelp() {
		f.PrintHelp()
	}

	if f.ShowVersion() {
		f.PrintVersion()
	}

	return true
}
