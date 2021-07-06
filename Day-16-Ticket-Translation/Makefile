# Go parameters
GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

all: clean deps test run
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
run:
		$(GORUN) main.go
deps:
		$(GOGET) github.com/stretchr/testify/assert
