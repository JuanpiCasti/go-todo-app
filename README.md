# go-todo-app

This is a project template for a simple todo app using Go with the Gin framework and a PostgreSQL database for the persistence.

The project is structured following a layered architecture, aiming to provide a clean separation of concerns and allowing easy testing and maintenance.

The layers, from top to bottom, are:

- **Middleware**: Contains the middleware functions that are applied to the routes, executing before the handlers.
- **Handler**: Handles HTTP requests and responses. Interactions are preferably made with DTOs.
- **Service**: Contains the business logic.
- **Model**: Contains the domain entities.
- **Repository**: Handles the data access.

These layers are all found under the `/app` directory. Routing is defined in the `/router/router.go` file.

Initialization scripts for the database are found under the `/scripts/db` directory.

An example openapi specification is provided under the `/docs` directory, describing the example endpoints.

## Running the app

The app is containerized using Docker.

To run the app locally with the PostgreSQL database, just run `docker-compose up`.

If you have Air installed, you can run the app with `air` for hot reloading. Remember to start the database and run the initialization scripts before, and to set up the environment variables.

## Environment variables

The following environment variables are required:

```
DATABASE_USER=root
DATABASE_PASSWORD=root
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_SSL_MODE=disable
DATABASE_NAME=todo_app
SSLROOTCERT=/
SERVER_PORT=8080
GO_MODE=debug
ALLOWED_ORIGINS=*
JWT_SECRET=9e7e07a6d1c7a1e8b45a3ab5600c5dde307478d847b708d02698b7c0c2373367
TOKEN_DURATION_MINUTES=5
TRUSTED_PROXY_IPS=127.0.0.1,192.168.1.1
```

These are just examples, you should set them according to your environment.

## Most important used packages

- **Gin**: Web framework.
- **SQLX**: Extension for the standard `database/sql` package.
- **Zerolog**: Logging library.

## Used and included middleware

- **CORS**: To allow requests from the specified origins in the environment variables, separated by commas.
- **Logger**: Logging of requests/responses with metadata like latency and status code.
- **Auth**: Middleware for JWT authentication. It checks the Authorization header for a valid JWT token and sets claims in the context.
- **Role**: Middleware for role-based authorization. It checks the claims set by the Auth middleware and compares them with the required role for the route.

## Basic entities

- **User**: Represents a user of the application. It has an ID, a username, a password, a role, and an "active" field.
- **Role**: Two roles are defined: `admin` and `user`.
- **Todo**: Example domain entity.

## Testing

`go get github.com/stretchr/testify`
`go get github.com/golang/mock/mockgen`
`go generate ./...` for generating mocks.

Mocked interfaces should have their file annotated with `//go:generate go run github.com/golang/mock/mockgen -destination=mock_todo_repository.go -package=repository github.com/juanpicasti/go-todo-app/internal/app/repository TodoRepository`

### Run unit tests
`go test ./internal/...`

### Run integration tests only
`go test ./tests/integration/...`

### Run all tests
`go test ./...`

### Run tests with coverage
`go test -coverprofile=coverage.out ./...`
`go tool cover -html=coverage.out`