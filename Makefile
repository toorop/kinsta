build:
	go build -o dist/kinsta

run: build
	dist/kinsta