package nitrogo

const (
	callHomeURL    = "/nitro/v1/config/callhome"
	installURL     = "/nitro/v1/config/install"
	pingURL        = "/nitro/v1/config/ping"
	ping6URL       = "/nitro/v1/config/ping6"
	raidURL        = "/nitro/v1/config/raid"
	techSupportURL = "/nitro/v1/config/techsupport"
	traceRouteURL  = "/nitro/v1/config/traceroute"
	traceRoute6URL = "/nitro/v1/config/traceroute6"
)

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
