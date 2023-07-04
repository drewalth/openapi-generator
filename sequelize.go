package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
)

// parseSequelizeModels parses Sequelize models and returns a map of model names and their properties
func parseSequelizeModels(path string) (map[string]map[string]string, error) {
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

		reModelName := regexp.MustCompile(`sequelize\.define\(['"](\w+)['"]`)
		matchModelName := reModelName.FindStringSubmatch(string(content))
		if len(matchModelName) < 2 {
			continue
		}

		modelName := matchModelName[1]

		re := regexp.MustCompile(`(\w+):\s*\{\s*type:\s*Sequelize\.(\w+)(\([^)]*\))?\s*`)
		matches := re.FindAllStringSubmatch(string(content), -1)

		properties := make(map[string]string)
		for _, match := range matches {
			propName := match[1]
			dataType := match[2]

			// Map Sequelize data type to OpenAPI data type
			switch dataType {
			case "STRING", "TEXT":
				properties[propName] = "string"
			case "INTEGER", "BIGINT":
				properties[propName] = "integer"
			case "FLOAT", "REAL", "DOUBLE":
				properties[propName] = "number"
			case "BOOLEAN":
				properties[propName] = "boolean"
			case "DATE":
				properties[propName] = "string" // You can also add "format": "date-time" in OpenAPI spec
			default:
				properties[propName] = "string" // Default to string if no match
			}
		}
		models[modelName] = properties
	}

	return models, nil
}
