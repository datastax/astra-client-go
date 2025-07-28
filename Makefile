
generate: generate-client generate-streaming-client generate-rest-api-client

tools:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1

generate-client: tools
	oapi-codegen -old-config-style -generate types,client -package astra -o astra/astra.gen.go openapi/astra-devops-api.yaml

generate-streaming-client: tools
	oapi-codegen -old-config-style -generate types,client -package astrastreaming -o astra-streaming/astra-streaming.gen.go openapi/streaming-devops-api.yaml

generate-rest-api-client: tools
	oapi-codegen -old-config-style -generate types,client -package astrarestapi -o astra-rest-api/astra-rest-api.gen.go openapi/stargate-api.yaml

build: generate
	go build ./...

.PHONY: tools generate generate-client generate-streaming-client generate-rest-api-client build
