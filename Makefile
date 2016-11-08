BINARY=epp

VERSION=1.0.0
GIT_COMMIT=`git rev-parse @`
LDFLAGS=-X main.Version=$(VERSION) -X main.GitCommit=$(GIT_COMMIT)

build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY)

release:
	go build -ldflags "-s -w $(LDFLAGS)" -o $(BINARY)
