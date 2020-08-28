package generate_secret

import (
	kpModel "argocd-plugin-key-protect/models/keyprotect_secret"
	"argocd-plugin-key-protect/models/kubernetes"
	"argocd-plugin-key-protect/models/metadata"
	kp "argocd-plugin-key-protect/pkg/key_protect"
	"encoding/base64"
	"github.com/imdario/mergo"
)

func GenerateSecret(kp kpModel.Secret) kubernetes.Secret {
	var values map[string]string
	var annotations map[string]string

	values = make(map[string]string)
	annotations = make(map[string]string)

	addKeyProtectInstanceId(&annotations)

	specValues := kp.Spec.Values
	for _, kps := range specValues {
		values[kps.Name] = convertValue(kps, &annotations)
	}

	mergo.Merge(&annotations, kp.Spec.Annotations)

	return kubernetes.New(metadata.New(kp.Metadata.Name, kp.Spec.Labels, annotations), values)
}

func convertValue(kp kpModel.SecretValue, annotations *map[string]string) string {
	var result string

	if kp.KeyId != "" {
		result = lookupValueFromKeyProtect(kp.KeyId)

		a := *annotations

		a["key-protect.key-id/" + kp.Name] = kp.KeyId
	} else if kp.Value != "" {
		result = base64.StdEncoding.EncodeToString([]byte(kp.Value))
	} else {
		result = kp.B64Value
	}

	return result
}

func lookupValueFromKeyProtect(keyId string) string {
	return kp.GetKey(keyId)
}

func addKeyProtectInstanceId(annotations *map[string]string) {
	kp.AddKeyProtectInstanceId(annotations)
}