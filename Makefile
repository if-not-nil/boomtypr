BINARY_NAME=boomtypr
GO=go
LDFLAGS := -s -w \
  -X main.version=$(VERSION) \
  -X main.commit=$(shell git rev-parse --short HEAD)

# all: build
all: build

build:
	CGO_ENABLED=0 $(GO) build -o $(BINARY_NAME) -ldflags "$(LDFLAGS)" .

run:
	$(GO) run .

test:
	$(GO) test ./...

clean:
	rm -f $(BINARY_NAME)

vet:
	$(GO) vet ./...

.PHONY: all build run test clean fmt vet
