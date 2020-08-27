package generate_secrets_from_files

import (
	"argocd-plugin-key-protect/models/keyprotect_secret"
	"argocd-plugin-key-protect/models/kubernetes"
	"argocd-plugin-key-protect/pkg/generate_secrets"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GenerateSecretsFromFiles(rootPath string) string {
	kpSecrets := readYamlFiles(rootPath)

	secrets := generate_secrets.GenerateSecrets(kpSecrets)

	return kubernetes.AsYaml(secrets)
}

func readYamlFiles(rootPath string) []keyprotect_secret.Secret {
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

func readFilesAsSecrets(paths []string) []keyprotect_secret.Secret {
	var result []keyprotect_secret.Secret

	result = []keyprotect_secret.Secret{}

	for _, path := range paths {
		result = append(result, readFileAsSecret(path))
	}

	return result
}

func readFileAsSecret(path string) keyprotect_secret.Secret {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return keyprotect_secret.FromYaml(dat)
}
