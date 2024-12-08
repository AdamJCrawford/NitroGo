package nitrogo

// Utilities.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/utility/utility
type UtilityService struct {
	client *Client
}

// callhome
func (s *UtilityService) UpdateCallHome() {}
func (s *UtilityService) UnsetCallHome()  {}
func (s *UtilityService) GetAllCallHome() {}

// install
func (s *UtilityService) InstallInstall() {}

// ping
func (s *UtilityService) PingPing() {}

// ping6
func (s *UtilityService) Ping6Ping6() {}

// raid
func (s *UtilityService) GetAllRAID() {}

// techsupport
func (s *UtilityService) GetAllTechSupport() {}

// traceroute
func (s *UtilityService) TraceRouteTraceRoute() {}

// traceroute6
func (s *UtilityService) TraceRoute6TraceRoute6() {}
