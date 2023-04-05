.PHONY: deps build clean install

export LDFLAGS +=

install:
	go install -ldflags='$(LDFLAGS)' ./...

uninstall:
	rm -f `which xml-cli`

build: deps
	go build -o bin/ -ldflags='$(LDFLAGS)' ./...

deps:
	go mod tidy
	go mod vendor -v

clean:
	go clean -x ./...