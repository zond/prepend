package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	glob := flag.String("glob", "", "What files to prepend to.")
	prefix := flag.String("prefix", "", "What to ensure they are prepended with.")
	dryrun := flag.Bool("dryrun", false, "Whether to only print the files that would have been prepended.")

	if *glob == "" || *prefix == "" {
		flag.Usage()
		os.Exit(1)
	}

	paths, err := filepath.Glob(*glob)
	if err != nil {
		panic(err)
	}
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		if !strings.HasPrefix(string(data), *prefix) {
			if *dryrun {
				fmt.Println(path)
			} else {
				data = append([]byte(*prefix), data...)
				if err := os.WriteFile(path, data); err != nil {
					panic(err)
				}
			}
		}
	}
}
