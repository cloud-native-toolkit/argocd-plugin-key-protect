package generate_secrets

import (
	kpModel "github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/secret_template"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/kubernetes"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/generate_secret"
)

func GenerateSecrets(kp []kpModel.SecretTemplate) []kubernetes.Secret {
	var results []kubernetes.Secret

	results = []kubernetes.Secret{}

	for _, s := range kp {
		results = append(results, generate_secret.GenerateSecret(s))
	}

	return results
}
