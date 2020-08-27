package keyprotect_secret

import (
	"argocd-plugin-key-protect/models/metadata"
	"gopkg.in/yaml.v2"
	"log"
)

type SecretValue struct {
	Name string
	Value string `yaml:"value,omitempty"`
	B64Value string `yaml:"b64value,omitempty"`
	KeyId string `yaml:"keyId,omitempty"`
}

type Secret struct {
	ApiVersion string
	Kind string
	Metadata metadata.Metadata
	Spec []SecretValue
}

func FromYaml(data []byte) Secret {
	value := Secret{}

	err := yaml.Unmarshal(data, &value)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return value
}
