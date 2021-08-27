.PHONY: build
build: clean
	@go build -o build/test cmd/main.go

.PHONY: run
run: build
	@docker-compose up d
	@./build/test

.PHONY: clean
clean:
	@dokcer-compose down --remove-orphans
	@rm -rf build/