GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_LOC=bin
BINARY_NAME=homie
VERSION=0.0.1

all: build
setup:
	$(GOGET) -v ./...
build:
	$(GOBUILD) -o ./$(BINARY_LOC)/$(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -rf $(BINARY_LOC)
run: clean build
	./$(BINARY_LOC)/$(BINARY_NAME)
