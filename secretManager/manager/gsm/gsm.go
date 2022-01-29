package gsm

import (
	"context"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/spf13/viper"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func init() {
	viper.AutomaticEnv()

	dir, _ := os.Getwd()
	viper.SetDefault("GOOGLE_APPLICATION_CREDENTIALS", dir+"/secretManager/manager/gsm/key.json")
}

func Get(secretName string) string {

	client := getClient()

	// Build the request.
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: getFullSecretName(secretName),
	}

	ctx := context.Background()
	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Printf("failed to access secret version: %v", err)
		return ""
	}

	defer client.Close()
	return string(result.Payload.Data)

}

func getFullSecretName(secretName string) string {
	return fmt.Sprintf("projects/%s/secrets/%s/versions/latest", viper.Get("GCP_PROJECT_ID"), secretName)
}

func getClient() secretmanager.Client {
	ctx := context.Background()

	secretClient, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}

	return *secretClient
}
