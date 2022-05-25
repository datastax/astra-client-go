.PHONY: generate-client tools

tools:
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

generate-client: tools
	oapi-codegen -old-config-style -generate types,client -package astra -o astra/astra.gen.go swagger.yaml

generate-streaming-client: tools
	oapi-codegen -old-config-style -generate types,client -package astrastreaming -o astra-streaming/astra-streaming.gen.go streaming-swagger.yaml
