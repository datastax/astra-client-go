
generate: generate-client generate-streaming-client generate-rest-api-client

tools:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11

generate-client: tools
	oapi-codegen -old-config-style -generate types,client -package astra -o astra/astra.gen.go swagger.yaml

generate-streaming-client: tools
	oapi-codegen -old-config-style -generate types,client -package astrastreaming -o astra-streaming/astra-streaming.gen.go streaming-swagger.yaml

generate-rest-api-client: tools
	oapi-codegen -old-config-style -generate types,client -package astrarestapi -o astra-rest-api/astra-rest-api.gen.go rest-api-swagger.yaml

build: generate
	go build ./...

.PHONY: tools generate generate-client generate-streaming-client generate-rest-api-client build
