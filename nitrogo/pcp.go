package nitrogo

// Port Control Protocol Configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/pcp/pcp
type PCPService struct {
	client *Client
}

// pcpmap
// Configuration for server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/pcp/pcpmap
func (s *PCPService) GetAllPCPMap() {}
func (s *PCPService) CountPCPMap()  {}

// pcpprofile
// Configuration for PCP Profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/pcp/pcpprofile
func (s *PCPService) AddPCPProfile()    {}
func (s *PCPService) DeletePCPProfile() {}
func (s *PCPService) UpdatePCPProfile() {}
func (s *PCPService) UnsetPCPProfile()  {}
func (s *PCPService) GetAllPCPProfile() {}
func (s *PCPService) GetPCPProfile()    {}
func (s *PCPService) CountPCPProfile()  {}

// pcpserver
// Configuration for server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/pcp/pcpserver
func (s *PCPService) AddPCPServer()    {}
func (s *PCPService) DeletePCPServer() {}
func (s *PCPService) UpdatePCPServer() {}
func (s *PCPService) UnsetPCPServer()  {}
func (s *PCPService) GetAllPCPServer() {}
func (s *PCPService) GetPCPServer()    {}
func (s *PCPService) CountPCPServer()  {}
