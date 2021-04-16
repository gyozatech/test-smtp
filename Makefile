.PHONY: clean build

clean:
	rm -rf build

build:
	go build -o build/receiver main.go