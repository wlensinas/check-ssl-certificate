test:
	go test -v -cover ./...

build:
	go build -ldflags "-w -s"

run:
	go run main.go expiration.go

.PHONY: test build check
