package generate_secret

import (
	"encoding/base64"
	"fmt"
	kpModel "github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/secret_template"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/key_management"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/key_management/factory"
	"github.com/imdario/mergo"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newMetadata(name string, labels *map[string]string, annotations *map[string]string) *metav1.ObjectMeta {
	return &metav1.ObjectMeta{
		Name: name,
		Labels: *labels,
		Annotations: *annotations,
	}
}

func newSecret(metadata *metav1.ObjectMeta, data *map[string][]byte) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: *metadata,
		Data: *data,
	}
}

func GenerateSecret(secretTemplate *kpModel.SecretTemplate) *corev1.Secret {
	var values map[string][]byte
	var annotations map[string]string

	values = make(map[string][]byte)
	annotations = make(map[string]string)

	kp := *secretTemplate

	keyManager := factory.LoadKeyManager(kp.ObjectMeta.Annotations)

	(*keyManager).PopulateMetadata(&annotations)

	specValues := kp.Spec.Values
	for _, kps := range specValues {
		values[kps.Name] = convertValue(keyManager, &kps, &annotations)
	}

	mergo.Merge(&annotations, kp.Spec.Annotations)

	return newSecret(newMetadata(kp.ObjectMeta.Name, &kp.Spec.Labels, &annotations), &values)
}

func convertValue(keyManager *key_management.KeyManager, keyValue *kpModel.SecretTemplateValue, annotations *map[string]string) []byte {
	var result []byte

	km := *keyManager
	kp := *keyValue

	if kp.KeyId != "" {
		result = []byte(km.GetKey(kp.KeyId))

		a := *annotations

		a[fmt.Sprintf("%s.keyId/%s", km.Id(), kp.Name)] = kp.KeyId
	} else if kp.Value != "" {
		result = []byte(base64.StdEncoding.EncodeToString([]byte(kp.Value)))
	} else {
		result = []byte(kp.B64Value)
	}

	return result
}
