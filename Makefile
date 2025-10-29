VERSION  = $(or $(shell git tag --points-at HEAD | grep -o 'v[0-9.]*' | sed 's/v//'), unknown)
REVISION = $(shell git rev-parse HEAD)

REPOSITORY := github.com/mrlhansen/idrac_exporter
LDFLAGS    := -X $(REPOSITORY)/internal/version.Version=$(VERSION)
LDFLAGS    += -X $(REPOSITORY)/internal/version.Revision=$(REVISION)
GOFLAGS    := -ldflags "$(LDFLAGS)"
RUNFLAGS   ?= -config config.yml -verbose

# Build for current platform
build:
	go build $(GOFLAGS) -o idrac_exporter ./cmd/idrac_exporter

# Run the exporter
run:
	go run ./cmd/idrac_exporter $(RUNFLAGS)

# Cross-compilation targets
build-all: build-linux build-windows build-darwin

build-linux:
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -o bin/idrac_exporter-linux-amd64 ./cmd/idrac_exporter
	GOOS=linux GOARCH=arm64 go build $(GOFLAGS) -o bin/idrac_exporter-linux-arm64 ./cmd/idrac_exporter
	GOOS=linux GOARCH=arm GOARM=7 go build $(GOFLAGS) -o bin/idrac_exporter-linux-armv7 ./cmd/idrac_exporter

build-windows:
	GOOS=windows GOARCH=amd64 go build $(GOFLAGS) -o bin/idrac_exporter-windows-amd64.exe ./cmd/idrac_exporter

build-darwin:
	GOOS=darwin GOARCH=amd64 go build $(GOFLAGS) -o bin/idrac_exporter-darwin-amd64 ./cmd/idrac_exporter
	GOOS=darwin GOARCH=arm64 go build $(GOFLAGS) -o bin/idrac_exporter-darwin-arm64 ./cmd/idrac_exporter

# Build for a specific platform
# Usage: make build-platform GOOS=linux GOARCH=amd64
build-platform:
	go build $(GOFLAGS) -o bin/idrac_exporter-$(GOOS)-$(GOARCH)$(if $(filter windows,$(GOOS)),.exe) ./cmd/idrac_exporter
