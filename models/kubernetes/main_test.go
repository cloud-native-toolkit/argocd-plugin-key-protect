package kubernetes

import (
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/metadata"
	testSupport "github.com/ibm-garage-cloud/argocd-plugin-key-protect/util/test_support"
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
	metadata := metadata.New(name, labels, annotations)

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
	testSupport.ExpectEqual(t, expected, got)
}
