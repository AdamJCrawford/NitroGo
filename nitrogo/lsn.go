package nitrogo

const (
	lsnAppsAttributesURL                      = "/nitro/v1/config/lsnappsattributes"
	lsnAppsProfileURL                         = "/nitro/v1/config/lsnappsprofile"
	lsnAppsProfileBindingURL                  = "/nitro/v1/config/lsnappsprofile_binding"
	lsnAppsProfileLSNAppsAttributesBindingURL = "/nitro/v1/config/lsnappsprofile_lsnappsattributes_binding"
	lsnAppsProfilePortBindingURL              = "/nitro/v1/config/lsnappsprofile_port_binding"
	lsnClientURL                              = "/nitro/v1/config/lsnclient"
	lsnClientBindingURL                       = "/nitro/v1/config/lsnclient_binding"
	lsnClientNetwork6BindingURL               = "/nitro/v1/config/lsnclient_network6_binding"
	lsnClientNetworkBindingURL                = "/nitro/v1/config/lsnclient_network_binding"
	lsnClientNSACL6BindingURL                 = "/nitro/v1/config/lsnclient_nsacl6_binding"
	lsnClientNSACLBindingURL                  = "/nitro/v1/config/lsnclient_nsacl_binding"
	lsnDeterministicNATURL                    = "/nitro/v1/config/lsndeterministicnat"
	lsnGroupURL                               = "/nitro/v1/config/lsngroup"
	lsnGroupBindingURL                        = "/nitro/v1/config/lsngroup_binding"
	lsnGroupIPSecALGProfileBindingURL         = "/nitro/v1/config/lsngroup_ipsecalgprofile_binding"
	lsnGroupLSNAppsProfileBindingURL          = "/nitro/v1/config/lsngroup_lsnappsprofile_binding"
	lsnGroupLSNHTTPHDRLogProfileBindingURL    = "/nitro/v1/config/lsngroup_lsnhttphdrlogprofile_binding"
	lsnGroupLSNLogProfileBindingURL           = "/nitro/v1/config/lsngroup_lsnlogprofile_binding"
	lsnGroupLSNPoolBindingURL                 = "/nitro/v1/config/lsngroup_lsnpool_binding"
	lsnGroupLSNRTSPALGProfileBindingURL       = "/nitro/v1/config/lsngroup_lsnrtspalgprofile_binding"
	lsnGroupLSNSIPALGProfileBindingURL        = "/nitro/v1/config/lsngroup_lsnsipalgprofile_binding"
	lsnGroupLSNTransportProfileBindingURL     = "/nitro/v1/config/lsngroup_lsntransportprofile_binding"
	lsnGroupPCPServerBindingURL               = "/nitro/v1/config/lsngroup_pcpserver_binding"
	lsnHTTPHDRLogProfileURL                   = "/nitro/v1/config/lsnhttphdrlogprofile"
	lsnIP6ProfileURL                          = "/nitro/v1/config/lsnip6profile"
	lsnLogProfileURL                          = "/nitro/v1/config/lsnlogprofile"
	lsnParameterURL                           = "/nitro/v1/config/lsnparameter"
	lsnPoolURL                                = "/nitro/v1/config/lsnpool"
	lsnPoolBindingURL                         = "/nitro/v1/config/lsnpool_binding"
	lsnPoolLSNIPBindingURL                    = "/nitro/v1/config/lsnpool_lsnip_binding"
	lsnRTSPALGProfileURL                      = "/nitro/v1/config/lsnrtspalgprofile"
	lsnRTSPALGSessionURL                      = "/nitro/v1/config/lsnrtspalgsession"
	lsnRTSPALGSessionBindingURL               = "/nitro/v1/config/lsnrtspalgsession_binding"
	lsnRTSPALGSessionDataChannelBindingURL    = "/nitro/v1/config/lsnrtspalgsession_datachannel_binding"
	lsnSessionURL                             = "/nitro/v1/config/lsnsession"
	lsnSIPALGCallURL                          = "/nitro/v1/config/lsnsipalgcall"
	lsnSIPALGCallBindingURL                   = "/nitro/v1/config/lsnsipalgcall_binding"
	lsnSIPALGCallControlChannelBindingURL     = "/nitro/v1/config/lsnsipalgcall_controlchannel_binding"
	lsnSIPALGCallDataChannelBindingURL        = "/nitro/v1/config/lsnsipalgcall_datachannel_binding"
	lsnSIPALGProfileURL                       = "/nitro/v1/config/lsnsipalgprofile"
	lsnStaticURL                              = "/nitro/v1/config/lsnstatic"
	lsnTransportProfileURL                    = "/nitro/v1/config/lsntransportprofile"
)

// Large Scale NAT commands
type LSNService struct {
	client *Client
}

// lsnappsattributes
func (s *LSNService) AddLSNAppsAttributes()    {}
func (s *LSNService) DeleteLSNAppsAttributes() {}
func (s *LSNService) UpdateLSNAppsAttributes() {}
func (s *LSNService) UnsetLSNAppsAttributes()  {}
func (s *LSNService) GetAllLSNAppsAttributes() {}
func (s *LSNService) GetLSNAppsAttributes()    {}
func (s *LSNService) CountLSNAppsAttributes()  {}

// lsnappsprofile
func (s *LSNService) AddLSNAppsProfile()    {}
func (s *LSNService) DeleteLSNAppsProfile() {}
func (s *LSNService) UpdateLSNAppsProfile() {}
func (s *LSNService) UnsetLSNAppsProfile()  {}
func (s *LSNService) GetAllLSNAppsProfile() {}
func (s *LSNService) GetLSNAppsProfile()    {}
func (s *LSNService) CountLSNAppsProfile()  {}

// lsnappsprofile_binding
func (s *LSNService) GetAllLSNAppsProfileBinding() {}
func (s *LSNService) GetLSNAppsProfileBinding()    {}

// lsnappsprofile_lsnappsattributes_binding
func (s *LSNService) AddLSNAppsProfileLSNAppsAttributesBinding()    {}
func (s *LSNService) DeleteLSNAppsProfileLSNAppsAttributesBinding() {}
func (s *LSNService) GetAllLSNAppsProfileLSNAppsAttributesBinding() {}
func (s *LSNService) GetLSNAppsProfileLSNAppsAttributesBinding()    {}
func (s *LSNService) CountLSNAppsProfileLSNAppsAttributesBinding()  {}

// lsnappsprofile_port_binding
func (s *LSNService) AddLSNAppsProfilePortBinding()    {}
func (s *LSNService) DeleteLSNAppsProfilePortBinding() {}
func (s *LSNService) GetAllLSNAppsProfilePortBinding() {}
func (s *LSNService) GetLSNAppsProfilePortBinding()    {}
func (s *LSNService) CountLSNAppsProfilePortBinding()  {}

// lsnclient
func (s *LSNService) AddLSNClient()    {}
func (s *LSNService) DeleteLSNClient() {}
func (s *LSNService) GetAllLSNClient() {}
func (s *LSNService) GetLSNClient()    {}
func (s *LSNService) CountLSNClient()  {}

// lsnclient_binding
func (s *LSNService) GetAllLSNClientBinding() {}
func (s *LSNService) GetLSNClientBinding()    {}

// lsnclient_network6_binding
func (s *LSNService) AddLSNClientNetwork6Binding()    {}
func (s *LSNService) DeleteLSNClientNetwork6Binding() {}
func (s *LSNService) GetAllLSNClientNetwork6Binding() {}
func (s *LSNService) GetLSNClientNetwork6Binding()    {}
func (s *LSNService) CountLSNClientNetwork6Binding()  {}

// lsnclient_network_binding
func (s *LSNService) AddLSNClientNetworkBinding()    {}
func (s *LSNService) DeleteLSNClientNetworkBinding() {}
func (s *LSNService) GetAllLSNClientNetworkBinding() {}
func (s *LSNService) GetLSNClientNetworkBinding()    {}
func (s *LSNService) CountLSNClientNetworkBinding()  {}

// lsnclient_nsacl6_binding
func (s *LSNService) AddLSNClientNSACL6Binding()    {}
func (s *LSNService) DeleteLSNClientNSACL6Binding() {}
func (s *LSNService) GetAllLSNClientNSACL6Binding() {}
func (s *LSNService) GetLSNClientNSACL6Binding()    {}
func (s *LSNService) CountLSNClientNSACL6Binding()  {}

// lsnclient_nsacl_binding
func (s *LSNService) AddLSNClientNSACLBinding()    {}
func (s *LSNService) DeleteLSNClientNSACLBinding() {}
func (s *LSNService) GetAllLSNClientNSACLBinding() {}
func (s *LSNService) GetLSNClientNSACLBinding()    {}
func (s *LSNService) CountLSNClientNSACLBinding()  {}

// lsndeterministicnat
func (s *LSNService) GetAllLSNDeterministicNAT() {}
func (s *LSNService) CountLSNDeterministicNAT()  {}

// lsngroup
func (s *LSNService) AddLSNGroup()    {}
func (s *LSNService) DeleteLSNGroup() {}
func (s *LSNService) UpdateLSNGroup() {}
func (s *LSNService) UnsetLSNGroup()  {}
func (s *LSNService) GetAllLSNGroup() {}
func (s *LSNService) GetLSNGroup()    {}
func (s *LSNService) CountLSNGroup()  {}

// lsngroup_binding
func (s *LSNService) GetAllLSNGroupBinding() {}
func (s *LSNService) GetLSNGroupBinding()    {}

// lsngroup_ipsecalgprofile_binding
func (s *LSNService) AddLSNGroupIPSecALGProfileBinding()    {}
func (s *LSNService) DeleteLSNGroupIPSecALGProfileBinding() {}
func (s *LSNService) GetAllLSNGroupIPSecALGProfileBinding() {}
func (s *LSNService) GetLSNGroupIPSecALGProfileBinding()    {}
func (s *LSNService) CountLSNGroupIPSecALGProfileBinding()  {}

// lsngroup_lsnappsprofile_binding
func (s *LSNService) AddLSNGroupLSNAppsProfileBinding()    {}
func (s *LSNService) DeleteLSNGroupLSNAppsProfileBinding() {}
func (s *LSNService) GetAllLSNGroupLSNAppsProfileBinding() {}
func (s *LSNService) GetLSNGroupLSNAppsProfileBinding()    {}
func (s *LSNService) CountLSNGroupLSNAppsProfileBinding()  {}

// lsngroup_lsnhttphdrlogprofile_binding
func (s *LSNService) AddLSNGroupLSNHTTPHDRLogProfileBinding()    {}
func (s *LSNService) DeleteLSNGroupLSNHTTPHDRLogProfileBinding() {}
func (s *LSNService) GetAllLSNGroupLSNHTTPHDRLogProfileBinding() {}
func (s *LSNService) GetLSNGroupLSNHTTPHDRLogProfileBinding()    {}
func (s *LSNService) CountLSNGroupLSNHTTPHDRLogProfileBinding()  {}

// lsngroup_lsnlogprofile_binding
func (s *LSNService) AddLSNGroupLSNLogProfileBinding()    {}
func (s *LSNService) DeleteLSNGroupLSNLogProfileBinding() {}
func (s *LSNService) GetAllLSNGroupLSNLogProfileBinding() {}
func (s *LSNService) GetLSNGroupLSNLogProfileBinding()    {}
func (s *LSNService) CountLSNGroupLSNLogProfileBinding()  {}

// lsngroup_lsnpool_binding
func (s *LSNService) AddLSNGroupLSNPoolBinding()    {}
func (s *LSNService) DeleteLSNGroupLSNPoolBinding() {}
func (s *LSNService) GetAllLSNGroupLSNPoolBinding() {}
func (s *LSNService) GetLSNGroupLSNPoolBinding()    {}
func (s *LSNService) CountLSNGroupLSNPoolBinding()  {}

// lsngroup_lsnrtspalgprofile_binding
func (s *LSNService) AddLSNGroupLSNRTSPALGProfileBinding()    {}
func (s *LSNService) DeleteLSNGroupLSNRTSPALGProfileBinding() {}
func (s *LSNService) GetAllLSNGroupLSNRTSPALGProfileBinding() {}
func (s *LSNService) GetLSNGroupLSNRTSPALGProfileBinding()    {}
func (s *LSNService) CountLSNGroupLSNRTSPALGProfileBinding()  {}

// lsngroup_lsnsipalgprofile_binding
func (s *LSNService) AddKSBGroupLSNSIPALGProfileBinding()    {}
func (s *LSNService) DeleteKSBGroupLSNSIPALGProfileBinding() {}
func (s *LSNService) GetAllKSBGroupLSNSIPALGProfileBinding() {}
func (s *LSNService) GetKSBGroupLSNSIPALGProfileBinding()    {}
func (s *LSNService) CountKSBGroupLSNSIPALGProfileBinding()  {}

// lsngroup_lsntransportprofile_binding
func (s *LSNService) AddLSNGroupLSNTransportProfileBinding()    {}
func (s *LSNService) DeleteLSNGroupLSNTransportProfileBinding() {}
func (s *LSNService) GetAllLSNGroupLSNTransportProfileBinding() {}
func (s *LSNService) GetLSNGroupLSNTransportProfileBinding()    {}
func (s *LSNService) CountLSNGroupLSNTransportProfileBinding()  {}

// lsngroup_pcpserver_binding
func (s *LSNService) AddLSNGroupPCPServerBinding()    {}
func (s *LSNService) DeleteLSNGroupPCPServerBinding() {}
func (s *LSNService) GetAllLSNGroupPCPServerBinding() {}
func (s *LSNService) GetLSNGroupPCPServerBinding()    {}
func (s *LSNService) CountLSNGroupPCPServerBinding()  {}

// lsnhttphdrlogprofile
func (s *LSNService) AddLSNHTTPHDRLogProfile()    {}
func (s *LSNService) DeleteLSNHTTPHDRLogProfile() {}
func (s *LSNService) UpdateLSNHTTPHDRLogProfile() {}
func (s *LSNService) UnsetLSNHTTPHDRLogProfile()  {}
func (s *LSNService) GetAllLSNHTTPHDRLogProfile() {}
func (s *LSNService) GetLSNHTTPHDRLogProfile()    {}
func (s *LSNService) CountLSNHTTPHDRLogProfile()  {}

// lsnip6profile
func (s *LSNService) AddLSNIP6Profile()    {}
func (s *LSNService) DeleteLSNIP6Profile() {}
func (s *LSNService) GetAllLSNIP6Profile() {}
func (s *LSNService) GetLSNIP6Profile()    {}
func (s *LSNService) CountLSNIP6Profile()  {}

// lsnlogprofile
func (s *LSNService) AddLSNLogProfile()    {}
func (s *LSNService) DeleteLSNLogProfile() {}
func (s *LSNService) UpdateLSNLogProfile() {}
func (s *LSNService) UnsetLSNLogProfile()  {}
func (s *LSNService) GetAllLSNLogProfile() {}
func (s *LSNService) GetLSNLogProfile()    {}
func (s *LSNService) CountLSNLogProfile()  {}

// lsnparameter
func (s *LSNService) UpdateLSNParameter() {}
func (s *LSNService) UnsetLSNParameter()  {}
func (s *LSNService) GetAllLSNParameter() {}

// lsnpool
func (s *LSNService) AddLSNPool()    {}
func (s *LSNService) DeleteLSNPool() {}
func (s *LSNService) UpdateLSNPool() {}
func (s *LSNService) UnsetLSNPool()  {}
func (s *LSNService) GetAllLSNPool() {}
func (s *LSNService) GetLSNPool()    {}
func (s *LSNService) CountLSNPool()  {}

// lsnpool_binding
func (s *LSNService) GetAllLSNPoolBinding() {}
func (s *LSNService) GetLSNPoolBinding()    {}

// lsnpool_lsnip_binding
func (s *LSNService) AddLSNPoolLSNIPBinding()    {}
func (s *LSNService) DeleteLSNPoolLSNIPBinding() {}
func (s *LSNService) GetAllLSNPoolLSNIPBinding() {}
func (s *LSNService) GetLSNPoolLSNIPBinding()    {}
func (s *LSNService) CountLSNPoolLSNIPBinding()  {}

// lsnrtspalgprofile
func (s *LSNService) AddLSNRTSPALGProfile()    {}
func (s *LSNService) DeleteLSNRTSPALGProfile() {}
func (s *LSNService) UpdateLSNRTSPALGProfile() {}
func (s *LSNService) UnsetLSNRTSPALGProfile()  {}
func (s *LSNService) GetAllLSNRTSPALGProfile() {}
func (s *LSNService) GetLSNRTSPALGProfile()    {}
func (s *LSNService) CountLSNRTSPALGProfile()  {}

// lsnrtspalgsession
func (s *LSNService) GetAllLSNRTSPALGSession() {}
func (s *LSNService) GetLSNRTSPALGSession()    {}
func (s *LSNService) CountLSNRTSPALGSession()  {}
func (s *LSNService) FlushLSNRTSPALGSession()  {}

// lsnrtspalgsession_binding
func (s *LSNService) GetAllLSNRTSPALGSessionBinding() {}
func (s *LSNService) GetLSNRTSPALGSessionBinding()    {}

// lsnrtspalgsession_datachannel_binding
func (s *LSNService) GetAllLSNRTSPALGSessionDataChannelBinding() {}
func (s *LSNService) GetLSNRTSPALGSessionDataChannelBinding()    {}
func (s *LSNService) CountLSNRTSPALGSessionDataChannelBinding()  {}

// lsnsession
func (s *LSNService) GetAllLSNSession() {}
func (s *LSNService) CountLSNSession()  {}
func (s *LSNService) FlushLSNSession()  {}

// lsnsipalgcall
func (s *LSNService) GetAllLSNSIPALGCall() {}
func (s *LSNService) GetLSNSIPALGCall()    {}
func (s *LSNService) CountLSNSIPALGCall()  {}
func (s *LSNService) FlushLSNSIPALGCall()  {}

// lsnsipalgcall_binding
func (s *LSNService) GetAllLSNSIPALGCallBinding() {}
func (s *LSNService) GetLSNSIPALGCallBinding()    {}

// lsnsipalgcall_controlchannel_binding
func (s *LSNService) GetAllLSNSIPALGCallControlChannelBinding() {}
func (s *LSNService) GetLSNSIPALGCallControlChannelBinding()    {}
func (s *LSNService) CountLSNSIPALGCallControlChannelBinding()  {}

// lsnsipalgcall_datachannel_binding
func (s *LSNService) GetAllLSNSIPALGCallDataChannelBinding() {}
func (s *LSNService) GetLSNSIPALGCallDataChannelBinding()    {}
func (s *LSNService) CountLSNSIPALGCallDataChannelBinding()  {}

// lsnsipalgprofile
func (s *LSNService) AddLSNSIPALGProfile()    {}
func (s *LSNService) DeleteLSNSIPALGProfile() {}
func (s *LSNService) UpdateLSNSIPALGProfile() {}
func (s *LSNService) UnsetLSNSIPALGProfile()  {}
func (s *LSNService) GetAllLSNSIPALGProfile() {}
func (s *LSNService) GetLSNSIPALGProfile()    {}
func (s *LSNService) CountLSNSIPALGProfile()  {}

// lsnstatic
func (s *LSNService) AddLSNStatic()    {}
func (s *LSNService) DeleteLSNStatic() {}
func (s *LSNService) GetAllLSNStatic() {}
func (s *LSNService) GetLSNStatic()    {}
func (s *LSNService) CountLSNStatic()  {}

// lsntransportprofile
func (s *LSNService) AddLSNTransportProfile()    {}
func (s *LSNService) DeleteLSNTransportProfile() {}
func (s *LSNService) UpdateLSNTransportProfile() {}
func (s *LSNService) UnsetLSNTransportProfile()  {}
func (s *LSNService) GetAllLSNTransportProfile() {}
func (s *LSNService) GetLSNTransportProfile()    {}
func (s *LSNService) CountLSNTransportProfile()  {}
