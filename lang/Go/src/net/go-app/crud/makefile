build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o ./hello

run: build
	./hello

site: build
	go test -v .
