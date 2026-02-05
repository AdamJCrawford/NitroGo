package nitrogo

const (
	azureApplicationURL = "/nitro/v1/config/azureapplication"
	azureKeyVaultURL    = "/nitro/v1/config/azurekeyvault"
)

// Azure configuration
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/azure/azure
type AzureService struct {
	client *Client
}

// azureapplication
// Configuration for Azure Application resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/azure/azureapplication
func (s *AzureService) AddAzureApplication()    {}
func (s *AzureService) DeleteAzureApplication() {}
func (s *AzureService) GetAllAzureApplication() {}
func (s *AzureService) GetAzureApplication()    {}
func (s *AzureService) CountAzureApplication()  {}

// azurekeyvault
// Configuration for Azure Key Vault entity resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/azure/azurekeyvault
func (s *AzureService) AddAzureKeyVault()    {}
func (s *AzureService) DeleteAzureKeyVault() {}
func (s *AzureService) GetAllAzureKeyVault() {}
func (s *AzureService) GetAzureKeyVault()    {}
func (s *AzureService) CountAzureKeyVault()  {}
