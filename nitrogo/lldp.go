package nitrogo

const (
	lldpNeighborsURL = "/nitro/v1/config/lldpneighbors"
	lldpParamURL     = "/nitro/v1/config/lldpparam"
)

// Link Layer Discovery Protocol.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/lldp/lldp
type LLDPService struct {
	client *Client
}

// lldpneighbors
// Configuration for lldp neighbors resource
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/lldp/lldpneighbors
func (s *LLDPService) GetAllLLDPNeighbors() {}
func (s *LLDPService) GetLLDPNeighbors()    {}
func (s *LLDPService) CountLLDPNeighbors()  {}
func (s *LLDPService) ClearLLDPNeighbors()  {}

// lldpparam
// Configuration for lldp params resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/lldp/lldpparam
func (s *LLDPService) UpdateLLDPParam() {}
func (s *LLDPService) UnsetLLDPParam()  {}
func (s *LLDPService) GetAllLLDPParam() {}
