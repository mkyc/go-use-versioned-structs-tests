build:
	go build -x -o ./output/producer github.com/mkyc/go-use-versioned-structs-tests/cmd/producer

clean:
	rm -rf ./output
