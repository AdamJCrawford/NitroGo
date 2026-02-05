package nitrogo

const (
	tunnelGlobalBindingURL                    = "/nitro/v1/config/tunnelglobal_binding"
	tunnelGlobalTunnelTrafficPolicyBindingURL = "/nitro/v1/config/tunnelglobal_tunneltrafficpolicy_binding"
	tunnelTrafficPolicyURL                    = "/nitro/v1/config/tunneltrafficpolicy"
	tunnelTrafficPolicyBindingURL             = "/nitro/v1/config/tunneltrafficpolicy_binding"
	tunnelTrafficPolicyTunnelGlobalBindingURL = "/nitro/v1/config/tunneltrafficpolicy_tunnelglobal_binding"
)

// SSL VPN Tunnel Configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tunnel/tunnel
type TunnelService struct {
	client *Client
}

// tunnelglobal_binding
// Binding object which returns the resources bound to tunnelglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tunnel/tunnelglobal_binding
func (s *TunnelService) GetTunnelGlobalBinding() {}

// tunnelglobal_tunneltrafficpolicy_binding
// Binding object showing the tunneltrafficpolicy that can be bound to tunnelglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tunnel/tunnelglobal_tunneltrafficpolicy_binding
func (s *TunnelService) AddTunnelGlobalTunnelTrafficPolicyBinding()    {}
func (s *TunnelService) DeleteTunnelGlobalTunnelTrafficPolicyBinding() {}
func (s *TunnelService) GetTunnelGlobalTunnelTrafficPolicyBinding()    {}
func (s *TunnelService) CountTunnelGlobalTunnelTrafficPolicyBinding()  {}

// tunneltrafficpolicy
// Configuration for tunnel policy resource
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tunnel/tunneltrafficpolicy
func (s *TunnelService) AddTunnelTrafficPolicy()    {}
func (s *TunnelService) DeleteTunnelTrafficPolicy() {}
func (s *TunnelService) UpdateTunnelTrafficPolicy() {}
func (s *TunnelService) UnsetTunnelTrafficPolicy()  {}
func (s *TunnelService) GetAllTunnelTrafficPolicy() {}
func (s *TunnelService) GetTunnelTrafficPolicy()    {}
func (s *TunnelService) CountTunnelTrafficPolicy()  {}
func (s *TunnelService) RenameTunnelTrafficPolicy() {}

// tunneltrafficpolicy_binding
// Binding object which returns the resources bound to tunneltrafficpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tunnel/tunneltrafficpolicy_binding
func (s *TunnelService) GetAllTunnelTrafficPolicyBinding() {}
func (s *TunnelService) GetTunnelTrafficPolicyBinding()    {}

// tunneltrafficpolicy_tunnelglobal_binding
// Binding object showing the tunnelglobal that can be bound to tunneltrafficpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tunnel/tunneltrafficpolicy_tunnelglobal_binding
func (s *TunnelService) GetAllTunnelTrafficPolicyTunnelGlobalBinding() {}
func (s *TunnelService) GetTunnelTrafficPolicyTunnelGlobalBinding()    {}
func (s *TunnelService) CountTunnelTrafficPolicyTunnelGlobalBinding()  {}
