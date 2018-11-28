package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func resolveEnvironment(variableSpec string) map[string]interface{} {
	// split and trim the input variables
	specs := strings.Split(variableSpec, ",")
	for i, s := range specs {
		specs[i] = strings.Trim(s, " ")
	}

	// filter the environment by the matching variables
	m := make(map[string]interface{})
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if envMatchesPatterns(pair[0], specs) {
			m[pair[0]] = pair[1]
		}
	}

	// and return the filtered array
	return m
}

// checks if a variables matches a pattern
func envMatchesPatterns(env string, patterns []string) bool {
	for _, p := range patterns {
		if p != "" {
			// if the pattern matches, return true
			matched, err := filepath.Match(p, env)
			if err == nil && matched {
				return true
			}
		}
	}
	return false
}

func main() {
	// read all the arguments
	bindAddress := flag.String("bind", "0.0.0.0:3000", "address to bind to")
	envPath := flag.String("path", "/", "server path to serve the resulting json on")
	variables := flag.String("vars", "", "comma-seperated list of variables and globs to use")

	// parse them
	flag.Parse()

	// filter the environment and marshal into json
	exposed := resolveEnvironment(*variables)
	output, err := json.Marshal(exposed)
	if err != nil {
		log.Fatal(err)
	}

	// add a handler
	http.HandleFunc(*envPath, func(w http.ResponseWriter, r *http.Request) {
		callback := r.FormValue("callback")
		if callback != "" {
			w.Header().Set("Content-Type", "application/javascript")
			fmt.Fprintf(w, "%s(%s)", callback, output)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(output)
		}
	})

	// pretty print what we are going to do
	prettyOutput, err := json.MarshalIndent(exposed, "", "    ")
	if err == nil {
		fmt.Printf("Exposing %s ", prettyOutput)
	}
	fmt.Printf("on '%s', at path '%s'\n", *bindAddress, *envPath)

	// and bind
	http.ListenAndServe(*bindAddress, nil)
}
