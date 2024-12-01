package nitrogo

// RDP configuration
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/rdp/rdp
type RDPService struct {
	client *Client
}

// rdpclientprofile
// Configuration for RDP clientprofile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/rdp/rdpclientprofile
func (s *RDPService) AddRDPClientProfile()    {}
func (s *RDPService) DeleteRDPClientProfile() {}
func (s *RDPService) UpdateRDPClientProfile() {}
func (s *RDPService) UnsetRDPClientProfile()  {}
func (s *RDPService) GetAllRDPClientProfile() {}
func (s *RDPService) GetRDPClientProfile()    {}
func (s *RDPService) CountRDPClientProfile()  {}

// rdpconnections
// Configuration for active rdp connections resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/rdp/rdpconnections
func (s *RDPService) GetAllRDPConnections() {}
func (s *RDPService) CountRDPConnections()  {}
func (s *RDPService) KillRDPConnections()   {}

// rspserverprofile
// Configuration for RDP serverprofile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/rdp/rdpserverprofile
func (s *RDPService) AddRDPServerProfile()    {}
func (s *RDPService) DeleteRDPServerProfile() {}
func (s *RDPService) UpdateRDPServerProfile() {}
func (s *RDPService) UnsetRDPServerProfile()  {}
func (s *RDPService) GetAllRDPServerProfile() {}
func (s *RDPService) GetRDPServerProfile()    {}
func (s *RDPService) CountRDPServerProfile()  {}
