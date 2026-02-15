package models

// ssl configuration structs
type SslcertkeybundleIntermediatecertlinksBinding struct {
	Certkeybundlename   string `json:"certkeybundlename,omitempty"`
	Clientcertnotafter  string `json:"clientcertnotafter,omitempty"`
	Clientcertnotbefore string `json:"clientcertnotbefore,omitempty"`
	Daystoexpiration    int    `json:"daystoexpiration,omitempty"`
	Issuer              string `json:"issuer,omitempty"`
	Publickey           string `json:"publickey,omitempty"`
	Publickeysize       int    `json:"publickeysize,omitempty"`
	Sandns              string `json:"sandns,omitempty"`
	Sanipadd            string `json:"sanipadd,omitempty"`
	Serial              string `json:"serial,omitempty"`
	Signaturealg        string `json:"signaturealg,omitempty"`
	Status              string `json:"status,omitempty"`
	Subject             string `json:"subject,omitempty"`
}

type Sslhpkekey struct {
	Count              float64 `json:"__count,omitempty"`
	Dhkem              string  `json:"dhkem,omitempty"`
	File               string  `json:"file,omitempty"`
	Hpkekeyname        string  `json:"hpkekeyname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type SslpolicyCsvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type SslcipherBinding struct {
	Ciphergroupname                  string        `json:"ciphergroupname,omitempty"`
	SslcipherIndividualcipherBinding []interface{} `json:"sslcipher_individualcipher_binding,omitempty"`
	SslcipherSslciphersuiteBinding   []interface{} `json:"sslcipher_sslciphersuite_binding,omitempty"`
	SslcipherSslprofileBinding       []interface{} `json:"sslcipher_sslprofile_binding,omitempty"`
}

type SslpolicySslpolicylabelBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type SslvserverSslcertkeyBinding struct {
	Ca            bool   `json:"ca,omitempty"`
	Certkeyname   string `json:"certkeyname,omitempty"`
	Cleartextport int    `json:"cleartextport,omitempty"`
	Crlcheck      string `json:"crlcheck,omitempty"`
	Ocspcheck     string `json:"ocspcheck,omitempty"`
	Skipcaname    bool   `json:"skipcaname,omitempty"`
	Snicert       bool   `json:"snicert,omitempty"`
	Vservername   string `json:"vservername,omitempty"`
}

type SslcrlSerialnumberBinding struct {
	Crlname string `json:"crlname,omitempty"`
	Date    string `json:"date,omitempty"`
	Number  string `json:"number,omitempty"`
}

type Sslfipskey struct {
	Count              float64 `json:"__count,omitempty"`
	Curve              string  `json:"curve,omitempty"`
	Exponent           string  `json:"exponent,omitempty"`
	Fipskeyname        string  `json:"fipskeyname,omitempty"`
	Inform             string  `json:"inform,omitempty"`
	Iv                 string  `json:"iv,omitempty"`
	Key                string  `json:"key,omitempty"`
	Keytype            string  `json:"keytype,omitempty"`
	Modulus            int     `json:"modulus,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Size               int     `json:"size,omitempty"`
	Wrapkeyname        string  `json:"wrapkeyname,omitempty"`
}

type SslcipherSslprofileBinding struct {
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
	Description     string `json:"description,omitempty"`
	Sslprofile      string `json:"sslprofile,omitempty"`
}

type Sslfipssimtarget struct {
	Certfile     string `json:"certfile,omitempty"`
	Keyvector    string `json:"keyvector,omitempty"`
	Sourcesecret string `json:"sourcesecret,omitempty"`
	Targetsecret string `json:"targetsecret,omitempty"`
}

type SslserviceSslpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Polinherit             int    `json:"polinherit,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Servicename            string `json:"servicename,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type SslprofileSslciphersuiteBinding struct {
	Ciphername     string `json:"ciphername,omitempty"`
	Cipherpriority int    `json:"cipherpriority,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
}

type SslpolicySslserviceBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type SslvserverBinding struct {
	SslvserverEcccurveBinding         []interface{} `json:"sslvserver_ecccurve_binding,omitempty"`
	SslvserverHashicorpBinding        []interface{} `json:"sslvserver_hashicorp_binding,omitempty"`
	SslvserverSslcacertbundleBinding  []interface{} `json:"sslvserver_sslcacertbundle_binding,omitempty"`
	SslvserverSslcertkeyBinding       []interface{} `json:"sslvserver_sslcertkey_binding,omitempty"`
	SslvserverSslcertkeybundleBinding []interface{} `json:"sslvserver_sslcertkeybundle_binding,omitempty"`
	SslvserverSslcipherBinding        []interface{} `json:"sslvserver_sslcipher_binding,omitempty"`
	SslvserverSslciphersuiteBinding   []interface{} `json:"sslvserver_sslciphersuite_binding,omitempty"`
	SslvserverSslpolicyBinding        []interface{} `json:"sslvserver_sslpolicy_binding,omitempty"`
	Vservername                       string        `json:"vservername,omitempty"`
}

type Sslecdsakey struct {
	Aes256   bool   `json:"aes256,omitempty"`
	Curve    string `json:"curve,omitempty"`
	Des      bool   `json:"des,omitempty"`
	Des3     bool   `json:"des3,omitempty"`
	Keyfile  string `json:"keyfile,omitempty"`
	Keyform  string `json:"keyform,omitempty"`
	Password string `json:"password,omitempty"`
	Pkcs8    bool   `json:"pkcs8,omitempty"`
}

type SslservicegroupSslcacertbundleBinding struct {
	Cacertbundlename string `json:"cacertbundlename,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type Sslhsmkey struct {
	Count              float64 `json:"__count,omitempty"`
	Hsmkeyname         string  `json:"hsmkeyname,omitempty"`
	Hsmtype            string  `json:"hsmtype,omitempty"`
	Key                string  `json:"key,omitempty"`
	Keystore           string  `json:"keystore,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Serialnum          string  `json:"serialnum,omitempty"`
	State              string  `json:"state,omitempty"`
}

type SslcacertgroupBinding struct {
	Cacertgroupname                 string        `json:"cacertgroupname,omitempty"`
	SslcacertgroupSslcertkeyBinding []interface{} `json:"sslcacertgroup_sslcertkey_binding,omitempty"`
}

type Sslciphersuite struct {
	Ciphername         string  `json:"ciphername,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Description        string  `json:"description,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type SslcrlBinding struct {
	Crlname                   string        `json:"crlname,omitempty"`
	SslcrlSerialnumberBinding []interface{} `json:"sslcrl_serialnumber_binding,omitempty"`
}

type SslpolicylabelBinding struct {
	Labelname                      string        `json:"labelname,omitempty"`
	SslpolicylabelSslpolicyBinding []interface{} `json:"sslpolicylabel_sslpolicy_binding,omitempty"`
}

type SslservicegroupBinding struct {
	Servicegroupname                      string        `json:"servicegroupname,omitempty"`
	SslservicegroupEcccurveBinding        []interface{} `json:"sslservicegroup_ecccurve_binding,omitempty"`
	SslservicegroupSslcacertbundleBinding []interface{} `json:"sslservicegroup_sslcacertbundle_binding,omitempty"`
	SslservicegroupSslcertkeyBinding      []interface{} `json:"sslservicegroup_sslcertkey_binding,omitempty"`
	SslservicegroupSslcipherBinding       []interface{} `json:"sslservicegroup_sslcipher_binding,omitempty"`
	SslservicegroupSslciphersuiteBinding  []interface{} `json:"sslservicegroup_sslciphersuite_binding,omitempty"`
}

type Sslwrapkey struct {
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Salt               string  `json:"salt,omitempty"`
	Wrapkeyname        string  `json:"wrapkeyname,omitempty"`
}

type Sslocspresponder struct {
	Batchingdelay         int     `json:"batchingdelay,omitempty"`
	Batchingdepth         int     `json:"batchingdepth,omitempty"`
	Cache                 string  `json:"cache,omitempty"`
	Cachetimeout          int     `json:"cachetimeout,omitempty"`
	Count                 float64 `json:"__count,omitempty"`
	Httpmethod            string  `json:"httpmethod,omitempty"`
	Insertclientcert      string  `json:"insertclientcert,omitempty"`
	Name                  string  `json:"name,omitempty"`
	Nextgenapiresource    string  `json:"_nextgenapiresource,omitempty"`
	Ocspaiarefcount       int     `json:"ocspaiarefcount,omitempty"`
	Ocspipaddrstr         string  `json:"ocspipaddrstr,omitempty"`
	Ocspurlresolvetimeout int     `json:"ocspurlresolvetimeout,omitempty"`
	Port                  int     `json:"port,omitempty"`
	Producedattimeskew    int     `json:"producedattimeskew,omitempty"`
	Respondercert         string  `json:"respondercert,omitempty"`
	Resptimeout           int     `json:"resptimeout,omitempty"`
	Signingcert           string  `json:"signingcert,omitempty"`
	Trustresponder        bool    `json:"trustresponder,omitempty"`
	Url                   string  `json:"url,omitempty"`
	Usenonce              string  `json:"usenonce,omitempty"`
}

type Sslcrl struct {
	Basedn             string  `json:"basedn,omitempty"`
	Binary             string  `json:"binary,omitempty"`
	Binddn             string  `json:"binddn,omitempty"`
	Cacert             string  `json:"cacert,omitempty"`
	Cacertfile         string  `json:"cacertfile,omitempty"`
	Cakeyfile          string  `json:"cakeyfile,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Crlname            string  `json:"crlname,omitempty"`
	Crlpath            string  `json:"crlpath,omitempty"`
	Day                int     `json:"day,omitempty"`
	Daystoexpiration   int     `json:"daystoexpiration,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	Gencrl             string  `json:"gencrl,omitempty"`
	Indexfile          string  `json:"indexfile,omitempty"`
	Inform             string  `json:"inform,omitempty"`
	Interval           string  `json:"interval,omitempty"`
	Issuer             string  `json:"issuer,omitempty"`
	Lastupdate         string  `json:"lastupdate,omitempty"`
	Lastupdatetime     int     `json:"lastupdatetime,omitempty"`
	Method             string  `json:"method,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nextupdate         string  `json:"nextupdate,omitempty"`
	Password           string  `json:"password,omitempty"`
	Port               int     `json:"port,omitempty"`
	Refresh            string  `json:"refresh,omitempty"`
	Revoke             string  `json:"revoke,omitempty"`
	Scope              string  `json:"scope,omitempty"`
	Server             string  `json:"server,omitempty"`
	Signaturealgo      string  `json:"signaturealgo,omitempty"`
	Time               string  `json:"time,omitempty"`
	Url                string  `json:"url,omitempty"`
	Version            int     `json:"version,omitempty"`
}

type SslcertkeybundleBinding struct {
	Certkeybundlename                            string        `json:"certkeybundlename,omitempty"`
	SslcertkeybundleIntermediatecertlinksBinding []interface{} `json:"sslcertkeybundle_intermediatecertlinks_binding,omitempty"`
	SslcertkeybundleSslvserverBinding            []interface{} `json:"sslcertkeybundle_sslvserver_binding,omitempty"`
}

type SslprofileSslcipherBinding struct {
	Cipheraliasname string `json:"cipheraliasname,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
	Description     string `json:"description,omitempty"`
	Name            string `json:"name,omitempty"`
}

type SslpolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Sslpkcs8 struct {
	Keyfile   string `json:"keyfile,omitempty"`
	Keyform   string `json:"keyform,omitempty"`
	Password  string `json:"password,omitempty"`
	Pkcs8file string `json:"pkcs8file,omitempty"`
}

type SslvserverSslcacertbundleBinding struct {
	Cacertbundlename string `json:"cacertbundlename,omitempty"`
	Skipcacertbundle bool   `json:"skipcacertbundle,omitempty"`
	Vservername      string `json:"vservername,omitempty"`
}

type SslprofileBinding struct {
	Name                            string        `json:"name,omitempty"`
	SslprofileEcccurveBinding       []interface{} `json:"sslprofile_ecccurve_binding,omitempty"`
	SslprofileSslcertkeyBinding     []interface{} `json:"sslprofile_sslcertkey_binding,omitempty"`
	SslprofileSslcipherBinding      []interface{} `json:"sslprofile_sslcipher_binding,omitempty"`
	SslprofileSslciphersuiteBinding []interface{} `json:"sslprofile_sslciphersuite_binding,omitempty"`
	SslprofileSslechconfigBinding   []interface{} `json:"sslprofile_sslechconfig_binding,omitempty"`
	SslprofileSslvserverBinding     []interface{} `json:"sslprofile_sslvserver_binding,omitempty"`
}

type SslcertkeyServiceBinding struct {
	Ca               bool   `json:"ca,omitempty"`
	Certkey          string `json:"certkey,omitempty"`
	Data             int    `json:"data,omitempty"`
	Service          bool   `json:"service,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Version          int    `json:"version,omitempty"`
}

type SslvserverEcccurveBinding struct {
	Ecccurvename string `json:"ecccurvename,omitempty"`
	Vservername  string `json:"vservername,omitempty"`
}

type Sslcacertgroup struct {
	Cacertgroupname       string  `json:"cacertgroupname,omitempty"`
	Cacertgroupreferences int     `json:"cacertgroupreferences,omitempty"`
	Count                 float64 `json:"__count,omitempty"`
	Crlcheck              string  `json:"crlcheck,omitempty"`
	Nextgenapiresource    string  `json:"_nextgenapiresource,omitempty"`
	Ocspcheck             string  `json:"ocspcheck,omitempty"`
}

type SslserviceSslcipherBinding struct {
	Cipheraliasname string `json:"cipheraliasname,omitempty"`
	Cipherdefaulton int    `json:"cipherdefaulton,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
	Description     string `json:"description,omitempty"`
	Servicename     string `json:"servicename,omitempty"`
}

type SslvserverSslcipherBinding struct {
	Cipheraliasname string `json:"cipheraliasname,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
	Description     string `json:"description,omitempty"`
	Vservername     string `json:"vservername,omitempty"`
}

type Ssldefaultprofile struct {
}

type SslcertkeyCrldistributionBinding struct {
	Ca      bool   `json:"ca,omitempty"`
	Certkey string `json:"certkey,omitempty"`
	Issuer  string `json:"issuer,omitempty"`
}

type Sslcacertbundle struct {
	Bundlefile         string  `json:"bundlefile,omitempty"`
	Cacertbundlename   string  `json:"cacertbundlename,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Servername         string  `json:"servername,omitempty"`
}

type Ssldtlsprofile struct {
	Builtin              []string `json:"builtin,omitempty"`
	Count                float64  `json:"__count,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	Helloverifyrequest   string   `json:"helloverifyrequest,omitempty"`
	Initialretrytimeout  int      `json:"initialretrytimeout,omitempty"`
	Maxbadmacignorecount int      `json:"maxbadmacignorecount,omitempty"`
	Maxholdqlen          int      `json:"maxholdqlen,omitempty"`
	Maxpacketsize        int      `json:"maxpacketsize,omitempty"`
	Maxrecordsize        int      `json:"maxrecordsize,omitempty"`
	Maxretrytime         int      `json:"maxretrytime,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Nextgenapiresource   string   `json:"_nextgenapiresource,omitempty"`
	Pmtudiscovery        string   `json:"pmtudiscovery,omitempty"`
	Terminatesession     string   `json:"terminatesession,omitempty"`
}

type Sslcert struct {
	Cacert         string `json:"cacert,omitempty"`
	Cacertform     string `json:"cacertform,omitempty"`
	Cakey          string `json:"cakey,omitempty"`
	Cakeyform      string `json:"cakeyform,omitempty"`
	Caserial       string `json:"caserial,omitempty"`
	Certfile       string `json:"certfile,omitempty"`
	Certform       string `json:"certform,omitempty"`
	Certtype       string `json:"certtype,omitempty"`
	Days           int    `json:"days,omitempty"`
	Keyfile        string `json:"keyfile,omitempty"`
	Keyform        string `json:"keyform,omitempty"`
	Pempassphrase  string `json:"pempassphrase,omitempty"`
	Reqfile        string `json:"reqfile,omitempty"`
	Subjectaltname string `json:"subjectaltname,omitempty"`
}

type SslpolicySslglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Sslrsakey struct {
	Aes256   bool   `json:"aes256,omitempty"`
	Bits     int    `json:"bits,omitempty"`
	Des      bool   `json:"des,omitempty"`
	Des3     bool   `json:"des3,omitempty"`
	Exponent string `json:"exponent,omitempty"`
	Keyfile  string `json:"keyfile,omitempty"`
	Keyform  string `json:"keyform,omitempty"`
	Password string `json:"password,omitempty"`
	Pkcs8    bool   `json:"pkcs8,omitempty"`
}

type SslserviceEcccurveBinding struct {
	Ecccurvename string `json:"ecccurvename,omitempty"`
	Servicename  string `json:"servicename,omitempty"`
}

type SslserviceSslcertkeyBinding struct {
	Ca            bool   `json:"ca,omitempty"`
	Certkeyname   string `json:"certkeyname,omitempty"`
	Cleartextport int    `json:"cleartextport,omitempty"`
	Crlcheck      string `json:"crlcheck,omitempty"`
	Ocspcheck     string `json:"ocspcheck,omitempty"`
	Servicename   string `json:"servicename,omitempty"`
	Skipcaname    bool   `json:"skipcaname,omitempty"`
	Snicert       bool   `json:"snicert,omitempty"`
}

type Sslservicegroup struct {
	Ca                   bool    `json:"ca,omitempty"`
	Cipherredirect       string  `json:"cipherredirect,omitempty"`
	Cipherurl            string  `json:"cipherurl,omitempty"`
	Cleartextport        int     `json:"cleartextport,omitempty"`
	Clientauth           string  `json:"clientauth,omitempty"`
	Clientcert           string  `json:"clientcert,omitempty"`
	Commonname           string  `json:"commonname,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	Crlcheck             string  `json:"crlcheck,omitempty"`
	Dh                   string  `json:"dh,omitempty"`
	Dhcount              int     `json:"dhcount,omitempty"`
	Dhfile               string  `json:"dhfile,omitempty"`
	Dhkeyexpsizelimit    string  `json:"dhkeyexpsizelimit,omitempty"`
	Ersa                 string  `json:"ersa,omitempty"`
	Ersacount            int     `json:"ersacount,omitempty"`
	Nextgenapiresource   string  `json:"_nextgenapiresource,omitempty"`
	Nonfipsciphers       string  `json:"nonfipsciphers,omitempty"`
	Ocspcheck            string  `json:"ocspcheck,omitempty"`
	Ocspstapling         string  `json:"ocspstapling,omitempty"`
	Quicflag             bool    `json:"quicflag,omitempty"`
	Redirectportrewrite  string  `json:"redirectportrewrite,omitempty"`
	Sendclosenotify      string  `json:"sendclosenotify,omitempty"`
	Serverauth           string  `json:"serverauth,omitempty"`
	Servicegroupname     string  `json:"servicegroupname,omitempty"`
	Servicename          string  `json:"servicename,omitempty"`
	Sessreuse            string  `json:"sessreuse,omitempty"`
	Sesstimeout          int     `json:"sesstimeout,omitempty"`
	Snicert              bool    `json:"snicert,omitempty"`
	Snienable            string  `json:"snienable,omitempty"`
	Ssl2                 string  `json:"ssl2,omitempty"`
	Ssl3                 string  `json:"ssl3,omitempty"`
	Sslclientlogs        string  `json:"sslclientlogs,omitempty"`
	Sslprofile           string  `json:"sslprofile,omitempty"`
	Sslredirect          string  `json:"sslredirect,omitempty"`
	Sslv2redirect        string  `json:"sslv2redirect,omitempty"`
	Sslv2url             string  `json:"sslv2url,omitempty"`
	Strictsigdigestcheck string  `json:"strictsigdigestcheck,omitempty"`
	Tls1                 string  `json:"tls1,omitempty"`
	Tls11                string  `json:"tls11,omitempty"`
	Tls12                string  `json:"tls12,omitempty"`
	Tls13                string  `json:"tls13,omitempty"`
}

type SslprofileSslvserverBinding struct {
	Cipherpriority int    `json:"cipherpriority,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	Servicename    string `json:"servicename,omitempty"`
}

type Sslcertlink struct {
	Certkeyname        string  `json:"certkeyname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Linkcertkeyname    string  `json:"linkcertkeyname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type SslglobalSslpolicyBinding struct {
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type Sslcrlfile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type SslcertchainSslcertkeyBinding struct {
	Addsubject      bool   `json:"addsubject,omitempty"`
	Certkeyname     string `json:"certkeyname,omitempty"`
	Isca            bool   `json:"isca,omitempty"`
	Islinked        bool   `json:"islinked,omitempty"`
	Linkcertkeyname string `json:"linkcertkeyname,omitempty"`
}

type Sslparameter struct {
	Crlmemorysizemb          int      `json:"crlmemorysizemb,omitempty"`
	Cryptodevdisablelimit    int      `json:"cryptodevdisablelimit,omitempty"`
	Defaultprofile           string   `json:"defaultprofile,omitempty"`
	Denysslreneg             string   `json:"denysslreneg,omitempty"`
	Dropreqwithnohostheader  string   `json:"dropreqwithnohostheader,omitempty"`
	Encrypttriggerpktcount   int      `json:"encrypttriggerpktcount,omitempty"`
	Heterogeneoussslhw       string   `json:"heterogeneoussslhw,omitempty"`
	Hybridfipsmode           string   `json:"hybridfipsmode,omitempty"`
	Insertcertspace          string   `json:"insertcertspace,omitempty"`
	Insertionencoding        string   `json:"insertionencoding,omitempty"`
	Montls1112disable        string   `json:"montls1112disable,omitempty"`
	Ndcppcompliancecertcheck string   `json:"ndcppcompliancecertcheck,omitempty"`
	Nextgenapiresource       string   `json:"_nextgenapiresource,omitempty"`
	Ocspcachesize            int      `json:"ocspcachesize,omitempty"`
	Operationqueuelimit      int      `json:"operationqueuelimit,omitempty"`
	Pushenctriggertimeout    int      `json:"pushenctriggertimeout,omitempty"`
	Pushflag                 int      `json:"pushflag,omitempty"`
	Quantumsize              string   `json:"quantumsize,omitempty"`
	Sendclosenotify          string   `json:"sendclosenotify,omitempty"`
	Sigdigesttype            []string `json:"sigdigesttype,omitempty"`
	Snihttphostmatch         string   `json:"snihttphostmatch,omitempty"`
	Softwarecryptothreshold  int      `json:"softwarecryptothreshold,omitempty"`
	Sslierrorcache           string   `json:"sslierrorcache,omitempty"`
	Sslimaxerrorcachemem     int      `json:"sslimaxerrorcachemem,omitempty"`
	Ssltriggertimeout        int      `json:"ssltriggertimeout,omitempty"`
	Strictcachecks           string   `json:"strictcachecks,omitempty"`
	Svctls1112disable        string   `json:"svctls1112disable,omitempty"`
	Undefactioncontrol       string   `json:"undefactioncontrol,omitempty"`
	Undefactiondata          string   `json:"undefactiondata,omitempty"`
}

type Ssldhparam struct {
	Bits   int    `json:"bits,omitempty"`
	Dhfile string `json:"dhfile,omitempty"`
	Gen    string `json:"gen,omitempty"`
}

type Sslpolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Reqaction          string   `json:"reqaction,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Undefaction        string   `json:"undefaction,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type Sslcertbundle struct {
	Count              float64 `json:"__count,omitempty"`
	Inuse              string  `json:"inuse,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type SslcertkeySslvserverBinding struct {
	Ca          bool   `json:"ca,omitempty"`
	Certkey     string `json:"certkey,omitempty"`
	Data        int    `json:"data,omitempty"`
	Servername  string `json:"servername,omitempty"`
	Version     int    `json:"version,omitempty"`
	Vserver     bool   `json:"vserver,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}

type SslcertkeybundleSslvserverBinding struct {
	Certkeybundlename string `json:"certkeybundlename,omitempty"`
	Servername        string `json:"servername,omitempty"`
}

type SslservicegroupEcccurveBinding struct {
	Ecccurvename     string `json:"ecccurvename,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type SslprofileSslcertkeyBinding struct {
	Certkeyname      string `json:"certkeyname,omitempty"`
	Cipherpriority   int    `json:"cipherpriority,omitempty"`
	Forgingcacertkey bool   `json:"forgingcacertkey,omitempty"`
	Name             string `json:"name,omitempty"`
	Sslicacertkey    string `json:"sslicacertkey,omitempty"`
}

type Sslcertfile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type SslcacertgroupSslcertkeyBinding struct {
	Cacertgroupname string `json:"cacertgroupname,omitempty"`
	Certkeyname     string `json:"certkeyname,omitempty"`
	Crlcheck        string `json:"crlcheck,omitempty"`
	Ocspcheck       string `json:"ocspcheck,omitempty"`
}

type SslservicegroupSslcipherBinding struct {
	Cipheraliasname  string `json:"cipheraliasname,omitempty"`
	Ciphername       string `json:"ciphername,omitempty"`
	Description      string `json:"description,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type Sslcertkey struct {
	Builtin                     []string `json:"builtin,omitempty"`
	Bundle                      string   `json:"bundle,omitempty"`
	Cert                        string   `json:"cert,omitempty"`
	Certificatesource           string   `json:"certificatesource,omitempty"`
	Certificatetype             []string `json:"certificatetype,omitempty"`
	Certkey                     string   `json:"certkey,omitempty"`
	Certkeydigest               string   `json:"certkeydigest,omitempty"`
	Certkeystatus               string   `json:"certkeystatus,omitempty"`
	Clientcertnotafter          string   `json:"clientcertnotafter,omitempty"`
	Clientcertnotbefore         string   `json:"clientcertnotbefore,omitempty"`
	Count                       float64  `json:"__count,omitempty"`
	Data                        int      `json:"data,omitempty"`
	Daystoexpiration            int      `json:"daystoexpiration,omitempty"`
	Deletecertkeyfilesonremoval string   `json:"deletecertkeyfilesonremoval,omitempty"`
	Deletefromdevice            bool     `json:"deletefromdevice,omitempty"`
	Expirymonitor               string   `json:"expirymonitor,omitempty"`
	Feature                     string   `json:"feature,omitempty"`
	Fipskey                     string   `json:"fipskey,omitempty"`
	Hsmkey                      string   `json:"hsmkey,omitempty"`
	Inform                      string   `json:"inform,omitempty"`
	Issuer                      string   `json:"issuer,omitempty"`
	Key                         string   `json:"key,omitempty"`
	Linkcertkeyname             string   `json:"linkcertkeyname,omitempty"`
	Nextgenapiresource          string   `json:"_nextgenapiresource,omitempty"`
	Nodomaincheck               bool     `json:"nodomaincheck,omitempty"`
	Notificationperiod          int      `json:"notificationperiod,omitempty"`
	Ocspresponsestatus          string   `json:"ocspresponsestatus,omitempty"`
	Ocspstaplingcache           bool     `json:"ocspstaplingcache,omitempty"`
	Passcrypt                   string   `json:"passcrypt,omitempty"`
	Passplain                   string   `json:"passplain,omitempty"`
	Password                    bool     `json:"password,omitempty"`
	Priority                    int      `json:"priority,omitempty"`
	Publickey                   string   `json:"publickey,omitempty"`
	Publickeysize               int      `json:"publickeysize,omitempty"`
	Sandns                      string   `json:"sandns,omitempty"`
	Sanipadd                    string   `json:"sanipadd,omitempty"`
	Serial                      string   `json:"serial,omitempty"`
	Servicename                 string   `json:"servicename,omitempty"`
	Signaturealg                string   `json:"signaturealg,omitempty"`
	Status                      string   `json:"status,omitempty"`
	Subject                     string   `json:"subject,omitempty"`
	Version                     int      `json:"version,omitempty"`
}

type SslprofileSslechconfigBinding struct {
	Cipherpriority int    `json:"cipherpriority,omitempty"`
	Echconfigname  string `json:"echconfigname,omitempty"`
	Name           string `json:"name,omitempty"`
}

type SslserviceSslciphersuiteBinding struct {
	Cipherdefaulton int    `json:"cipherdefaulton,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
	Description     string `json:"description,omitempty"`
	Servicename     string `json:"servicename,omitempty"`
}

type Sslcertkeybundle struct {
	Bundlefile         string  `json:"bundlefile,omitempty"`
	Certkeybundlename  string  `json:"certkeybundlename,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Passplain          string  `json:"passplain,omitempty"`
}

type SslcertchainBinding struct {
	Certkeyname                   string        `json:"certkeyname,omitempty"`
	SslcertchainSslcertkeyBinding []interface{} `json:"sslcertchain_sslcertkey_binding,omitempty"`
}

type Sslcertificatechain struct {
	Certkeyname        string   `json:"certkeyname,omitempty"`
	Chaincomplete      int      `json:"chaincomplete,omitempty"`
	Chainissuer        string   `json:"chainissuer,omitempty"`
	Chainlinked        []string `json:"chainlinked,omitempty"`
	Chainpossiblelinks []string `json:"chainpossiblelinks,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type SslprofileEcccurveBinding struct {
	Cipherpriority int    `json:"cipherpriority,omitempty"`
	Ecccurvename   string `json:"ecccurvename,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Sslservice struct {
	Cipherredirect       string  `json:"cipherredirect,omitempty"`
	Cipherurl            string  `json:"cipherurl,omitempty"`
	Clientauth           string  `json:"clientauth,omitempty"`
	Clientcert           string  `json:"clientcert,omitempty"`
	Commonname           string  `json:"commonname,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	Dh                   string  `json:"dh,omitempty"`
	Dhcount              int     `json:"dhcount,omitempty"`
	Dhfile               string  `json:"dhfile,omitempty"`
	Dhkeyexpsizelimit    string  `json:"dhkeyexpsizelimit,omitempty"`
	Dtls1                string  `json:"dtls1,omitempty"`
	Dtls12               string  `json:"dtls12,omitempty"`
	Dtlsflag             bool    `json:"dtlsflag,omitempty"`
	Dtlsprofilename      string  `json:"dtlsprofilename,omitempty"`
	Ersa                 string  `json:"ersa,omitempty"`
	Ersacount            int     `json:"ersacount,omitempty"`
	Nextgenapiresource   string  `json:"_nextgenapiresource,omitempty"`
	Nonfipsciphers       string  `json:"nonfipsciphers,omitempty"`
	Ocspstapling         string  `json:"ocspstapling,omitempty"`
	Pushenctrigger       string  `json:"pushenctrigger,omitempty"`
	Quicflag             bool    `json:"quicflag,omitempty"`
	Redirectportrewrite  string  `json:"redirectportrewrite,omitempty"`
	Sendclosenotify      string  `json:"sendclosenotify,omitempty"`
	Serverauth           string  `json:"serverauth,omitempty"`
	Service              int     `json:"service,omitempty"`
	Servicename          string  `json:"servicename,omitempty"`
	Sessreuse            string  `json:"sessreuse,omitempty"`
	Sesstimeout          int     `json:"sesstimeout,omitempty"`
	Skipcacertbundle     bool    `json:"skipcacertbundle,omitempty"`
	Skipcaname           bool    `json:"skipcaname,omitempty"`
	Snienable            string  `json:"snienable,omitempty"`
	Ssl2                 string  `json:"ssl2,omitempty"`
	Ssl3                 string  `json:"ssl3,omitempty"`
	Sslclientlogs        string  `json:"sslclientlogs,omitempty"`
	Sslprofile           string  `json:"sslprofile,omitempty"`
	Sslredirect          string  `json:"sslredirect,omitempty"`
	Sslv2redirect        string  `json:"sslv2redirect,omitempty"`
	Sslv2url             string  `json:"sslv2url,omitempty"`
	Strictsigdigestcheck string  `json:"strictsigdigestcheck,omitempty"`
	Tls1                 string  `json:"tls1,omitempty"`
	Tls11                string  `json:"tls11,omitempty"`
	Tls12                string  `json:"tls12,omitempty"`
	Tls13                string  `json:"tls13,omitempty"`
}

type Sslprofile struct {
	Allowextendedmastersecret         string   `json:"allowextendedmastersecret,omitempty"`
	Allowlegacykdf                    string   `json:"allowlegacykdf,omitempty"`
	Allowunknownsni                   string   `json:"allowunknownsni,omitempty"`
	Alpnprotocol                      string   `json:"alpnprotocol,omitempty"`
	Builtin                           []string `json:"builtin,omitempty"`
	Ciphername                        string   `json:"ciphername,omitempty"`
	Cipherpriority                    int      `json:"cipherpriority,omitempty"`
	Cipherredirect                    string   `json:"cipherredirect,omitempty"`
	Cipherurl                         string   `json:"cipherurl,omitempty"`
	Cleartextport                     int      `json:"cleartextport,omitempty"`
	Clientauth                        string   `json:"clientauth,omitempty"`
	Clientauthuseboundcachain         string   `json:"clientauthuseboundcachain,omitempty"`
	Clientcert                        string   `json:"clientcert,omitempty"`
	Commonname                        string   `json:"commonname,omitempty"`
	Count                             float64  `json:"__count,omitempty"`
	Crlcheck                          string   `json:"crlcheck,omitempty"`
	Defaultsni                        string   `json:"defaultsni,omitempty"`
	Denysslreneg                      string   `json:"denysslreneg,omitempty"`
	Dh                                string   `json:"dh,omitempty"`
	Dhcount                           int      `json:"dhcount,omitempty"`
	Dhekeyexchangewithpsk             string   `json:"dhekeyexchangewithpsk,omitempty"`
	Dhfile                            string   `json:"dhfile,omitempty"`
	Dhkeyexpsizelimit                 string   `json:"dhkeyexpsizelimit,omitempty"`
	Dropreqwithnohostheader           string   `json:"dropreqwithnohostheader,omitempty"`
	Dynamicclientcert                 string   `json:"dynamicclientcert,omitempty"`
	Encryptedclienthello              string   `json:"encryptedclienthello,omitempty"`
	Encrypttriggerpktcount            int      `json:"encrypttriggerpktcount,omitempty"`
	Ersa                              string   `json:"ersa,omitempty"`
	Ersacount                         int      `json:"ersacount,omitempty"`
	Feature                           string   `json:"feature,omitempty"`
	Hsts                              string   `json:"hsts,omitempty"`
	Includesubdomains                 string   `json:"includesubdomains,omitempty"`
	Insertionencoding                 string   `json:"insertionencoding,omitempty"`
	Invoke                            bool     `json:"invoke,omitempty"`
	Labeltype                         string   `json:"labeltype,omitempty"`
	Maxage                            int      `json:"maxage,omitempty"`
	Maxrenegrate                      int      `json:"maxrenegrate,omitempty"`
	Name                              string   `json:"name,omitempty"`
	Nextgenapiresource                string   `json:"_nextgenapiresource,omitempty"`
	Nodefaultbindings                 string   `json:"nodefaultbindings,omitempty"`
	Nonfipsciphers                    string   `json:"nonfipsciphers,omitempty"`
	Ocspcheck                         string   `json:"ocspcheck,omitempty"`
	Ocspstapling                      string   `json:"ocspstapling,omitempty"`
	Preload                           string   `json:"preload,omitempty"`
	Prevsessionkeylifetime            int      `json:"prevsessionkeylifetime,omitempty"`
	Pushenctrigger                    string   `json:"pushenctrigger,omitempty"`
	Pushenctriggertimeout             int      `json:"pushenctriggertimeout,omitempty"`
	Pushflag                          int      `json:"pushflag,omitempty"`
	Quantumsize                       string   `json:"quantumsize,omitempty"`
	Redirectportrewrite               string   `json:"redirectportrewrite,omitempty"`
	Sendclosenotify                   string   `json:"sendclosenotify,omitempty"`
	Serverauth                        string   `json:"serverauth,omitempty"`
	Service                           int      `json:"service,omitempty"`
	Sessionkeylifetime                int      `json:"sessionkeylifetime,omitempty"`
	Sessionticket                     string   `json:"sessionticket,omitempty"`
	Sessionticketkeydata              string   `json:"sessionticketkeydata,omitempty"`
	Sessionticketkeyrefresh           string   `json:"sessionticketkeyrefresh,omitempty"`
	Sessionticketlifetime             int      `json:"sessionticketlifetime,omitempty"`
	Sessreuse                         string   `json:"sessreuse,omitempty"`
	Sesstimeout                       int      `json:"sesstimeout,omitempty"`
	Skipcaname                        bool     `json:"skipcaname,omitempty"`
	Skipclientcertpolicycheck         string   `json:"skipclientcertpolicycheck,omitempty"`
	Snicert                           bool     `json:"snicert,omitempty"`
	Snienable                         string   `json:"snienable,omitempty"`
	Snihttphostmatch                  string   `json:"snihttphostmatch,omitempty"`
	Ssl3                              string   `json:"ssl3,omitempty"`
	Sslclientlogs                     string   `json:"sslclientlogs,omitempty"`
	Sslimaxsessperserver              int      `json:"sslimaxsessperserver,omitempty"`
	Sslinterception                   string   `json:"sslinterception,omitempty"`
	Ssliocspcheck                     string   `json:"ssliocspcheck,omitempty"`
	Sslireneg                         string   `json:"sslireneg,omitempty"`
	Ssliverifyservercertforreuse      string   `json:"ssliverifyservercertforreuse,omitempty"`
	Ssllogprofile                     string   `json:"ssllogprofile,omitempty"`
	Sslpfobjecttype                   int      `json:"sslpfobjecttype,omitempty"`
	Sslprofiletype                    string   `json:"sslprofiletype,omitempty"`
	Sslredirect                       string   `json:"sslredirect,omitempty"`
	Ssltriggertimeout                 int      `json:"ssltriggertimeout,omitempty"`
	Strictcachecks                    string   `json:"strictcachecks,omitempty"`
	Strictsigdigestcheck              string   `json:"strictsigdigestcheck,omitempty"`
	Tls1                              string   `json:"tls1,omitempty"`
	Tls11                             string   `json:"tls11,omitempty"`
	Tls12                             string   `json:"tls12,omitempty"`
	Tls13                             string   `json:"tls13,omitempty"`
	Tls13sessionticketsperauthcontext int      `json:"tls13sessionticketsperauthcontext,omitempty"`
	Zerorttearlydata                  string   `json:"zerorttearlydata,omitempty"`
}

type Sslzerotouchparam struct {
	Admconnectivitystatus  string `json:"admconnectivitystatus,omitempty"`
	Httpstatuscode         string `json:"httpstatuscode,omitempty"`
	Keyfilename            string `json:"keyfilename,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
	Nextrequesttime        string `json:"nextrequesttime,omitempty"`
	Ocspbatchingdelay      int    `json:"ocspbatchingdelay,omitempty"`
	Ocspbatchingdepth      int    `json:"ocspbatchingdepth,omitempty"`
	Ocspcachetimeout       int    `json:"ocspcachetimeout,omitempty"`
	Ocsphttpmethod         string `json:"ocsphttpmethod,omitempty"`
	Ocspproducedattimeskew int    `json:"ocspproducedattimeskew,omitempty"`
	Ocspresptimeout        int    `json:"ocspresptimeout,omitempty"`
	Ocsptrustresponder     string `json:"ocsptrustresponder,omitempty"`
	Ocspurlresolvetimeout  int    `json:"ocspurlresolvetimeout,omitempty"`
	Ocspusenonce           string `json:"ocspusenonce,omitempty"`
	Passphrase             string `json:"passphrase,omitempty"`
	Remoteserverip         string `json:"remoteserverip,omitempty"`
	Requesttimestamp       string `json:"requesttimestamp,omitempty"`
	Requesttype            string `json:"requesttype,omitempty"`
	Zerotouch              string `json:"zerotouch,omitempty"`
}

type SslpolicyBinding struct {
	Name                           string        `json:"name,omitempty"`
	SslpolicyCsvserverBinding      []interface{} `json:"sslpolicy_csvserver_binding,omitempty"`
	SslpolicyLbvserverBinding      []interface{} `json:"sslpolicy_lbvserver_binding,omitempty"`
	SslpolicySslglobalBinding      []interface{} `json:"sslpolicy_sslglobal_binding,omitempty"`
	SslpolicySslpolicylabelBinding []interface{} `json:"sslpolicy_sslpolicylabel_binding,omitempty"`
	SslpolicySslserviceBinding     []interface{} `json:"sslpolicy_sslservice_binding,omitempty"`
	SslpolicySslvserverBinding     []interface{} `json:"sslpolicy_sslvserver_binding,omitempty"`
}

type Sslpkcs12 struct {
	Aes256        bool   `json:"aes256,omitempty"`
	Certfile      string `json:"certfile,omitempty"`
	Des           bool   `json:"des,omitempty"`
	Des3          bool   `json:"des3,omitempty"`
	Export        bool   `json:"export,omitempty"`
	Import        bool   `json:"Import,omitempty"`
	Keyfile       string `json:"keyfile,omitempty"`
	Outfile       string `json:"outfile,omitempty"`
	Password      string `json:"password,omitempty"`
	Pempassphrase string `json:"pempassphrase,omitempty"`
	Pkcs12file    string `json:"pkcs12file,omitempty"`
}

type SslpolicySslvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type SslcipherSslciphersuiteBinding struct {
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
	Description     string `json:"description,omitempty"`
}

type Sslkeyfile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type SslpolicylabelSslpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Ssldynamicclientcertcache struct {
}

type SslglobalBinding struct {
	SslglobalSslpolicyBinding []interface{} `json:"sslglobal_sslpolicy_binding,omitempty"`
}

type SslserviceSslcacertbundleBinding struct {
	Cacertbundlename string `json:"cacertbundlename,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Skipcacertbundle bool   `json:"skipcacertbundle,omitempty"`
}

type SslcipherIndividualcipherBinding struct {
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
	Description     string `json:"description,omitempty"`
}

type SslvserverSslcertkeybundleBinding struct {
	Certkeybundlename string `json:"certkeybundlename,omitempty"`
	Snicertkeybundle  bool   `json:"snicertkeybundle,omitempty"`
	Vservername       string `json:"vservername,omitempty"`
}

type SslcertkeySslocspresponderBinding struct {
	Ca            bool   `json:"ca,omitempty"`
	Certkey       string `json:"certkey,omitempty"`
	Ocspresponder string `json:"ocspresponder,omitempty"`
	Priority      int    `json:"priority,omitempty"`
}

type Ssldhfile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type SslvserverSslciphersuiteBinding struct {
	Ciphername  string `json:"ciphername,omitempty"`
	Description string `json:"description,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}

type SslvserverHashicorpBinding struct {
	Vault       string `json:"vault,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}

type Sslfips struct {
	Coresenabled        int    `json:"coresenabled,omitempty"`
	Coresmax            int    `json:"coresmax,omitempty"`
	Erasedata           string `json:"erasedata,omitempty"`
	Fipsfw              string `json:"fipsfw,omitempty"`
	Fipshwmajorversion  int    `json:"fipshwmajorversion,omitempty"`
	Fipshwminorversion  int    `json:"fipshwminorversion,omitempty"`
	Fipshwversionstring string `json:"fipshwversionstring,omitempty"`
	Firmwarereleasedate string `json:"firmwarereleasedate,omitempty"`
	Flag                int    `json:"flag,omitempty"`
	Flashmemoryfree     int    `json:"flashmemoryfree,omitempty"`
	Flashmemorytotal    int    `json:"flashmemorytotal,omitempty"`
	Hsmlabel            string `json:"hsmlabel,omitempty"`
	Inithsm             string `json:"inithsm,omitempty"`
	Majorversion        int    `json:"majorversion,omitempty"`
	Minorversion        int    `json:"minorversion,omitempty"`
	Model               string `json:"model,omitempty"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
	Oldsopassword       string `json:"oldsopassword,omitempty"`
	Serial              int    `json:"serial,omitempty"`
	Serialno            string `json:"serialno,omitempty"`
	Sopassword          string `json:"sopassword,omitempty"`
	Sramfree            int    `json:"sramfree,omitempty"`
	Sramtotal           int    `json:"sramtotal,omitempty"`
	State               int    `json:"state,omitempty"`
	Status              int    `json:"status,omitempty"`
	Userpassword        string `json:"userpassword,omitempty"`
}

type Sslaction struct {
	Builtin                []string `json:"builtin,omitempty"`
	Cacertgrpname          string   `json:"cacertgrpname,omitempty"`
	Certfingerprintdigest  string   `json:"certfingerprintdigest,omitempty"`
	Certfingerprintheader  string   `json:"certfingerprintheader,omitempty"`
	Certhashheader         string   `json:"certhashheader,omitempty"`
	Certheader             string   `json:"certheader,omitempty"`
	Certissuerheader       string   `json:"certissuerheader,omitempty"`
	Certnotafterheader     string   `json:"certnotafterheader,omitempty"`
	Certnotbeforeheader    string   `json:"certnotbeforeheader,omitempty"`
	Certserialheader       string   `json:"certserialheader,omitempty"`
	Certsubjectheader      string   `json:"certsubjectheader,omitempty"`
	Cipher                 string   `json:"cipher,omitempty"`
	Cipherheader           string   `json:"cipherheader,omitempty"`
	Clientauth             string   `json:"clientauth,omitempty"`
	Clientcert             string   `json:"clientcert,omitempty"`
	Clientcertfingerprint  string   `json:"clientcertfingerprint,omitempty"`
	Clientcerthash         string   `json:"clientcerthash,omitempty"`
	Clientcertissuer       string   `json:"clientcertissuer,omitempty"`
	Clientcertnotafter     string   `json:"clientcertnotafter,omitempty"`
	Clientcertnotbefore    string   `json:"clientcertnotbefore,omitempty"`
	Clientcertserialnumber string   `json:"clientcertserialnumber,omitempty"`
	Clientcertsubject      string   `json:"clientcertsubject,omitempty"`
	Clientcertverification string   `json:"clientcertverification,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Description            string   `json:"description,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Forward                string   `json:"forward,omitempty"`
	Hits                   int      `json:"hits,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
	Ocspcache              string   `json:"ocspcache,omitempty"`
	Ocspcertvalidation     string   `json:"ocspcertvalidation,omitempty"`
	Ocspstapling           string   `json:"ocspstapling,omitempty"`
	Owasupport             string   `json:"owasupport,omitempty"`
	Referencecount         int      `json:"referencecount,omitempty"`
	Sessionid              string   `json:"sessionid,omitempty"`
	Sessionidheader        string   `json:"sessionidheader,omitempty"`
	Ssllogprofile          string   `json:"ssllogprofile,omitempty"`
	Undefhits              int      `json:"undefhits,omitempty"`
}

type SslservicegroupSslciphersuiteBinding struct {
	Ciphername       string `json:"ciphername,omitempty"`
	Description      string `json:"description,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type Ssllogprofile struct {
	Count                float64 `json:"__count,omitempty"`
	Name                 string  `json:"name,omitempty"`
	Nextgenapiresource   string  `json:"_nextgenapiresource,omitempty"`
	Ssllogclauth         string  `json:"ssllogclauth,omitempty"`
	Ssllogclauthfailures string  `json:"ssllogclauthfailures,omitempty"`
	Sslloghs             string  `json:"sslloghs,omitempty"`
	Sslloghsfailures     string  `json:"sslloghsfailures,omitempty"`
}

type Sslfipssimsource struct {
	Certfile     string `json:"certfile,omitempty"`
	Sourcesecret string `json:"sourcesecret,omitempty"`
	Targetsecret string `json:"targetsecret,omitempty"`
}

type SslservicegroupSslcertkeyBinding struct {
	Ca               bool   `json:"ca,omitempty"`
	Certkeyname      string `json:"certkeyname,omitempty"`
	Crlcheck         string `json:"crlcheck,omitempty"`
	Ocspcheck        string `json:"ocspcheck,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Snicert          bool   `json:"snicert,omitempty"`
}

type Sslcertreq struct {
	Challengepassword    string `json:"challengepassword,omitempty"`
	Commonname           string `json:"commonname,omitempty"`
	Companyname          string `json:"companyname,omitempty"`
	Countryname          string `json:"countryname,omitempty"`
	Digestmethod         string `json:"digestmethod,omitempty"`
	Emailaddress         string `json:"emailaddress,omitempty"`
	Fipskeyname          string `json:"fipskeyname,omitempty"`
	Keyfile              string `json:"keyfile,omitempty"`
	Keyform              string `json:"keyform,omitempty"`
	Localityname         string `json:"localityname,omitempty"`
	Organizationname     string `json:"organizationname,omitempty"`
	Organizationunitname string `json:"organizationunitname,omitempty"`
	Pempassphrase        string `json:"pempassphrase,omitempty"`
	Reqfile              string `json:"reqfile,omitempty"`
	Statename            string `json:"statename,omitempty"`
	Subjectaltname       string `json:"subjectaltname,omitempty"`
}

type SslcacertbundleBinding struct {
	Cacertbundlename                             string        `json:"cacertbundlename,omitempty"`
	SslcacertbundleIntermediatecacertlistBinding []interface{} `json:"sslcacertbundle_intermediatecacertlist_binding,omitempty"`
}

type SslserviceBinding struct {
	Servicename                      string        `json:"servicename,omitempty"`
	SslserviceEcccurveBinding        []interface{} `json:"sslservice_ecccurve_binding,omitempty"`
	SslserviceSslcacertbundleBinding []interface{} `json:"sslservice_sslcacertbundle_binding,omitempty"`
	SslserviceSslcertkeyBinding      []interface{} `json:"sslservice_sslcertkey_binding,omitempty"`
	SslserviceSslcipherBinding       []interface{} `json:"sslservice_sslcipher_binding,omitempty"`
	SslserviceSslciphersuiteBinding  []interface{} `json:"sslservice_sslciphersuite_binding,omitempty"`
	SslserviceSslpolicyBinding       []interface{} `json:"sslservice_sslpolicy_binding,omitempty"`
}

type Sslcipher struct {
	Ciphergroupname    string  `json:"ciphergroupname,omitempty"`
	Ciphername         string  `json:"ciphername,omitempty"`
	Cipherpriority     int     `json:"cipherpriority,omitempty"`
	Ciphgrpalias       string  `json:"ciphgrpalias,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Sslprofile         string  `json:"sslprofile,omitempty"`
}

type Sslcertchain struct {
	Certkeyname        string  `json:"certkeyname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type SslcertkeyBinding struct {
	Certkey                           string        `json:"certkey,omitempty"`
	SslcertkeyCrldistributionBinding  []interface{} `json:"sslcertkey_crldistribution_binding,omitempty"`
	SslcertkeyServiceBinding          []interface{} `json:"sslcertkey_service_binding,omitempty"`
	SslcertkeySslocspresponderBinding []interface{} `json:"sslcertkey_sslocspresponder_binding,omitempty"`
	SslcertkeySslprofileBinding       []interface{} `json:"sslcertkey_sslprofile_binding,omitempty"`
	SslcertkeySslvserverBinding       []interface{} `json:"sslcertkey_sslvserver_binding,omitempty"`
}

type Sslpolicylabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
	Flowtype               int     `json:"flowtype,omitempty"`
	Gotopriorityexpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	Invoke                 bool    `json:"invoke,omitempty"`
	InvokeLabelname        string  `json:"invoke_labelname,omitempty"`
	Labelname              string  `json:"labelname,omitempty"`
	Labeltype              string  `json:"labeltype,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Numpol                 int     `json:"numpol,omitempty"`
	Policyname             string  `json:"policyname,omitempty"`
	Priority               int     `json:"priority,omitempty"`
	TypeField              string  `json:"type,omitempty"`
}

type SslcertkeySslprofileBinding struct {
	Ca         bool   `json:"ca,omitempty"`
	Certkey    string `json:"certkey,omitempty"`
	Sslprofile string `json:"sslprofile,omitempty"`
}

type SslvserverSslpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Polinherit             int    `json:"polinherit,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Vservername            string `json:"vservername,omitempty"`
}

type Sslechconfig struct {
	Count              float64 `json:"__count,omitempty"`
	Echcipher          string  `json:"echcipher,omitempty"`
	Echconfigid        int     `json:"echconfigid,omitempty"`
	Echconfigname      string  `json:"echconfigname,omitempty"`
	Echpublicname      string  `json:"echpublicname,omitempty"`
	Hpkekeyname        string  `json:"hpkekeyname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Version            int     `json:"version,omitempty"`
}

type SslcacertbundleIntermediatecacertlistBinding struct {
	Cacertbundlename    string `json:"cacertbundlename,omitempty"`
	Clientcertnotafter  string `json:"clientcertnotafter,omitempty"`
	Clientcertnotbefore string `json:"clientcertnotbefore,omitempty"`
	Daystoexpiration    int    `json:"daystoexpiration,omitempty"`
	Issuer              string `json:"issuer,omitempty"`
	Publickey           string `json:"publickey,omitempty"`
	Publickeysize       int    `json:"publickeysize,omitempty"`
	Sandns              string `json:"sandns,omitempty"`
	Sanipadd            string `json:"sanipadd,omitempty"`
	Serial              string `json:"serial,omitempty"`
	Signaturealg        string `json:"signaturealg,omitempty"`
	Status              string `json:"status,omitempty"`
	Subject             string `json:"subject,omitempty"`
}

type Sslvserver struct {
	Ca                                bool    `json:"ca,omitempty"`
	Cipherredirect                    string  `json:"cipherredirect,omitempty"`
	Cipherurl                         string  `json:"cipherurl,omitempty"`
	Cleartextport                     int     `json:"cleartextport,omitempty"`
	Clientauth                        string  `json:"clientauth,omitempty"`
	Clientcert                        string  `json:"clientcert,omitempty"`
	Count                             float64 `json:"__count,omitempty"`
	Crlcheck                          string  `json:"crlcheck,omitempty"`
	Defaultsni                        string  `json:"defaultsni,omitempty"`
	Dh                                string  `json:"dh,omitempty"`
	Dhcount                           int     `json:"dhcount,omitempty"`
	Dhekeyexchangewithpsk             string  `json:"dhekeyexchangewithpsk,omitempty"`
	Dhfile                            string  `json:"dhfile,omitempty"`
	Dhkeyexpsizelimit                 string  `json:"dhkeyexpsizelimit,omitempty"`
	Dtls1                             string  `json:"dtls1,omitempty"`
	Dtls12                            string  `json:"dtls12,omitempty"`
	Dtlsflag                          bool    `json:"dtlsflag,omitempty"`
	Dtlsprofilename                   string  `json:"dtlsprofilename,omitempty"`
	Ersa                              string  `json:"ersa,omitempty"`
	Ersacount                         int     `json:"ersacount,omitempty"`
	Hsts                              string  `json:"hsts,omitempty"`
	Includesubdomains                 string  `json:"includesubdomains,omitempty"`
	Maxage                            int     `json:"maxage,omitempty"`
	Nextgenapiresource                string  `json:"_nextgenapiresource,omitempty"`
	Nonfipsciphers                    string  `json:"nonfipsciphers,omitempty"`
	Ocspcheck                         string  `json:"ocspcheck,omitempty"`
	Ocspstapling                      string  `json:"ocspstapling,omitempty"`
	Preload                           string  `json:"preload,omitempty"`
	Pushenctrigger                    string  `json:"pushenctrigger,omitempty"`
	Quicflag                          bool    `json:"quicflag,omitempty"`
	Redirectportrewrite               string  `json:"redirectportrewrite,omitempty"`
	Sendclosenotify                   string  `json:"sendclosenotify,omitempty"`
	Service                           int     `json:"service,omitempty"`
	Sessreuse                         string  `json:"sessreuse,omitempty"`
	Sesstimeout                       int     `json:"sesstimeout,omitempty"`
	Skipcacertbundle                  bool    `json:"skipcacertbundle,omitempty"`
	Skipcaname                        bool    `json:"skipcaname,omitempty"`
	Snicert                           bool    `json:"snicert,omitempty"`
	Snienable                         string  `json:"snienable,omitempty"`
	Ssl2                              string  `json:"ssl2,omitempty"`
	Ssl3                              string  `json:"ssl3,omitempty"`
	Sslclientlogs                     string  `json:"sslclientlogs,omitempty"`
	Sslprofile                        string  `json:"sslprofile,omitempty"`
	Sslredirect                       string  `json:"sslredirect,omitempty"`
	Sslv2redirect                     string  `json:"sslv2redirect,omitempty"`
	Sslv2url                          string  `json:"sslv2url,omitempty"`
	Strictsigdigestcheck              string  `json:"strictsigdigestcheck,omitempty"`
	Tls1                              string  `json:"tls1,omitempty"`
	Tls11                             string  `json:"tls11,omitempty"`
	Tls12                             string  `json:"tls12,omitempty"`
	Tls13                             string  `json:"tls13,omitempty"`
	Tls13sessionticketsperauthcontext int     `json:"tls13sessionticketsperauthcontext,omitempty"`
	Vservername                       string  `json:"vservername,omitempty"`
	Zerorttearlydata                  string  `json:"zerorttearlydata,omitempty"`
}
