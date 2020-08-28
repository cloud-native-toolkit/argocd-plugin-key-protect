package metadata

type Metadata struct {
	Name string
	Labels map[string]string `yaml:"labels,omitempty"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
}

func New(name string, labels map[string]string, annotations map[string]string) Metadata {
	return Metadata{name, labels, annotations}
}
