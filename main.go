package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go ...<extensions>\nExample: go run main.go .txt .bat .py .sh .vb .vbs .ps1 .ps2 .csv .ini .env")
		os.Exit(0)
	}

	for i, v := range os.Args {
		os.Args[i] = strings.ToLower(v)
	}

	exts := os.Args[1:]

	terms := []string{
		"password",
		"credential",
	}

	scanFiles(exts, terms)
	recursiveWalk(terms)
}

func scanFiles(exts []string, terms []string) {

	output, _ := os.Create("scan.output")
	defer output.Close()

	filepath.Walk(".", func(path string, info os.FileInfo, _ error) error {
		for _, e := range exts {
			if filepath.Ext(strings.ToLower(path)) == e {

				line := 1

				file, _ := os.Open(path)
				defer file.Close()

				scan := bufio.NewScanner(file)

				for scan.Scan() {
					for _, t := range terms {
						if strings.Contains(strings.ToLower(scan.Text()), t) {
							fmt.Fprintln(output, "["+strconv.Itoa(line)+"] - "+path+" - "+scan.Text())
						}
					}
					line++
				}
			}
		}
		return nil
	})
}

func recursiveWalk(terms []string) {

	output, _ := os.Create("walk.output")
	defer output.Close()

	filepath.Walk(".", func(path string, _ os.FileInfo, _ error) error {
		for _, t := range terms {
			if strings.Contains(strings.ToLower(path), t) {
				fmt.Fprintln(output, path)
			}
		}
		return nil
	})
}
