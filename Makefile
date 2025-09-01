.PHONY: build
build:
	templ generate internal/view
	go build -o ./bin/myapp cmd/goservice/main.go

.PHONY: run
run: build
	./bin/myapp