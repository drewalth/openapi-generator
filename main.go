package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "openapi-generator",
		Usage: "Generate OpenAPI Specification from data models using ORM",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "path",
				Aliases:  []string{"p"},
				Usage:    "Path to the directory where the models are defined",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "orm",
				Aliases:  []string{"o"},
				Usage:    "The ORM being used (e.g. Eloquent, Sequelize)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			path := replaceTildeWithHome(c.String("path"))

			return generateOpenAPISpec(path, c.String("orm"))
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
