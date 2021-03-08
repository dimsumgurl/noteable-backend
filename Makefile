build:
	go build -o bin/noteable cmd/noteable/main.go 
test:
unit-test:
integration-test:
clean:
tools:
lint:
build-api-server:
	swagger validate pkg/api/noteable.yaml
	swagger generate server --exclude-main -f pkg/api/noteable.yaml -A noteable-backend -t pkg/api/