# JerseyHUB E-Commerce Backend Rest API

JerseyHUB E-Commerce Backend Rest API is a feature-rich backend solution for E-commerce applications developed using Golang with the Gin web framework. This API is designed to efficiently handle routing and HTTP requests while following best practices in code architecture and dependency management.

## Key Features

- **Clean Code Architecture**: The project follows clean code architecture principles, making it maintainable and scalable.
- **Dependency Injection**: Utilizes the Dependency Injection design pattern for flexible component integration.
- **Compile-Time Dependency Injection**: Dependencies are managed using Wire for compile-time injection.
- **Database**: Leverages PostgreSQL for efficient and relational data storage.
- **AWS Integration**: Integrates with AWS S3 for cloud-based storage solutions.
- **E-Commerce Features**: Implements a wide range of features commonly found in e-commerce applications, including cart management, wishlist, wallet, offers, and coupon management.
- **Code Quality**: Includes unit tests to maintain code quality and implements Continuous Integration/Continuous Deployment (CI/CD) with GitHub Actions.

## Deployment

The application is hosted on AWS EC2 and is served using Nginx, ensuring reliability and scalability.

- **Containerization**: The project is containerized using Docker, and successful deployment and scaling are tested on Kubernetes.

## API Documentation

For interactive API documentation, Swagger is implemented. You can explore and test the API endpoints in real-time.

## Security

Security is a top priority for the project:

- **OTP Verification**: Twilio API is integrated for OTP verification.
- **Payment Integration**: Razorpay API is used for payment processing.
- **Refresh Tokens**: Enhances security and extends user sessions using refresh tokens.

## Mobile Application

In collaboration with a team of developers, a mobile application is being built using the Flutter framework to complement this backend API.

## Getting Started

To run the project locally, you can follow these steps:

1. Clone the repository.
2. Set up your environment with the required dependencies, including Golang, PostgreSQL, Docker, and Wire.
3. Configure your environment variables (e.g., database credentials, AWS keys, Twilio credentials).
4. Build and run the project.

## License

This project is licensed under the [LICENSE_NAME] license - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Wire for Dependency Injection](https://github.com/google/wire)
- [PostgreSQL](https://www.postgresql.org/)
- [AWS S3](https://aws.amazon.com/s3/)
- [Swagger API Documentation](https://swagger.io/)
- [Twilio](https://www.twilio.com/)
- [Razorpay](https://razorpay.com/)
- [Flutter Framework](https://flutter.dev/)


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
mock                           Generate Mock Files using mockgen
help                           Display this help screen
```

# Environment Variables

Before running the project, you need to set the following environment variables with your corresponding values:

## PostgreSQL

- `DB_HOST`: Database host
- `DB_NAME`: Database name
- `DB_USER`: Database user
- `DB_PORT`: Database port
- `DB_PASSWORD`: Database password

## Twilio

- `DB_AUTHTOKEN`: Twilio authentication token
- `DB_ACCOUNTSID`: Twilio account SID
- `DB_SERVICESID`: Twilio services ID

## AWS

- `AWS_REGION`: AWS region
- `AWS_ACCESS_KEY_ID`: AWS access key ID
- `AWS_SECRET_ACCESS_KEY`: AWS secret access key

Make sure to provide the appropriate values for these environment variables to configure the project correctly.
