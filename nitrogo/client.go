package nitrogo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Credential struct {
	username string
	password string
}

type Client struct {
	baseURL    string
	hostname   string
	proxiedURL string

	credential Credential

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
	Autoscale         *AutoscaleService
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
	c := &Client{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

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
	c.Autoscale = &AutoscaleService{client: c}
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

	c.credential.username = username
	c.credential.password = password

	return c, nil
}

// NewRequest - creates the http.Request and applies the relevant authorization
func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	url = fmt.Sprintf("%s/%s", c.baseURL, url)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if c.proxiedURL == "" {
		req.Header.Set("X-NITRO-USER", c.credential.username)
		req.Header.Set("X-NITRO-PASS", c.credential.password)
	} else {
		req.SetBasicAuth(c.credential.username, c.credential.password)
		req.Header.Set("_MPS_API_PROXY_MANAGED_INSTANCE_IP", c.proxiedURL)
	}

	return req, nil
}

// Do - executes the HTTP request and unmarshals the result into r.
func (c *Client) Do(req *http.Request) (map[string]any, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, errors.New("request returned non-200 status code")
	}

	v := map[string]any{}

	return v, json.Unmarshal(respByte, &v)
}

// Credential - getter for client's credentials
func (c *Client) Credential() Credential {
	return c.credential
}

// Hostname - getter for client's hostname
func (c *Client) Hostname() string {
	return c.hostname
}

// URL - getter for the client's baseURL
func (c *Client) URL() string {
	return c.baseURL
}

// SetCredential - setter for client's credential. Handles clearing and creating a new Nitro token if present.
func (c *Client) SetCredential(cred Credential) error {
	c.credential = cred
	return nil
}

// Credential - getter for client's credentials
func (c *Client) SetHostname(hostname string) {
	c.hostname = hostname
}

// SetProxiedURL - setter for client's proxy URL. Handles clearing and creating a new Nitro token if present.
func (c *Client) SetProxiedURL(proxy string) {
	c.proxiedURL = proxy
}

// URL - getter for the client's baseURL
func (c *Client) SetURL(url string) {
	c.baseURL = url
}
