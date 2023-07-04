package main

import (
	"fmt"
	"strings"
)

func generateOpenAPISpecContent(models map[string]map[string]string) string {
	spec := `openapi: 3.1.0
info:
  title: Auto-generated API Spec
  version: 1.0.0
paths:
`

	for modelName, properties := range models {
		lowercaseModel := strings.ToLower(modelName)
		spec += fmt.Sprintf(`  /%s:
    get:
      summary: List all %s
      responses:
        '200':
          description: Successful response
    post:
      summary: Create a new %s
      requestBody:
        content:
          application/json:
            schema:
              type: object
              properties:
`, lowercaseModel, modelName, modelName)

		for propName, propType := range properties {
			spec += fmt.Sprintf(`                %s:
                  type: %s
`, propName, propType) // Map the property types accordingly
		}

		spec += `      responses:
        '201':
          description: Created
`

		spec += fmt.Sprintf(`  /%s/{id}:
    get:
      summary: Get a %s by ID
      responses:
        '200':
          description: Successful response
    put:
      summary: Update a %s
      responses:
        '200':
          description: Successful response
    delete:
      summary: Delete a %s
      responses:
        '204':
          description: No content
`, lowercaseModel, modelName, modelName, modelName)
	}

	return spec
}
