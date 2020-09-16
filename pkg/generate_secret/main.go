package generate_secret

import (
	kpModel "github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/secret_template"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/kubernetes"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/metadata"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/key_management"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/key_management/factory"
	"encoding/base64"
	"fmt"
	"github.com/imdario/mergo"
)

func GenerateSecret(kp kpModel.SecretTemplate) kubernetes.Secret {
	var values map[string]string
	var annotations map[string]string

	values = make(map[string]string)
	annotations = make(map[string]string)

	keyManager := factory.LoadKeyManager(kp.Metadata.Annotations)

	keyManager.PopulateMetadata(&annotations)

	specValues := kp.Spec.Values
	for _, kps := range specValues {
		values[kps.Name] = convertValue(keyManager, kps, &annotations)
	}

	mergo.Merge(&annotations, kp.Spec.Annotations)

	return kubernetes.New(metadata.New(kp.Metadata.Name, kp.Spec.Labels, annotations), values)
}

func convertValue(keyManager key_management.KeyManager, kp kpModel.SecretTemplateValue, annotations *map[string]string, ) string {
	var result string

	if kp.KeyId != "" {
		result = keyManager.GetKey(kp.KeyId)

		a := *annotations

		a[fmt.Sprintf("%s.keyId/%s", keyManager.Id(), kp.Name)] = kp.KeyId
	} else if kp.Value != "" {
		result = base64.StdEncoding.EncodeToString([]byte(kp.Value))
	} else {
		result = kp.B64Value
	}

	return result
}
