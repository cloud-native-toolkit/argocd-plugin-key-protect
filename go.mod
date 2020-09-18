module github.com/ibm-garage-cloud/argocd-plugin-key-protect

go 1.13

require (
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/ibm-garage-cloud/key-management-operator v0.9.3
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/urfave/cli/v2 v2.2.0
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.19.1
	k8s.io/apimachinery v0.19.1
)

replace github.com/ibm-garage-cloud/argocd-plugin-key-protect => ./
