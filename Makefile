test:
	go test -v -cover ./...

build:
	go build -ldflags "-w -s"

check:
	go run main.go expiration.go www.google.com.ar

.PHONY: test build check
