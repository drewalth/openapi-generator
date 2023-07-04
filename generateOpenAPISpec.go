package main

import (
	"fmt"
)

// generateOpenAPISpec generates the OpenAPI Specification
func generateOpenAPISpec(path string, orm string) error {
	var models map[string]map[string]string
	var err error
	switch orm {
	case "Eloquent":
		models, err = parseEloquentModels(path)
	case "Sequelize":
		models, err = parseSequelizeModels(path)
	case "SpringBoot":
		models, err = parseSpringBootModels(path)
	default:
		return fmt.Errorf("Unsupported ORM: %s", orm)
	}

	if err != nil {
		return err
	}

	// Generate OpenAPI Specification
	spec := generateOpenAPISpecContent(models)

	// Save to file
	err = saveSpecToFile(spec)
	return err
}
