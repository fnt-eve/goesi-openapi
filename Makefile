.PHONY: build test generate lint download-spec

download-spec:
	@./scripts/get_latest_esi_spec.sh

generate:
	@java -jar openapi-generator-cli.jar generate \
		-i esi-openapi-spec.json \
		-c openapi-generator-config.yaml \
		-g go \
		-o esi/ \
		>/dev/null 2>&1
	@./scripts/fix-generated-code.sh
	@go generate ./...
	@# Update ESI_VERSION with current spec version
	@jq -r '.info.version' esi-openapi-spec.json > ESI_VERSION
	@echo "Updated ESI_VERSION to $$(cat ESI_VERSION)"

build:
	@go build ./...

test:
	@go test ./...

lint:
	@golangci-lint run
