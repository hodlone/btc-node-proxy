include .prod.env

ORGNAME := nodehodl
PROJECTNAME := $(shell basename "$(PWD)")
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
IMAGE := $(ORGNAME)/$(PROJECTNAME)

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

setup:
	PORT = 4000
	BTC_NODE_ZMQ_ADDR = tcp://0.0.0.0:29000
	NATS_ADDR = nats://localhost:4222
	NATS_NAME = test-test-test

## build: Build binary into local bin folder. : make build
build:
	@echo "  >  Building binary..."
	go build -o $(GOBIN)/main main.go
	chmod +x $(GOBIN)/main

## run: Runs binary in bin folder. : make run
run:
	@echo "  >  Running binary..."
	$(GOBIN)/main

## build-run: Builds && runs binary in bin folder. : make build-run
build-run: build run

## watch: Runs binary with hot-reloader(reflex) monitoring source files. : make watch
watch:
	@echo "  >  Starting reflex watcher..."
	ulimit -n 1000 #increase the file watch limit
	reflex -s -r '\.go$$' make build-run

## build-container: Builds container, argument [dev|prod] : make build-container target=dev
build-container:
	@echo "  >  Building container..."
	@IMAGE=$(IMAGE) TARGET=$(target) bash scripts/build-container.sh

## run-container: Runs container, argument [dev|prod] : make run-container target=dev
run-container:
	@echo "  >  Running container..."
	@IMAGE=$(IMAGE) TARGET=$(target) GOBASE=$(GOBASE) \
	PORT=$(PORT) BTC_NODE_ZMQ_ADDR=$(BTC_NODE_ZMQ_ADDR) \
	NATS_ADDR=$(NATS_ADDR) NATS_NAME=$(NATS_NAME) \
	bash scripts/run-container.sh

## tele-watch: Start telepresence with watcher enabled : make tele-watch
tele-watch:
	@echo "  >  Telepresence watcher..."
	@PROJECTNAME=$(PROJECTNAME) GOBASE=$(GOBASE) \
	IMAGE=$(IMAGE) PORT=$(PORT) \
 	bash scripts/telepresence-watcher.sh

## clean: Clean build files. Runs `go clean` internally. : make clean
clean:
	@echo "  >  Cleaning build cache & $(GOBIN)/main"
	@-rm $(GOBIN)/main 2> /dev/null
	go clean

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo