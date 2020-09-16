package generate_secrets_from_files

import (
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/secret_template"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/models/kubernetes"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/generate_secrets"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GenerateSecretsFromFiles(rootPath string) string {
	kpSecrets := readYamlFiles(rootPath)

	secrets := generate_secrets.GenerateSecrets(kpSecrets)

	return kubernetes.AsYaml(secrets)
}

func readYamlFiles(rootPath string) []secret_template.Secret {
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

func readFilesAsSecrets(paths []string) []secret_template.Secret {
	var result []secret_template.Secret

	result = []secret_template.Secret{}

	for _, path := range paths {
		result = append(result, readFileAsSecret(path))
	}

	return result
}

func readFileAsSecret(path string) secret_template.Secret {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return secret_template.FromYaml(dat)
}
