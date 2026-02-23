package nitrogo

const (
	ipsecALGProfileURL = "/nitro/v1/config/ipsecalgprofile"
	ipsecALGSessionURL = "/nitro/v1/config/ipsecalgsession"
)

// IPSECALG configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ipsecalg/ipsecalg
type IPSECALGService struct {
	client *Client
}

// ipsecalgprofile
// Configuration for IPSEC ALG profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ipsecalg/ipsecalgprofile
func (s *IPSECALGService) AddIPSECALGProfile()    {}
func (s *IPSECALGService) DeleteIPSECALGProfile() {}
func (s *IPSECALGService) UpdateIPSECALGProfile() {}
func (s *IPSECALGService) UnsetIPSECALGProfile()  {}
func (s *IPSECALGService) GetAllIPSECALGProfile() {}
func (s *IPSECALGService) GetIPSECALGProfile()    {}
func (s *IPSECALGService) CountIPSECALGProfile()  {}

// ipsecalgsession
// Configuration for IPSEC ALG session resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ipsecalg/ipsecalgsession
func (s *IPSECALGService) GetAllIPSECALGSession() {}
func (s *IPSECALGService) CountIPSECALGSession()  {}
func (s *IPSECALGService) FlushIPSECALGSession()  {}
