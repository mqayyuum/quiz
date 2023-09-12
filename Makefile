# Define the name of your Go binary
BINARY_NAME := quizapp

# Define Go compiler and flags
GO := go
GOFLAGS :=

# Define build and run targets
build:
	$(GO) build $(GOFLAGS) -o $(BINARY_NAME) .

run:
	./$(BINARY_NAME)

# Define clean target
clean:
	rm -f $(BINARY_NAME)

# Define the default target (build)
.DEFAULT_GOAL := build
