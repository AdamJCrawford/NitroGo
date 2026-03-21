package models

// policy configuration structs
type PolicyPatsetPatternBinding struct {
	Builtin []string `json:"builtin,omitempty"`
	Charset string   `json:"charset,omitempty"`
	Comment string   `json:"comment,omitempty"`
	Feature string   `json:"feature,omitempty"`
	Index   int      `json:"index,omitempty"`
	Name    string   `json:"name,omitempty"`
	String  string   `json:"String,omitempty"`
}

type PolicyStringMapBinding struct {
	Name                          string `json:"name,omitempty"`
	PolicyStringMapPatternBinding []any  `json:"policystringmap_pattern_binding,omitempty"`
}

type PolicyDatasetBinding struct {
	Name                      string `json:"name,omitempty"`
	PolicyDatasetValueBinding []any  `json:"policydataset_value_binding,omitempty"`
}

type PolicyStringMapPatternBinding struct {
	Comment string `json:"comment,omitempty"`
	Key     string `json:"key,omitempty"`
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`
}

type PolicyPatsetFile struct {
	BindStatus         string  `json:"bindstatus,omitempty"`
	BindStatusCode     int     `json:"bindstatuscode,omitempty"`
	BoundPatterns      int     `json:"boundpatterns,omitempty"`
	Charset            string  `json:"charset,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Delimiter          string  `json:"delimiter,omitempty"`
	Imported           bool    `json:"imported,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool    `json:"overwrite,omitempty"`
	PatsetName         string  `json:"patsetname,omitempty"`
	Src                string  `json:"src,omitempty"`
	TotalPatterns      int     `json:"totalpatterns,omitempty"`
}

type PolicyStringMap struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type PolicyPatsetBinding struct {
	Name                       string `json:"name,omitempty"`
	PolicyPatsetPatternBinding []any  `json:"policypatset_pattern_binding,omitempty"`
}

type PolicyDatasetValueBinding struct {
	Comment  string `json:"comment,omitempty"`
	EndRange string `json:"endrange,omitempty"`
	Index    int    `json:"index,omitempty"`
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
}

type PolicyTracing struct {
	CaptureSSLHandshakePolicies string   `json:"capturesslhandshakepolicies,omitempty"`
	ClientIP                    string   `json:"clientip,omitempty"`
	Count                       float64  `json:"__count,omitempty"`
	DestIP                      string   `json:"destip,omitempty"`
	DestPort                    int      `json:"destport,omitempty"`
	Detail                      string   `json:"detail,omitempty"`
	FilterExpr                  string   `json:"filterexpr,omitempty"`
	IsResponse                  int      `json:"isresponse,omitempty"`
	IsUndefPolicy               int      `json:"isundefpolicy,omitempty"`
	NextGenAPIResource          string   `json:"_nextgenapiresource,omitempty"`
	NodeID                      int      `json:"nodeid,omitempty"`
	PacketEngineID              int      `json:"packetengineid,omitempty"`
	PolicyNames                 []string `json:"policynames,omitempty"`
	PolicyTracingModule         string   `json:"policytracingmodule,omitempty"`
	PolicyTracingRecordCount    int      `json:"policytracingrecordcount,omitempty"`
	ProtocolType                string   `json:"protocoltype,omitempty"`
	SrcPort                     int      `json:"srcport,omitempty"`
	TransactionID               string   `json:"transactionid,omitempty"`
	TransactionTime             string   `json:"transactiontime,omitempty"`
	URL                         string   `json:"url,omitempty"`
}

type PolicyURLSet struct {
	CanaryURL           string  `json:"canaryurl,omitempty"`
	Comment             string  `json:"comment,omitempty"`
	Count               float64 `json:"__count,omitempty"`
	Delimiter           string  `json:"delimiter,omitempty"`
	Imported            bool    `json:"imported,omitempty"`
	Interval            int     `json:"interval,omitempty"`
	MatchedID           int     `json:"matchedid,omitempty"`
	Name                string  `json:"name,omitempty"`
	NextGenAPIResource  string  `json:"_nextgenapiresource,omitempty"`
	Overwrite           bool    `json:"overwrite,omitempty"`
	PatternCount        int     `json:"patterncount,omitempty"`
	PrivateSet          bool    `json:"privateset,omitempty"`
	RowSeparator        string  `json:"rowseparator,omitempty"`
	SubdomainExactMatch bool    `json:"subdomainexactmatch,omitempty"`
	URL                 string  `json:"url,omitempty"`
}

type PolicyEvaluation struct {
	Action                     string   `json:"action,omitempty"`
	Count                      float64  `json:"__count,omitempty"`
	Expression                 string   `json:"expression,omitempty"`
	Input                      string   `json:"input,omitempty"`
	IsTruncatedRefResult       bool     `json:"istruncatedrefresult,omitempty"`
	NextGenAPIResource         string   `json:"_nextgenapiresource,omitempty"`
	PitActionErrorResult       string   `json:"pitactionerrorresult,omitempty"`
	PitActionEvalTime          int      `json:"pitactionevaltime,omitempty"`
	PitBoolErrorResult         string   `json:"pitboolerrorresult,omitempty"`
	PitBoolEvalTime            int      `json:"pitboolevaltime,omitempty"`
	PitBoolResult              bool     `json:"pitboolresult,omitempty"`
	PitDoubleErrorResult       string   `json:"pitdoubleerrorresult,omitempty"`
	PitDoubleEvalTime          int      `json:"pitdoubleevaltime,omitempty"`
	PitDoubleResult            float64  `json:"pitdoubleresult,omitempty"`
	PitModifiedInputData       string   `json:"pitmodifiedinputdata,omitempty"`
	PitNewOffsetArray          []any    `json:"pitnewoffsetarray,omitempty"`
	PitNumErrorResult          string   `json:"pitnumerrorresult,omitempty"`
	PitNumEvalTime             int      `json:"pitnumevaltime,omitempty"`
	PitNumResult               int      `json:"pitnumresult,omitempty"`
	PitOffsetErrorResult       string   `json:"pitoffseterrorresult,omitempty"`
	PitOffsetEvalTime          int      `json:"pitoffsetevaltime,omitempty"`
	PitOffsetLengthArray       []any    `json:"pitoffsetlengtharray,omitempty"`
	PitOffsetNewLengthArray    []any    `json:"pitoffsetnewlengtharray,omitempty"`
	PitOffsetResult            int      `json:"pitoffsetresult,omitempty"`
	PitOffsetResultLen         int      `json:"pitoffsetresultlen,omitempty"`
	PitOldOffsetArray          []any    `json:"pitoldoffsetarray,omitempty"`
	PitOperationPerformerArray []string `json:"pitoperationperformerarray,omitempty"`
	PitRefErrorResult          string   `json:"pitreferrorresult,omitempty"`
	PitRefEvalTime             int      `json:"pitrefevaltime,omitempty"`
	PitRefResult               string   `json:"pitrefresult,omitempty"`
	PitUlongErrorResult        string   `json:"pitulongerrorresult,omitempty"`
	PitUlongEvalTime           int      `json:"pitulongevaltime,omitempty"`
	PitUlongResult             int      `json:"pitulongresult,omitempty"`
	TypeField                  string   `json:"type,omitempty"`
}

type PolicyHTTPCallout struct {
	BodyExpr           string   `json:"bodyexpr,omitempty"`
	CacheForSecs       int      `json:"cacheforsecs,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	EffectiveState     string   `json:"effectivestate,omitempty"`
	FullReqExpr        string   `json:"fullreqexpr,omitempty"`
	Headers            []string `json:"headers,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	HostExpr           string   `json:"hostexpr,omitempty"`
	HTTPMethod         string   `json:"httpmethod,omitempty"`
	IPAddress          string   `json:"ipaddress,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Parameters         []string `json:"parameters,omitempty"`
	Port               int      `json:"port,omitempty"`
	RecursiveCallout   int      `json:"recursivecallout,omitempty"`
	ResultExpr         string   `json:"resultexpr,omitempty"`
	ReturnType         string   `json:"returntype,omitempty"`
	Scheme             string   `json:"scheme,omitempty"`
	SvrState           string   `json:"svrstate,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
	UndefReason        string   `json:"undefreason,omitempty"`
	URLStemExpr        string   `json:"urlstemexpr,omitempty"`
	VServer            string   `json:"vserver,omitempty"`
}

type PolicyDataset struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Dynamic            string  `json:"dynamic,omitempty"`
	DynamicOnly        bool    `json:"dynamiconly,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PatsetFile         string  `json:"patsetfile,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type PolicyMap struct {
	Count              float64 `json:"__count,omitempty"`
	MapPolicyName      string  `json:"mappolicyname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	SD                 string  `json:"sd,omitempty"`
	SU                 string  `json:"su,omitempty"`
	TargetName         string  `json:"targetname,omitempty"`
	TD                 string  `json:"td,omitempty"`
	TU                 string  `json:"tu,omitempty"`
}

type PolicyExpression struct {
	Builtin               []string `json:"builtin,omitempty"`
	ClientSecurityMessage string   `json:"clientsecuritymessage,omitempty"`
	Comment               string   `json:"comment,omitempty"`
	Count                 float64  `json:"__count,omitempty"`
	Feature               string   `json:"feature,omitempty"`
	Hits                  int      `json:"hits,omitempty"`
	IsDefault             bool     `json:"isdefault,omitempty"`
	Name                  string   `json:"name,omitempty"`
	NextGenAPIResource    string   `json:"_nextgenapiresource,omitempty"`
	PIHits                int      `json:"pihits,omitempty"`
	Type1                 string   `json:"type1,omitempty"`
	TypeField             string   `json:"type,omitempty"`
	Value                 string   `json:"value,omitempty"`
}

type PolicyPatset struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Dynamic            string  `json:"dynamic,omitempty"`
	DynamicOnly        bool    `json:"dynamiconly,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PatsetFile         string  `json:"patsetfile,omitempty"`
}

type PolicyParam struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}
