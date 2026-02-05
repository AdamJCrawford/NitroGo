package nitrogo

const (
	smppParamURL = "/nitro/v1/config/smppparam"
	smppUserURL  = "/nitro/v1/config/smppuser"
)

// All the commands associated with SMPP.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/smpp/smpp
type SMPPService struct {
	client *Client
}

// smppparam
// Configuration for SMPP configuration parameters resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/smpp/smppparam
func (s *SMPPService) UpdateSMPPParam() {}
func (s *SMPPService) UnsetSMPPParam()  {}
func (s *SMPPService) GetAllSMPPParam() {}

// smppuser
// Configuration for SMPP user resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/smpp/smppuser
func (s *SMPPService) AddSMPPUser()    {}
func (s *SMPPService) DeleteSMPPUser() {}
func (s *SMPPService) UpdateSMPPUser() {}
func (s *SMPPService) GetAllSMPPUser() {}
func (s *SMPPService) GetSMPPUser()    {}
func (s *SMPPService) CountSMPPUser()  {}
