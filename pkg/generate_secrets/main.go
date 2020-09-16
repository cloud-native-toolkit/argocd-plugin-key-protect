package generate_secrets

import (
	kpModel "github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/secret_template"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/generate_secret"
	corev1 "k8s.io/api/core/v1"
)

func GenerateSecrets(secretTemplates *[]kpModel.SecretTemplate) *[]corev1.Secret {
	var results []corev1.Secret

	results = []corev1.Secret{}

	for _, secretTemplate := range *secretTemplates {
		results = append(results, *generate_secret.GenerateSecret(&secretTemplate))
	}

	return &results
}
