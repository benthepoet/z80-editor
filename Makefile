main:
	go build -ldflags="-s -w" -gcflags="-m" -o bin/main cmd/main.go