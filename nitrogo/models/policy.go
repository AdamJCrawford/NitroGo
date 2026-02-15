package models

// policy configuration structs
type PolicypatsetPatternBinding struct {
	Builtin []string `json:"builtin,omitempty"`
	Charset string   `json:"charset,omitempty"`
	Comment string   `json:"comment,omitempty"`
	Feature string   `json:"feature,omitempty"`
	Index   int      `json:"index,omitempty"`
	Name    string   `json:"name,omitempty"`
	String  string   `json:"String,omitempty"`
}

type PolicystringmapBinding struct {
	Name                          string        `json:"name,omitempty"`
	PolicystringmapPatternBinding []interface{} `json:"policystringmap_pattern_binding,omitempty"`
}

type PolicydatasetBinding struct {
	Name                      string        `json:"name,omitempty"`
	PolicydatasetValueBinding []interface{} `json:"policydataset_value_binding,omitempty"`
}

type PolicystringmapPatternBinding struct {
	Comment string `json:"comment,omitempty"`
	Key     string `json:"key,omitempty"`
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`
}

type Policypatsetfile struct {
	Bindstatus         string  `json:"bindstatus,omitempty"`
	Bindstatuscode     int     `json:"bindstatuscode,omitempty"`
	Boundpatterns      int     `json:"boundpatterns,omitempty"`
	Charset            string  `json:"charset,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Delimiter          string  `json:"delimiter,omitempty"`
	Imported           bool    `json:"imported,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool    `json:"overwrite,omitempty"`
	Patsetname         string  `json:"patsetname,omitempty"`
	Src                string  `json:"src,omitempty"`
	Totalpatterns      int     `json:"totalpatterns,omitempty"`
}

type Policystringmap struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type PolicypatsetBinding struct {
	Name                       string        `json:"name,omitempty"`
	PolicypatsetPatternBinding []interface{} `json:"policypatset_pattern_binding,omitempty"`
}

type PolicydatasetValueBinding struct {
	Comment  string `json:"comment,omitempty"`
	Endrange string `json:"endrange,omitempty"`
	Index    int    `json:"index,omitempty"`
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
}

type Policytracing struct {
	Capturesslhandshakepolicies string   `json:"capturesslhandshakepolicies,omitempty"`
	Clientip                    string   `json:"clientip,omitempty"`
	Count                       float64  `json:"__count,omitempty"`
	Destip                      string   `json:"destip,omitempty"`
	Destport                    int      `json:"destport,omitempty"`
	Detail                      string   `json:"detail,omitempty"`
	Filterexpr                  string   `json:"filterexpr,omitempty"`
	Isresponse                  int      `json:"isresponse,omitempty"`
	Isundefpolicy               int      `json:"isundefpolicy,omitempty"`
	Nextgenapiresource          string   `json:"_nextgenapiresource,omitempty"`
	Nodeid                      int      `json:"nodeid,omitempty"`
	Packetengineid              int      `json:"packetengineid,omitempty"`
	Policynames                 []string `json:"policynames,omitempty"`
	Policytracingmodule         string   `json:"policytracingmodule,omitempty"`
	Policytracingrecordcount    int      `json:"policytracingrecordcount,omitempty"`
	Protocoltype                string   `json:"protocoltype,omitempty"`
	Srcport                     int      `json:"srcport,omitempty"`
	Transactionid               string   `json:"transactionid,omitempty"`
	Transactiontime             string   `json:"transactiontime,omitempty"`
	Url                         string   `json:"url,omitempty"`
}

type Policyurlset struct {
	Canaryurl           string  `json:"canaryurl,omitempty"`
	Comment             string  `json:"comment,omitempty"`
	Count               float64 `json:"__count,omitempty"`
	Delimiter           string  `json:"delimiter,omitempty"`
	Imported            bool    `json:"imported,omitempty"`
	Interval            int     `json:"interval,omitempty"`
	Matchedid           int     `json:"matchedid,omitempty"`
	Name                string  `json:"name,omitempty"`
	Nextgenapiresource  string  `json:"_nextgenapiresource,omitempty"`
	Overwrite           bool    `json:"overwrite,omitempty"`
	Patterncount        int     `json:"patterncount,omitempty"`
	Privateset          bool    `json:"privateset,omitempty"`
	Rowseparator        string  `json:"rowseparator,omitempty"`
	Subdomainexactmatch bool    `json:"subdomainexactmatch,omitempty"`
	Url                 string  `json:"url,omitempty"`
}

type Policyevaluation struct {
	Action                     string        `json:"action,omitempty"`
	Count                      float64       `json:"__count,omitempty"`
	Expression                 string        `json:"expression,omitempty"`
	Input                      string        `json:"input,omitempty"`
	Istruncatedrefresult       bool          `json:"istruncatedrefresult,omitempty"`
	Nextgenapiresource         string        `json:"_nextgenapiresource,omitempty"`
	Pitactionerrorresult       string        `json:"pitactionerrorresult,omitempty"`
	Pitactionevaltime          int           `json:"pitactionevaltime,omitempty"`
	Pitboolerrorresult         string        `json:"pitboolerrorresult,omitempty"`
	Pitboolevaltime            int           `json:"pitboolevaltime,omitempty"`
	Pitboolresult              bool          `json:"pitboolresult,omitempty"`
	Pitdoubleerrorresult       string        `json:"pitdoubleerrorresult,omitempty"`
	Pitdoubleevaltime          int           `json:"pitdoubleevaltime,omitempty"`
	Pitdoubleresult            float64       `json:"pitdoubleresult,omitempty"`
	Pitmodifiedinputdata       string        `json:"pitmodifiedinputdata,omitempty"`
	Pitnewoffsetarray          []interface{} `json:"pitnewoffsetarray,omitempty"`
	Pitnumerrorresult          string        `json:"pitnumerrorresult,omitempty"`
	Pitnumevaltime             int           `json:"pitnumevaltime,omitempty"`
	Pitnumresult               int           `json:"pitnumresult,omitempty"`
	Pitoffseterrorresult       string        `json:"pitoffseterrorresult,omitempty"`
	Pitoffsetevaltime          int           `json:"pitoffsetevaltime,omitempty"`
	Pitoffsetlengtharray       []interface{} `json:"pitoffsetlengtharray,omitempty"`
	Pitoffsetnewlengtharray    []interface{} `json:"pitoffsetnewlengtharray,omitempty"`
	Pitoffsetresult            int           `json:"pitoffsetresult,omitempty"`
	Pitoffsetresultlen         int           `json:"pitoffsetresultlen,omitempty"`
	Pitoldoffsetarray          []interface{} `json:"pitoldoffsetarray,omitempty"`
	Pitoperationperformerarray []string      `json:"pitoperationperformerarray,omitempty"`
	Pitreferrorresult          string        `json:"pitreferrorresult,omitempty"`
	Pitrefevaltime             int           `json:"pitrefevaltime,omitempty"`
	Pitrefresult               string        `json:"pitrefresult,omitempty"`
	Pitulongerrorresult        string        `json:"pitulongerrorresult,omitempty"`
	Pitulongevaltime           int           `json:"pitulongevaltime,omitempty"`
	Pitulongresult             int           `json:"pitulongresult,omitempty"`
	TypeField                  string        `json:"type,omitempty"`
}

type Policyhttpcallout struct {
	Bodyexpr           string   `json:"bodyexpr,omitempty"`
	Cacheforsecs       int      `json:"cacheforsecs,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Effectivestate     string   `json:"effectivestate,omitempty"`
	Fullreqexpr        string   `json:"fullreqexpr,omitempty"`
	Headers            []string `json:"headers,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Hostexpr           string   `json:"hostexpr,omitempty"`
	Httpmethod         string   `json:"httpmethod,omitempty"`
	Ipaddress          string   `json:"ipaddress,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Parameters         []string `json:"parameters,omitempty"`
	Port               int      `json:"port,omitempty"`
	Recursivecallout   int      `json:"recursivecallout,omitempty"`
	Resultexpr         string   `json:"resultexpr,omitempty"`
	Returntype         string   `json:"returntype,omitempty"`
	Scheme             string   `json:"scheme,omitempty"`
	Svrstate           string   `json:"svrstate,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
	Undefreason        string   `json:"undefreason,omitempty"`
	Urlstemexpr        string   `json:"urlstemexpr,omitempty"`
	Vserver            string   `json:"vserver,omitempty"`
}

type Policydataset struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Dynamic            string  `json:"dynamic,omitempty"`
	Dynamiconly        bool    `json:"dynamiconly,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Patsetfile         string  `json:"patsetfile,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type Policymap struct {
	Count              float64 `json:"__count,omitempty"`
	Mappolicyname      string  `json:"mappolicyname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Sd                 string  `json:"sd,omitempty"`
	Su                 string  `json:"su,omitempty"`
	Targetname         string  `json:"targetname,omitempty"`
	Td                 string  `json:"td,omitempty"`
	Tu                 string  `json:"tu,omitempty"`
}

type Policyexpression struct {
	Builtin               []string `json:"builtin,omitempty"`
	Clientsecuritymessage string   `json:"clientsecuritymessage,omitempty"`
	Comment               string   `json:"comment,omitempty"`
	Count                 float64  `json:"__count,omitempty"`
	Feature               string   `json:"feature,omitempty"`
	Hits                  int      `json:"hits,omitempty"`
	Isdefault             bool     `json:"isdefault,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Nextgenapiresource    string   `json:"_nextgenapiresource,omitempty"`
	Pihits                int      `json:"pihits,omitempty"`
	Type1                 string   `json:"type1,omitempty"`
	TypeField             string   `json:"type,omitempty"`
	Value                 string   `json:"value,omitempty"`
}

type Policypatset struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Dynamic            string  `json:"dynamic,omitempty"`
	Dynamiconly        bool    `json:"dynamiconly,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Patsetfile         string  `json:"patsetfile,omitempty"`
}

type Policyparam struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}
