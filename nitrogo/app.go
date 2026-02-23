package nitrogo

const (
	applicationURL = "/nitro/v1/config/application"
)

// Application
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/app/app
type AppService struct {
	client *Client
}

// application
// Configuration for application resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/app/application
func (s *AppService) ImportApplication() {}
func (s *AppService) ExportApplication() {}
func (s *AppService) DeleteApplication() {}
