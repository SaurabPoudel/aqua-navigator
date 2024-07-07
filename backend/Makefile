# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=aqua-navigator

# Default target executed when no arguments are provided to make
all: build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Run the application
run: build
	./$(BINARY_NAME)

# Test the application
test:
	$(GOTEST) -v ./...

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

.PHONY: all build clean run test
