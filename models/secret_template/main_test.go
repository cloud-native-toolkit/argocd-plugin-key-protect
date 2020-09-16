package secret_template

import (
	testSupport "github.com/ibm-garage-cloud/argocd-plugin-key-protect/util/test_support"
	"io/ioutil"
	"testing"
)

func TestCanary(t *testing.T) {
	if false {
		t.Errorf("Canary failed")
	}
}

func TestFromYaml(t *testing.T) {

	yamlBytes, err := ioutil.ReadFile("../../test/secret-in.yaml")
	if err != nil {
		panic(err)
	}

	got := FromYaml(yamlBytes)
	testSupport.ExpectEqual(t, "mysecret", got.Metadata.Name)
	testSupport.ExpectNotEmpty(t, &got.Metadata.Annotations, "annotations")
	testSupport.ExpectEqualInt(t, len(got.Spec.Values), 3)
}
