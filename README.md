# Clean Architecture using Golang with Gin framework

## Template Structure

- [Gin](github.com/gin-gonic/gin) is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
- [JWT](github.com/golang-jwt/jwt) A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.
- [GORM](https://gorm.io/index.html) with [PostgresSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)The fantastic ORM library for Golang aims to be developer friendly.
- [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.
- [Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.
- [swag](https://github.com/swaggo/swag) converts Go annotations to Swagger Documentation 2.0 with [gin-swagger](https://github.com/swaggo/gin-swagger) and [swaggo files](github.com/swaggo/files)

## Using `go-gin-clean-arch` project

To use `go-gin-clean-arch` project, follow these steps:

```bash
# Navigate into the project
cd ./go-gin-clean-arch

# Install dependencies
make deps

# Generate wire_gen.go for dependency injection
# Please make sure you exported the env for GOPATH
make wire

# Run the project in Development Mode
make run
```

Additional commands:


```bash
➔ make help
build                          Compile the code, build Executable File
run                            Start application
test                           Run tests
test-coverage                  Run tests and generate coverage file
deps                           Install dependencies
deps-cleancache                Clear cache in Go module
wire                           Generate wire_gen.go
swag                           Generate swagger docs
help                           Display this help screen
```

# Environment variables needed:
# fill with your corresponding values:

- postgres
--------
DB_HOST=?
DB_NAME=?
DB_USER=?
DB_PORT=?
DB_PASSWORD=?

- twilio
--------
DB_AUTHTOKEN=?
DB_ACCOUNTSID=?
DB_SERVICESID=?

- AWS
--------
AWS_REGION=?
AWS_ACCESS_KEY_ID=?
AWS_SECRET_ACCESS_KEY=?

