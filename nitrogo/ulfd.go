package nitrogo

// ULFD server configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ulfd/ulfd
type ULFDService struct {
	client *Client
}

// ulfdserver
// Configuration for ulfd server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ulfd/ulfdserver
func (s *ULFDService) AddULFDServer()    {}
func (s *ULFDService) DeleteULFDServer() {}
func (s *ULFDService) GetAllULFDServer() {}
func (s *ULFDService) GetULFDServer()    {}
func (s *ULFDService) CountULFDServer()  {}
