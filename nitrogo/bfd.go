package nitrogo

const (
	bfdSessionURL = "/nitro/v1/config/bfdsession"
)

// Bidirectional Forwarding Detection
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/bfd/bfd
type BFDService struct {
	client *Client
}

// bfdsession
// Configuration for BFD configuration resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/bfd/bfdsession
func (s *BFDService) GetAllBFDSession() {}
func (s *BFDService) CountBFDSession()  {}
