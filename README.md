# go-todo-app

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