package nitrogo

// Subscriber configuration
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/subscriber/subscriber
type SubscriberService struct {
	client *Client
}

// subscribergxinterface
// Configuration for Gx interface Parameters resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/subscriber/subscribergxinterface
func (s *SubscriberService) UpdateSubscriberGxInterface() {}
func (s *SubscriberService) UnsetSubscriberGxInterface()  {}
func (s *SubscriberService) GetAllSubscriberGxInterface() {}

// subscriberparam
// Configuration for Subscriber Params resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/subscriber/subscriberparam
func (s *SubscriberService) UpdateSubscriberParam() {}
func (s *SubscriberService) UnsetSubscriberParam()  {}
func (s *SubscriberService) GetAllSubscriberParam() {}

// subscriberprofile
// Configuration for Subscriber Profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/subscriber/subscriberprofile
func (s *SubscriberService) AddSubscriberProfile()    {}
func (s *SubscriberService) DeleteSubscriberProfile() {}
func (s *SubscriberService) UpdateSubscriberProfile() {}
func (s *SubscriberService) UnsetSubscriberProfile()  {}
func (s *SubscriberService) GetAllSubscriberProfile() {}
func (s *SubscriberService) CountSubscriberProfile()  {}

// subscriverradiusinterface
// Configuration for RADIUS interface Parameters resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/subscriber/subscriberradiusinterface
func (s *SubscriberService) UpdateSubscriberRADIUSInterface() {}
func (s *SubscriberService) UnsetSubscriberRADIUSInterface()  {}
func (s *SubscriberService) GetAllSubscriberRADIUSInterface() {}

// subscribersessions
// Configuration for subscriber sesions resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/subscriber/subscribersessions
func (s *SubscriberService) ClearSubscriberSessions()  {}
func (s *SubscriberService) GetAllSubscriberSessions() {}
func (s *SubscriberService) CountSubscriberSessions()  {}
