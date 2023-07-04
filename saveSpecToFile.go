package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// saveSpecToFile saves the specification to a yaml file
func saveSpecToFile(spec string) error {
	// Save the specification in 'openapi' directory
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	outputDir := filepath.Join(cwd, "openapi")
	os.MkdirAll(outputDir, os.ModePerm)

	filePath := filepath.Join(outputDir, "openapi.yaml")
	return ioutil.WriteFile(filePath, []byte(spec), 0644)
}
