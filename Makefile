ORGNAME := nodehodl
PROJECTNAME := $(shell basename "$(CURDIR)")
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
IMAGE := $(ORGNAME)/$(PROJECTNAME)

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# build: Build binary into local bin folder. : make build
build:
	@echo "  >  Building binary..."
	go build -o $(GOBIN)/main *.go
	chmod +x $(GOBIN)/main

# run: Runs binary in local bin folder. : make run
run:
	@echo "  >  Running binary..."
	$(GOBIN)/main

# build-run: Builds && runs binary in bin folder. : make build-run
build-run: build run

## build-container: Builds container, argument [dev|prod] : make build-container target=dev
build-container:
	@echo "  >  Building container..."
	@IMAGE=$(IMAGE) TARGET=$(target) \
	bash scripts/build-container.sh

## run-container: Runs container locally in dev mode : make run-container
run-container:
	@echo "  >  Starting container locally..."
	@IMAGE=$(IMAGE) GOBASE=$(GOBASE) TARGET=dev \
	PORT=4000 \
	bash scripts/run-container.sh

# watch: Meant to be used by the container on dev target : make watch
watch:
	@echo "  >  Starting reflex watcher..."
	ulimit -n 1000
	reflex -s -r '\.go$$' make build-run

## tele-watch: Start telepresence(swaps container) with watcher enabled : make tele-watch
## : it always uses the image built with the dev target
tele-watch:
	@echo "  >  Telepresence watcher..."
	@PROJECTNAME=$(PROJECTNAME) GOBASE=$(GOBASE) IMAGE=$(IMAGE) TARGET=dev \
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