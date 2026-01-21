package v2

//go:generate go tool oapi-codegen --config config.yaml --generate client,types -o client.generated.go -package v2 ../../openapispec.yaml
//go:generate go tool generator ../../openapispec.yaml client.extension.go v2 v2

const (
	StateNotReady = "ClusterIsNotReady"
	StateReady    = "ClusterIsReady"
)
