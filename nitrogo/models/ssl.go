// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Sslcacertbundle struct {
	Cacertbundlename   string `json:"cacertbundlename,omitempty"`
	Bundlefile         string `json:"bundlefile,omitempty"`
	Servername         string `json:"servername,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslcertificatechain struct {
	Certkeyname        string `json:"certkeyname,omitempty"`
	Chainlinked        string `json:"chainlinked,omitempty"`
	Chainpossiblelinks string `json:"chainpossiblelinks,omitempty"`
	Chainissuer        string `json:"chainissuer,omitempty"`
	Chaincomplete      string `json:"chaincomplete,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslcertkeyocspresponderbinding struct {
	Ocspresponder string `json:"ocspresponder,omitempty"`
	Priority      uint32 `json:"priority,omitempty"`
	Certkey       string `json:"certkey,omitempty"`
	Ca            bool   `json:"ca,omitempty"`
}

type Sslcipher struct {
	Ciphergroupname    string `json:"ciphergroupname,omitempty"`
	Ciphgrpalias       string `json:"ciphgrpalias,omitempty"`
	Ciphername         string `json:"ciphername,omitempty"`
	Cipherpriority     int    `json:"cipherpriority,omitempty"`
	Sslprofile         string `json:"sslprofile,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslpolicylbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Sslpolicyservicebinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Sslcertkeysslprofilebinding struct {
	Sslprofile string `json:"sslprofile,omitempty"`
	Certkey    string `json:"certkey,omitempty"`
	Ca         bool   `json:"ca,omitempty"`
}

type Sslechconfig struct {
	Echconfigname      string `json:"echconfigname,omitempty"`
	Echcipher          string `json:"echcipher,omitempty"`
	Hpkekeyname        string `json:"hpkekeyname,omitempty"`
	Echpublicname      string `json:"echpublicname,omitempty"`
	Echconfigid        int    `json:"echconfigid,omitempty"`
	Version            int    `json:"version,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslocspresponder struct {
	Name                  string `json:"name,omitempty"`
	Url                   string `json:"url,omitempty"`
	Cache                 string `json:"cache,omitempty"`
	Cachetimeout          int    `json:"cachetimeout,omitempty"`
	Batchingdepth         int    `json:"batchingdepth,omitempty"`
	Batchingdelay         int    `json:"batchingdelay,omitempty"`
	Resptimeout           int    `json:"resptimeout,omitempty"`
	Ocspurlresolvetimeout int    `json:"ocspurlresolvetimeout,omitempty"`
	Respondercert         string `json:"respondercert,omitempty"`
	Trustresponder        bool   `json:"trustresponder,omitempty"`
	Producedattimeskew    int    `json:"producedattimeskew,omitempty"`
	Signingcert           string `json:"signingcert,omitempty"`
	Usenonce              string `json:"usenonce,omitempty"`
	Insertclientcert      string `json:"insertclientcert,omitempty"`
	Httpmethod            string `json:"httpmethod,omitempty"`
	Ocspaiarefcount       string `json:"ocspaiarefcount,omitempty"`
	Ocspipaddrstr         string `json:"ocspipaddrstr,omitempty"`
	Port                  string `json:"port,omitempty"`
	Nextgenapiresource    string `json:"_nextgenapiresource,omitempty"`
}

type Sslprofilesslcertkeybinding struct {
	Sslicacertkey  string `json:"sslicacertkey,omitempty"`
	Name           string `json:"name,omitempty"`
	Cipherpriority int    `json:"cipherpriority,omitempty"`
}

type Sslprofilevserverbinding struct {
	Servicename    string `json:"servicename,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	Cipherpriority uint32 `json:"cipherpriority,omitempty"`
}

type Sslrsakey struct {
	Keyfile  string `json:"keyfile,omitempty"`
	Bits     int    `json:"bits,omitempty"`
	Exponent string `json:"exponent,omitempty"`
	Keyform  string `json:"keyform,omitempty"`
	Des      bool   `json:"des,omitempty"`
	Des3     bool   `json:"des3,omitempty"`
	Aes256   bool   `json:"aes256,omitempty"`
	Password string `json:"password,omitempty"`
	Pkcs8    bool   `json:"pkcs8,omitempty"`
}

type Sslservice struct {
	Servicename          string `json:"servicename,omitempty"`
	Dh                   string `json:"dh,omitempty"`
	Dhfile               string `json:"dhfile,omitempty"`
	Dhcount              int    `json:"dhcount,omitempty"`
	Dhkeyexpsizelimit    string `json:"dhkeyexpsizelimit,omitempty"`
	Ersa                 string `json:"ersa,omitempty"`
	Ersacount            int    `json:"ersacount,omitempty"`
	Sessreuse            string `json:"sessreuse,omitempty"`
	Sesstimeout          int    `json:"sesstimeout,omitempty"`
	Cipherredirect       string `json:"cipherredirect,omitempty"`
	Cipherurl            string `json:"cipherurl,omitempty"`
	Sslv2redirect        string `json:"sslv2redirect,omitempty"`
	Sslv2url             string `json:"sslv2url,omitempty"`
	Clientauth           string `json:"clientauth,omitempty"`
	Clientcert           string `json:"clientcert,omitempty"`
	Sslredirect          string `json:"sslredirect,omitempty"`
	Redirectportrewrite  string `json:"redirectportrewrite,omitempty"`
	Ssl2                 string `json:"ssl2,omitempty"`
	Ssl3                 string `json:"ssl3,omitempty"`
	Tls1                 string `json:"tls1,omitempty"`
	Tls11                string `json:"tls11,omitempty"`
	Tls12                string `json:"tls12,omitempty"`
	Tls13                string `json:"tls13,omitempty"`
	Dtls1                string `json:"dtls1,omitempty"`
	Dtls12               string `json:"dtls12,omitempty"`
	Snienable            string `json:"snienable,omitempty"`
	Ocspstapling         string `json:"ocspstapling,omitempty"`
	Serverauth           string `json:"serverauth,omitempty"`
	Commonname           string `json:"commonname,omitempty"`
	Pushenctrigger       string `json:"pushenctrigger,omitempty"`
	Sendclosenotify      string `json:"sendclosenotify,omitempty"`
	Dtlsprofilename      string `json:"dtlsprofilename,omitempty"`
	Sslprofile           string `json:"sslprofile,omitempty"`
	Strictsigdigestcheck string `json:"strictsigdigestcheck,omitempty"`
	Sslclientlogs        string `json:"sslclientlogs,omitempty"`
	Nonfipsciphers       string `json:"nonfipsciphers,omitempty"`
	Service              string `json:"service,omitempty"`
	Skipcaname           string `json:"skipcaname,omitempty"`
	Dtlsflag             string `json:"dtlsflag,omitempty"`
	Quicflag             string `json:"quicflag,omitempty"`
	Skipcacertbundle     string `json:"skipcacertbundle,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Sslciphersslciphersuitebinding struct {
	Ciphername      string `json:"ciphername,omitempty"`
	Description     string `json:"description,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
}

type Ssldhfile struct {
	Name               string `json:"name,omitempty"`
	Src                string `json:"src,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Ssldhparam struct {
	Dhfile string `json:"dhfile,omitempty"`
	Bits   int    `json:"bits,omitempty"`
	Gen    string `json:"gen,omitempty"`
}

type Sslpolicylabel struct {
	Labelname              string `json:"labelname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 string `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Flowtype               string `json:"flowtype,omitempty"`
	Description            string `json:"description,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Sslprofilecertkeybinding struct {
	Sslicacertkey  string `json:"sslicacertkey,omitempty"`
	Name           string `json:"name,omitempty"`
	Cipherpriority uint32 `json:"cipherpriority,omitempty"`
}

type Sslvserversslcipherbinding struct {
	Cipheraliasname string `json:"cipheraliasname,omitempty"`
	Description     string `json:"description,omitempty"`
	Vservername     string `json:"vservername,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
}

type Sslcrlbinding struct {
	Crlname string `json:"crlname,omitempty"`
}

type Sslfipssimtarget struct {
	Keyvector    string `json:"keyvector,omitempty"`
	Sourcesecret string `json:"sourcesecret,omitempty"`
	Certfile     string `json:"certfile,omitempty"`
	Targetsecret string `json:"targetsecret,omitempty"`
}

type Sslglobalpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Sslpkcs12 struct {
	Outfile       string `json:"outfile,omitempty"`
	Import        bool   `json:"Import,omitempty"`
	Pkcs12file    string `json:"pkcs12file,omitempty"`
	Des           bool   `json:"des,omitempty"`
	Des3          bool   `json:"des3,omitempty"`
	Aes256        bool   `json:"aes256,omitempty"`
	Export        bool   `json:"export,omitempty"`
	Certfile      string `json:"certfile,omitempty"`
	Keyfile       string `json:"keyfile,omitempty"`
	Password      string `json:"password,omitempty"`
	Pempassphrase string `json:"pempassphrase,omitempty"`
}

type Sslcertkeysslocspresponderbinding struct {
	Ocspresponder string `json:"ocspresponder,omitempty"`
	Priority      int    `json:"priority,omitempty"`
	Certkey       string `json:"certkey,omitempty"`
	Ca            bool   `json:"ca,omitempty"`
}

type Sslcertkeybundlesslvserverbinding struct {
	Servername        string `json:"servername,omitempty"`
	Certkeybundlename string `json:"certkeybundlename,omitempty"`
}

type Sslprofileecccurvebinding struct {
	Ecccurvename   string `json:"ecccurvename,omitempty"`
	Name           string `json:"name,omitempty"`
	Cipherpriority int    `json:"cipherpriority,omitempty"`
}

type Sslvserverhashicorpbinding struct {
	Vault       string `json:"vault,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}

type Sslcert struct {
	Certfile       string `json:"certfile,omitempty"`
	Reqfile        string `json:"reqfile,omitempty"`
	Certtype       string `json:"certtype,omitempty"`
	Keyfile        string `json:"keyfile,omitempty"`
	Keyform        string `json:"keyform,omitempty"`
	Pempassphrase  string `json:"pempassphrase,omitempty"`
	Days           int    `json:"days,omitempty"`
	Subjectaltname string `json:"subjectaltname,omitempty"`
	Certform       string `json:"certform,omitempty"`
	Cacert         string `json:"cacert,omitempty"`
	Cacertform     string `json:"cacertform,omitempty"`
	Cakey          string `json:"cakey,omitempty"`
	Cakeyform      string `json:"cakeyform,omitempty"`
	Caserial       string `json:"caserial,omitempty"`
}

type Sslcertkeysslvserverbinding struct {
	Servername  string `json:"servername,omitempty"`
	Data        int    `json:"data,omitempty"`
	Version     int    `json:"version,omitempty"`
	Certkey     string `json:"certkey,omitempty"`
	Vservername string `json:"vservername,omitempty"`
	Vserver     bool   `json:"vserver,omitempty"`
	Ca          bool   `json:"ca,omitempty"`
}

type Sslcipherservicebinding struct {
	Ciphergroupname  string `json:"ciphergroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Service          bool   `json:"service,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicegroup     bool   `json:"servicegroup,omitempty"`
	Cipheroperation  string `json:"cipheroperation,omitempty"`
	Ciphgrpals       string `json:"ciphgrpals,omitempty"`
	Cipherpriority   int    `json:"cipherpriority,omitempty"`
}

type Ssldtlsprofile struct {
	Name                 string `json:"name,omitempty"`
	Pmtudiscovery        string `json:"pmtudiscovery,omitempty"`
	Maxrecordsize        int    `json:"maxrecordsize,omitempty"`
	Maxretrytime         int    `json:"maxretrytime,omitempty"`
	Helloverifyrequest   string `json:"helloverifyrequest,omitempty"`
	Terminatesession     string `json:"terminatesession,omitempty"`
	Maxpacketsize        int    `json:"maxpacketsize,omitempty"`
	Maxholdqlen          int    `json:"maxholdqlen,omitempty"`
	Maxbadmacignorecount int    `json:"maxbadmacignorecount,omitempty"`
	Initialretrytimeout  int    `json:"initialretrytimeout,omitempty"`
	Builtin              string `json:"builtin,omitempty"`
	Feature              string `json:"feature,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Sslfipssimsource struct {
	Targetsecret string `json:"targetsecret,omitempty"`
	Sourcesecret string `json:"sourcesecret,omitempty"`
	Certfile     string `json:"certfile,omitempty"`
}

type Sslhpkekey struct {
	Hpkekeyname        string `json:"hpkekeyname,omitempty"`
	File               string `json:"file,omitempty"`
	Dhkem              string `json:"dhkem,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslpolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}

type Sslvservercertkeybinding struct {
	Certkeyname   string `json:"certkeyname,omitempty"`
	Crlcheck      string `json:"crlcheck,omitempty"`
	Ocspcheck     string `json:"ocspcheck,omitempty"`
	Cleartextport int32  `json:"cleartextport,omitempty"`
	Ca            bool   `json:"ca,omitempty"`
	Snicert       bool   `json:"snicert,omitempty"`
	Skipcaname    bool   `json:"skipcaname,omitempty"`
	Vservername   string `json:"vservername,omitempty"`
}

type Sslcertchainsslcertkeybinding struct {
	Linkcertkeyname string `json:"linkcertkeyname,omitempty"`
	Islinked        bool   `json:"islinked,omitempty"`
	Isca            bool   `json:"isca,omitempty"`
	Addsubject      bool   `json:"addsubject,omitempty"`
	Certkeyname     string `json:"certkeyname,omitempty"`
}

type Sslcacertbundlebinding struct {
	Cacertbundlename string `json:"cacertbundlename,omitempty"`
}

type Sslcrl struct {
	Crlname            string `json:"crlname,omitempty"`
	Crlpath            string `json:"crlpath,omitempty"`
	Inform             string `json:"inform,omitempty"`
	Refresh            string `json:"refresh,omitempty"`
	Cacert             string `json:"cacert,omitempty"`
	Method             string `json:"method,omitempty"`
	Server             string `json:"server,omitempty"`
	Url                string `json:"url,omitempty"`
	Port               int    `json:"port,omitempty"`
	Basedn             string `json:"basedn,omitempty"`
	Scope              string `json:"scope,omitempty"`
	Interval           string `json:"interval,omitempty"`
	Day                int    `json:"day,omitempty"`
	Time               string `json:"time,omitempty"`
	Binddn             string `json:"binddn,omitempty"`
	Password           string `json:"password,omitempty"`
	Binary             string `json:"binary,omitempty"`
	Cacertfile         string `json:"cacertfile,omitempty"`
	Cakeyfile          string `json:"cakeyfile,omitempty"`
	Indexfile          string `json:"indexfile,omitempty"`
	Revoke             string `json:"revoke,omitempty"`
	Gencrl             string `json:"gencrl,omitempty"`
	Flags              string `json:"flags,omitempty"`
	Lastupdatetime     string `json:"lastupdatetime,omitempty"`
	Version            string `json:"version,omitempty"`
	Signaturealgo      string `json:"signaturealgo,omitempty"`
	Issuer             string `json:"issuer,omitempty"`
	Lastupdate         string `json:"lastupdate,omitempty"`
	Nextupdate         string `json:"nextupdate,omitempty"`
	Daystoexpiration   string `json:"daystoexpiration,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslpolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Sslvserversslcertkeybinding struct {
	Certkeyname   string `json:"certkeyname,omitempty"`
	Crlcheck      string `json:"crlcheck,omitempty"`
	Ocspcheck     string `json:"ocspcheck,omitempty"`
	Cleartextport int    `json:"cleartextport,omitempty"`
	Ca            bool   `json:"ca,omitempty"`
	Snicert       bool   `json:"snicert,omitempty"`
	Skipcaname    bool   `json:"skipcaname,omitempty"`
	Vservername   string `json:"vservername,omitempty"`
}

type Sslcertfile struct {
	Name               string `json:"name,omitempty"`
	Src                string `json:"src,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslprofilesslechconfigbinding struct {
	Echconfigname  string `json:"echconfigname,omitempty"`
	Name           string `json:"name,omitempty"`
	Cipherpriority int    `json:"cipherpriority,omitempty"`
}

type Sslserviceecccurvebinding struct {
	Ecccurvename string `json:"ecccurvename,omitempty"`
	Servicename  string `json:"servicename,omitempty"`
}

type Sslcertkeyprofilebinding struct {
	Sslprofile string `json:"sslprofile,omitempty"`
	Certkey    string `json:"certkey,omitempty"`
	Ca         bool   `json:"ca,omitempty"`
}

type Sslcertchaincertkeybinding struct {
	Linkcertkeyname string `json:"linkcertkeyname,omitempty"`
	Islinked        bool   `json:"islinked,omitempty"`
	Isca            bool   `json:"isca,omitempty"`
	Addsubject      bool   `json:"addsubject,omitempty"`
	Certkeyname     string `json:"certkeyname,omitempty"`
}

type Sslcertkeycrldistributionbinding struct {
	Issuer  string `json:"issuer,omitempty"`
	Certkey string `json:"certkey,omitempty"`
	Ca      bool   `json:"ca,omitempty"`
}

type Sslhsmkey struct {
	Hsmkeyname         string `json:"hsmkeyname,omitempty"`
	Hsmtype            string `json:"hsmtype,omitempty"`
	Key                string `json:"key,omitempty"`
	Serialnum          string `json:"serialnum,omitempty"`
	Password           string `json:"password,omitempty"`
	Keystore           string `json:"keystore,omitempty"`
	State              string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslpolicysslservicebinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Sslservicegroup struct {
	Servicegroupname     string `json:"servicegroupname,omitempty"`
	Sslprofile           string `json:"sslprofile,omitempty"`
	Sessreuse            string `json:"sessreuse,omitempty"`
	Sesstimeout          int    `json:"sesstimeout,omitempty"`
	Ssl3                 string `json:"ssl3,omitempty"`
	Tls1                 string `json:"tls1,omitempty"`
	Tls11                string `json:"tls11,omitempty"`
	Tls12                string `json:"tls12,omitempty"`
	Tls13                string `json:"tls13,omitempty"`
	Snienable            string `json:"snienable,omitempty"`
	Ocspstapling         string `json:"ocspstapling,omitempty"`
	Serverauth           string `json:"serverauth,omitempty"`
	Commonname           string `json:"commonname,omitempty"`
	Sendclosenotify      string `json:"sendclosenotify,omitempty"`
	Strictsigdigestcheck string `json:"strictsigdigestcheck,omitempty"`
	Sslclientlogs        string `json:"sslclientlogs,omitempty"`
	Dh                   string `json:"dh,omitempty"`
	Dhfile               string `json:"dhfile,omitempty"`
	Dhcount              string `json:"dhcount,omitempty"`
	Dhkeyexpsizelimit    string `json:"dhkeyexpsizelimit,omitempty"`
	Ersa                 string `json:"ersa,omitempty"`
	Ersacount            string `json:"ersacount,omitempty"`
	Cipherredirect       string `json:"cipherredirect,omitempty"`
	Cipherurl            string `json:"cipherurl,omitempty"`
	Sslv2redirect        string `json:"sslv2redirect,omitempty"`
	Sslv2url             string `json:"sslv2url,omitempty"`
	Clientauth           string `json:"clientauth,omitempty"`
	Clientcert           string `json:"clientcert,omitempty"`
	Sslredirect          string `json:"sslredirect,omitempty"`
	Redirectportrewrite  string `json:"redirectportrewrite,omitempty"`
	Nonfipsciphers       string `json:"nonfipsciphers,omitempty"`
	Ssl2                 string `json:"ssl2,omitempty"`
	Ocspcheck            string `json:"ocspcheck,omitempty"`
	Crlcheck             string `json:"crlcheck,omitempty"`
	Cleartextport        string `json:"cleartextport,omitempty"`
	Servicename          string `json:"servicename,omitempty"`
	Ca                   string `json:"ca,omitempty"`
	Snicert              string `json:"snicert,omitempty"`
	Quicflag             string `json:"quicflag,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Sslcipherbinding struct {
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
}

type Sslpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Sslpolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Sslservicecertkeybinding struct {
	Certkeyname   string `json:"certkeyname,omitempty"`
	Cleartextport int32  `json:"cleartextport,omitempty"`
	Crlcheck      string `json:"crlcheck,omitempty"`
	Ocspcheck     string `json:"ocspcheck,omitempty"`
	Ca            bool   `json:"ca,omitempty"`
	Snicert       bool   `json:"snicert,omitempty"`
	Skipcaname    bool   `json:"skipcaname,omitempty"`
	Servicename   string `json:"servicename,omitempty"`
}

type Sslservicesslpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Polinherit             int    `json:"polinherit,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Servicename            string `json:"servicename,omitempty"`
}

type Sslservicegroupsslcacertbundlebinding struct {
	Cacertbundlename string `json:"cacertbundlename,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type Sslcacertgroup struct {
	Cacertgroupname       string `json:"cacertgroupname,omitempty"`
	Cacertgroupreferences string `json:"cacertgroupreferences,omitempty"`
	Ocspcheck             string `json:"ocspcheck,omitempty"`
	Crlcheck              string `json:"crlcheck,omitempty"`
	Nextgenapiresource    string `json:"_nextgenapiresource,omitempty"`
}

type Sslprofileciphersuitebinding struct {
	Ciphername     string `json:"ciphername,omitempty"`
	Cipherpriority uint32 `json:"cipherpriority,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Sslservicegroupecccurvebinding struct {
	Ecccurvename     string `json:"ecccurvename,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type Sslcertchain struct {
	Certkeyname        string `json:"certkeyname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslparameter struct {
	Quantumsize              string   `json:"quantumsize,omitempty"`
	Crlmemorysizemb          int      `json:"crlmemorysizemb,omitempty"`
	Strictcachecks           string   `json:"strictcachecks,omitempty"`
	Ssltriggertimeout        int      `json:"ssltriggertimeout,omitempty"`
	Sendclosenotify          string   `json:"sendclosenotify,omitempty"`
	Encrypttriggerpktcount   int      `json:"encrypttriggerpktcount,omitempty"`
	Denysslreneg             string   `json:"denysslreneg,omitempty"`
	Insertionencoding        string   `json:"insertionencoding,omitempty"`
	Ocspcachesize            int      `json:"ocspcachesize,omitempty"`
	Pushflag                 int      `json:"pushflag,omitempty"`
	Dropreqwithnohostheader  string   `json:"dropreqwithnohostheader,omitempty"`
	Snihttphostmatch         string   `json:"snihttphostmatch,omitempty"`
	Pushenctriggertimeout    int      `json:"pushenctriggertimeout,omitempty"`
	Cryptodevdisablelimit    int      `json:"cryptodevdisablelimit,omitempty"`
	Undefactioncontrol       string   `json:"undefactioncontrol,omitempty"`
	Undefactiondata          string   `json:"undefactiondata,omitempty"`
	Defaultprofile           string   `json:"defaultprofile,omitempty"`
	Softwarecryptothreshold  int      `json:"softwarecryptothreshold,omitempty"`
	Hybridfipsmode           string   `json:"hybridfipsmode,omitempty"`
	Sigdigesttype            []string `json:"sigdigesttype,omitempty"`
	Sslierrorcache           string   `json:"sslierrorcache,omitempty"`
	Sslimaxerrorcachemem     int      `json:"sslimaxerrorcachemem,omitempty"`
	Insertcertspace          string   `json:"insertcertspace,omitempty"`
	Ndcppcompliancecertcheck string   `json:"ndcppcompliancecertcheck,omitempty"`
	Heterogeneoussslhw       string   `json:"heterogeneoussslhw,omitempty"`
	Operationqueuelimit      int      `json:"operationqueuelimit,omitempty"`
	Svctls1112disable        string   `json:"svctls1112disable,omitempty"`
	Montls1112disable        string   `json:"montls1112disable,omitempty"`
	Nextgenapiresource       string   `json:"_nextgenapiresource,omitempty"`
}

type Sslvserverbinding struct {
	Vservername string `json:"vservername,omitempty"`
}

type Sslfipskey struct {
	Fipskeyname        string `json:"fipskeyname,omitempty"`
	Keytype            string `json:"keytype,omitempty"`
	Exponent           string `json:"exponent,omitempty"`
	Modulus            int    `json:"modulus,omitempty"`
	Curve              string `json:"curve,omitempty"`
	Key                string `json:"key,omitempty"`
	Inform             string `json:"inform,omitempty"`
	Wrapkeyname        string `json:"wrapkeyname,omitempty"`
	Iv                 string `json:"iv,omitempty"`
	Size               string `json:"size,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslcertchainbinding struct {
	Certkeyname string `json:"certkeyname,omitempty"`
}

type Sslcertkeybinding struct {
	Certkey string `json:"certkey,omitempty"`
}

type Sslpolicypolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Sslvserversslcacertbundlebinding struct {
	Cacertbundlename string `json:"cacertbundlename,omitempty"`
	Skipcacertbundle bool   `json:"skipcacertbundle,omitempty"`
	Vservername      string `json:"vservername,omitempty"`
}

type Sslciphersslprofilebinding struct {
	Sslprofile      string `json:"sslprofile,omitempty"`
	Description     string `json:"description,omitempty"`
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
}

type Sslprofilecipherbinding struct {
	Cipheraliasname string `json:"cipheraliasname,omitempty"`
	Cipherpriority  uint32 `json:"cipherpriority,omitempty"`
	Description     string `json:"description,omitempty"`
	Name            string `json:"name,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
}

type Sslcertreq struct {
	Reqfile              string `json:"reqfile,omitempty"`
	Keyfile              string `json:"keyfile,omitempty"`
	Subjectaltname       string `json:"subjectaltname,omitempty"`
	Fipskeyname          string `json:"fipskeyname,omitempty"`
	Keyform              string `json:"keyform,omitempty"`
	Pempassphrase        string `json:"pempassphrase,omitempty"`
	Countryname          string `json:"countryname,omitempty"`
	Statename            string `json:"statename,omitempty"`
	Organizationname     string `json:"organizationname,omitempty"`
	Organizationunitname string `json:"organizationunitname,omitempty"`
	Localityname         string `json:"localityname,omitempty"`
	Commonname           string `json:"commonname,omitempty"`
	Emailaddress         string `json:"emailaddress,omitempty"`
	Challengepassword    string `json:"challengepassword,omitempty"`
	Companyname          string `json:"companyname,omitempty"`
	Digestmethod         string `json:"digestmethod,omitempty"`
}

type Sslpolicysslvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Sslcertlink struct {
	Certkeyname        string `json:"certkeyname,omitempty"`
	Linkcertkeyname    string `json:"linkcertkeyname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslcipherprofilebinding struct {
	Sslprofile      string `json:"sslprofile,omitempty"`
	Description     string `json:"description,omitempty"`
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
	Cipherpriority  uint32 `json:"cipherpriority,omitempty"`
}

type Sslcertkeybundleintermediatecertlinksbinding struct {
	Subject             string `json:"subject,omitempty"`
	Serial              string `json:"serial,omitempty"`
	Issuer              string `json:"issuer,omitempty"`
	Publickey           string `json:"publickey,omitempty"`
	Publickeysize       int    `json:"publickeysize,omitempty"`
	Sandns              string `json:"sandns,omitempty"`
	Sanipadd            string `json:"sanipadd,omitempty"`
	Clientcertnotbefore string `json:"clientcertnotbefore,omitempty"`
	Clientcertnotafter  string `json:"clientcertnotafter,omitempty"`
	Daystoexpiration    int    `json:"daystoexpiration,omitempty"`
	Signaturealg        string `json:"signaturealg,omitempty"`
	Status              string `json:"status,omitempty"`
	Certkeybundlename   string `json:"certkeybundlename,omitempty"`
}

type Sslcipherindividualcipherbinding struct {
	Ciphername      string `json:"ciphername,omitempty"`
	Description     string `json:"description,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
}

type Sslglobalsslpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Ssllogprofile struct {
	Name                 string `json:"name,omitempty"`
	Ssllogclauth         string `json:"ssllogclauth,omitempty"`
	Ssllogclauthfailures string `json:"ssllogclauthfailures,omitempty"`
	Sslloghs             string `json:"sslloghs,omitempty"`
	Sslloghsfailures     string `json:"sslloghsfailures,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Sslservicesslcacertbundlebinding struct {
	Cacertbundlename string `json:"cacertbundlename,omitempty"`
	Skipcacertbundle bool   `json:"skipcacertbundle,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
}

type Sslservicegroupcipherbinding struct {
	Cipheraliasname  string `json:"cipheraliasname,omitempty"`
	Description      string `json:"description,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Ciphername       string `json:"ciphername,omitempty"`
}

type Sslciphersuite struct {
	Ciphername         string `json:"ciphername,omitempty"`
	Description        string `json:"description,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslprofilebinding struct {
	Name string `json:"name,omitempty"`
}

type Sslservicegroupsslcertkeybinding struct {
	Certkeyname      string `json:"certkeyname,omitempty"`
	Crlcheck         string `json:"crlcheck,omitempty"`
	Ocspcheck        string `json:"ocspcheck,omitempty"`
	Ca               bool   `json:"ca,omitempty"`
	Snicert          bool   `json:"snicert,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type Sslcacertgroupcertkeybinding struct {
	Certkeyname     string `json:"certkeyname,omitempty"`
	Crlcheck        string `json:"crlcheck,omitempty"`
	Ocspcheck       string `json:"ocspcheck,omitempty"`
	Cacertgroupname string `json:"cacertgroupname,omitempty"`
}

type Sslcacertgroupsslcertkeybinding struct {
	Certkeyname     string `json:"certkeyname,omitempty"`
	Crlcheck        string `json:"crlcheck,omitempty"`
	Ocspcheck       string `json:"ocspcheck,omitempty"`
	Cacertgroupname string `json:"cacertgroupname,omitempty"`
}

type Sslcertkeybundle struct {
	Certkeybundlename  string `json:"certkeybundlename,omitempty"`
	Bundlefile         string `json:"bundlefile,omitempty"`
	Passplain          string `json:"passplain,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslservicesslciphersuitebinding struct {
	Ciphername      string `json:"ciphername,omitempty"`
	Description     string `json:"description,omitempty"`
	Cipherdefaulton int    `json:"cipherdefaulton,omitempty"`
	Servicename     string `json:"servicename,omitempty"`
}

type Sslvserverecccurvebinding struct {
	Ecccurvename string `json:"ecccurvename,omitempty"`
	Vservername  string `json:"vservername,omitempty"`
}

type Sslvserversslpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Type                   string `json:"type,omitempty"`
	Polinherit             int    `json:"polinherit,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Vservername            string `json:"vservername,omitempty"`
}

type Sslcertkeyvserverbinding struct {
	Servername  string `json:"servername,omitempty"`
	Data        uint32 `json:"data,omitempty"`
	Version     int32  `json:"version,omitempty"`
	Certkey     string `json:"certkey,omitempty"`
	Vservername string `json:"vservername,omitempty"`
	Vserver     bool   `json:"vserver,omitempty"`
	Ca          bool   `json:"ca,omitempty"`
}

type Sslpolicysslpolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Sslcipherciphersuitebinding struct {
	Ciphername      string `json:"ciphername,omitempty"`
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Description     string `json:"description,omitempty"`
	Cipherpriority  uint32 `json:"cipherpriority,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
}

type Ssldefaultprofile struct {
}

type Sslservicegroupsslcipherbinding struct {
	Cipheraliasname  string `json:"cipheraliasname,omitempty"`
	Description      string `json:"description,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Ciphername       string `json:"ciphername,omitempty"`
}

type Sslvserversslciphersuitebinding struct {
	Ciphername  string `json:"ciphername,omitempty"`
	Description string `json:"description,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}

type Sslcrlfile struct {
	Name               string `json:"name,omitempty"`
	Src                string `json:"src,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslservicegroupbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type Sslpolicylabelsslpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
}

type Sslprofilesslciphersuitebinding struct {
	Ciphername     string `json:"ciphername,omitempty"`
	Cipherpriority int    `json:"cipherpriority,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Sslservicesslcertkeybinding struct {
	Certkeyname   string `json:"certkeyname,omitempty"`
	Cleartextport int    `json:"cleartextport,omitempty"`
	Crlcheck      string `json:"crlcheck,omitempty"`
	Ocspcheck     string `json:"ocspcheck,omitempty"`
	Ca            bool   `json:"ca,omitempty"`
	Snicert       bool   `json:"snicert,omitempty"`
	Skipcaname    bool   `json:"skipcaname,omitempty"`
	Servicename   string `json:"servicename,omitempty"`
}

type Sslvserversslcertkeybundlebinding struct {
	Certkeybundlename string `json:"certkeybundlename,omitempty"`
	Snicertkeybundle  bool   `json:"snicertkeybundle,omitempty"`
	Vservername       string `json:"vservername,omitempty"`
}

type Sslwrapkey struct {
	Wrapkeyname        string `json:"wrapkeyname,omitempty"`
	Password           string `json:"password,omitempty"`
	Salt               string `json:"salt,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslcertkeybundlebinding struct {
	Certkeybundlename string `json:"certkeybundlename,omitempty"`
}

type Sslpkcs8 struct {
	Pkcs8file string `json:"pkcs8file,omitempty"`
	Keyfile   string `json:"keyfile,omitempty"`
	Keyform   string `json:"keyform,omitempty"`
	Password  string `json:"password,omitempty"`
}

type Sslpolicysslglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Sslserviceciphersuitebinding struct {
	Ciphername  string `json:"ciphername,omitempty"`
	Description string `json:"description,omitempty"`
	Servicename string `json:"servicename,omitempty"`
}

type Sslcacertbundleintermediatecacertlistbinding struct {
	Subject             string `json:"subject,omitempty"`
	Serial              string `json:"serial,omitempty"`
	Issuer              string `json:"issuer,omitempty"`
	Publickey           string `json:"publickey,omitempty"`
	Publickeysize       int    `json:"publickeysize,omitempty"`
	Sandns              string `json:"sandns,omitempty"`
	Sanipadd            string `json:"sanipadd,omitempty"`
	Clientcertnotbefore string `json:"clientcertnotbefore,omitempty"`
	Clientcertnotafter  string `json:"clientcertnotafter,omitempty"`
	Daystoexpiration    int    `json:"daystoexpiration,omitempty"`
	Signaturealg        string `json:"signaturealg,omitempty"`
	Status              string `json:"status,omitempty"`
	Cacertbundlename    string `json:"cacertbundlename,omitempty"`
}

type Sslcipherservicegroupbinding struct {
	Ciphergroupname  string `json:"ciphergroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Service          bool   `json:"service,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicegroup     bool   `json:"servicegroup,omitempty"`
	Cipheroperation  string `json:"cipheroperation,omitempty"`
	Ciphgrpals       string `json:"ciphgrpals,omitempty"`
	Cipherpriority   int    `json:"cipherpriority,omitempty"`
}

type Sslvserver struct {
	Vservername                       string `json:"vservername,omitempty"`
	Cleartextport                     int    `json:"cleartextport,omitempty"`
	Dh                                string `json:"dh,omitempty"`
	Dhfile                            string `json:"dhfile,omitempty"`
	Dhcount                           int    `json:"dhcount,omitempty"`
	Dhkeyexpsizelimit                 string `json:"dhkeyexpsizelimit,omitempty"`
	Ersa                              string `json:"ersa,omitempty"`
	Ersacount                         int    `json:"ersacount,omitempty"`
	Sessreuse                         string `json:"sessreuse,omitempty"`
	Sesstimeout                       int    `json:"sesstimeout,omitempty"`
	Cipherredirect                    string `json:"cipherredirect,omitempty"`
	Cipherurl                         string `json:"cipherurl,omitempty"`
	Sslv2redirect                     string `json:"sslv2redirect,omitempty"`
	Sslv2url                          string `json:"sslv2url,omitempty"`
	Clientauth                        string `json:"clientauth,omitempty"`
	Clientcert                        string `json:"clientcert,omitempty"`
	Sslredirect                       string `json:"sslredirect,omitempty"`
	Redirectportrewrite               string `json:"redirectportrewrite,omitempty"`
	Ssl2                              string `json:"ssl2,omitempty"`
	Ssl3                              string `json:"ssl3,omitempty"`
	Tls1                              string `json:"tls1,omitempty"`
	Tls11                             string `json:"tls11,omitempty"`
	Tls12                             string `json:"tls12,omitempty"`
	Tls13                             string `json:"tls13,omitempty"`
	Dtls1                             string `json:"dtls1,omitempty"`
	Dtls12                            string `json:"dtls12,omitempty"`
	Snienable                         string `json:"snienable,omitempty"`
	Ocspstapling                      string `json:"ocspstapling,omitempty"`
	Pushenctrigger                    string `json:"pushenctrigger,omitempty"`
	Sendclosenotify                   string `json:"sendclosenotify,omitempty"`
	Dtlsprofilename                   string `json:"dtlsprofilename,omitempty"`
	Sslprofile                        string `json:"sslprofile,omitempty"`
	Hsts                              string `json:"hsts,omitempty"`
	Maxage                            int    `json:"maxage,omitempty"`
	Includesubdomains                 string `json:"includesubdomains,omitempty"`
	Preload                           string `json:"preload,omitempty"`
	Strictsigdigestcheck              string `json:"strictsigdigestcheck,omitempty"`
	Zerorttearlydata                  string `json:"zerorttearlydata,omitempty"`
	Tls13sessionticketsperauthcontext int    `json:"tls13sessionticketsperauthcontext,omitempty"`
	Dhekeyexchangewithpsk             string `json:"dhekeyexchangewithpsk,omitempty"`
	Defaultsni                        string `json:"defaultsni,omitempty"`
	Sslclientlogs                     string `json:"sslclientlogs,omitempty"`
	Crlcheck                          string `json:"crlcheck,omitempty"`
	Nonfipsciphers                    string `json:"nonfipsciphers,omitempty"`
	Service                           string `json:"service,omitempty"`
	Ocspcheck                         string `json:"ocspcheck,omitempty"`
	Ca                                string `json:"ca,omitempty"`
	Snicert                           string `json:"snicert,omitempty"`
	Skipcaname                        string `json:"skipcaname,omitempty"`
	Dtlsflag                          string `json:"dtlsflag,omitempty"`
	Quicflag                          string `json:"quicflag,omitempty"`
	Skipcacertbundle                  string `json:"skipcacertbundle,omitempty"`
	Nextgenapiresource                string `json:"_nextgenapiresource,omitempty"`
}

type Sslvservercipherbinding struct {
	Cipheraliasname string `json:"cipheraliasname,omitempty"`
	Description     string `json:"description,omitempty"`
	Vservername     string `json:"vservername,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
}

type Sslcertkeyservicebinding struct {
	Servicename      string `json:"servicename,omitempty"`
	Data             int    `json:"data,omitempty"`
	Version          int    `json:"version,omitempty"`
	Certkey          string `json:"certkey,omitempty"`
	Service          bool   `json:"service,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Ca               bool   `json:"ca,omitempty"`
}

type Sslecdsakey struct {
	Keyfile  string `json:"keyfile,omitempty"`
	Curve    string `json:"curve,omitempty"`
	Keyform  string `json:"keyform,omitempty"`
	Des      bool   `json:"des,omitempty"`
	Des3     bool   `json:"des3,omitempty"`
	Aes256   bool   `json:"aes256,omitempty"`
	Password string `json:"password,omitempty"`
	Pkcs8    bool   `json:"pkcs8,omitempty"`
}

type Sslkeyfile struct {
	Name               string `json:"name,omitempty"`
	Src                string `json:"src,omitempty"`
	Password           string `json:"password,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslservicecipherbinding struct {
	Cipheraliasname string `json:"cipheraliasname,omitempty"`
	Description     string `json:"description,omitempty"`
	Servicename     string `json:"servicename,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
}

type Sslaction struct {
	Name                   string `json:"name,omitempty"`
	Clientauth             string `json:"clientauth,omitempty"`
	Clientcertverification string `json:"clientcertverification,omitempty"`
	Ssllogprofile          string `json:"ssllogprofile,omitempty"`
	Clientcert             string `json:"clientcert,omitempty"`
	Certheader             string `json:"certheader,omitempty"`
	Clientcertserialnumber string `json:"clientcertserialnumber,omitempty"`
	Certserialheader       string `json:"certserialheader,omitempty"`
	Clientcertsubject      string `json:"clientcertsubject,omitempty"`
	Certsubjectheader      string `json:"certsubjectheader,omitempty"`
	Clientcerthash         string `json:"clientcerthash,omitempty"`
	Certhashheader         string `json:"certhashheader,omitempty"`
	Clientcertfingerprint  string `json:"clientcertfingerprint,omitempty"`
	Certfingerprintheader  string `json:"certfingerprintheader,omitempty"`
	Certfingerprintdigest  string `json:"certfingerprintdigest,omitempty"`
	Clientcertissuer       string `json:"clientcertissuer,omitempty"`
	Certissuerheader       string `json:"certissuerheader,omitempty"`
	Sessionid              string `json:"sessionid,omitempty"`
	Sessionidheader        string `json:"sessionidheader,omitempty"`
	Cipher                 string `json:"cipher,omitempty"`
	Cipherheader           string `json:"cipherheader,omitempty"`
	Clientcertnotbefore    string `json:"clientcertnotbefore,omitempty"`
	Certnotbeforeheader    string `json:"certnotbeforeheader,omitempty"`
	Clientcertnotafter     string `json:"clientcertnotafter,omitempty"`
	Certnotafterheader     string `json:"certnotafterheader,omitempty"`
	Owasupport             string `json:"owasupport,omitempty"`
	Forward                string `json:"forward,omitempty"`
	Cacertgrpname          string `json:"cacertgrpname,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Undefhits              string `json:"undefhits,omitempty"`
	Referencecount         string `json:"referencecount,omitempty"`
	Description            string `json:"description,omitempty"`
	Builtin                string `json:"builtin,omitempty"`
	Feature                string `json:"feature,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Sslglobalbinding struct {
}

type Sslservicegroupcertkeybinding struct {
	Certkeyname      string `json:"certkeyname,omitempty"`
	Crlcheck         string `json:"crlcheck,omitempty"`
	Ocspcheck        string `json:"ocspcheck,omitempty"`
	Ca               bool   `json:"ca,omitempty"`
	Snicert          bool   `json:"snicert,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type Sslpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Reqaction          string `json:"reqaction,omitempty"`
	Action             string `json:"action,omitempty"`
	Undefaction        string `json:"undefaction,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Description        string `json:"description,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslpolicylabelpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
}

type Sslservicesslcipherbinding struct {
	Cipheraliasname string `json:"cipheraliasname,omitempty"`
	Description     string `json:"description,omitempty"`
	Cipherdefaulton int    `json:"cipherdefaulton,omitempty"`
	Servicename     string `json:"servicename,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
}

type Sslservicegroupsslciphersuitebinding struct {
	Ciphername       string `json:"ciphername,omitempty"`
	Description      string `json:"description,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type Sslvserverciphersuitebinding struct {
	Ciphername  string `json:"ciphername,omitempty"`
	Description string `json:"description,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}

type Sslcrlserialnumberbinding struct {
	Number  string `json:"number,omitempty"`
	Date    string `json:"date,omitempty"`
	Crlname string `json:"crlname,omitempty"`
}

type Sslpolicycsvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Sslprofilesslcipherbinding struct {
	Cipheraliasname string `json:"cipheraliasname,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
	Description     string `json:"description,omitempty"`
	Name            string `json:"name,omitempty"`
	Ciphername      string `json:"ciphername,omitempty"`
}

type Sslcacertgroupbinding struct {
	Cacertgroupname string `json:"cacertgroupname,omitempty"`
}

type Sslciphersslvserverbinding struct {
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Vservername     string `json:"vservername,omitempty"`
	Vserver         bool   `json:"vserver,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
	Cipherpriority  int    `json:"cipherpriority,omitempty"`
}

type Sslprofile struct {
	Name                              string `json:"name,omitempty"`
	Sslprofiletype                    string `json:"sslprofiletype,omitempty"`
	Ssllogprofile                     string `json:"ssllogprofile,omitempty"`
	Dhcount                           int    `json:"dhcount,omitempty"`
	Dh                                string `json:"dh,omitempty"`
	Dhfile                            string `json:"dhfile,omitempty"`
	Ersa                              string `json:"ersa,omitempty"`
	Ersacount                         int    `json:"ersacount,omitempty"`
	Sessreuse                         string `json:"sessreuse,omitempty"`
	Sesstimeout                       int    `json:"sesstimeout,omitempty"`
	Cipherredirect                    string `json:"cipherredirect,omitempty"`
	Cipherurl                         string `json:"cipherurl,omitempty"`
	Clientauth                        string `json:"clientauth,omitempty"`
	Clientcert                        string `json:"clientcert,omitempty"`
	Dhkeyexpsizelimit                 string `json:"dhkeyexpsizelimit,omitempty"`
	Sslredirect                       string `json:"sslredirect,omitempty"`
	Redirectportrewrite               string `json:"redirectportrewrite,omitempty"`
	Ssl3                              string `json:"ssl3,omitempty"`
	Tls1                              string `json:"tls1,omitempty"`
	Tls11                             string `json:"tls11,omitempty"`
	Tls12                             string `json:"tls12,omitempty"`
	Tls13                             string `json:"tls13,omitempty"`
	Snienable                         string `json:"snienable,omitempty"`
	Allowunknownsni                   string `json:"allowunknownsni,omitempty"`
	Ocspstapling                      string `json:"ocspstapling,omitempty"`
	Serverauth                        string `json:"serverauth,omitempty"`
	Commonname                        string `json:"commonname,omitempty"`
	Pushenctrigger                    string `json:"pushenctrigger,omitempty"`
	Sendclosenotify                   string `json:"sendclosenotify,omitempty"`
	Cleartextport                     int    `json:"cleartextport,omitempty"`
	Insertionencoding                 string `json:"insertionencoding,omitempty"`
	Denysslreneg                      string `json:"denysslreneg,omitempty"`
	Maxrenegrate                      int    `json:"maxrenegrate,omitempty"`
	Quantumsize                       string `json:"quantumsize,omitempty"`
	Strictcachecks                    string `json:"strictcachecks,omitempty"`
	Encrypttriggerpktcount            int    `json:"encrypttriggerpktcount,omitempty"`
	Pushflag                          int    `json:"pushflag,omitempty"`
	Dropreqwithnohostheader           string `json:"dropreqwithnohostheader,omitempty"`
	Snihttphostmatch                  string `json:"snihttphostmatch,omitempty"`
	Pushenctriggertimeout             int    `json:"pushenctriggertimeout,omitempty"`
	Ssltriggertimeout                 int    `json:"ssltriggertimeout,omitempty"`
	Clientauthuseboundcachain         string `json:"clientauthuseboundcachain,omitempty"`
	Sslinterception                   string `json:"sslinterception,omitempty"`
	Sslireneg                         string `json:"sslireneg,omitempty"`
	Ssliocspcheck                     string `json:"ssliocspcheck,omitempty"`
	Sslimaxsessperserver              int    `json:"sslimaxsessperserver,omitempty"`
	Sessionticket                     string `json:"sessionticket,omitempty"`
	Sessionticketlifetime             int    `json:"sessionticketlifetime,omitempty"`
	Sessionticketkeyrefresh           string `json:"sessionticketkeyrefresh,omitempty"`
	Sessionticketkeydata              string `json:"sessionticketkeydata,omitempty"`
	Sessionkeylifetime                int    `json:"sessionkeylifetime,omitempty"`
	Prevsessionkeylifetime            int    `json:"prevsessionkeylifetime,omitempty"`
	Hsts                              string `json:"hsts,omitempty"`
	Maxage                            int    `json:"maxage,omitempty"`
	Includesubdomains                 string `json:"includesubdomains,omitempty"`
	Preload                           string `json:"preload,omitempty"`
	Skipclientcertpolicycheck         string `json:"skipclientcertpolicycheck,omitempty"`
	Zerorttearlydata                  string `json:"zerorttearlydata,omitempty"`
	Tls13sessionticketsperauthcontext int    `json:"tls13sessionticketsperauthcontext,omitempty"`
	Dhekeyexchangewithpsk             string `json:"dhekeyexchangewithpsk,omitempty"`
	Allowextendedmastersecret         string `json:"allowextendedmastersecret,omitempty"`
	Alpnprotocol                      string `json:"alpnprotocol,omitempty"`
	Encryptedclienthello              string `json:"encryptedclienthello,omitempty"`
	Defaultsni                        string `json:"defaultsni,omitempty"`
	Sslclientlogs                     string `json:"sslclientlogs,omitempty"`
	Ciphername                        string `json:"ciphername,omitempty"`
	Cipherpriority                    int    `json:"cipherpriority,omitempty"`
	Strictsigdigestcheck              string `json:"strictsigdigestcheck,omitempty"`
	Nonfipsciphers                    string `json:"nonfipsciphers,omitempty"`
	Crlcheck                          string `json:"crlcheck,omitempty"`
	Ocspcheck                         string `json:"ocspcheck,omitempty"`
	Snicert                           string `json:"snicert,omitempty"`
	Skipcaname                        string `json:"skipcaname,omitempty"`
	Invoke                            string `json:"invoke,omitempty"`
	Labeltype                         string `json:"labeltype,omitempty"`
	Service                           string `json:"service,omitempty"`
	Builtin                           string `json:"builtin,omitempty"`
	Feature                           string `json:"feature,omitempty"`
	Sslpfobjecttype                   string `json:"sslpfobjecttype,omitempty"`
	Ssliverifyservercertforreuse      string `json:"ssliverifyservercertforreuse,omitempty"`
	Nodefaultbindings                 string `json:"nodefaultbindings,omitempty"`
	Nextgenapiresource                string `json:"_nextgenapiresource,omitempty"`
}

type Sslprofilesslvserverbinding struct {
	Servicename    string `json:"servicename,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	Cipherpriority int    `json:"cipherpriority,omitempty"`
}

type Sslservicebinding struct {
	Servicename string `json:"servicename,omitempty"`
}

type Sslvserverpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Type                   string `json:"type,omitempty"`
	Polinherit             uint32 `json:"polinherit,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Vservername            string `json:"vservername,omitempty"`
}

type Sslfips struct {
	Inithsm             string `json:"inithsm,omitempty"`
	Sopassword          string `json:"sopassword,omitempty"`
	Oldsopassword       string `json:"oldsopassword,omitempty"`
	Userpassword        string `json:"userpassword,omitempty"`
	Hsmlabel            string `json:"hsmlabel,omitempty"`
	Fipsfw              string `json:"fipsfw,omitempty"`
	Erasedata           string `json:"erasedata,omitempty"`
	Serial              string `json:"serial,omitempty"`
	Majorversion        string `json:"majorversion,omitempty"`
	Minorversion        string `json:"minorversion,omitempty"`
	Fipshwmajorversion  string `json:"fipshwmajorversion,omitempty"`
	Fipshwminorversion  string `json:"fipshwminorversion,omitempty"`
	Fipshwversionstring string `json:"fipshwversionstring,omitempty"`
	Flashmemorytotal    string `json:"flashmemorytotal,omitempty"`
	Flashmemoryfree     string `json:"flashmemoryfree,omitempty"`
	Sramtotal           string `json:"sramtotal,omitempty"`
	Sramfree            string `json:"sramfree,omitempty"`
	Status              string `json:"status,omitempty"`
	Flag                string `json:"flag,omitempty"`
	Serialno            string `json:"serialno,omitempty"`
	Model               string `json:"model,omitempty"`
	State               string `json:"state,omitempty"`
	Firmwarereleasedate string `json:"firmwarereleasedate,omitempty"`
	Coresmax            string `json:"coresmax,omitempty"`
	Coresenabled        string `json:"coresenabled,omitempty"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
}

type Sslservicepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Polinherit             uint32 `json:"polinherit,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Servicename            string `json:"servicename,omitempty"`
}

type Sslcertbundle struct {
	Name               string `json:"name,omitempty"`
	Src                string `json:"src,omitempty"`
	Inuse              string `json:"inuse,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Sslcertkey struct {
	Certkey                     string `json:"certkey,omitempty"`
	Cert                        string `json:"cert,omitempty"`
	Key                         string `json:"key,omitempty"`
	Password                    bool   `json:"password,omitempty"`
	Fipskey                     string `json:"fipskey,omitempty"`
	Hsmkey                      string `json:"hsmkey,omitempty"`
	Inform                      string `json:"inform,omitempty"`
	Passplain                   string `json:"passplain,omitempty"`
	Expirymonitor               string `json:"expirymonitor,omitempty"`
	Notificationperiod          int    `json:"notificationperiod,omitempty"`
	Bundle                      string `json:"bundle,omitempty"`
	Deletecertkeyfilesonremoval string `json:"deletecertkeyfilesonremoval,omitempty"`
	Deletefromdevice            bool   `json:"deletefromdevice,omitempty"`
	Linkcertkeyname             string `json:"linkcertkeyname,omitempty"`
	Nodomaincheck               bool   `json:"nodomaincheck,omitempty"`
	Ocspstaplingcache           bool   `json:"ocspstaplingcache,omitempty"`
	Signaturealg                string `json:"signaturealg,omitempty"`
	Certificatetype             string `json:"certificatetype,omitempty"`
	Serial                      string `json:"serial,omitempty"`
	Issuer                      string `json:"issuer,omitempty"`
	Clientcertnotbefore         string `json:"clientcertnotbefore,omitempty"`
	Clientcertnotafter          string `json:"clientcertnotafter,omitempty"`
	Daystoexpiration            string `json:"daystoexpiration,omitempty"`
	Subject                     string `json:"subject,omitempty"`
	Publickey                   string `json:"publickey,omitempty"`
	Publickeysize               string `json:"publickeysize,omitempty"`
	Version                     string `json:"version,omitempty"`
	Priority                    string `json:"priority,omitempty"`
	Status                      string `json:"status,omitempty"`
	Passcrypt                   string `json:"passcrypt,omitempty"`
	Data                        string `json:"data,omitempty"`
	Servicename                 string `json:"servicename,omitempty"`
	Sandns                      string `json:"sandns,omitempty"`
	Sanipadd                    string `json:"sanipadd,omitempty"`
	Ocspresponsestatus          string `json:"ocspresponsestatus,omitempty"`
	Builtin                     string `json:"builtin,omitempty"`
	Feature                     string `json:"feature,omitempty"`
	Certkeydigest               string `json:"certkeydigest,omitempty"`
	Certificatesource           string `json:"certificatesource,omitempty"`
	Certkeystatus               string `json:"certkeystatus,omitempty"`
	Nextgenapiresource          string `json:"_nextgenapiresource,omitempty"`
}

type Sslciphervserverbinding struct {
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	Vservername     string `json:"vservername,omitempty"`
	Vserver         bool   `json:"vserver,omitempty"`
	Cipheroperation string `json:"cipheroperation,omitempty"`
	Ciphgrpals      string `json:"ciphgrpals,omitempty"`
	Cipherpriority  uint32 `json:"cipherpriority,omitempty"`
}

type Sslservicegroupciphersuitebinding struct {
	Ciphername       string `json:"ciphername,omitempty"`
	Description      string `json:"description,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}
