package models

// ssl configuration structs
type SSLCertKeyBundleIntermediateCertLinksBinding struct {
	CertKeyBundleName   string `json:"certkeybundlename,omitempty"`
	ClientCertNotAfter  string `json:"clientcertnotafter,omitempty"`
	ClientCertNotBefore string `json:"clientcertnotbefore,omitempty"`
	DaysToExpiration    int    `json:"daystoexpiration,omitempty"`
	Issuer              string `json:"issuer,omitempty"`
	PublicKey           string `json:"publickey,omitempty"`
	PublicKeySize       int    `json:"publickeysize,omitempty"`
	SANDNS              string `json:"sandns,omitempty"`
	SANIPAdd            string `json:"sanipadd,omitempty"`
	Serial              string `json:"serial,omitempty"`
	SignatureAlg        string `json:"signaturealg,omitempty"`
	Status              string `json:"status,omitempty"`
	Subject             string `json:"subject,omitempty"`
}

type SSLHPKEKey struct {
	Count              float64 `json:"__count,omitempty"`
	DHKEM              string  `json:"dhkem,omitempty"`
	File               string  `json:"file,omitempty"`
	HPKEKeyName        string  `json:"hpkekeyname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type SSLPolicyCSVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type SSLCipherBinding struct {
	CipherGroupName                  string        `json:"ciphergroupname,omitempty"`
	SSLCipherIndividualCipherBinding []interface{} `json:"sslcipher_individualcipher_binding,omitempty"`
	SSLCipherSSLCipherSuiteBinding   []interface{} `json:"sslcipher_sslciphersuite_binding,omitempty"`
	SSLCipherSSLProfileBinding       []interface{} `json:"sslcipher_sslprofile_binding,omitempty"`
}

type SSLPolicySSLPolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type SSLVServerSSLCertKeyBinding struct {
	CA            bool   `json:"ca,omitempty"`
	CertKeyName   string `json:"certkeyname,omitempty"`
	ClearTextPort int    `json:"cleartextport,omitempty"`
	CRLCheck      string `json:"crlcheck,omitempty"`
	OCSPCheck     string `json:"ocspcheck,omitempty"`
	SkipCAName    bool   `json:"skipcaname,omitempty"`
	SNICert       bool   `json:"snicert,omitempty"`
	VServerName   string `json:"vservername,omitempty"`
}

type SSLCRLSerialNumberBinding struct {
	CRLName string `json:"crlname,omitempty"`
	Date    string `json:"date,omitempty"`
	Number  string `json:"number,omitempty"`
}

type SSLFIPSKey struct {
	Count              float64 `json:"__count,omitempty"`
	Curve              string  `json:"curve,omitempty"`
	Exponent           string  `json:"exponent,omitempty"`
	FIPSKeyName        string  `json:"fipskeyname,omitempty"`
	Inform             string  `json:"inform,omitempty"`
	IV                 string  `json:"iv,omitempty"`
	Key                string  `json:"key,omitempty"`
	KeyType            string  `json:"keytype,omitempty"`
	Modulus            int     `json:"modulus,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Size               int     `json:"size,omitempty"`
	WrapKeyName        string  `json:"wrapkeyname,omitempty"`
}

type SSLCipherSSLProfileBinding struct {
	CipherGroupName string `json:"ciphergroupname,omitempty"`
	CipherOperation string `json:"cipheroperation,omitempty"`
	CipherPriority  int    `json:"cipherpriority,omitempty"`
	CiphGrpAls      string `json:"ciphgrpals,omitempty"`
	Description     string `json:"description,omitempty"`
	SSLProfile      string `json:"sslprofile,omitempty"`
}

type SSLFIPSSimTarget struct {
	CertFile     string `json:"certfile,omitempty"`
	KeyVector    string `json:"keyvector,omitempty"`
	SourceSecret string `json:"sourcesecret,omitempty"`
	TargetSecret string `json:"targetsecret,omitempty"`
}

type SSLServiceSSLPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	PolInherit             int    `json:"polinherit,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	ServiceName            string `json:"servicename,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type SSLProfileSSLCipherSuiteBinding struct {
	CipherName     string `json:"ciphername,omitempty"`
	CipherPriority int    `json:"cipherpriority,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
}

type SSLPolicySSLServiceBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type SSLVServerBinding struct {
	SSLVServerECCCurveBinding         []interface{} `json:"sslvserver_ecccurve_binding,omitempty"`
	SSLVServerHashicorpBinding        []interface{} `json:"sslvserver_hashicorp_binding,omitempty"`
	SSLVServerSSLCACertBundleBinding  []interface{} `json:"sslvserver_sslcacertbundle_binding,omitempty"`
	SSLVServerSSLCertKeyBinding       []interface{} `json:"sslvserver_sslcertkey_binding,omitempty"`
	SSLVServerSSLCertKeyBundleBinding []interface{} `json:"sslvserver_sslcertkeybundle_binding,omitempty"`
	SSLVServerSSLCipherBinding        []interface{} `json:"sslvserver_sslcipher_binding,omitempty"`
	SSLVServerSSLCipherSuiteBinding   []interface{} `json:"sslvserver_sslciphersuite_binding,omitempty"`
	SSLVServerSSLPolicyBinding        []interface{} `json:"sslvserver_sslpolicy_binding,omitempty"`
	VServerName                       string        `json:"vservername,omitempty"`
}

type SSLECDSAKey struct {
	AES256   bool   `json:"aes256,omitempty"`
	Curve    string `json:"curve,omitempty"`
	DES      bool   `json:"des,omitempty"`
	DES3     bool   `json:"des3,omitempty"`
	KeyFile  string `json:"keyfile,omitempty"`
	KeyForm  string `json:"keyform,omitempty"`
	Password string `json:"password,omitempty"`
	PKCS8    bool   `json:"pkcs8,omitempty"`
}

type SSLServiceGroupSSLCACertBundleBinding struct {
	CACertBundleName string `json:"cacertbundlename,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
}

type SSLHSMKey struct {
	Count              float64 `json:"__count,omitempty"`
	HSMKeyName         string  `json:"hsmkeyname,omitempty"`
	HSMType            string  `json:"hsmtype,omitempty"`
	Key                string  `json:"key,omitempty"`
	KeyStore           string  `json:"keystore,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	SerialNum          string  `json:"serialnum,omitempty"`
	State              string  `json:"state,omitempty"`
}

type SSLCACertGroupBinding struct {
	CACertGroupName                 string        `json:"cacertgroupname,omitempty"`
	SSLCACertGroupSSLCertKeyBinding []interface{} `json:"sslcacertgroup_sslcertkey_binding,omitempty"`
}

type SSLCipherSuite struct {
	CipherName         string  `json:"ciphername,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Description        string  `json:"description,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type SSLCRLBinding struct {
	CRLName                   string        `json:"crlname,omitempty"`
	SSLCRLSerialNumberBinding []interface{} `json:"sslcrl_serialnumber_binding,omitempty"`
}

type SSLPolicyLabelBinding struct {
	LabelName                      string        `json:"labelname,omitempty"`
	SSLPolicyLabelSSLPolicyBinding []interface{} `json:"sslpolicylabel_sslpolicy_binding,omitempty"`
}

type SSLServiceGroupBinding struct {
	ServiceGroupName                      string        `json:"servicegroupname,omitempty"`
	SSLServiceGroupECCCurveBinding        []interface{} `json:"sslservicegroup_ecccurve_binding,omitempty"`
	SSLServiceGroupSSLCACertBundleBinding []interface{} `json:"sslservicegroup_sslcacertbundle_binding,omitempty"`
	SSLServiceGroupSSLCertKeyBinding      []interface{} `json:"sslservicegroup_sslcertkey_binding,omitempty"`
	SSLServiceGroupSSLCipherBinding       []interface{} `json:"sslservicegroup_sslcipher_binding,omitempty"`
	SSLServiceGroupSSLCipherSuiteBinding  []interface{} `json:"sslservicegroup_sslciphersuite_binding,omitempty"`
}

type SSLWrapKey struct {
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Salt               string  `json:"salt,omitempty"`
	WrapKeyName        string  `json:"wrapkeyname,omitempty"`
}

type SSLOCSPResponder struct {
	BatchingDelay         int     `json:"batchingdelay,omitempty"`
	BatchingDepth         int     `json:"batchingdepth,omitempty"`
	Cache                 string  `json:"cache,omitempty"`
	CacheTimeout          int     `json:"cachetimeout,omitempty"`
	Count                 float64 `json:"__count,omitempty"`
	HTTPMethod            string  `json:"httpmethod,omitempty"`
	InsertClientCert      string  `json:"insertclientcert,omitempty"`
	Name                  string  `json:"name,omitempty"`
	NextGenAPIResource    string  `json:"_nextgenapiresource,omitempty"`
	OCSPAIARefCount       int     `json:"ocspaiarefcount,omitempty"`
	OCSPIPAddrStr         string  `json:"ocspipaddrstr,omitempty"`
	OCSPURLResolveTimeout int     `json:"ocspurlresolvetimeout,omitempty"`
	Port                  int     `json:"port,omitempty"`
	ProducedAtTimeSkew    int     `json:"producedattimeskew,omitempty"`
	ResponderCert         string  `json:"respondercert,omitempty"`
	RespTimeout           int     `json:"resptimeout,omitempty"`
	SigningCert           string  `json:"signingcert,omitempty"`
	TrustResponder        bool    `json:"trustresponder,omitempty"`
	URL                   string  `json:"url,omitempty"`
	UseNonce              string  `json:"usenonce,omitempty"`
}

type SSLCRL struct {
	BaseDN             string  `json:"basedn,omitempty"`
	Binary             string  `json:"binary,omitempty"`
	BindDN             string  `json:"binddn,omitempty"`
	CACert             string  `json:"cacert,omitempty"`
	CACertFile         string  `json:"cacertfile,omitempty"`
	CAKeyFile          string  `json:"cakeyfile,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	CRLName            string  `json:"crlname,omitempty"`
	CRLPath            string  `json:"crlpath,omitempty"`
	Day                int     `json:"day,omitempty"`
	DaysToExpiration   int     `json:"daystoexpiration,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	GenCRL             string  `json:"gencrl,omitempty"`
	IndexFile          string  `json:"indexfile,omitempty"`
	Inform             string  `json:"inform,omitempty"`
	Interval           string  `json:"interval,omitempty"`
	Issuer             string  `json:"issuer,omitempty"`
	LastUpdate         string  `json:"lastupdate,omitempty"`
	LastUpdateTime     int     `json:"lastupdatetime,omitempty"`
	Method             string  `json:"method,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NextUpdate         string  `json:"nextupdate,omitempty"`
	Password           string  `json:"password,omitempty"`
	Port               int     `json:"port,omitempty"`
	Refresh            string  `json:"refresh,omitempty"`
	Revoke             string  `json:"revoke,omitempty"`
	Scope              string  `json:"scope,omitempty"`
	Server             string  `json:"server,omitempty"`
	SignatureAlgo      string  `json:"signaturealgo,omitempty"`
	Time               string  `json:"time,omitempty"`
	URL                string  `json:"url,omitempty"`
	Version            int     `json:"version,omitempty"`
}

type SSLCertKeyBundleBinding struct {
	CertKeyBundleName                            string        `json:"certkeybundlename,omitempty"`
	SSLCertKeyBundleIntermediateCertLinksBinding []interface{} `json:"sslcertkeybundle_intermediatecertlinks_binding,omitempty"`
	SSLCertKeyBundleSSLVServerBinding            []interface{} `json:"sslcertkeybundle_sslvserver_binding,omitempty"`
}

type SSLProfileSSLCipherBinding struct {
	CipherAliasName string `json:"cipheraliasname,omitempty"`
	CipherName      string `json:"ciphername,omitempty"`
	CipherPriority  int    `json:"cipherpriority,omitempty"`
	Description     string `json:"description,omitempty"`
	Name            string `json:"name,omitempty"`
}

type SSLPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type SSLPKCS8 struct {
	KeyFile   string `json:"keyfile,omitempty"`
	KeyForm   string `json:"keyform,omitempty"`
	Password  string `json:"password,omitempty"`
	PKCS8File string `json:"pkcs8file,omitempty"`
}

type SSLVServerSSLCACertBundleBinding struct {
	CACertBundleName string `json:"cacertbundlename,omitempty"`
	SkipCACertBundle bool   `json:"skipcacertbundle,omitempty"`
	VServerName      string `json:"vservername,omitempty"`
}

type SSLProfileBinding struct {
	Name                            string        `json:"name,omitempty"`
	SSLProfileECCCurveBinding       []interface{} `json:"sslprofile_ecccurve_binding,omitempty"`
	SSLProfileSSLCertKeyBinding     []interface{} `json:"sslprofile_sslcertkey_binding,omitempty"`
	SSLProfileSSLCipherBinding      []interface{} `json:"sslprofile_sslcipher_binding,omitempty"`
	SSLProfileSSLCipherSuiteBinding []interface{} `json:"sslprofile_sslciphersuite_binding,omitempty"`
	SSLProfileSSLECHConfigBinding   []interface{} `json:"sslprofile_sslechconfig_binding,omitempty"`
	SSLProfileSSLVServerBinding     []interface{} `json:"sslprofile_sslvserver_binding,omitempty"`
}

type SSLCertKeyServiceBinding struct {
	CA               bool   `json:"ca,omitempty"`
	CertKey          string `json:"certkey,omitempty"`
	Data             int    `json:"data,omitempty"`
	Service          bool   `json:"service,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
	ServiceName      string `json:"servicename,omitempty"`
	Version          int    `json:"version,omitempty"`
}

type SSLVServerECCCurveBinding struct {
	ECCCurveName string `json:"ecccurvename,omitempty"`
	VServerName  string `json:"vservername,omitempty"`
}

type SSLCACertGroup struct {
	CACertGroupName       string  `json:"cacertgroupname,omitempty"`
	CACertGroupReferences int     `json:"cacertgroupreferences,omitempty"`
	Count                 float64 `json:"__count,omitempty"`
	CRLCheck              string  `json:"crlcheck,omitempty"`
	NextGenAPIResource    string  `json:"_nextgenapiresource,omitempty"`
	OCSPCheck             string  `json:"ocspcheck,omitempty"`
}

type SSLServiceSSLCipherBinding struct {
	CipherAliasName string `json:"cipheraliasname,omitempty"`
	CipherDefaultOn int    `json:"cipherdefaulton,omitempty"`
	CipherName      string `json:"ciphername,omitempty"`
	Description     string `json:"description,omitempty"`
	ServiceName     string `json:"servicename,omitempty"`
}

type SSLVServerSSLCipherBinding struct {
	CipherAliasName string `json:"cipheraliasname,omitempty"`
	CipherName      string `json:"ciphername,omitempty"`
	Description     string `json:"description,omitempty"`
	VServerName     string `json:"vservername,omitempty"`
}

type SSLDefaultProfile struct {
}

type SSLCertKeyCRLDistributionBinding struct {
	CA      bool   `json:"ca,omitempty"`
	CertKey string `json:"certkey,omitempty"`
	Issuer  string `json:"issuer,omitempty"`
}

type SSLCACertBundle struct {
	BundleFile         string  `json:"bundlefile,omitempty"`
	CACertBundleName   string  `json:"cacertbundlename,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ServerName         string  `json:"servername,omitempty"`
}

type SSLDTLSProfile struct {
	Builtin              []string `json:"builtin,omitempty"`
	Count                float64  `json:"__count,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	HelloVerifyRequest   string   `json:"helloverifyrequest,omitempty"`
	InitialRetryTimeout  int      `json:"initialretrytimeout,omitempty"`
	MaxBadMACIgnoreCount int      `json:"maxbadmacignorecount,omitempty"`
	MaxHoldQLen          int      `json:"maxholdqlen,omitempty"`
	MaxPacketSize        int      `json:"maxpacketsize,omitempty"`
	MaxRecordSize        int      `json:"maxrecordsize,omitempty"`
	MaxRetryTime         int      `json:"maxretrytime,omitempty"`
	Name                 string   `json:"name,omitempty"`
	NextGenAPIResource   string   `json:"_nextgenapiresource,omitempty"`
	PMTUDiscovery        string   `json:"pmtudiscovery,omitempty"`
	TerminateSession     string   `json:"terminatesession,omitempty"`
}

type SSLCert struct {
	CACert         string `json:"cacert,omitempty"`
	CACertForm     string `json:"cacertform,omitempty"`
	CAKey          string `json:"cakey,omitempty"`
	CAKeyForm      string `json:"cakeyform,omitempty"`
	CASerial       string `json:"caserial,omitempty"`
	CertFile       string `json:"certfile,omitempty"`
	CertForm       string `json:"certform,omitempty"`
	CertType       string `json:"certtype,omitempty"`
	Days           int    `json:"days,omitempty"`
	KeyFile        string `json:"keyfile,omitempty"`
	KeyForm        string `json:"keyform,omitempty"`
	PEMPassphrase  string `json:"pempassphrase,omitempty"`
	ReqFile        string `json:"reqfile,omitempty"`
	SubjectAltName string `json:"subjectaltname,omitempty"`
}

type SSLPolicySSLGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type SSLRSAKey struct {
	AES256   bool   `json:"aes256,omitempty"`
	Bits     int    `json:"bits,omitempty"`
	DES      bool   `json:"des,omitempty"`
	DES3     bool   `json:"des3,omitempty"`
	Exponent string `json:"exponent,omitempty"`
	KeyFile  string `json:"keyfile,omitempty"`
	KeyForm  string `json:"keyform,omitempty"`
	Password string `json:"password,omitempty"`
	PKCS8    bool   `json:"pkcs8,omitempty"`
}

type SSLServiceECCCurveBinding struct {
	ECCCurveName string `json:"ecccurvename,omitempty"`
	ServiceName  string `json:"servicename,omitempty"`
}

type SSLServiceSSLCertKeyBinding struct {
	CA            bool   `json:"ca,omitempty"`
	CertKeyName   string `json:"certkeyname,omitempty"`
	ClearTextPort int    `json:"cleartextport,omitempty"`
	CRLCheck      string `json:"crlcheck,omitempty"`
	OCSPCheck     string `json:"ocspcheck,omitempty"`
	ServiceName   string `json:"servicename,omitempty"`
	SkipCAName    bool   `json:"skipcaname,omitempty"`
	SNICert       bool   `json:"snicert,omitempty"`
}

type SSLServiceGroup struct {
	CA                   bool    `json:"ca,omitempty"`
	CipherRedirect       string  `json:"cipherredirect,omitempty"`
	CipherURL            string  `json:"cipherurl,omitempty"`
	ClearTextPort        int     `json:"cleartextport,omitempty"`
	ClientAuth           string  `json:"clientauth,omitempty"`
	ClientCert           string  `json:"clientcert,omitempty"`
	CommonName           string  `json:"commonname,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	CRLCheck             string  `json:"crlcheck,omitempty"`
	DH                   string  `json:"dh,omitempty"`
	DHCount              int     `json:"dhcount,omitempty"`
	DHFile               string  `json:"dhfile,omitempty"`
	DHKeyExpSizeLimit    string  `json:"dhkeyexpsizelimit,omitempty"`
	ERSA                 string  `json:"ersa,omitempty"`
	ERSACount            int     `json:"ersacount,omitempty"`
	NextGenAPIResource   string  `json:"_nextgenapiresource,omitempty"`
	NonFIPSCiphers       string  `json:"nonfipsciphers,omitempty"`
	OCSPCheck            string  `json:"ocspcheck,omitempty"`
	OCSPStapling         string  `json:"ocspstapling,omitempty"`
	QUICFlag             bool    `json:"quicflag,omitempty"`
	RedirectPortRewrite  string  `json:"redirectportrewrite,omitempty"`
	SendCloseNotify      string  `json:"sendclosenotify,omitempty"`
	ServerAuth           string  `json:"serverauth,omitempty"`
	ServiceGroupName     string  `json:"servicegroupname,omitempty"`
	ServiceName          string  `json:"servicename,omitempty"`
	SessReuse            string  `json:"sessreuse,omitempty"`
	SessTimeout          int     `json:"sesstimeout,omitempty"`
	SNICert              bool    `json:"snicert,omitempty"`
	SNIEnable            string  `json:"snienable,omitempty"`
	SSL2                 string  `json:"ssl2,omitempty"`
	SSL3                 string  `json:"ssl3,omitempty"`
	SSLClientLogs        string  `json:"sslclientlogs,omitempty"`
	SSLProfile           string  `json:"sslprofile,omitempty"`
	SSLRedirect          string  `json:"sslredirect,omitempty"`
	SSLv2Redirect        string  `json:"sslv2redirect,omitempty"`
	SSLv2URL             string  `json:"sslv2url,omitempty"`
	StrictSigDigestCheck string  `json:"strictsigdigestcheck,omitempty"`
	TLS1                 string  `json:"tls1,omitempty"`
	TLS11                string  `json:"tls11,omitempty"`
	TLS12                string  `json:"tls12,omitempty"`
	TLS13                string  `json:"tls13,omitempty"`
}

type SSLProfileSSLVServerBinding struct {
	CipherPriority int    `json:"cipherpriority,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	ServiceName    string `json:"servicename,omitempty"`
}

type SSLCertLink struct {
	CertKeyName        string  `json:"certkeyname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	LinkCertKeyName    string  `json:"linkcertkeyname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type SSLGlobalSSLPolicyBinding struct {
	GlobalBindType         string `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type SSLCRLFile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type SSLCertChainSSLCertKeyBinding struct {
	AddSubject      bool   `json:"addsubject,omitempty"`
	CertKeyName     string `json:"certkeyname,omitempty"`
	IsCA            bool   `json:"isca,omitempty"`
	IsLinked        bool   `json:"islinked,omitempty"`
	LinkCertKeyName string `json:"linkcertkeyname,omitempty"`
}

type SSLParameter struct {
	CRLMemorySizeMB          int      `json:"crlmemorysizemb,omitempty"`
	CryptoDevDisableLimit    int      `json:"cryptodevdisablelimit,omitempty"`
	DefaultProfile           string   `json:"defaultprofile,omitempty"`
	DenySSLReneg             string   `json:"denysslreneg,omitempty"`
	DropReqWithNoHostHeader  string   `json:"dropreqwithnohostheader,omitempty"`
	EncryptTriggerPktCount   int      `json:"encrypttriggerpktcount,omitempty"`
	HeterogeneousSSLHW       string   `json:"heterogeneoussslhw,omitempty"`
	HybridFIPSMode           string   `json:"hybridfipsmode,omitempty"`
	InsertCertSpace          string   `json:"insertcertspace,omitempty"`
	InsertionEncoding        string   `json:"insertionencoding,omitempty"`
	MonTLS1112Disable        string   `json:"montls1112disable,omitempty"`
	NDCPPComplianceCertCheck string   `json:"ndcppcompliancecertcheck,omitempty"`
	NextGenAPIResource       string   `json:"_nextgenapiresource,omitempty"`
	OCSPCachesize            int      `json:"ocspcachesize,omitempty"`
	OperationQueueLimit      int      `json:"operationqueuelimit,omitempty"`
	PushEncTriggerTimeout    int      `json:"pushenctriggertimeout,omitempty"`
	PushFlag                 int      `json:"pushflag,omitempty"`
	QuantumSize              string   `json:"quantumsize,omitempty"`
	SendCloseNotify          string   `json:"sendclosenotify,omitempty"`
	SigDigestType            []string `json:"sigdigesttype,omitempty"`
	SNIHTTPHostMatch         string   `json:"snihttphostmatch,omitempty"`
	SoftwareCryptoThreshold  int      `json:"softwarecryptothreshold,omitempty"`
	SSLIErrorCache           string   `json:"sslierrorcache,omitempty"`
	SSLIMaxErrorCacheMem     int      `json:"sslimaxerrorcachemem,omitempty"`
	SSLTriggerTimeout        int      `json:"ssltriggertimeout,omitempty"`
	StrictCAChecks           string   `json:"strictcachecks,omitempty"`
	SvcTLS1112Disable        string   `json:"svctls1112disable,omitempty"`
	UndefActionControl       string   `json:"undefactioncontrol,omitempty"`
	UndefActionData          string   `json:"undefactiondata,omitempty"`
}

type SSLDHParam struct {
	Bits   int    `json:"bits,omitempty"`
	DHFile string `json:"dhfile,omitempty"`
	Gen    string `json:"gen,omitempty"`
}

type SSLPolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ReqAction          string   `json:"reqaction,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	UndefAction        string   `json:"undefaction,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type SSLCertBundle struct {
	Count              float64 `json:"__count,omitempty"`
	InUse              string  `json:"inuse,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type SSLCertKeySSLVServerBinding struct {
	CA          bool   `json:"ca,omitempty"`
	CertKey     string `json:"certkey,omitempty"`
	Data        int    `json:"data,omitempty"`
	ServerName  string `json:"servername,omitempty"`
	Version     int    `json:"version,omitempty"`
	VServer     bool   `json:"vserver,omitempty"`
	VServerName string `json:"vservername,omitempty"`
}

type SSLCertKeyBundleSSLVServerBinding struct {
	CertKeyBundleName string `json:"certkeybundlename,omitempty"`
	ServerName        string `json:"servername,omitempty"`
}

type SSLServiceGroupECCCurveBinding struct {
	ECCCurveName     string `json:"ecccurvename,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
}

type SSLProfileSSLCertKeyBinding struct {
	CertKeyName      string `json:"certkeyname,omitempty"`
	CipherPriority   int    `json:"cipherpriority,omitempty"`
	ForgingCACertKey bool   `json:"forgingcacertkey,omitempty"`
	Name             string `json:"name,omitempty"`
	SSLICACertKey    string `json:"sslicacertkey,omitempty"`
}

type SSLCertFile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type SSLCACertGroupSSLCertKeyBinding struct {
	CACertGroupName string `json:"cacertgroupname,omitempty"`
	CertKeyName     string `json:"certkeyname,omitempty"`
	CRLCheck        string `json:"crlcheck,omitempty"`
	OCSPCheck       string `json:"ocspcheck,omitempty"`
}

type SSLServiceGroupSSLCipherBinding struct {
	CipherAliasName  string `json:"cipheraliasname,omitempty"`
	CipherName       string `json:"ciphername,omitempty"`
	Description      string `json:"description,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
}

type SSLCertKey struct {
	Builtin                     []string `json:"builtin,omitempty"`
	Bundle                      string   `json:"bundle,omitempty"`
	Cert                        string   `json:"cert,omitempty"`
	CertificateSource           string   `json:"certificatesource,omitempty"`
	CertificateType             []string `json:"certificatetype,omitempty"`
	CertKey                     string   `json:"certkey,omitempty"`
	CertKeyDigest               string   `json:"certkeydigest,omitempty"`
	CertKeyStatus               string   `json:"certkeystatus,omitempty"`
	ClientCertNotAfter          string   `json:"clientcertnotafter,omitempty"`
	ClientCertNotBefore         string   `json:"clientcertnotbefore,omitempty"`
	Count                       float64  `json:"__count,omitempty"`
	Data                        int      `json:"data,omitempty"`
	DaysToExpiration            int      `json:"daystoexpiration,omitempty"`
	DeleteCertKeyFilesOnRemoval string   `json:"deletecertkeyfilesonremoval,omitempty"`
	DeleteFromDevice            bool     `json:"deletefromdevice,omitempty"`
	ExpiryMonitor               string   `json:"expirymonitor,omitempty"`
	Feature                     string   `json:"feature,omitempty"`
	FIPSKey                     string   `json:"fipskey,omitempty"`
	HSMKey                      string   `json:"hsmkey,omitempty"`
	Inform                      string   `json:"inform,omitempty"`
	Issuer                      string   `json:"issuer,omitempty"`
	Key                         string   `json:"key,omitempty"`
	LinkCertKeyName             string   `json:"linkcertkeyname,omitempty"`
	NextGenAPIResource          string   `json:"_nextgenapiresource,omitempty"`
	NoDomainCheck               bool     `json:"nodomaincheck,omitempty"`
	NotificationPeriod          int      `json:"notificationperiod,omitempty"`
	OCSPResponseStatus          string   `json:"ocspresponsestatus,omitempty"`
	OCSPStaplingCache           bool     `json:"ocspstaplingcache,omitempty"`
	PassCrypt                   string   `json:"passcrypt,omitempty"`
	PassPlain                   string   `json:"passplain,omitempty"`
	Password                    bool     `json:"password,omitempty"`
	Priority                    int      `json:"priority,omitempty"`
	PublicKey                   string   `json:"publickey,omitempty"`
	PublicKeySize               int      `json:"publickeysize,omitempty"`
	SANDNS                      string   `json:"sandns,omitempty"`
	SANIPAdd                    string   `json:"sanipadd,omitempty"`
	Serial                      string   `json:"serial,omitempty"`
	ServiceName                 string   `json:"servicename,omitempty"`
	SignatureAlg                string   `json:"signaturealg,omitempty"`
	Status                      string   `json:"status,omitempty"`
	Subject                     string   `json:"subject,omitempty"`
	Version                     int      `json:"version,omitempty"`
}

type SSLProfileSSLECHConfigBinding struct {
	CipherPriority int    `json:"cipherpriority,omitempty"`
	ECHConfigName  string `json:"echconfigname,omitempty"`
	Name           string `json:"name,omitempty"`
}

type SSLServiceSSLCipherSuiteBinding struct {
	CipherDefaultOn int    `json:"cipherdefaulton,omitempty"`
	CipherName      string `json:"ciphername,omitempty"`
	Description     string `json:"description,omitempty"`
	ServiceName     string `json:"servicename,omitempty"`
}

type SSLCertKeyBundle struct {
	BundleFile         string  `json:"bundlefile,omitempty"`
	CertKeyBundleName  string  `json:"certkeybundlename,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PassPlain          string  `json:"passplain,omitempty"`
}

type SSLCertChainBinding struct {
	CertKeyName                   string        `json:"certkeyname,omitempty"`
	SSLCertChainSSLCertKeyBinding []interface{} `json:"sslcertchain_sslcertkey_binding,omitempty"`
}

type SSLCertificateChain struct {
	CertKeyName        string   `json:"certkeyname,omitempty"`
	ChainComplete      int      `json:"chaincomplete,omitempty"`
	ChainIssuer        string   `json:"chainissuer,omitempty"`
	ChainLinked        []string `json:"chainlinked,omitempty"`
	ChainPossibleLinks []string `json:"chainpossiblelinks,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
}

type SSLProfileECCCurveBinding struct {
	CipherPriority int    `json:"cipherpriority,omitempty"`
	ECCCurveName   string `json:"ecccurvename,omitempty"`
	Name           string `json:"name,omitempty"`
}

type SSLService struct {
	CipherRedirect       string  `json:"cipherredirect,omitempty"`
	CipherURL            string  `json:"cipherurl,omitempty"`
	ClientAuth           string  `json:"clientauth,omitempty"`
	ClientCert           string  `json:"clientcert,omitempty"`
	CommonName           string  `json:"commonname,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	DH                   string  `json:"dh,omitempty"`
	DHCount              int     `json:"dhcount,omitempty"`
	DHFile               string  `json:"dhfile,omitempty"`
	DHKeyExpSizeLimit    string  `json:"dhkeyexpsizelimit,omitempty"`
	DTLS1                string  `json:"dtls1,omitempty"`
	DTLS12               string  `json:"dtls12,omitempty"`
	DTLSFlag             bool    `json:"dtlsflag,omitempty"`
	DTLSProfileName      string  `json:"dtlsprofilename,omitempty"`
	ERSA                 string  `json:"ersa,omitempty"`
	ERSACount            int     `json:"ersacount,omitempty"`
	NextGenAPIResource   string  `json:"_nextgenapiresource,omitempty"`
	NonFIPSCiphers       string  `json:"nonfipsciphers,omitempty"`
	OCSPStapling         string  `json:"ocspstapling,omitempty"`
	PushEncTrigger       string  `json:"pushenctrigger,omitempty"`
	QUICFlag             bool    `json:"quicflag,omitempty"`
	RedirectPortRewrite  string  `json:"redirectportrewrite,omitempty"`
	SendCloseNotify      string  `json:"sendclosenotify,omitempty"`
	ServerAuth           string  `json:"serverauth,omitempty"`
	Service              int     `json:"service,omitempty"`
	ServiceName          string  `json:"servicename,omitempty"`
	SessReuse            string  `json:"sessreuse,omitempty"`
	SessTimeout          int     `json:"sesstimeout,omitempty"`
	SkipCACertBundle     bool    `json:"skipcacertbundle,omitempty"`
	SkipCAName           bool    `json:"skipcaname,omitempty"`
	SNIEnable            string  `json:"snienable,omitempty"`
	SSL2                 string  `json:"ssl2,omitempty"`
	SSL3                 string  `json:"ssl3,omitempty"`
	SSLClientLogs        string  `json:"sslclientlogs,omitempty"`
	SSLProfile           string  `json:"sslprofile,omitempty"`
	SSLRedirect          string  `json:"sslredirect,omitempty"`
	SSLv2Redirect        string  `json:"sslv2redirect,omitempty"`
	SSLv2URL             string  `json:"sslv2url,omitempty"`
	StrictSigDigestCheck string  `json:"strictsigdigestcheck,omitempty"`
	TLS1                 string  `json:"tls1,omitempty"`
	TLS11                string  `json:"tls11,omitempty"`
	TLS12                string  `json:"tls12,omitempty"`
	TLS13                string  `json:"tls13,omitempty"`
}

type SSLProfile struct {
	AllowExtendedMasterSecret         string   `json:"allowextendedmastersecret,omitempty"`
	AllowLegacyKDF                    string   `json:"allowlegacykdf,omitempty"`
	AllowUnknownSNI                   string   `json:"allowunknownsni,omitempty"`
	ALPNProtocol                      string   `json:"alpnprotocol,omitempty"`
	Builtin                           []string `json:"builtin,omitempty"`
	CipherName                        string   `json:"ciphername,omitempty"`
	CipherPriority                    int      `json:"cipherpriority,omitempty"`
	CipherRedirect                    string   `json:"cipherredirect,omitempty"`
	CipherURL                         string   `json:"cipherurl,omitempty"`
	ClearTextPort                     int      `json:"cleartextport,omitempty"`
	ClientAuth                        string   `json:"clientauth,omitempty"`
	ClientAuthUseBoundCAChain         string   `json:"clientauthuseboundcachain,omitempty"`
	ClientCert                        string   `json:"clientcert,omitempty"`
	CommonName                        string   `json:"commonname,omitempty"`
	Count                             float64  `json:"__count,omitempty"`
	CRLCheck                          string   `json:"crlcheck,omitempty"`
	DefaultSNI                        string   `json:"defaultsni,omitempty"`
	DenySSLReneg                      string   `json:"denysslreneg,omitempty"`
	DH                                string   `json:"dh,omitempty"`
	DHCount                           int      `json:"dhcount,omitempty"`
	DHEKeyExchangeWithPSK             string   `json:"dhekeyexchangewithpsk,omitempty"`
	DHFile                            string   `json:"dhfile,omitempty"`
	DHKeyExpSizeLimit                 string   `json:"dhkeyexpsizelimit,omitempty"`
	DropReqWithNoHostHeader           string   `json:"dropreqwithnohostheader,omitempty"`
	DynamicClientCert                 string   `json:"dynamicclientcert,omitempty"`
	EncryptedClientHello              string   `json:"encryptedclienthello,omitempty"`
	EncryptTriggerPktCount            int      `json:"encrypttriggerpktcount,omitempty"`
	ERSA                              string   `json:"ersa,omitempty"`
	ERSACount                         int      `json:"ersacount,omitempty"`
	Feature                           string   `json:"feature,omitempty"`
	HSTS                              string   `json:"hsts,omitempty"`
	IncludeSubdomains                 string   `json:"includesubdomains,omitempty"`
	InsertionEncoding                 string   `json:"insertionencoding,omitempty"`
	Invoke                            bool     `json:"invoke,omitempty"`
	LabelType                         string   `json:"labeltype,omitempty"`
	MaxAge                            int      `json:"maxage,omitempty"`
	MaxRenegRate                      int      `json:"maxrenegrate,omitempty"`
	Name                              string   `json:"name,omitempty"`
	NextGenAPIResource                string   `json:"_nextgenapiresource,omitempty"`
	NoDefaultBindings                 string   `json:"nodefaultbindings,omitempty"`
	NonFIPSCiphers                    string   `json:"nonfipsciphers,omitempty"`
	OCSPCheck                         string   `json:"ocspcheck,omitempty"`
	OCSPStapling                      string   `json:"ocspstapling,omitempty"`
	Preload                           string   `json:"preload,omitempty"`
	PrevSessionKeyLifetime            int      `json:"prevsessionkeylifetime,omitempty"`
	PushEncTrigger                    string   `json:"pushenctrigger,omitempty"`
	PushEncTriggerTimeout             int      `json:"pushenctriggertimeout,omitempty"`
	PushFlag                          int      `json:"pushflag,omitempty"`
	QuantumSize                       string   `json:"quantumsize,omitempty"`
	RedirectPortRewrite               string   `json:"redirectportrewrite,omitempty"`
	SendCloseNotify                   string   `json:"sendclosenotify,omitempty"`
	ServerAuth                        string   `json:"serverauth,omitempty"`
	Service                           int      `json:"service,omitempty"`
	SessionKeyLifetime                int      `json:"sessionkeylifetime,omitempty"`
	SessionTicket                     string   `json:"sessionticket,omitempty"`
	SessionTicketKeyData              string   `json:"sessionticketkeydata,omitempty"`
	SessionTicketKeyRefresh           string   `json:"sessionticketkeyrefresh,omitempty"`
	SessionTicketLifetime             int      `json:"sessionticketlifetime,omitempty"`
	SessReuse                         string   `json:"sessreuse,omitempty"`
	SessTimeout                       int      `json:"sesstimeout,omitempty"`
	SkipCAName                        bool     `json:"skipcaname,omitempty"`
	SkipClientCertPolicyCheck         string   `json:"skipclientcertpolicycheck,omitempty"`
	SNICert                           bool     `json:"snicert,omitempty"`
	SNIEnable                         string   `json:"snienable,omitempty"`
	SNIHTTPHostMatch                  string   `json:"snihttphostmatch,omitempty"`
	SSL3                              string   `json:"ssl3,omitempty"`
	SSLClientLogs                     string   `json:"sslclientlogs,omitempty"`
	SSLIMaxSessPerServer              int      `json:"sslimaxsessperserver,omitempty"`
	SSLInterception                   string   `json:"sslinterception,omitempty"`
	SSLIOCSPCheck                     string   `json:"ssliocspcheck,omitempty"`
	SSLIReneg                         string   `json:"sslireneg,omitempty"`
	SSLIVerifyServerCertForReuse      string   `json:"ssliverifyservercertforreuse,omitempty"`
	SSLLogProfile                     string   `json:"ssllogprofile,omitempty"`
	SSLPFObjectType                   int      `json:"sslpfobjecttype,omitempty"`
	SSLProfileType                    string   `json:"sslprofiletype,omitempty"`
	SSLRedirect                       string   `json:"sslredirect,omitempty"`
	SSLTriggerTimeout                 int      `json:"ssltriggertimeout,omitempty"`
	StrictCAChecks                    string   `json:"strictcachecks,omitempty"`
	StrictSigDigestCheck              string   `json:"strictsigdigestcheck,omitempty"`
	TLS1                              string   `json:"tls1,omitempty"`
	TLS11                             string   `json:"tls11,omitempty"`
	TLS12                             string   `json:"tls12,omitempty"`
	TLS13                             string   `json:"tls13,omitempty"`
	TLS13SessionTicketsPerAuthContext int      `json:"tls13sessionticketsperauthcontext,omitempty"`
	ZeroRTTEarlyData                  string   `json:"zerorttearlydata,omitempty"`
}

type SSLZeroTouchParam struct {
	ADMConnectivityStatus  string `json:"admconnectivitystatus,omitempty"`
	HTTPStatusCode         string `json:"httpstatuscode,omitempty"`
	KeyFileName            string `json:"keyfilename,omitempty"`
	NextGenAPIResource     string `json:"_nextgenapiresource,omitempty"`
	NextRequestTime        string `json:"nextrequesttime,omitempty"`
	OCSPBatchingDelay      int    `json:"ocspbatchingdelay,omitempty"`
	OCSPBatchingDepth      int    `json:"ocspbatchingdepth,omitempty"`
	OCSPCacheTimeout       int    `json:"ocspcachetimeout,omitempty"`
	OCSPHTTPMethod         string `json:"ocsphttpmethod,omitempty"`
	OCSPProducedAtTimeSkew int    `json:"ocspproducedattimeskew,omitempty"`
	OCSPRespTimeout        int    `json:"ocspresptimeout,omitempty"`
	OCSPTrustResponder     string `json:"ocsptrustresponder,omitempty"`
	OCSPURLResolveTimeout  int    `json:"ocspurlresolvetimeout,omitempty"`
	OCSPUseNonce           string `json:"ocspusenonce,omitempty"`
	Passphrase             string `json:"passphrase,omitempty"`
	RemoteServerIP         string `json:"remoteserverip,omitempty"`
	RequestTimestamp       string `json:"requesttimestamp,omitempty"`
	RequestType            string `json:"requesttype,omitempty"`
	ZeroTouch              string `json:"zerotouch,omitempty"`
}

type SSLPolicyBinding struct {
	Name                           string        `json:"name,omitempty"`
	SSLPolicyCSVServerBinding      []interface{} `json:"sslpolicy_csvserver_binding,omitempty"`
	SSLPolicyLBVServerBinding      []interface{} `json:"sslpolicy_lbvserver_binding,omitempty"`
	SSLPolicySSLGlobalBinding      []interface{} `json:"sslpolicy_sslglobal_binding,omitempty"`
	SSLPolicySSLPolicyLabelBinding []interface{} `json:"sslpolicy_sslpolicylabel_binding,omitempty"`
	SSLPolicySSLServiceBinding     []interface{} `json:"sslpolicy_sslservice_binding,omitempty"`
	SSLPolicySSLVServerBinding     []interface{} `json:"sslpolicy_sslvserver_binding,omitempty"`
}

type SSLPKCS12 struct {
	AES256        bool   `json:"aes256,omitempty"`
	CertFile      string `json:"certfile,omitempty"`
	DES           bool   `json:"des,omitempty"`
	DES3          bool   `json:"des3,omitempty"`
	Export        bool   `json:"export,omitempty"`
	Import        bool   `json:"Import,omitempty"`
	KeyFile       string `json:"keyfile,omitempty"`
	OutFile       string `json:"outfile,omitempty"`
	Password      string `json:"password,omitempty"`
	PEMPassphrase string `json:"pempassphrase,omitempty"`
	PKCS12File    string `json:"pkcs12file,omitempty"`
}

type SSLPolicySSLVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type SSLCipherSSLCipherSuiteBinding struct {
	CipherGroupName string `json:"ciphergroupname,omitempty"`
	CipherName      string `json:"ciphername,omitempty"`
	CipherOperation string `json:"cipheroperation,omitempty"`
	CipherPriority  int    `json:"cipherpriority,omitempty"`
	CiphGrpAls      string `json:"ciphgrpals,omitempty"`
	Description     string `json:"description,omitempty"`
}

type SSLKeyFile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type SSLPolicyLabelSSLPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type SSLDynamicClientCertCache struct {
}

type SSLGlobalBinding struct {
	SSLGlobalSSLPolicyBinding []interface{} `json:"sslglobal_sslpolicy_binding,omitempty"`
}

type SSLServiceSSLCACertBundleBinding struct {
	CACertBundleName string `json:"cacertbundlename,omitempty"`
	ServiceName      string `json:"servicename,omitempty"`
	SkipCACertBundle bool   `json:"skipcacertbundle,omitempty"`
}

type SSLCipherIndividualCipherBinding struct {
	CipherGroupName string `json:"ciphergroupname,omitempty"`
	CipherName      string `json:"ciphername,omitempty"`
	CipherOperation string `json:"cipheroperation,omitempty"`
	CipherPriority  int    `json:"cipherpriority,omitempty"`
	CiphGrpAls      string `json:"ciphgrpals,omitempty"`
	Description     string `json:"description,omitempty"`
}

type SSLVServerSSLCertKeyBundleBinding struct {
	CertKeyBundleName string `json:"certkeybundlename,omitempty"`
	SNICertKeyBundle  bool   `json:"snicertkeybundle,omitempty"`
	VServerName       string `json:"vservername,omitempty"`
}

type SSLCertKeySSLOCSPResponderBinding struct {
	CA            bool   `json:"ca,omitempty"`
	CertKey       string `json:"certkey,omitempty"`
	OCSPResponder string `json:"ocspresponder,omitempty"`
	Priority      int    `json:"priority,omitempty"`
}

type SSLDHFile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type SSLVServerSSLCipherSuiteBinding struct {
	CipherName  string `json:"ciphername,omitempty"`
	Description string `json:"description,omitempty"`
	VServerName string `json:"vservername,omitempty"`
}

type SSLVServerHashicorpBinding struct {
	Vault       string `json:"vault,omitempty"`
	VServerName string `json:"vservername,omitempty"`
}

type SSLFIPS struct {
	CoresEnabled        int    `json:"coresenabled,omitempty"`
	CoresMax            int    `json:"coresmax,omitempty"`
	EraseData           string `json:"erasedata,omitempty"`
	FIPSFW              string `json:"fipsfw,omitempty"`
	FIPSHWMajorVersion  int    `json:"fipshwmajorversion,omitempty"`
	FIPSHWMinorVersion  int    `json:"fipshwminorversion,omitempty"`
	FIPSHWVersionString string `json:"fipshwversionstring,omitempty"`
	FirmwareReleaseDate string `json:"firmwarereleasedate,omitempty"`
	Flag                int    `json:"flag,omitempty"`
	FlashMemoryFree     int    `json:"flashmemoryfree,omitempty"`
	FlashMemoryTotal    int    `json:"flashmemorytotal,omitempty"`
	HSMLabel            string `json:"hsmlabel,omitempty"`
	InitHSM             string `json:"inithsm,omitempty"`
	MajorVersion        int    `json:"majorversion,omitempty"`
	MinorVersion        int    `json:"minorversion,omitempty"`
	Model               string `json:"model,omitempty"`
	NextGenAPIResource  string `json:"_nextgenapiresource,omitempty"`
	OldSOPassword       string `json:"oldsopassword,omitempty"`
	Serial              int    `json:"serial,omitempty"`
	SerialNo            string `json:"serialno,omitempty"`
	SOPassword          string `json:"sopassword,omitempty"`
	SRAMFree            int    `json:"sramfree,omitempty"`
	SRAMTotal           int    `json:"sramtotal,omitempty"`
	State               int    `json:"state,omitempty"`
	Status              int    `json:"status,omitempty"`
	UserPassword        string `json:"userpassword,omitempty"`
}

type SSLAction struct {
	Builtin                []string `json:"builtin,omitempty"`
	CACertGrpName          string   `json:"cacertgrpname,omitempty"`
	CertFingerprintDigest  string   `json:"certfingerprintdigest,omitempty"`
	CertFingerprintHeader  string   `json:"certfingerprintheader,omitempty"`
	CertHashHeader         string   `json:"certhashheader,omitempty"`
	CertHeader             string   `json:"certheader,omitempty"`
	CertIssuerHeader       string   `json:"certissuerheader,omitempty"`
	CertNotAfterHeader     string   `json:"certnotafterheader,omitempty"`
	CertNotBeforeHeader    string   `json:"certnotbeforeheader,omitempty"`
	CertSerialHeader       string   `json:"certserialheader,omitempty"`
	CertSubjectHeader      string   `json:"certsubjectheader,omitempty"`
	Cipher                 string   `json:"cipher,omitempty"`
	CipherHeader           string   `json:"cipherheader,omitempty"`
	ClientAuth             string   `json:"clientauth,omitempty"`
	ClientCert             string   `json:"clientcert,omitempty"`
	ClientCertFingerprint  string   `json:"clientcertfingerprint,omitempty"`
	ClientCertHash         string   `json:"clientcerthash,omitempty"`
	ClientCertIssuer       string   `json:"clientcertissuer,omitempty"`
	ClientCertNotAfter     string   `json:"clientcertnotafter,omitempty"`
	ClientCertNotBefore    string   `json:"clientcertnotbefore,omitempty"`
	ClientCertSerialNumber string   `json:"clientcertserialnumber,omitempty"`
	ClientCertSubject      string   `json:"clientcertsubject,omitempty"`
	ClientCertVerification string   `json:"clientcertverification,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Description            string   `json:"description,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Forward                string   `json:"forward,omitempty"`
	Hits                   int      `json:"hits,omitempty"`
	Name                   string   `json:"name,omitempty"`
	NextGenAPIResource     string   `json:"_nextgenapiresource,omitempty"`
	OCSPCache              string   `json:"ocspcache,omitempty"`
	OCSPCertValidation     string   `json:"ocspcertvalidation,omitempty"`
	OCSPStapling           string   `json:"ocspstapling,omitempty"`
	OWASupport             string   `json:"owasupport,omitempty"`
	ReferenceCount         int      `json:"referencecount,omitempty"`
	SessionID              string   `json:"sessionid,omitempty"`
	SessionIDHeader        string   `json:"sessionidheader,omitempty"`
	SSLLogProfile          string   `json:"ssllogprofile,omitempty"`
	UndefHits              int      `json:"undefhits,omitempty"`
}

type SSLServiceGroupSSLCipherSuiteBinding struct {
	CipherName       string `json:"ciphername,omitempty"`
	Description      string `json:"description,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
}

type SSLLogProfile struct {
	Count                float64 `json:"__count,omitempty"`
	Name                 string  `json:"name,omitempty"`
	NextGenAPIResource   string  `json:"_nextgenapiresource,omitempty"`
	SSLLogClAuth         string  `json:"ssllogclauth,omitempty"`
	SSLLogClAuthFailures string  `json:"ssllogclauthfailures,omitempty"`
	SSLLogHS             string  `json:"sslloghs,omitempty"`
	SSLLogHSFailures     string  `json:"sslloghsfailures,omitempty"`
}

type SSLFIPSSimSource struct {
	CertFile     string `json:"certfile,omitempty"`
	SourceSecret string `json:"sourcesecret,omitempty"`
	TargetSecret string `json:"targetsecret,omitempty"`
}

type SSLServiceGroupSSLCertKeyBinding struct {
	CA               bool   `json:"ca,omitempty"`
	CertKeyName      string `json:"certkeyname,omitempty"`
	CRLCheck         string `json:"crlcheck,omitempty"`
	OCSPCheck        string `json:"ocspcheck,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
	SNICert          bool   `json:"snicert,omitempty"`
}

type SSLCertReq struct {
	ChallengePassword    string `json:"challengepassword,omitempty"`
	CommonName           string `json:"commonname,omitempty"`
	CompanyName          string `json:"companyname,omitempty"`
	CountryName          string `json:"countryname,omitempty"`
	DigestMethod         string `json:"digestmethod,omitempty"`
	EmailAddress         string `json:"emailaddress,omitempty"`
	FIPSKeyName          string `json:"fipskeyname,omitempty"`
	KeyFile              string `json:"keyfile,omitempty"`
	KeyForm              string `json:"keyform,omitempty"`
	LocalityName         string `json:"localityname,omitempty"`
	OrganizationName     string `json:"organizationname,omitempty"`
	OrganizationUnitName string `json:"organizationunitname,omitempty"`
	PEMPassphrase        string `json:"pempassphrase,omitempty"`
	ReqFile              string `json:"reqfile,omitempty"`
	StateName            string `json:"statename,omitempty"`
	SubjectAltName       string `json:"subjectaltname,omitempty"`
}

type SSLCACertBundleBinding struct {
	CACertBundleName                             string        `json:"cacertbundlename,omitempty"`
	SSLCACertBundleIntermediateCACertListBinding []interface{} `json:"sslcacertbundle_intermediatecacertlist_binding,omitempty"`
}

type SSLServiceBinding struct {
	ServiceName                      string        `json:"servicename,omitempty"`
	SSLServiceECCCurveBinding        []interface{} `json:"sslservice_ecccurve_binding,omitempty"`
	SSLServiceSSLCACertBundleBinding []interface{} `json:"sslservice_sslcacertbundle_binding,omitempty"`
	SSLServiceSSLCertKeyBinding      []interface{} `json:"sslservice_sslcertkey_binding,omitempty"`
	SSLServiceSSLCipherBinding       []interface{} `json:"sslservice_sslcipher_binding,omitempty"`
	SSLServiceSSLCipherSuiteBinding  []interface{} `json:"sslservice_sslciphersuite_binding,omitempty"`
	SSLServiceSSLPolicyBinding       []interface{} `json:"sslservice_sslpolicy_binding,omitempty"`
}

type SSLCipher struct {
	CipherGroupName    string  `json:"ciphergroupname,omitempty"`
	CipherName         string  `json:"ciphername,omitempty"`
	CipherPriority     int     `json:"cipherpriority,omitempty"`
	CiphGrpAlias       string  `json:"ciphgrpalias,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	SSLProfile         string  `json:"sslprofile,omitempty"`
}

type SSLCertChain struct {
	CertKeyName        string  `json:"certkeyname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type SSLCertKeyBinding struct {
	CertKey                           string        `json:"certkey,omitempty"`
	SSLCertKeyCRLDistributionBinding  []interface{} `json:"sslcertkey_crldistribution_binding,omitempty"`
	SSLCertKeyServiceBinding          []interface{} `json:"sslcertkey_service_binding,omitempty"`
	SSLCertKeySSLOCSPResponderBinding []interface{} `json:"sslcertkey_sslocspresponder_binding,omitempty"`
	SSLCertKeySSLProfileBinding       []interface{} `json:"sslcertkey_sslprofile_binding,omitempty"`
	SSLCertKeySSLVServerBinding       []interface{} `json:"sslcertkey_sslvserver_binding,omitempty"`
}

type SSLPolicyLabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
	FlowType               int     `json:"flowtype,omitempty"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	Invoke                 bool    `json:"invoke,omitempty"`
	InvokeLabelName        string  `json:"invoke_labelname,omitempty"`
	LabelName              string  `json:"labelname,omitempty"`
	LabelType              string  `json:"labeltype,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	NumPol                 int     `json:"numpol,omitempty"`
	PolicyName             string  `json:"policyname,omitempty"`
	Priority               int     `json:"priority,omitempty"`
	TypeField              string  `json:"type,omitempty"`
}

type SSLCertKeySSLProfileBinding struct {
	CA         bool   `json:"ca,omitempty"`
	CertKey    string `json:"certkey,omitempty"`
	SSLProfile string `json:"sslprofile,omitempty"`
}

type SSLVServerSSLPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	PolInherit             int    `json:"polinherit,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	VServerName            string `json:"vservername,omitempty"`
}

type SSLECHConfig struct {
	Count              float64 `json:"__count,omitempty"`
	ECHCipher          string  `json:"echcipher,omitempty"`
	ECHConfigID        int     `json:"echconfigid,omitempty"`
	ECHConfigName      string  `json:"echconfigname,omitempty"`
	ECHPublicName      string  `json:"echpublicname,omitempty"`
	HPKEKeyName        string  `json:"hpkekeyname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Version            int     `json:"version,omitempty"`
}

type SSLCACertBundleIntermediateCACertListBinding struct {
	CACertBundleName    string `json:"cacertbundlename,omitempty"`
	ClientCertNotAfter  string `json:"clientcertnotafter,omitempty"`
	ClientCertNotBefore string `json:"clientcertnotbefore,omitempty"`
	DaysToExpiration    int    `json:"daystoexpiration,omitempty"`
	Issuer              string `json:"issuer,omitempty"`
	PublicKey           string `json:"publickey,omitempty"`
	PublicKeySize       int    `json:"publickeysize,omitempty"`
	SANDNS              string `json:"sandns,omitempty"`
	SANIPAdd            string `json:"sanipadd,omitempty"`
	Serial              string `json:"serial,omitempty"`
	SignatureAlg        string `json:"signaturealg,omitempty"`
	Status              string `json:"status,omitempty"`
	Subject             string `json:"subject,omitempty"`
}

type SSLVServer struct {
	CA                                bool    `json:"ca,omitempty"`
	CipherRedirect                    string  `json:"cipherredirect,omitempty"`
	CipherURL                         string  `json:"cipherurl,omitempty"`
	ClearTextPort                     int     `json:"cleartextport,omitempty"`
	ClientAuth                        string  `json:"clientauth,omitempty"`
	ClientCert                        string  `json:"clientcert,omitempty"`
	Count                             float64 `json:"__count,omitempty"`
	CRLCheck                          string  `json:"crlcheck,omitempty"`
	DefaultSNI                        string  `json:"defaultsni,omitempty"`
	DH                                string  `json:"dh,omitempty"`
	DHCount                           int     `json:"dhcount,omitempty"`
	DHEKeyExchangeWithPSK             string  `json:"dhekeyexchangewithpsk,omitempty"`
	DHFile                            string  `json:"dhfile,omitempty"`
	DHKeyExpSizeLimit                 string  `json:"dhkeyexpsizelimit,omitempty"`
	DTLS1                             string  `json:"dtls1,omitempty"`
	DTLS12                            string  `json:"dtls12,omitempty"`
	DTLSFlag                          bool    `json:"dtlsflag,omitempty"`
	DTLSProfileName                   string  `json:"dtlsprofilename,omitempty"`
	ERSA                              string  `json:"ersa,omitempty"`
	ERSACount                         int     `json:"ersacount,omitempty"`
	HSTS                              string  `json:"hsts,omitempty"`
	IncludeSubdomains                 string  `json:"includesubdomains,omitempty"`
	MaxAge                            int     `json:"maxage,omitempty"`
	NextGenAPIResource                string  `json:"_nextgenapiresource,omitempty"`
	NonFIPSCiphers                    string  `json:"nonfipsciphers,omitempty"`
	OCSPCheck                         string  `json:"ocspcheck,omitempty"`
	OCSPStapling                      string  `json:"ocspstapling,omitempty"`
	Preload                           string  `json:"preload,omitempty"`
	PushEncTrigger                    string  `json:"pushenctrigger,omitempty"`
	QUICFlag                          bool    `json:"quicflag,omitempty"`
	RedirectPortRewrite               string  `json:"redirectportrewrite,omitempty"`
	SendCloseNotify                   string  `json:"sendclosenotify,omitempty"`
	Service                           int     `json:"service,omitempty"`
	SessReuse                         string  `json:"sessreuse,omitempty"`
	SessTimeout                       int     `json:"sesstimeout,omitempty"`
	SkipCACertBundle                  bool    `json:"skipcacertbundle,omitempty"`
	SkipCAName                        bool    `json:"skipcaname,omitempty"`
	SNICert                           bool    `json:"snicert,omitempty"`
	SNIEnable                         string  `json:"snienable,omitempty"`
	SSL2                              string  `json:"ssl2,omitempty"`
	SSL3                              string  `json:"ssl3,omitempty"`
	SSLClientLogs                     string  `json:"sslclientlogs,omitempty"`
	SSLProfile                        string  `json:"sslprofile,omitempty"`
	SSLRedirect                       string  `json:"sslredirect,omitempty"`
	SSLv2Redirect                     string  `json:"sslv2redirect,omitempty"`
	SSLv2URL                          string  `json:"sslv2url,omitempty"`
	StrictSigDigestCheck              string  `json:"strictsigdigestcheck,omitempty"`
	TLS1                              string  `json:"tls1,omitempty"`
	TLS11                             string  `json:"tls11,omitempty"`
	TLS12                             string  `json:"tls12,omitempty"`
	TLS13                             string  `json:"tls13,omitempty"`
	TLS13SessionTicketsPerAuthContext int     `json:"tls13sessionticketsperauthcontext,omitempty"`
	VServerName                       string  `json:"vservername,omitempty"`
	ZeroRTTEarlyData                  string  `json:"zerorttearlydata,omitempty"`
}
