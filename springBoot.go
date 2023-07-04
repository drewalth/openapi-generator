package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

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
			properties[propName] = propType
		}
		models[className] = properties
	}

	return models, nil
}
