package nitrogo

// IPSEC configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ipsec/ipsec
type IPSECService struct {
	client *Client
}

// ipsecparameter
// Configuration for IPSEC paramter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ipsec/ipsecparameter
func (s *IPSECService) UpdateIPSECParamter() {}
func (s *IPSECService) UnsetIPSECParamter()  {}
func (s *IPSECService) GetAllIPSECParamter() {}

// ipsecprofile
// Configuration for IPSEC profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ipsec/ipsecprofile
func (s *IPSECService) AddIPSECProfile()    {}
func (s *IPSECService) DeleteIPSECProfile() {}
func (s *IPSECService) GetAllIPSECProfile() {}
func (s *IPSECService) GetIPSECProfile()    {}
func (s *IPSECService) CountIPSECProfile()  {}
