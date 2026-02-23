package nitrogo

const (
	reputationSettingsURL = "/nitro/v1/config/reputationsettings"
)

// Reputation services configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/reputation/reputation
type ReputationService struct {
	client *Client
}

// reputationsettings
// Configuration for Reputation service settings resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/reputation/reputationsettings
func (s *ReputationService) UpdateReputationSettings() {}
func (s *ReputationService) UnsetReputationSettings()  {}
func (s *ReputationService) GetAllReputationSettings() {}
