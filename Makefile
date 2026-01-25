BINARY_NAME=boomtypr
GO=go

# all: build
all: run

build:
	$(GO) build -o $(BINARY_NAME) .

run:
	$(GO) run .

test:
	$(GO) test ./...

clean:
	rm -f $(BINARY_NAME)

vet:
	$(GO) vet ./...

.PHONY: all build run test clean fmt vet
