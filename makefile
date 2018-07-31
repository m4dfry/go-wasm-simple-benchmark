# More a script than a makefile

# Go parameters
GOCMD=go
GORUN=$(GOCMD) run
GOBETACMD=go1.11beta2
GOBETABUILD=$(GOBETACMD) build

wasm:
	GOARCH=wasm GOOS=js $(GOBETABUILD) -o fibonacci.wasm fibonacci.go
run:
	$(GORUN) server_improved.go
