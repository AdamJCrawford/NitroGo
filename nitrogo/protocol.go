package nitrogo

const (
	protocolHTTPBandURL = "/nitro/v1/config/protocolhttpband"
)

// Protocol Configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/protocol/protocol
type ProtocolService struct {
	client *Client
}

// protocolhttpband
// Configuration for HTTP request/response band resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/protocol/protocolhttpband
func (s *ProtocolService) UpdateProtocolHTTPBand() {}
func (s *ProtocolService) UnsetProtocolHTTPBand()  {}
func (s *ProtocolService) GetAllProtocolHTTPBand() {}
func (s *ProtocolService) ClearProtocolHTTPBand()  {}
