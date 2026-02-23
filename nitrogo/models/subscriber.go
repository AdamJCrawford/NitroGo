package models

// subscriber configuration structs
type SubscriberSessions struct {
	AVPDisplayBuffer    string   `json:"avpdisplaybuffer,omitempty"`
	Count               float64  `json:"__count,omitempty"`
	Flags               int      `json:"flags,omitempty"`
	IdleTTL             int      `json:"idlettl,omitempty"`
	IP                  string   `json:"ip,omitempty"`
	NextGenAPIResource  string   `json:"_nextgenapiresource,omitempty"`
	NodeID              int      `json:"nodeid,omitempty"`
	ServicePath         string   `json:"servicepath,omitempty"`
	SubscriberRules     []string `json:"subscriberrules,omitempty"`
	SubscriptionIDType  string   `json:"subscriptionidtype,omitempty"`
	SubscriptionIDValue string   `json:"subscriptionidvalue,omitempty"`
	TTL                 int      `json:"ttl,omitempty"`
	VLAN                int      `json:"vlan,omitempty"`
}

type SubscriberGxInterface struct {
	CerRequestTimeout         int           `json:"cerrequesttimeout,omitempty"`
	GxReportingAVP1           []interface{} `json:"gxreportingavp1,omitempty"`
	GxReportingAVP1Type       string        `json:"gxreportingavp1type,omitempty"`
	GxReportingAVP1VendorID   int           `json:"gxreportingavp1vendorid,omitempty"`
	GxReportingAVP2           []interface{} `json:"gxreportingavp2,omitempty"`
	GxReportingAVP2Type       string        `json:"gxreportingavp2type,omitempty"`
	GxReportingAVP2VendorID   int           `json:"gxreportingavp2vendorid,omitempty"`
	GxReportingAVP3           []interface{} `json:"gxreportingavp3,omitempty"`
	GxReportingAVP3Type       string        `json:"gxreportingavp3type,omitempty"`
	GxReportingAVP3VendorID   int           `json:"gxreportingavp3vendorid,omitempty"`
	GxReportingAVP4           []interface{} `json:"gxreportingavp4,omitempty"`
	GxReportingAVP4Type       string        `json:"gxreportingavp4type,omitempty"`
	GxReportingAVP4VendorID   int           `json:"gxreportingavp4vendorid,omitempty"`
	GxReportingAVP5           []interface{} `json:"gxreportingavp5,omitempty"`
	GxReportingAVP5Type       string        `json:"gxreportingavp5type,omitempty"`
	GxReportingAVP5VendorID   int           `json:"gxreportingavp5vendorid,omitempty"`
	HealthCheck               string        `json:"healthcheck,omitempty"`
	HealthCheckTTL            int           `json:"healthcheckttl,omitempty"`
	HoldOnSubscriberAbsence   string        `json:"holdonsubscriberabsence,omitempty"`
	Identity                  string        `json:"identity,omitempty"`
	IdleTTL                   int           `json:"idlettl,omitempty"`
	NegativeTTL               int           `json:"negativettl,omitempty"`
	NegativeTTLLimitedSuccess string        `json:"negativettllimitedsuccess,omitempty"`
	NextGenAPIResource        string        `json:"_nextgenapiresource,omitempty"`
	NodeID                    int           `json:"nodeid,omitempty"`
	PCRFRealm                 string        `json:"pcrfrealm,omitempty"`
	PurgeSDBOnGxFailure       string        `json:"purgesdbongxfailure,omitempty"`
	Realm                     string        `json:"realm,omitempty"`
	RequestRetryAttempts      int           `json:"requestretryattempts,omitempty"`
	RequestTimeout            int           `json:"requesttimeout,omitempty"`
	RevalidationTimeout       int           `json:"revalidationtimeout,omitempty"`
	Service                   string        `json:"service,omitempty"`
	ServicePathAVP            []interface{} `json:"servicepathavp,omitempty"`
	ServicePathInfoMode       string        `json:"servicepathinfomode,omitempty"`
	ServicePathVendorID       int           `json:"servicepathvendorid,omitempty"`
	Status                    string        `json:"status,omitempty"`
	SvrState                  string        `json:"svrstate,omitempty"`
	VServer                   string        `json:"vserver,omitempty"`
}

type SubscriberProfile struct {
	AVPDisplayBuffer    string   `json:"avpdisplaybuffer,omitempty"`
	Count               float64  `json:"__count,omitempty"`
	Flags               int      `json:"flags,omitempty"`
	IP                  string   `json:"ip,omitempty"`
	NextGenAPIResource  string   `json:"_nextgenapiresource,omitempty"`
	ServicePath         string   `json:"servicepath,omitempty"`
	SubscriberRules     []string `json:"subscriberrules,omitempty"`
	SubscriptionIDType  string   `json:"subscriptionidtype,omitempty"`
	SubscriptionIDValue string   `json:"subscriptionidvalue,omitempty"`
	TTL                 int      `json:"ttl,omitempty"`
	VLAN                int      `json:"vlan,omitempty"`
}

type SubscriberParam struct {
	Builtin              []string      `json:"builtin,omitempty"`
	Feature              string        `json:"feature,omitempty"`
	IdleAction           string        `json:"idleaction,omitempty"`
	IdleTTL              int           `json:"idlettl,omitempty"`
	InterfaceType        string        `json:"interfacetype,omitempty"`
	IPv6PrefixLookupList []interface{} `json:"ipv6prefixlookuplist,omitempty"`
	KeyType              string        `json:"keytype,omitempty"`
	NextGenAPIResource   string        `json:"_nextgenapiresource,omitempty"`
}

type SubscriberRadiusInterface struct {
	ListeningService     string `json:"listeningservice,omitempty"`
	NextGenAPIResource   string `json:"_nextgenapiresource,omitempty"`
	RadiusInterimAsStart string `json:"radiusinterimasstart,omitempty"`
	SvrState             string `json:"svrstate,omitempty"`
}
