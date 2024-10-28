package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
)

func main() {
	files, err := filepath.Glob("*.mid.hex")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err := convert(file, file[:len(file)-4])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func convert(src, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	data = rxRemoveComments.ReplaceAll(data, nil)
	data = rxRemoveWhitespace.ReplaceAll(data, nil)

	data, err = hex.DecodeString(string(data))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dst, data, 0644)
}

var (
	rxRemoveComments   = regexp.MustCompile(`//.*`)
	rxRemoveWhitespace = regexp.MustCompile(`(?m)\s+`)
)
