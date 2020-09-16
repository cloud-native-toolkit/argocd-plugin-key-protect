package secret_template

import (
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/metadata"
	"gopkg.in/yaml.v2"
	"log"
)

type SecretTemplateValue struct {
	Name string
	Value string `yaml:"value,omitempty"`
	B64Value string `yaml:"b64value,omitempty"`
	KeyId string `yaml:"keyId,omitempty"`
}

type SecretTemplateSpec struct {
	Labels map[string]string `yaml:"labels,omitempty"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
	Values []SecretTemplateValue
}

type SecretTemplate struct {
	ApiVersion string
	Kind       string
	Metadata   metadata.Metadata
	Spec       SecretTemplateSpec
}

func FromYaml(data []byte) SecretTemplate {
	value := SecretTemplate{}

	err := yaml.Unmarshal(data, &value)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return value
}
