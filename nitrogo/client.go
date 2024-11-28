package nitrogo

import (
	"io"
	"net/http"
)

type Client struct {
	baseURL    string
	proxiedURL string

	credential struct {
		username string
		password string
	}

	httpClient *http.Client

	AAA               *AAAService
	ADM               *ADMService
	Analytics         *AnalyticsService
	App               *AppService
	AppFlow           *AppFlowService
	AppFW             *AppFWService
	AppQOE            *AppQOEService
	Audit             *AuditService
	Authentication    *AuthenticationService
	Authorization     *AuthorizationService
	AutoScale         *AutoScaleService
	Azure             *AzureService
	Basic             *BasicService
	BFD               *BFDService
	Bot               *BotService
	Cache             *CacheService
	Cloud             *CloudService
	Cluster           *ClusterService
	CMP               *CMPService
	ContentInspection *ContentInspectionService
	CR                *CRService
	CS                *CSService
	DB                *DBService
	DNS               *DNSService
	FEO               *FEOService
	GLSB              *GLSBService
	HA                *HAService
	ICA               *ICAService
	IPSEC             *IPSECService
	IPSECALG          *IPSECALGService
	LB                *LBService
	LLDP              *LLDPService
	LSN               *LSNService
	Network           *NetworkService
	NS                *NSService
	NTP               *NTPService
	PCP               *PCPService
	Policy            *PolicyService
	Protocol          *ProtocolService
	RDP               *RDPService
	Reputation        *ReputationService
	Responder         *ResponderService
	Rewrite           *RewriteService
	Router            *RouterService
	SMPP              *SMPPService
	SNMP              *SNMPService
	Spillover         *SpilloverService
	SSL               *SSLService
	Stream            *StreamService
	Subscriber        *SubscriberService
	System            *SystemService
	TM                *TMService
	Transform         *TransformService
	Tunnel            *TunnelService
	ULFD              *ULFDService
	URLFiltering      *URLFilteringService
	User              *UserService
	Utility           *UtilityService
	VideoOptimization *VideoOptimizationService
	VPN               *VPNService
}

func NewNitroClient(username, password string, options ...ClientOptionFunc) (*Client, error) {
	client, err := newClient(options...)
	if err != nil {
		return nil, err
	}

	client.credential.username = username
	client.credential.password = password

	return client, nil
}

func newClient(options ...ClientOptionFunc) (*Client, error) {
	c := &Client{}

	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	c.AAA = &AAAService{client: c}
	c.ADM = &ADMService{client: c}
	c.Analytics = &AnalyticsService{client: c}
	c.App = &AppService{client: c}
	c.AppFlow = &AppFlowService{client: c}
	c.AppFW = &AppFWService{client: c}
	c.AppQOE = &AppQOEService{client: c}
	c.Audit = &AuditService{client: c}
	c.Authentication = &AuthenticationService{client: c}
	c.Authorization = &AuthorizationService{client: c}
	c.AutoScale = &AutoScaleService{client: c}
	c.Azure = &AzureService{client: c}
	c.Basic = &BasicService{client: c}
	c.BFD = &BFDService{client: c}
	c.Bot = &BotService{client: c}
	c.Cache = &CacheService{client: c}
	c.Cloud = &CloudService{client: c}
	c.Cluster = &ClusterService{client: c}
	c.CMP = &CMPService{client: c}
	c.ContentInspection = &ContentInspectionService{client: c}
	c.CR = &CRService{client: c}
	c.CS = &CSService{client: c}
	c.DB = &DBService{client: c}
	c.DNS = &DNSService{client: c}
	c.FEO = &FEOService{client: c}
	c.GLSB = &GLSBService{client: c}
	c.HA = &HAService{client: c}
	c.ICA = &ICAService{client: c}
	c.IPSEC = &IPSECService{client: c}
	c.IPSECALG = &IPSECALGService{client: c}
	c.LB = &LBService{client: c}
	c.LLDP = &LLDPService{client: c}
	c.LSN = &LSNService{client: c}
	c.Network = &NetworkService{client: c}
	c.NS = &NSService{client: c}
	c.NTP = &NTPService{client: c}
	c.PCP = &PCPService{client: c}
	c.Policy = &PolicyService{client: c}
	c.Protocol = &ProtocolService{client: c}
	c.RDP = &RDPService{client: c}
	c.Reputation = &ReputationService{client: c}
	c.Responder = &ResponderService{client: c}
	c.Rewrite = &RewriteService{client: c}
	c.Router = &RouterService{client: c}
	c.SMPP = &SMPPService{client: c}
	c.SNMP = &SNMPService{client: c}
	c.Spillover = &SpilloverService{client: c}
	c.SSL = &SSLService{client: c}
	c.Stream = &StreamService{client: c}
	c.Subscriber = &SubscriberService{client: c}
	c.System = &SystemService{client: c}
	c.TM = &TMService{client: c}
	c.Transform = &TransformService{client: c}
	c.Tunnel = &TunnelService{client: c}
	c.ULFD = &ULFDService{client: c}
	c.URLFiltering = &URLFilteringService{client: c}
	c.User = &UserService{client: c}
	c.Utility = &UtilityService{client: c}
	c.VideoOptimization = &VideoOptimizationService{client: c}
	c.VPN = &VPNService{client: c}

	return c, nil
}

func (client *Client) newRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if client.proxiedURL == "" {
		req.Header.Set("X-NITRO-USER", client.credential.username)
		req.Header.Set("X-NITRO-PASS", client.credential.password)
	} else {
		req.SetBasicAuth(client.credential.username, client.credential.password)
		req.Header.Set("_MPS_API_PROXY_MANAGED_INSTANCE_IP", client.proxiedURL)
	}

	return req, nil
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	return client.httpClient.Do(req)
}
