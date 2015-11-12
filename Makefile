.PHONY: default run groom build

default: build run

run: build
	@- echo; echo "--- Running ..."
	${GOPATH}/bin/go-snake

build: groom
	@#- echo; echo "--- Testing ..."
	@# go test -v
	@- echo; echo "--- Building and installing ..."
	go install -p 2 -gcflags="-m"

groom:
	@- echo; echo "--- Linting ..."
	${GOPATH}/bin/golint
	@- echo; echo "--- Formatting ..."
	go fmt
	@- echo; echo "--- Vetting ..."
	go vet
	@# ${GOPATH}/bin/goimports -w .
