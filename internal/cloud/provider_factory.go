package cloud

import (
	"fmt"
	"os"
)

// ProviderFactory creates and returns a cloud provider based on the given name
func ProviderFactory(providerName string) (Provider, error) {
	switch providerName {
	case "AWS":
		return NewAWSProvider()
	case "Azure":
		subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
		if subscriptionID == "" {
			return nil, fmt.Errorf("AZURE_SUBSCRIPTION_ID environment variable is not set")
		}
		return nil, fmt.Errorf("currently unsupported %s coming soon in new release", providerName)
		//return NewAzureProvider(subscriptionID)
	case "GCP":
		projectID := os.Getenv("GCP_PROJECT_ID")
		credentialsFile := os.Getenv("GCP_CREDENTIALS_FILE")
		if projectID == "" || credentialsFile == "" {
			return nil, fmt.Errorf("GCP_PROJECT_ID or GCP_CREDENTIALS_FILE environment variable is not set")
		}
		return nil, fmt.Errorf("currently unsupported %s coming soon in new release", providerName)
		//return NewGCPProvider(projectID, credentialsFile)
	default:
		return nil, fmt.Errorf("unsupported cloud provider: %s", providerName)
	}
}
