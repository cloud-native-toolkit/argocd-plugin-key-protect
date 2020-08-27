package kubernetes

import (
	"argocd-plugin-key-protect/models/metadata"
	"fmt"
	"gopkg.in/yaml.v2"
)

type Secret struct {
	ApiVersion string
	Kind string
	Metadata metadata.Metadata
	Values map[string]string
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
