// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Policydataset struct {
	Name               string `json:"name,omitempty"`
	Type               string `json:"type,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Patsetfile         string `json:"patsetfile,omitempty"`
	Dynamic            string `json:"dynamic,omitempty"`
	Dynamiconly        bool   `json:"dynamiconly,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Policyhttpcallout struct {
	Name               string   `json:"name,omitempty"`
	Ipaddress          string   `json:"ipaddress,omitempty"`
	Port               int      `json:"port,omitempty"`
	Vserver            string   `json:"vserver,omitempty"`
	Returntype         string   `json:"returntype,omitempty"`
	Httpmethod         string   `json:"httpmethod,omitempty"`
	Hostexpr           string   `json:"hostexpr,omitempty"`
	Urlstemexpr        string   `json:"urlstemexpr,omitempty"`
	Headers            []string `json:"headers,omitempty"`
	Parameters         []string `json:"parameters,omitempty"`
	Bodyexpr           string   `json:"bodyexpr,omitempty"`
	Fullreqexpr        string   `json:"fullreqexpr,omitempty"`
	Scheme             string   `json:"scheme,omitempty"`
	Resultexpr         string   `json:"resultexpr,omitempty"`
	Cacheforsecs       int      `json:"cacheforsecs,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Hits               string   `json:"hits,omitempty"`
	Undefhits          string   `json:"undefhits,omitempty"`
	Svrstate           string   `json:"svrstate,omitempty"`
	Effectivestate     string   `json:"effectivestate,omitempty"`
	Undefreason        string   `json:"undefreason,omitempty"`
	Recursivecallout   string   `json:"recursivecallout,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Policypatsetpatternbinding struct {
	String  string   `json:"String,omitempty"`
	Index   int      `json:"index,omitempty"`
	Charset string   `json:"charset,omitempty"`
	Comment string   `json:"comment,omitempty"`
	Builtin []string `json:"builtin,omitempty"`
	Feature string   `json:"feature,omitempty"`
	Name    string   `json:"name,omitempty"`
}

type Policypatsetfile struct {
	Src                string `json:"src,omitempty"`
	Name               string `json:"name,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Delimiter          string `json:"delimiter,omitempty"`
	Charset            string `json:"charset,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Imported           bool   `json:"imported,omitempty"`
	Totalpatterns      string `json:"totalpatterns,omitempty"`
	Boundpatterns      string `json:"boundpatterns,omitempty"`
	Patsetname         string `json:"patsetname,omitempty"`
	Bindstatuscode     string `json:"bindstatuscode,omitempty"`
	Bindstatus         string `json:"bindstatus,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Policystringmappatternbinding struct {
	Key     string `json:"key,omitempty"`
	Value   string `json:"value,omitempty"`
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Policydatasetbinding struct {
	Name string `json:"name,omitempty"`
}

type Policyevaluation struct {
	Expression                 string `json:"expression,omitempty"`
	Action                     string `json:"action,omitempty"`
	Type                       string `json:"type,omitempty"`
	Input                      string `json:"input,omitempty"`
	Pitmodifiedinputdata       string `json:"pitmodifiedinputdata,omitempty"`
	Pitboolresult              string `json:"pitboolresult,omitempty"`
	Pitnumresult               string `json:"pitnumresult,omitempty"`
	Pitdoubleresult            string `json:"pitdoubleresult,omitempty"`
	Pitulongresult             string `json:"pitulongresult,omitempty"`
	Pitrefresult               string `json:"pitrefresult,omitempty"`
	Pitoffsetresult            string `json:"pitoffsetresult,omitempty"`
	Pitoffsetresultlen         string `json:"pitoffsetresultlen,omitempty"`
	Istruncatedrefresult       string `json:"istruncatedrefresult,omitempty"`
	Pitboolevaltime            string `json:"pitboolevaltime,omitempty"`
	Pitnumevaltime             string `json:"pitnumevaltime,omitempty"`
	Pitdoubleevaltime          string `json:"pitdoubleevaltime,omitempty"`
	Pitulongevaltime           string `json:"pitulongevaltime,omitempty"`
	Pitrefevaltime             string `json:"pitrefevaltime,omitempty"`
	Pitoffsetevaltime          string `json:"pitoffsetevaltime,omitempty"`
	Pitactionevaltime          string `json:"pitactionevaltime,omitempty"`
	Pitoperationperformerarray string `json:"pitoperationperformerarray,omitempty"`
	Pitoldoffsetarray          string `json:"pitoldoffsetarray,omitempty"`
	Pitnewoffsetarray          string `json:"pitnewoffsetarray,omitempty"`
	Pitoffsetlengtharray       string `json:"pitoffsetlengtharray,omitempty"`
	Pitoffsetnewlengtharray    string `json:"pitoffsetnewlengtharray,omitempty"`
	Pitboolerrorresult         string `json:"pitboolerrorresult,omitempty"`
	Pitnumerrorresult          string `json:"pitnumerrorresult,omitempty"`
	Pitdoubleerrorresult       string `json:"pitdoubleerrorresult,omitempty"`
	Pitulongerrorresult        string `json:"pitulongerrorresult,omitempty"`
	Pitreferrorresult          string `json:"pitreferrorresult,omitempty"`
	Pitoffseterrorresult       string `json:"pitoffseterrorresult,omitempty"`
	Pitactionerrorresult       string `json:"pitactionerrorresult,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Policyexpression struct {
	Name                  string `json:"name,omitempty"`
	Value                 string `json:"value,omitempty"`
	Comment               string `json:"comment,omitempty"`
	Clientsecuritymessage string `json:"clientsecuritymessage,omitempty"`
	Type                  string `json:"type,omitempty"`
	Hits                  string `json:"hits,omitempty"`
	Pihits                string `json:"pihits,omitempty"`
	Type1                 string `json:"type1,omitempty"`
	Isdefault             string `json:"isdefault,omitempty"`
	Builtin               string `json:"builtin,omitempty"`
	Feature               string `json:"feature,omitempty"`
	Nextgenapiresource    string `json:"_nextgenapiresource,omitempty"`
}

type Policyparam struct {
	Timeout            int    `json:"timeout,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Policypatset struct {
	Name               string `json:"name,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Patsetfile         string `json:"patsetfile,omitempty"`
	Dynamic            string `json:"dynamic,omitempty"`
	Dynamiconly        bool   `json:"dynamiconly,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Policymap struct {
	Mappolicyname      string `json:"mappolicyname,omitempty"`
	Sd                 string `json:"sd,omitempty"`
	Su                 string `json:"su,omitempty"`
	Td                 string `json:"td,omitempty"`
	Tu                 string `json:"tu,omitempty"`
	Targetname         string `json:"targetname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Policypatsetbinding struct {
	Name string `json:"name,omitempty"`
}

type Policystringmap struct {
	Name               string `json:"name,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Policytracing struct {
	Filterexpr                  string `json:"filterexpr,omitempty"`
	Protocoltype                string `json:"protocoltype,omitempty"`
	Capturesslhandshakepolicies string `json:"capturesslhandshakepolicies,omitempty"`
	Transactionid               string `json:"transactionid,omitempty"`
	Detail                      string `json:"detail,omitempty"`
	Nodeid                      int    `json:"nodeid,omitempty"`
	Packetengineid              string `json:"packetengineid,omitempty"`
	Clientip                    string `json:"clientip,omitempty"`
	Destip                      string `json:"destip,omitempty"`
	Srcport                     string `json:"srcport,omitempty"`
	Destport                    string `json:"destport,omitempty"`
	Transactiontime             string `json:"transactiontime,omitempty"`
	Policytracingmodule         string `json:"policytracingmodule,omitempty"`
	Url                         string `json:"url,omitempty"`
	Policynames                 string `json:"policynames,omitempty"`
	Isresponse                  string `json:"isresponse,omitempty"`
	Isundefpolicy               string `json:"isundefpolicy,omitempty"`
	Policytracingrecordcount    string `json:"policytracingrecordcount,omitempty"`
	Nextgenapiresource          string `json:"_nextgenapiresource,omitempty"`
}

type Policyurlset struct {
	Name                string `json:"name,omitempty"`
	Comment             string `json:"comment,omitempty"`
	Imported            bool   `json:"imported,omitempty"`
	Overwrite           bool   `json:"overwrite,omitempty"`
	Delimiter           string `json:"delimiter,omitempty"`
	Rowseparator        string `json:"rowseparator,omitempty"`
	Url                 string `json:"url,omitempty"`
	Interval            int    `json:"interval,omitempty"`
	Privateset          bool   `json:"privateset,omitempty"`
	Subdomainexactmatch bool   `json:"subdomainexactmatch,omitempty"`
	Matchedid           int    `json:"matchedid,omitempty"`
	Canaryurl           string `json:"canaryurl,omitempty"`
	Patterncount        string `json:"patterncount,omitempty"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
}

type Policydatasetvaluebinding struct {
	Value    string `json:"value,omitempty"`
	Index    int    `json:"index,omitempty"`
	Comment  string `json:"comment,omitempty"`
	Endrange string `json:"endrange,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Policystringmapbinding struct {
	Name string `json:"name,omitempty"`
}
