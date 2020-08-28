package generate_secret

import (
	kpModel "argocd-plugin-key-protect/models/keyprotect_secret"
	"argocd-plugin-key-protect/models/kubernetes"
	"argocd-plugin-key-protect/models/metadata"
	kp "argocd-plugin-key-protect/pkg/key_protect"
	"encoding/base64"
)

func GenerateSecret(kp kpModel.Secret) kubernetes.Secret {
	var values map[string]string

	values = make(map[string]string)

	specValues := kp.Spec.Values
	for _, kps := range specValues {
		values[kps.Name] = convertValue(kps)
	}

	return kubernetes.New(metadata.New(kp.Metadata.Name, kp.Spec.Labels, kp.Spec.Annotations), values)
}

func convertValue(kp kpModel.SecretValue) string {
	var result string

	if kp.KeyId != "" {
		result = lookupValueFromKeyId(kp.KeyId)
	} else if kp.Value != "" {
		result = base64.StdEncoding.EncodeToString([]byte(kp.Value))
	} else {
		result = kp.B64Value
	}

	return result
}

func lookupValueFromKeyId(keyId string) string {
	return kp.GetKey(keyId)
}
