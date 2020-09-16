package factory

import (
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/key_management"
	"github.com/ibm-garage-cloud/argocd-plugin-key-protect/pkg/key_management/key_protect"
	"fmt"
)

func LoadKeyManager(annotations map[string]string) key_management.KeyManager {
	keyManager, ok := annotations["key-manager"]
	if !ok {
		keyManager = "key-protect"
	}

	switch keyManager {
	case "key-protect":
		return key_protect.New(annotations)
	default:
		fmt.Printf("Key manager not found: %s", keyManager)
		panic("Key manager not found")
	}
}
