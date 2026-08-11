.PHONY: build test generate lint download-spec

download-spec:
	@./scripts/get_latest_esi_spec.sh

generate:
	@rm -rf esi
	@mkdir -p esi
	@cp openapi-generator-ignore esi/.openapi-generator-ignore
	@java -jar openapi-generator-cli.jar generate \
		-i esi-openapi-spec.json \
		-c openapi-generator-config.yaml \
		-g go \
		-o esi/ \
		>/tmp/openapi-generator.log 2>&1 \
		|| { echo "openapi-generator failed:"; cat /tmp/openapi-generator.log; exit 1; }
	@./scripts/fix-generated-code.sh
	@go generate ./...
	@jq -r '.info.version' esi-openapi-spec.json > ESI_VERSION
	@echo "Updated ESI_VERSION to $$(cat ESI_VERSION)"

build:
	@go build ./...

test:
	@go test ./...

lint:
	@golangci-lint run ./...
