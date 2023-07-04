package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
)

// parseEloquentModels parses Eloquent models and returns a map of model names and their properties
func parseEloquentModels(path string) (map[string]map[string]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	models := make(map[string]map[string]string)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		content, err := ioutil.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			return nil, err
		}

		reClass := regexp.MustCompile(`class\s+(\w+)\s+extends`)
		matchClass := reClass.FindStringSubmatch(string(content))

		if len(matchClass) < 2 {
			continue
		}

		className := matchClass[1]

		reProps := regexp.MustCompile(`\*\s+@var\s+(\w+)\s+\$([\w_]+)`)
		matchesProps := reProps.FindAllStringSubmatch(string(content), -1)

		properties := make(map[string]string)
		for _, match := range matchesProps {
			propType := match[1]
			propName := match[2]

			// Map Eloquent data type to OpenAPI data type
			switch propType {
			case "string":
				properties[propName] = "string"
			case "int":
				properties[propName] = "integer"
			case "float":
				properties[propName] = "number"
			case "bool":
				properties[propName] = "boolean"
			case "Carbon": // Carbon is a date-time library in Laravel
				properties[propName] = "string" // You can also add "format": "date-time" in OpenAPI spec
			default:
				properties[propName] = "string" // Default to string if no match
			}
		}
		models[className] = properties
	}

	return models, nil
}
