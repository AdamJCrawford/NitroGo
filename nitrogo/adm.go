package nitrogo

// ADM related configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/adm/adm
type ADMService struct {
	client *Client
}

// admparameter
// Configuration for ADM parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/adm/admparameter
func (s *ADMService) UpdateADMParamter() {}
func (s *ADMService) UnsetADMParamter()  {}
func (s *ADMService) GetAllADMParamter() {}
