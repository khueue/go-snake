.PHONY: default clean run groom build test

default: groom build groom run

run:
	@- echo; echo "--- Running ..."
	env GODEBUG="gctrace=1" ${GOPATH}/bin/go-snake > logs/run.stdout.log 2> logs/run.stderr.log

clean:
	@- echo; echo "--- Cleaning ..."
	go clean -x -i

build:
	@- echo; echo "--- Building ..."
	env GODEBUG="" time go install -gcflags="-m"

test:
	@- echo; echo "--- Testing ..."
	go test -v -cover

groom:
	@- mkdir -p logs
	@- echo; echo "--- Fixing formatting, imports and returns ..."
	${GOPATH}/bin/goreturns -w -l .
	@- echo; echo "--- Linting ..."
	${GOPATH}/bin/golint
	@- echo; echo "--- Vetting ..."
	go tool vet -v .
	@- echo; echo "--- Fixing ..."
	go tool fix .
