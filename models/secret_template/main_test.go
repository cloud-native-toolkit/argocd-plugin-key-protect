package secret_template

import (
	"argocd-plugin-key-protect/util/test_support"
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
	test_support.ExpectEqual(t, "mysecret", got.Metadata.Name)
	test_support.ExpectNotEmpty(t, &got.Metadata.Annotations, "annotations")
	test_support.ExpectEqualInt(t, len(got.Spec.Values), 3)
}
