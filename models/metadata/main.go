package metadata

type Metadata struct {
	Name string
	Labels map[string]string
	Annotations map[string]string
}

func New(name string, labels map[string]string, annotations map[string]string) Metadata {
	return Metadata{name, labels, annotations}
}
