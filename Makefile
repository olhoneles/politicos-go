setup:
	@go get github.com/codegangsta/gin
	@go get

setup-prod:
	@go get

run: run-api

run-api:
	@POLITICOS_DEBUG=True gin -b politicosapi -a 8888 -i

run-collector:
	@go run csv-collector/main.go

clean:
	@find . -name "*.swp" -delete

docs-update:
	@swag init
