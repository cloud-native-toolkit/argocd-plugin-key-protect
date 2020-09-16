package factory

import (
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/key_management"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/key_management/key_protect"
	"fmt"
)

func LoadKeyManager(annotations map[string]string) *key_management.KeyManager {
	keyManagerName, ok := annotations["key-manager"]
	if !ok {
		keyManagerName = "key-protect"
	}

	var keyManager key_management.KeyManager

	switch keyManagerName {
	case "key-protect":
		keyManager = key_protect.New(annotations)
		return &keyManager
	default:
		fmt.Printf("Key manager not found: %s", keyManager)
		panic("Key manager not found")
	}
}
