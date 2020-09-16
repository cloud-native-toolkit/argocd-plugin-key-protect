package key_management

type KeyManager interface {
	GetKey(keyId string) string
    PopulateMetadata(annotations *map[string]string)
	Id() string
}
