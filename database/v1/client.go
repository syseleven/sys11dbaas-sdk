package v1

//go:generate go tool oapi-codegen --config config.yaml --generate client,types -o client.generated.go -package v1 ../../openapispec.yaml
//go:generate go tool generator ../../openapispec.yaml client.extension.go v1 v1

const (
	StateNotReady = "ClusterIsNotReady"
	StateReady    = "ClusterIsReady"
)
