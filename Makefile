_build:
	go build -o bin/assistant ./cmd/aws_assistant/main.go
	sudo chmod +x ./bin/assistant

test:
	go test -v ./...
