package nitrogo

type AppService struct {
	client *Client
}

// application
func (s *AppService) ImportApplication() {}
func (s *AppService) ExportApplication() {}
func (s *AppService) DeleteApplication() {}
