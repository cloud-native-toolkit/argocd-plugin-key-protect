package generate_secrets_from_files

import (
	"fmt"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/secret_template"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/generate_secrets"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	"os"
	"path/filepath"
)

func secretsAsYaml(secrets *[]corev1.Secret) string {
	var result string

	result = ""

	for _, s := range *secrets {
		d, err := yaml.Marshal(&s)
		if err != nil {
			panic(err)
		}

		result = fmt.Sprintf("%s---\n%s\n", result, string(d))
	}

	return result
}

func GenerateSecretsFromFiles(rootPath string) string {
	kpSecrets := readYamlFiles(rootPath)

	secrets := generate_secrets.GenerateSecrets(&kpSecrets)

	return secretsAsYaml(secrets)
}

func readYamlFiles(rootPath string) []secret_template.SecretTemplate {
	yamlFiles := listYamlFiles(rootPath)

	return readFilesAsSecrets(yamlFiles)
}

func listYamlFiles(root string) []string {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}

func readFilesAsSecrets(paths []string) []secret_template.SecretTemplate {
	var result []secret_template.SecretTemplate

	result = []secret_template.SecretTemplate{}

	for _, path := range paths {
		result = append(result, readFileAsSecret(path))
	}

	return result
}

func readFileAsSecret(path string) secret_template.SecretTemplate {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return secret_template.FromYaml(dat)
}
