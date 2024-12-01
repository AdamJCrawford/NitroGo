package nitrogo

// Router configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/router/router
type RouterService struct {
	client *Client
}

// routerdynamicrouting
// Configuration for dynamic routing config resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/router/routerdynamicrouting
func (s *RouterService) AddRouterDynamicRouting()    {}
func (s *RouterService) DeleteRouterDynamicRouting() {}
func (s *RouterService) UpdateRouterDynamicRouting() {}
func (s *RouterService) UnsetRouterDynamicRouting()  {}
func (s *RouterService) GetAllRouterDynamicRouting() {}
func (s *RouterService) CountRouterDynamicRouting()  {}
func (s *RouterService) ApplyRouterDynamicRouting()  {}
