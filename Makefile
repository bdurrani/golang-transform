BINARY_NAME=main.out
 
all: build test
 
build:
		go build -o ./build/${BINARY_NAME}
 
test:
		go test ./...
 
run: build
		#./build/${BINARY_NAME}
		go run .
clean:
		go clean
		rm -rf ./build