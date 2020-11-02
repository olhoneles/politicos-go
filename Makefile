setup:
	@go get github.com/codegangsta/gin
	@go get

setup-prod:
	@go get

run: run-api

run-api:
	@POLITICOS_DEBUG=True gin -b politicosapi -a 8888 -i

build-collector:
	@go build -o collector csv-collector/main.go

clean:
	@find . -name "*.swp" -delete

docs-update:
	@swag init
