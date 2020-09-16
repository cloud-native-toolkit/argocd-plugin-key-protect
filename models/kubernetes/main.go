package kubernetes

import (
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/metadata"
	"fmt"
	"gopkg.in/yaml.v2"
)

type Secret struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   metadata.Metadata
	Data       map[string]string
}

func New(metadata metadata.Metadata, values map[string]string) Secret {
	s := Secret{"v1", "Secret", metadata, values}
	return s
}

func AsYaml(secrets []Secret) string {
	var result string

	result = ""

	for _, s := range secrets {
		d, err := yaml.Marshal(&s)
		if err != nil {
			panic(err)
		}

		result = fmt.Sprintf("%s---\n%s\n", result, string(d))
	}

	return result
}
