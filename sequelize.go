package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
)

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

		re := regexp.MustCompile(`(\w+):\s*Sequelize\.\w+`)
		matches := re.FindAllStringSubmatch(string(content), -1)

		properties := make(map[string]string)
		for _, match := range matches {
			propName := match[1]
			properties[propName] = "string" // Defaulting to string as a placeholder
		}
		models[file.Name()] = properties
	}

	return models, nil
}
