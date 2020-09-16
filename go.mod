module github.com/ibm-garage-cloud/argocd-plugin-key-protect

go 1.14

require (
	github.com/imdario/mergo v0.3.11
	github.com/urfave/cli/v2 v2.2.0
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.19.1
	k8s.io/apimachinery v0.19.1
)

replace github.com/ibm-garage-cloud/argocd-plugin-key-protect => ./
