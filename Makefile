.PHONY: default run groom build test setup

default: build groom run

run: build
	@- echo; echo "--- Running ..."
	env GODEBUG="gctrace=1" ${GOPATH}/bin/go-snake > logs/run.stdout.log 2> logs/run.stderr.log

build:
	@- echo; echo "--- Building and installing ..."
	env GODEBUG="" go install -gcflags="-m"

test: build
	@- echo; echo "--- Testing ..."
	go test -v -cover

groom: setup
	@- echo; echo "--- Linting ..."
	${GOPATH}/bin/golint
	@- echo; echo "--- Formatting ..."
	gofmt -e -s -w .
	@- echo; echo "--- Vetting ..."
	go tool vet -v .
	@- echo; echo "--- Fixing ..."
	go tool fix .
	@# ${GOPATH}/bin/goimports -w .

setup:
	@- mkdir -p logs
