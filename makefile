DOC_PORT:=8080
run:
	go run main.go

test:
	go test ./...

docs:
	@echo
	@echo "--- OPEN THIS LINK TO VIEW SWAGGER DOCS ---"
	@echo "http://localhost:${DOC_PORT}/?url=http://localhost:${DOC_PORT}/mnt/swagger.yaml?q=$(shell date +%s)"
	@echo
	docker run -e PORT=${DOC_PORT} -p ${DOC_PORT}:${DOC_PORT} -v `pwd`:/usr/share/nginx/html/mnt swaggerapi/swagger-editor
.PHONY: docs
