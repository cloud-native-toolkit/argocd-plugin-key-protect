package generate_secrets

import (
	kpModel "argocd-plugin-key-protect/models/keyprotect_secret"
	"argocd-plugin-key-protect/models/kubernetes"
	"argocd-plugin-key-protect/pkg/generate_secret"
)

func GenerateSecrets(kp []kpModel.Secret) []kubernetes.Secret {
	var results []kubernetes.Secret

	results = []kubernetes.Secret{}

	for _, s := range kp {
		results = append(results, generate_secret.GenerateSecret(s))
	}

	return results
}
