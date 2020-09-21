package generate_secrets_from_files

import (
	"fmt"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/generate_secrets"
 	yaml2 "github.com/ghodss/yaml"
	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/util/json"
	"io/ioutil"
	keymanagementv1 "github.com/ibm-garage-cloud/key-management-operator/pkg/apis/keymanagement/v1"
	corev1 "k8s.io/api/core/v1"
	"log"
	"os"
	"path/filepath"
)

func secretsAsYaml(secrets *[]corev1.Secret) string {
	var result string

	result = ""

	for _, s := range *secrets {
		jsonSecret, err := json.Marshal(&s)
		if err != nil {
			panic(err)
		}

		yamlSecret, err := yaml2.JSONToYAML(jsonSecret)

		result = fmt.Sprintf("%s---\n%s\n", result, string(yamlSecret))
	}

	return result
}

func secretFromYaml(data []byte) keymanagementv1.SecretTemplate {
	value := keymanagementv1.SecretTemplate{}

	err := yaml.Unmarshal(data, &value)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return value
}


func GenerateSecretsFromFiles(rootPath string) string {
	kpSecrets := readYamlFiles(rootPath)

	secrets := generate_secrets.GenerateSecrets(&kpSecrets)

	return secretsAsYaml(secrets)
}

func readYamlFiles(rootPath string) []keymanagementv1.SecretTemplate {
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

func readFilesAsSecrets(paths []string) []keymanagementv1.SecretTemplate {
	var result []keymanagementv1.SecretTemplate

	result = []keymanagementv1.SecretTemplate{}

	for _, path := range paths {
		result = append(result, readFileAsSecret(path))
	}

	return result
}

func readFileAsSecret(path string) keymanagementv1.SecretTemplate {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return secretFromYaml(dat)
}
