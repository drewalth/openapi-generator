# OpenAPI Specification Generator

This CLI tool generates an [OpenAPI Specification](https://swagger.io/specification/) from data models in a software project that uses an Object Relational
Mapper (ORM). The tool currently supports Laravel Eloquent (PHP), Sequelize (JavaScript), and Spring Boot (Java) ORMs.

## Why?

The OpenAPI Specification is a powerful tool for documenting REST APIs. However, it can be tedious to write the
specification manually. This tool aims to jumpstart the process by generating a basic starting point for your 
API documentation journey.

Once you have an OpenAPI specification, you can use it to generate documentation, client SDKs, and server stubs using
the [OpenAPI Generator](https://openapi-generator.tech/).

## Prerequisites

- Go 1.13 or higher

## Installation

1. Clone this repository:

```bash
git clone https://github.com/drewalth/openapi-generator.git
```

2. Navigate to the cloned repository:

```bash
cd openapi-generator
```

3. Run setup script. This will install project dependencies and configure git hooks.

```bash
make setup
```

## Usage

To run the script, execute the following command in your terminal:

```bash
go run *.go --path /path/to/models --orm ORMName
```

- Replace `/path/to/models` with the path to the directory containing the models of your project.
- Replace `ORMName` with one of the supported ORMs (`Eloquent`, `Sequelize`, or `SpringBoot`).

This will generate an OpenAPI specification in YAML format in a directory called `openapi` in the current working
directory.

## How It Works

1. The tool takes in the path to the directory containing your data models and the name of the ORM you are using as
   command line arguments.
2. It parses the data models in the specified directory. Parsing involves reading the files and extracting information
   such as class names and properties. Note that the parsing logic in this example is rudimentary and serves for
   illustration purposes.
3. The extracted information is used to generate a basic OpenAPI Specification in YAML format with CRUD endpoints (List,
   Create, Retrieve by ID, Update, Delete) for each model.
4. The generated specification is saved in an "openapi" directory in the current working directory.

## Limitations and Improvements

- The parsing logic is very basic and might not cover all cases.
- The type mappings from model properties to OpenAPI types are not handled; need to add logic to map the types
  correctly.
- The generated OpenAPI specification is very basic; in a real-world scenario, the specification could be much more
  detailed and complex.
- Error handling should be improved for production usage.
- This is just a starting point and you will need to build upon this to handle various cases (relationships, validation
  rules, etc.). This will only handle basic property types, and won't account for complex structures and relationships.
- Properly parsing PHP source code is a difficult task. Might want to look into using a PHP parser that can be called
  from Go.
- Consider using a more robust method for parsing Sequelize models such as writing a script in JavaScript or TypeScript
  that uses the abstract syntax tree (AST) for parsing.

## Contributing

Contributions to improve the functionality and robustness of this tool are welcome. Please feel free to open issues or
submit pull requests.

## License

MIT License
