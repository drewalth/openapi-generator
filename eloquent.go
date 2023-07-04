package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
)

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

		// Very basic parsing example (not production ready)
		re := regexp.MustCompile(`@property\s+(\w+)\s+\$(\w+)`)
		matches := re.FindAllStringSubmatch(string(content), -1)

		properties := make(map[string]string)
		for _, match := range matches {
			propType := match[1]
			propName := match[2]

			properties[propName] = propType
		}
		models[file.Name()] = properties
	}

	return models, nil
}
