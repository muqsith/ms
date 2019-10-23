.PHONY: build
build: 
	rm -rf bin && \
	mkdir bin && \
	go build -o bin/ms main.go