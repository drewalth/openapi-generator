package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

// parseSequelizeModels parses Sequelize models and returns a map of model names and their properties
func parseSpringBootModels(path string) (map[string]map[string]string, error) {
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

		if !strings.Contains(string(content), "@Entity") {
			continue
		}

		reClass := regexp.MustCompile(`public\s+class\s+(\w+)`)
		matchClass := reClass.FindStringSubmatch(string(content))

		if len(matchClass) < 2 {
			continue
		}

		className := matchClass[1]

		reProps := regexp.MustCompile(`private\s+(\w+)\s+(\w+);`)
		matchesProps := reProps.FindAllStringSubmatch(string(content), -1)

		properties := make(map[string]string)
		for _, match := range matchesProps {
			propType := match[1]
			propName := match[2]

			// Map Java data type to OpenAPI data type
			switch propType {
			case "String":
				properties[propName] = "string"
			case "int", "Integer", "long", "Long":
				properties[propName] = "integer"
			case "float", "double", "Float", "Double":
				properties[propName] = "number"
			case "boolean", "Boolean":
				properties[propName] = "boolean"
			case "Date":
				properties[propName] = "string" // You can also add "format": "date-time" in OpenAPI spec
			default:
				properties[propName] = "string" // Default to string if no match
			}
		}
		models[className] = properties
	}

	return models, nil
}
