package generate_secret

import (
	kpModel "argocd-plugin-key-protect/models/keyprotect_secret"
	"argocd-plugin-key-protect/models/kubernetes"
	kp "argocd-plugin-key-protect/pkg/key_protect"
	"encoding/base64"
)

func GenerateSecret(kp kpModel.Secret) kubernetes.Secret {
	var values map[string]string

	values = make(map[string]string)

	for _, kps := range kp.Spec {
		values[kps.Name] = convertValue(kps)
	}

	return kubernetes.New(kp.Metadata, values)
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
