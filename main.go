package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var flagPrompt = flag.Bool("p", false, "prompt for each file")

func main() {
	flag.Parse()
	from := []byte(flag.Arg(0))
	to := []byte(flag.Arg(1))

	if err := s(from, to, *flagPrompt); err != nil {
		log.Fatal(err)
	}
}

func s(from, to []byte, optPrompt bool) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("s: error getting working directory: %v", err)
	}
	if err := filepath.Walk(dir, sub(from, to, optPrompt)); err != nil {
		return fmt.Errorf("s: error walking filesystem %v", err)
	}
	return nil
}

func sub(from, to []byte, optPrompt bool) filepath.WalkFunc {
	f := func(path string, info os.FileInfo, err error) error {
		n := info.Name()
		ignore := n == "" || n[0] == '.' || n[0] == '_' || n == "vendor"

		if info.IsDir() {
			if ignore {
				return filepath.SkipDir
			}
			return nil
		}
		if !ignore {
			if !optPrompt || prompt(path) {
				return subf(path, from, to)
			}
		}
		return nil
	}
	return filepath.WalkFunc(f)
}

func prompt(name string) bool {
	fmt.Printf("update file: %s (Y/N) ", name)
	var confirm string
	fmt.Scanln(&confirm)
	switch confirm {
	case "y", "Y", "yes", "Yes", "YES":
		return true
	}
	return false
}

func subf(path string, old, new []byte) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("subf: error opening file: %v", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("subf: error reading file: %v", err)
	}
	bb := bytes.Replace(b, old, new, -1)
	if err := ioutil.WriteFile(path, bb, 0644); err != nil {
		return fmt.Errorf("subf: error writing file: %v", err)
	}
	return nil
}
