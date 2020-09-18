package generate_secrets

import (
	keymanagementv1 "github.com/ibm-garage-cloud/key-management-operator/pkg/apis/keymanagement/v1"
	generate_secret "github.com/ibm-garage-cloud/key-management-operator/pkg/service/generate_secret"
	corev1 "k8s.io/api/core/v1"
)

func GenerateSecrets(secretTemplates *[]keymanagementv1.SecretTemplate) *[]corev1.Secret {
	var results []corev1.Secret

	results = []corev1.Secret{}

	for _, secretTemplate := range *secretTemplates {
		results = append(results, *generate_secret.GenerateSecret(&secretTemplate))
	}

	return &results
}
