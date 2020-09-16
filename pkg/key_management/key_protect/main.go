package key_protect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type keyProtectResource struct {
	Payload string
}

type keyProtectResult struct {
	Resources []keyProtectResource
}

type tokenResult struct {
	AccessToken string `json:"access_token"`
}

type KeyProtectConfig struct {
	apiKey string
	instanceId string
	region string
}

func newConfig(apiKey string, instanceId string, region string) KeyProtectConfig {
	return KeyProtectConfig{apiKey, instanceId, region}
}

func buildConfig(config map[string]string) KeyProtectConfig {
	instanceId, instanceIdOk := config["key-protect/instanceId"]
	if !instanceIdOk {
		instanceId = os.Getenv("KP_INSTANCE_ID")
	}

	region, regionOk := config["key-protect/region"]
	if !regionOk {
		region = os.Getenv("REGION")
	}

	apiKey, apiKeyOk := os.LookupEnv("KP_API_KEY")
	if !apiKeyOk {
		apiKey = os.Getenv("API_KEY")
	}

	return newConfig(apiKey, instanceId, region)
}

type KeyProtect struct {
	id string
	config KeyProtectConfig
}

func New(config map[string]string) KeyProtect {
	return KeyProtect{"key-protect",buildConfig(config)}
}

func (k KeyProtect) GetKey(keyId string) string {
	config := k.config

	accessToken := getAccessToken(config.apiKey)

	url := fmt.Sprintf("https://%s.kms.cloud.ibm.com/api/v2/keys/%s", config.region, keyId)

	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("accept", "application/vnd.ibm.kms.key+json")
	request.Header.Set("bluemix-instance", config.instanceId)
	request.Header.Set("Authorization", "Bearer " + accessToken)

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result keyProtectResult
	json.Unmarshal(body, &result)

	return result.Resources[0].Payload
}

func getAccessToken(apiKey string) string {
	url := "https://iam.cloud.ibm.com/identity/token"

	client := http.Client{}

	bodyText := "grant_type=urn%3Aibm%3Aparams%3Aoauth%3Agrant-type%3Aapikey&apikey=" + apiKey

	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(bodyText)))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("accept", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result tokenResult
	json.Unmarshal(body, &result)

	return result.AccessToken
}

func (k KeyProtect) PopulateMetadata(annotations *map[string]string) {
	config := k.config

	a := *annotations

	a[fmt.Sprintf("%s.instanceId", k.Id())] = config.instanceId
	a[fmt.Sprintf("%s.region", k.Id())] = config.region
}

func (k KeyProtect) Id() string {
	return k.id
}
