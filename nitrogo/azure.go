package nitrogo

type AzureService struct {
	client *Client
}

// azureapplication
func (s *AzureService) AddAzureApplication()    {}
func (s *AzureService) DeleteAzureApplication() {}
func (s *AzureService) GetAllAzureApplication() {}
func (s *AzureService) GetAzureApplication()    {}
func (s *AzureService) CountAzureApplication()  {}

// azurekeyvault
func (s *AzureService) AddAzureKeyVault()    {}
func (s *AzureService) DeleteAzureKeyVault() {}
func (s *AzureService) GetAllAzureKeyVault() {}
func (s *AzureService) GetAzureKeyVault()    {}
func (s *AzureService) CountAzureKeyVault()  {}
