BINARY=epp

VERSION=1.0.0
GIT_COMMIT=$(git rev-parse @)
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$(GIT_COMMIT)"

build:
	go build $(LDFLAGS) -o $(BINARY)
