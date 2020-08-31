package kubernetes

import (
	metadata2 "argocd-plugin-key-protect/models/metadata"
	"argocd-plugin-key-protect/util/test_support"
	"testing"
)

func TestCanary(t *testing.T) {
	if false {
		t.Errorf("Canary failed")
	}
}

func TestAsYaml(t *testing.T) {

	name := "test"
	labels := make(map[string]string)
	annotations := make(map[string]string)
	metadata := metadata2.New(name, labels, annotations)

	values := make(map[string]string)

	values["test"] = "value"

	secrets := []Secret{New(metadata, values)}
	expected := `---
apiVersion: v1
kind: Secret
metadata:
  name: test
data:
  test: value
`

	got := AsYaml(secrets)
	test_support.ExpectEqual(t, expected, got)
}
