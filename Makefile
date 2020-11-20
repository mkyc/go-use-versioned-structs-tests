build:
	go build -x -o ./output/producer github.com/mkyc/go-use-versioned-structs-tests/cmd/producer
	go build -x -o ./output/consumer github.com/mkyc/go-use-versioned-structs-tests/cmd/consumer

clean:
	rm -rf ./output

run: build
	cd ./output/ && ./producer
	cd ./output/ && ./consumer
