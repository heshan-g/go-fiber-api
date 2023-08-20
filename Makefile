build:
	go build -o bin/server main.go

run: build
	./bin/server

dev:
	nodemon -e go --signal SIGTERM --exec 'make' run
