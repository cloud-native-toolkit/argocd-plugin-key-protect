# ArgoCD plugin - Key Protect secrets

ArgoCD plugin to generate secrets from values stored in Key Protect.

## Why?

A common issue when doing GitOps is how to handle sensitive information that should not be stored in the
Git repository (e.g. passwords, keys, etc). There are two different approaches to how to handle this issue:

1. Inject the values from another source into kubernetes Secret(s) at deployment time
2. Inject the values from another source in the pod at startup time via an InitContainer

The "other source" in this case would be a key management system that centralizes the storage and management
of sensitive information.

This repository addresses the first approach listed above by providing a plugin to ArgoCD to pull the 
sensitive values from **Key Protect** at deployment time and generate the appropriate kubernetes Secret(s).

## How it works

The plugin takes a directory containing template secret CR(s) as input, looks up the values of any sensitive information
for the provided keyIds from Key Protect, and generates a kubernetes Secret for each input template.

### Secret template

The input to the plugin is a directory that contains one or more "secret templates". In this case the "secret template" provides the 
structure of the desired template with placeholders for the values that will be pulled from the key management system. The following
provides the structure of the template:

```yaml
apiVersion: keymanagement.ibm/v1
kind: SecretTemplate
metadata:
  name: mysecret
  annotations:
    key-manager: key-protect
spec:
  labels: {}
  annotations: {}
  values:
    - name: url
      value: https://ibm.com
    - name: username
      b64value: dGVhbS1jYXA=
    - name: password
      keyId: 36397b07-d98d-4c0b-bd7a-d6c290163684
``` 

- The `metadata.annotations` value is optional and the only values supported currently for `key-manager` is `key-protect`
- The `metadata.name` value given will be used as the name for the Secret that will be generated.
- The information in `spec.labels` and `spec.annotations` will be copied over as the `labels` and `annotations` in the Secret that is generated
- The `spec.values` section contains the information that should be provided in the `data` section of the generated Secret. There are three prossible ways the values can be provided:

    - `value` - the actual value can be provided directly as clear text. This would be appropriate for information that is not sensitive but is required in the secret
    - `b64value` - a base64 encoded value can be provided to the secret. This can be used for large values that might present formatting issues or for information that is not sensitive but that might be obfuscated a bit (like a username)
    - `keyId` - the id (not the name) of the Standard Key that has been stored in Key Protect. The value stored in Key Protect can be anything

### Managing keys in Key Protect

Key Protect manages two different types of keys: `root keys` and `standard keys`. `Standard keys` are used to store any
kind of protected information. The Key Protect plugin reads the contents of a standard key, identified by a given key id, and
stores the key value into a secret in the cluster.

The following steps describe how to create a standard key:

1. Open the IBM Cloud console and navigate to the Key Protect service

2. Within Key Protect, select the **Manage Keys** tab

3. Press the `Add key` button to open the "Add a new key" dialog

4. Select the `Import your own key` radio button and `Standard key` from the drop down

5. Provide a descriptive name for the key and paste the base-64 encoded value of the key into the `Key material` field

    **Note:** A value can be encoded as base-64 from the terminal with the following command:
    
    ```shell script
    echo -n "{VALUE}" | base64
    ```
   
    If you need to encode a larger value, create the value in a file and encode the entire contents of the file with:
    
    ```shell script
    cat {file} | base64
    ```

6. Click **Import key** to create the key

7. Copy the value of the **ID**. This will be used later by the plugin

### Configuring an ArgoCD application

Assuming the [ArgoCD setup](#setting-up-argocd) has been completed for the plugin, the flow for ArgoCD would be as follows:

1. Create a directory in your GitOps repository for the secrets
2. Create a SecretTemplate yaml file in the directory for each Secret with the appropriate values
3. In ArgoCD, configure an Application for each secret directory created:

    - Provide the standard values for the **General**, **Source**, and **Destination** sections. The `path` in the **Source** section should refer to the directory containing the SecretTemplates
    - Select `Plugin` from the dropdown in the section at the bottom
    - Select `key-protect-secret` for the plugin name

Multiple directories can be defined within the GitOps repository to support whatever level of granularity is
desired for managing the secrets. Currently only one Key Protect instance is supported for a single ArgoCD instance.

## Setting up ArgoCD

### Key Protect credentials

In order to connect with Key Protect you will need three pieces of information:
- `IBM Cloud API Key` - an API Key that has `Reader` and `ReaderPlus` access to the Key Protect instance
- `Key Protect region` - the region where the Key Protect instance has been deployed
- `Key Protect instance id` - the GUID of the Key Protect instance where the secrets are stored

#### Get the instance id for Key Protect

1. Set the target resource group and region for the Key Protect instance.
    
    ```shell script
    ibmcloud target -g {RESOURCE_GROUP} -r {REGION}
    ```
  
2. List the available resources and find the name of the Key Protect instance.

    ```shell script
    ibmcloud resource service-instances
    ```
   
3. List the details for the Key Protect instance. The `Key Protect instance id` is listed as `GUID`.

    ```shell script
    ibmcloud resource service-instance {INSTANCE_NAME} 
    ```

#### Create the secret with the information needed to access Key Protect

Armed with the information from the previous step, run the following to create a secret:

```shell script
NAMESPACE="tools"
kubectl create secret generic -n ${NAMESPACE} key-protect-access \
  --from-literal=api-key=${API_KEY} \ 
  --from-literal=region=${REGION} \
  --from-literal=instance-id=${KP_INSTANCE_ID}
```

where:
- `NAMESPACE` should be the namespace where ArgoCD has been deployed

### Install the plugin dependencies

In order for the plugin to work, the `argocd-repo-server` deployment needs to have access to the command and
the credentials. To do this, the deployment needs to be patched with an InitContainer to install the command and
the environment variables from the secret generated in the previous step.

A script has been provided to simplify this step. To install the dependencies do the following:

1. Log into the cluster from the terminal

2. Clone the repository (if you haven't already)

3. Run the `install-plugin-dependencies.sh` script in the `hack` directory. (`hack` is a convention of Golang/operator repos. Don't blame us for the name of the folder...)

    ```shell script
    NAMESPACE="tools"
    ./hack/install-plugin-dependencies.sh ${NAMESPACE}
    ```
   
   where:
   - `NAMESPACE` is the namespace where ArgoCD has been deployed

When this script has completed, you should notice several new elements in the `argocd-repo-server` deployment:

- A new volume named `key-protect-install` defined in the `volumes` section
- A new initContainer named `key-protect-install`
- A new volume mount in the main conatiner named `key-protect-install`
- Three new environment variables in the `env` section referring to the `key-protect-access` secret

If you'd like to review further, the definition of the patch that has been applied can be found in `hack/install-plugin-dependencies-patch.json`

### Configure the plugin

The last step is to register the plugin with ArgoCD. To do this the `argocd-cm` ConfigMap must be updated.
However, if you are using the ArgoCD operator then the `argocd-cm` ConfigMap cannot be updated directly. Instead
the ArgoCD CR must be updated with the configuration values.

A script has been provided to patch the ArgoCD CR with the plugin configuration. To configure the plugin, do
the following:

1. Log into the cluster from the terminal

2. Clone the repository (if you haven't already)

3. Run the `install-plugin.sh` script in the `hack` directory.

    ```shell script
    NAMESPACE="tools"
    ./hack/install-plugin.sh ${NAMESPACE} [${ARGOCD_CR}]
    ```
   
   where:
   - `NAMESPACE` is the namespace where ArgoCD has been deployed
   - `ARGOCD_CR` is an optional value containing the name of the ArgoCD CR instance that will be patched. If not provided it defaults to `argocd`

When the script has completed, the ArgoCD CR should have a `configManagementPlugins` entry with the following
content, at a minimum:

```yaml
  configManagementPlugins: |
    - name: key-protect-secret
      generate:
        command: ["/key-protect/generate-secrets"]
        args: []
```

## Development

### Install dependencies

```shell script
make install
go mod vendor
```

### Test

```shell script
make test
```

### Build

```shell script
make compile
```

### Run

With the required information, the following will run the command:

```shell script
export API_KEY={IBM Cloud API Key}
export REGION={Key Protect region}
export KP_INSTANCE_ID={Key Protect instance id}
./bin/generate-secrets
```

### Publishing

The repository has been configured to use GitHub Actions to automate the build and release cycle. To release
a new version of the plugin, do the following:

1. Make your changes in a branch
2. Create a Pull Request with your changes. Be sure to add a change type label (`feature`, `bug`, or `chore`) to the PR
as well as a version label (`major`, `minor`, or `patch`)
3. If the build runs successfully and the changes look good, merge the PR. Pushing to master kicks off a chain of two other Actions:
    - The changes in master will be verified and, if passing, a release will be created
    - When the release is published, the changes will be built and published as assets on the release
