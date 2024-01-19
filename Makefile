build: clean
	go build -o build/genetic genetic.go

test:  clean build
	go test -v ./...

run-queens:
	go run genetic.go queens

run-catdog:
	go run genetic.go catdog

clean:
	rm -rf build