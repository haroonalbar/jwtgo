#install CompileDaemon on your device locally to run this.
daemon:
	@compiledaemon --build="make build" --command="make run"

build:
	@go build -o ./bin/jwtgo ./

run: build
	@./bin/jwtgo

test: 
	@go test ./... -v
