package main

import (
	"log"
	"path/filepath"
)

func main() {
	folder := "C:\\testing\\*-kvp.csv"
	log.Printf("Scanning folder %s for input files\n", folder)
	matches, err := filepath.Glob(folder)
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range matches {
		log.Printf("Found file: %s", m)
	}
}
