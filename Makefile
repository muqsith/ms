VERSION := 1.0.0

.PHONY: build
build: 
	go build -o bin/ms main.go

ms:
	go build -o bin/ms main.go

install: ms
	install -m 0755 bin/ms /usr/local/bin