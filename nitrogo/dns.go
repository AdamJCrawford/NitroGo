package nitrogo

const (
	dnsAAAARecURL                         = "/nitro/v1/config/dnsaaaarec"
	dnsActionURL                          = "/nitro/v1/config/dnsaction"
	dnsAction64URL                        = "/nitro/v1/config/dnsaction64"
	dnsAddRecURL                          = "/nitro/v1/config/dnsaddrec"
	dnsCNameRecURL                        = "/nitro/v1/config/dnscnamerec"
	dnsGlobalBindingURL                   = "/nitro/v1/config/dnsglobal_binding"
	dnsGlobalDNSPolicyBindingURL          = "/nitro/v1/config/dnsglobal_dnspolicy_binding"
	dnsKeyURL                             = "/nitro/v1/config/dnskey"
	dnsMXRecURL                           = "/nitro/v1/config/dnsmxrec"
	dnsNameServerURL                      = "/nitro/v1/config/dnsnameserver"
	dnsNAPTRRecURL                        = "/nitro/v1/config/dnsnaptrrec"
	dnsNSECRecURL                         = "/nitro/v1/config/dnsnsecrec"
	dnsNSRecURL                           = "/nitro/v1/config/dnsnsrec"
	dnsParameterURL                       = "/nitro/v1/config/dnsparameter"
	dnsPolicyURL                          = "/nitro/v1/config/dnspolicy"
	dnsPolicy64URL                        = "/nitro/v1/config/dnspolicy64"
	dnsPolicy64BindingURL                 = "/nitro/v1/config/dnspolicy64_binding"
	dnsPolicy64LBVServerBindingURL        = "/nitro/v1/config/dnspolicy64_lbvserver_binding"
	dnsPolicyLabelURL                     = "/nitro/v1/config/dnspolicylabel"
	dnsPolicyLabelBindingURL              = "/nitro/v1/config/dnspolicylabel_binding"
	dnsPolicyLabelDNSPolicyBindingURL     = "/nitro/v1/config/dnspolicylabel_dnspolicy_binding"
	dnsPolicyLabelPolicyBindingBindingURL = "/nitro/v1/config/dnspolicylabel_policybinding_binding"
	dnsPolicyBindingURL                   = "/nitro/v1/config/dnspolicy_binding"
	dnsPolicyDNSGlobalBindingURL          = "/nitro/v1/config/dnspolicy_dnsglobal_binding"
	dnsPolicyDNSPolicyLabelBindingURL     = "/nitro/v1/config/dnspolicy_dnspolicylabel_binding"
	dnsProfileURL                         = "/nitro/v1/config/dnsprofile"
	dnsProxyRecordsURL                    = "/nitro/v1/config/dnsproxyrecords"
	dnsPTRRecURL                          = "/nitro/v1/config/dnsptrrec"
	dnsSOARecURL                          = "/nitro/v1/config/dnssoarec"
	dnsSRVRecURL                          = "/nitro/v1/config/dnssrvrec"
	dnsSubnetCacheURL                     = "/nitro/v1/config/dnssubnetcache"
	dnsSuffixURL                          = "/nitro/v1/config/dnssuffix"
	dnsTXTRecURL                          = "/nitro/v1/config/dnstxtrec"
	dnsViewURL                            = "/nitro/v1/config/dnsview"
	dnsViewBindingURL                     = "/nitro/v1/config/dnsview_binding"
	dnsViewDNSPolicyBindingURL            = "/nitro/v1/config/dnsview_dnspolicy_binding"
	dnsViewGSLBServiceBindingURL          = "/nitro/v1/config/dnsview_gslbservice_binding"
	dnsZoneURL                            = "/nitro/v1/config/dnszone"
	dnsZoneBindingURL                     = "/nitro/v1/config/dnszone_binding"
	dnsZoneDNSKeyBindingURL               = "/nitro/v1/config/dnszone_dnskey_binding"
	dnsZoneDomainBindingURL               = "/nitro/v1/config/dnszone_domain_binding"
)

// Domain Name Service(DNS) configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/dns/dns
type DNSService struct {
	client *Client
}

// dnsaaaarec
func (s *DNSService) AddDNSAAAARec()    {}
func (s *DNSService) DeleteDNSAAAARec() {}
func (s *DNSService) GetAllDNSAAAARec() {}
func (s *DNSService) CountDNSAAAARec()  {}

// dnsaction
func (s *DNSService) AddDNSAction()    {}
func (s *DNSService) DeleteDNSAction() {}
func (s *DNSService) UpdateDNSAction() {}
func (s *DNSService) UnsetDNSAction()  {}
func (s *DNSService) GetAllDNSAction() {}
func (s *DNSService) GetDNSAction()    {}
func (s *DNSService) CountDNSAction()  {}

// dnsaction64
func (s *DNSService) AddDNSAction64()    {}
func (s *DNSService) DeleteDNSAction64() {}
func (s *DNSService) UpdateDNSAction64() {}
func (s *DNSService) UnsetDNSAction64()  {}
func (s *DNSService) GetAllDNSAction64() {}
func (s *DNSService) GetDNSAction64()    {}
func (s *DNSService) CountDNSAction64()  {}

// dnsaddrec
func (s *DNSService) AddDNSAddRec()    {}
func (s *DNSService) DeleteDNSAddRec() {}
func (s *DNSService) GetAllDNSAddRec() {}
func (s *DNSService) GetDNSAddRec()    {}
func (s *DNSService) CountDNSAddRec()  {}

// dnscnamerec
func (s *DNSService) AddDNSCNameRec()    {}
func (s *DNSService) DeleteDNSCNameRec() {}
func (s *DNSService) GetAllDNSCNameRec() {}
func (s *DNSService) GetDNSCNameRec()    {}
func (s *DNSService) CountDNSCNameRec()  {}

// dnsglobal_binding
func (s *DNSService) GetDNSGlobalBinding() {}

// dnsglobal_dnspolicy_binding
func (s *DNSService) AddDNSGlobalDNSPolicyBinding()    {}
func (s *DNSService) DeleteDNSGlobalDNSPolicyBinding() {}
func (s *DNSService) GetDNSGlobalDNSPolicyBinding()    {}
func (s *DNSService) CountDNSGlobalDNSPolicyBinding()  {}

// dnskey
func (s *DNSService) AddDNSKey()    {}
func (s *DNSService) DeleteDNSKey() {}
func (s *DNSService) UpdateDNSKey() {}
func (s *DNSService) UnsetDNSKey()  {}
func (s *DNSService) GetAllDNSKey() {}
func (s *DNSService) GetDNSKey()    {}
func (s *DNSService) CountDNSKey()  {}
func (s *DNSService) CreateDNSKey() {}
func (s *DNSService) ImportDNSKey() {}

// dnsmxrec
func (s *DNSService) AddDNSMXRec()    {}
func (s *DNSService) DeleteDNSMXRec() {}
func (s *DNSService) UpdateDNSMXRec() {}
func (s *DNSService) UnsetDNSMXRec()  {}
func (s *DNSService) GetAllDNSMXRec() {}
func (s *DNSService) GetDNSMXRec()    {}
func (s *DNSService) CountDNSMXRec()  {}

// dnsnameserver
func (s *DNSService) AddDNSNameServer()     {}
func (s *DNSService) DeleteDNSNameServer()  {}
func (s *DNSService) UpdateDNSNameServer()  {}
func (s *DNSService) UnsetDNSNameServer()   {}
func (s *DNSService) EnableDNSNameServer()  {}
func (s *DNSService) DisableDNSNameServer() {}
func (s *DNSService) GetAllDNSNameServer()  {}
func (s *DNSService) CountDNSNameServer()   {}

// dnsnaptrrec
func (s *DNSService) AddDNSNAPTRRec()    {}
func (s *DNSService) DeleteDNSNAPTRRec() {}
func (s *DNSService) GetAllDNSNAPTRRec() {}
func (s *DNSService) GetDNSNAPTRRec()    {}
func (s *DNSService) CountDNSNAPTRRec()  {}

// dnsnsecrec
func (s *DNSService) GetAllDNSNSECRec() {}
func (s *DNSService) GetDNSNSEcRec()    {}
func (s *DNSService) CountDNSNSECRec()  {}

// dnsnsrec
func (s *DNSService) AddDNSNSRec()    {}
func (s *DNSService) DeleteDNSNSRec() {}
func (s *DNSService) GetAllDNSNSRec() {}
func (s *DNSService) GetDNSNSRec()    {}
func (s *DNSService) CountDNSNSRec()  {}

// dnsparameter
func (s *DNSService) UpdateDNSParameter() {}
func (s *DNSService) UnsetDNSParameter()  {}
func (s *DNSService) GetAllDNSParameter() {}

// dnspolicy
func (s *DNSService) AddDNSPolicy()    {}
func (s *DNSService) DeleteDNSPolicy() {}
func (s *DNSService) UpdateDNSPolicy() {}
func (s *DNSService) UnsetDNSPolicy()  {}
func (s *DNSService) GetAllDNSPolicy() {}
func (s *DNSService) GetDNSPolicy()    {}
func (s *DNSService) CountDNSPolicy()  {}

// dnspolicy64
func (s *DNSService) AddDNSPolicy64()    {}
func (s *DNSService) DeleteDNSPolicy64() {}
func (s *DNSService) UpdateDNSPolicy64() {}
func (s *DNSService) GetAllDNSPolicy64() {}
func (s *DNSService) GetDNSPolicy64()    {}
func (s *DNSService) CountDNSPolicy64()  {}

// dnspolicy64_binding
func (s *DNSService) GetAllDNSPolicy64Binding() {}
func (s *DNSService) GetDNSPolicy64Binding()    {}

// dnspolicy64_lbvserver_binding
func (s *DNSService) GetAllDNSPolicy64LBVServerBinding() {}
func (s *DNSService) GetDNSPolicy64LBVServerBinding()    {}
func (s *DNSService) CountDNSPolicy64LBVServerBinding()  {}

// dnspolicylabel
func (s *DNSService) AddDNSPolicyLabel()    {}
func (s *DNSService) DeleteDNSPolicyLabel() {}
func (s *DNSService) GetAllDNSPolicyLabel() {}
func (s *DNSService) GetDNSPolicyLabel()    {}
func (s *DNSService) CountDNSPolicyLabel()  {}
func (s *DNSService) RenameDNSPolicyLabel() {}

// dnspolicylabel_binding
func (s *DNSService) GetAllDNSPolicyLabelBinding() {}
func (s *DNSService) GetDNSPolicyLabelBinding()    {}

// dnspolicylabel_dnspolicy_binding
func (s *DNSService) AddDNSPolicyLabelDNSPolicyBinding()    {}
func (s *DNSService) DeleteDNSPolicyLabelDNSPolicyBinding() {}
func (s *DNSService) GetAllDNSPolicyLabelDNSPolicyBinding() {}
func (s *DNSService) GetDNSPolicyLabelDNSPolicyBinding()    {}
func (s *DNSService) CountDNSPolicyLabelDNSPolicyBinding()  {}

// dnspolicylabel_policybinding_binding
func (s *DNSService) GetAllDNSPolicyLabelPolicyBindingBinding() {}
func (s *DNSService) GetDNSPolicyLabelPolicyBindingBinding()    {}
func (s *DNSService) CountDNSPolicyLabelPolicyBindingBinding()  {}

// dnspolicy_binding
func (s *DNSService) GetAllDNSPolicyBinding() {}
func (s *DNSService) GetDNSPolicyBinding()    {}

// dnspolicy_dnsglobal_binding
func (s *DNSService) GetAllDNSPolicyDNSGlobalBinding() {}
func (s *DNSService) GetDNSPolicyDNSGlobalBinding()    {}
func (s *DNSService) CountDNSPolicyDNSGlobalBinding()  {}

// dnspolicy_dnspolicylabel_binding
func (s *DNSService) GetAllDNSPolicyDNSPolicyLabelBinding() {}
func (s *DNSService) GetDNSPolicyDNSPolicyLabelBinding()    {}
func (s *DNSService) CountDNSPolicyDNSPolicyLabelBinding()  {}

// dnsprofile
func (s *DNSService) AddDNSProfile()    {}
func (s *DNSService) DeleteDNSProfile() {}
func (s *DNSService) UpdateDNSProfile() {}
func (s *DNSService) UnsetDNSProfile()  {}
func (s *DNSService) GetAllDNSProfile() {}
func (s *DNSService) GetDNSProfile()    {}
func (s *DNSService) CountDNSProfile()  {}

// dnsproxyrecords
func (s *DNSService) FlushDNSProxyRecords() {}

// dnsptrrec
func (s *DNSService) AddDNSPTRRec()    {}
func (s *DNSService) DeleteDNSPTRRec() {}
func (s *DNSService) GetAllDNSPTRRec() {}
func (s *DNSService) GetDNSPTRRec()    {}
func (s *DNSService) CountDNSPTRRec()  {}

// dnssoarec
func (s *DNSService) AddDNSSOARec()    {}
func (s *DNSService) DeleteDNSSOARec() {}
func (s *DNSService) UpdateDNSSOARec() {}
func (s *DNSService) UnsetDNSSOARec()  {}
func (s *DNSService) GetAllDNSSOARec() {}
func (s *DNSService) GetDNSSOARec()    {}
func (s *DNSService) CountDNSSOARec()  {}

// dnssrvrec
func (s *DNSService) AddDNSSRVRec()    {}
func (s *DNSService) DeleteDNSSRVRec() {}
func (s *DNSService) UpdateDNSSRVRec() {}
func (s *DNSService) UnsetDNSSRVRec()  {}
func (s *DNSService) GetAllDNSSRVRec() {}
func (s *DNSService) CountDNSSRVRec()  {}

// dnssubnetcache
func (s *DNSService) FlushDNSSubnetCache()  {}
func (s *DNSService) GetAllDNSSubnetCache() {}
func (s *DNSService) GetDNSSubnetCache()    {}
func (s *DNSService) CountDNSSubnetCache()  {}

// dnssuffix
func (s *DNSService) AddDNSSuffix()    {}
func (s *DNSService) DeleteDNSSuffix() {}
func (s *DNSService) GetAllDNSSuffix() {}
func (s *DNSService) GetDNSSuffix()    {}
func (s *DNSService) CountDNSSuffix()  {}

// dnstxtrec
func (s *DNSService) AddDNSTXTRec()    {}
func (s *DNSService) DeleteDNSTXTRec() {}
func (s *DNSService) GetAllDNSTXTRec() {}
func (s *DNSService) GetDNSTXTRec()    {}
func (s *DNSService) CountDNSTXTRec()  {}

// dnsview
func (s *DNSService) AddDNSView()    {}
func (s *DNSService) DeleteDNSView() {}
func (s *DNSService) GetAllDNSView() {}
func (s *DNSService) GetDNSView()    {}
func (s *DNSService) CountDNSView()  {}

// dnsview_binding
func (s *DNSService) GetAllDNSViewBinding() {}
func (s *DNSService) GetDNSViewBinding()    {}

// dnsview_dnspolicy_binding
func (s *DNSService) GetAllDNSViewDNSPolicyBinding() {}
func (s *DNSService) GetDNSViewDNSPolicyBinding()    {}
func (s *DNSService) CountDNSViewDNSPolicyBinding()  {}

// dnsview_gslbservice_binding
func (s *DNSService) GetAllDNSViewGSLBServiceBinding() {}
func (s *DNSService) GetDNSViewGSLBServiceBinding()    {}
func (s *DNSService) CountDNSViewGSLBServiceBinding()  {}

// dnszone
func (s *DNSService) AddDNSZone()    {}
func (s *DNSService) DeleteDNSZone() {}
func (s *DNSService) UpdateDNSZone() {}
func (s *DNSService) UnsetDNSZone()  {}
func (s *DNSService) SignDNSZone()   {}
func (s *DNSService) UnsignDNSZone() {}
func (s *DNSService) GetAllDNSZone() {}
func (s *DNSService) GetDNSZone()    {}
func (s *DNSService) CountDNSZone()  {}

// dnszone_binding
func (s *DNSService) GetAllDNSZoneBinding() {}
func (s *DNSService) GetDNSZoneBinding()    {}

// dnszone_dnskey_binding
func (s *DNSService) GetAllDNSZoneDNSKeyBinding() {}
func (s *DNSService) GetDNSZoneDNSKeyBinding()    {}
func (s *DNSService) CountDNSZoneDNSKeyBinding()  {}

// dnszone_domain_binding
func (s *DNSService) GetAllDNSZoneDomainBinding() {}
func (s *DNSService) GetDNSZoneDomainBinding()    {}
func (s *DNSService) CountDNSZoneDomainBinding()  {}
