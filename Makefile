
# All target
all: test build

# Build the project
build:
	@echo "Building..."
	@go build -o main cmd/main.go

# Run tests
test:
	@echo "Testing..."
	@go test -v ./...

# Clean build files
clean:
	@rm -rf main

generate:
	@echo "Generating..."
	@go generate ./...

# Run the application
run:
	./main