//go:build interfaces
// +build interfaces

//go:generate go tool github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=../../api/codegen_config/server.yaml ../../api/openapi.yaml
package api
