#TEST
.PHONY: build test

build:
		go build -i -o artifacts/compete github.com/LukeJelly/rotateimage

test:
		go test -race $(go list ./... | grep -v vendor)
