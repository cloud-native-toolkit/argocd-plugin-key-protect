package secret_template

import (
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

type SecretTemplateValue struct {
	Name string `json:"name" yaml:"name"`
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
	B64Value string `json:"b64value,omitempty" yaml:"b64value,omitempty"`
	KeyId string `json:"keyId,omitempty" yaml:"keyId,omitempty"`
}

type SecretTemplateSpec struct {
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Values []SecretTemplateValue `json:"values" yaml:"value"`
}

type SecretTemplate struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Spec  SecretTemplateSpec `json:"spec" yaml:"spec"`
}

func FromYaml(data []byte) SecretTemplate {
	value := SecretTemplate{}

	err := yaml.Unmarshal(data, &value)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return value
}
