package models

// cache configuration structs
type CachePolicyCSVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CacheGlobalBinding struct {
	CacheGlobalCachePolicyBinding []any `json:"cacheglobal_cachepolicy_binding,omitempty"`
}

type CachePolicyCachePolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CachePolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Flags              int      `json:"flags,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	InvalGroups        []string `json:"invalgroups,omitempty"`
	InvalObjects       []string `json:"invalobjects,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	PolicyName         string   `json:"policyname,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	StoreInGroup       string   `json:"storeingroup,omitempty"`
	UndefAction        string   `json:"undefaction,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type CachePolicyCacheGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CachePolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CacheParameter struct {
	CacheEvictionPolicy string `json:"cacheevictionpolicy,omitempty"`
	DiskLimit           int    `json:"disklimit,omitempty"`
	EnableBypass        string `json:"enablebypass,omitempty"`
	EnableHAObjPersist  string `json:"enablehaobjpersist,omitempty"`
	MaxDiskLimit        int    `json:"maxdisklimit,omitempty"`
	MaxMemLimit         int    `json:"maxmemlimit,omitempty"`
	MaxPostLen          int    `json:"maxpostlen,omitempty"`
	MemLimit            int    `json:"memlimit,omitempty"`
	MemLimitActive      int    `json:"memlimitactive,omitempty"`
	NextGenAPIResource  string `json:"_nextgenapiresource,omitempty"`
	PrefetchCur         int    `json:"prefetchcur,omitempty"`
	PrefetchMaxPending  int    `json:"prefetchmaxpending,omitempty"`
	UndefAction         string `json:"undefaction,omitempty"`
	VerifyUsing         string `json:"verifyusing,omitempty"`
	Via                 string `json:"via,omitempty"`
}

type CachePolicyLabelPolicyBindingBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CachePolicyBinding struct {
	CachePolicyCacheGlobalBinding      []any  `json:"cachepolicy_cacheglobal_binding,omitempty"`
	CachePolicyCachePolicyLabelBinding []any  `json:"cachepolicy_cachepolicylabel_binding,omitempty"`
	CachePolicyCSVServerBinding        []any  `json:"cachepolicy_csvserver_binding,omitempty"`
	CachePolicyLBVServerBinding        []any  `json:"cachepolicy_lbvserver_binding,omitempty"`
	PolicyName                         string `json:"policyname,omitempty"`
}

type CacheSelector struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Flags              int      `json:"flags,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               []string `json:"rule,omitempty"`
	SelectorName       string   `json:"selectorname,omitempty"`
}

type CachePolicyLabel struct {
	Builtin                []string `json:"builtin,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Evaluates              string   `json:"evaluates,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	FlowType               int      `json:"flowtype,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	Hits                   int      `json:"hits,omitempty"`
	InvokeLabelName        string   `json:"invoke_labelname,omitempty"`
	LabelName              string   `json:"labelname,omitempty"`
	LabelType              string   `json:"labeltype,omitempty"`
	NewName                string   `json:"newname,omitempty"`
	NextGenAPIResource     string   `json:"_nextgenapiresource,omitempty"`
	NumPol                 int      `json:"numpol,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type CacheGlobalCachePolicyBinding struct {
	FlowType               int    `json:"flowtype,omitempty"`
	GlobalBindType         string `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	NumPol                 int    `json:"numpol,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	PrecedeDefRules        string `json:"precededefrules,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type CacheObject struct {
	CacheCellAppFWMetadataExists string   `json:"cachecellappfwmetadataexists,omitempty"`
	CacheCellBaseFile            string   `json:"cachecellbasefile,omitempty"`
	CacheCellComplex             string   `json:"cachecellcomplex,omitempty"`
	CacheCellCompressionFormat   string   `json:"cachecellcompressionformat,omitempty"`
	CacheCellCurMisses           int      `json:"cachecellcurmisses,omitempty"`
	CacheCellCurReaders          int      `json:"cachecellcurreaders,omitempty"`
	CacheCellDestIPVerified      string   `json:"cachecelldestipverified,omitempty"`
	CacheCellDHits               int      `json:"cachecelldhits,omitempty"`
	CacheCellETagInserted        string   `json:"cachecelletaginserted,omitempty"`
	CacheCellExpires             int      `json:"cachecellexpires,omitempty"`
	CacheCellExpiresMillisec     int      `json:"cachecellexpiresmillisec,omitempty"`
	CacheCellFWPxyObj            string   `json:"cachecellfwpxyobj,omitempty"`
	CacheCellHits                int      `json:"cachecellhits,omitempty"`
	CacheCellHTTP11              string   `json:"cachecellhttp11,omitempty"`
	CacheCellMinHit              int      `json:"cachecellminhit,omitempty"`
	CacheCellMinHitFlag          string   `json:"cachecellminhitflag,omitempty"`
	CacheCellMisses              int      `json:"cachecellmisses,omitempty"`
	CacheCellPollEveryTime       string   `json:"cachecellpolleverytime,omitempty"`
	CacheCellReadyWithLastByte   string   `json:"cachecellreadywithlastbyte,omitempty"`
	CacheCellReqTime             int      `json:"cachecellreqtime,omitempty"`
	CacheCellResBadSize          string   `json:"cachecellresbadsize,omitempty"`
	CacheCellResTime             int      `json:"cachecellrestime,omitempty"`
	CacheCellWeakETag            string   `json:"cachecellweaketag,omitempty"`
	CacheControl                 string   `json:"cachecontrol,omitempty"`
	CacheCurAge                  int      `json:"cachecurage,omitempty"`
	CacheDirName                 string   `json:"cachedirname,omitempty"`
	CacheETag                    string   `json:"cacheetag,omitempty"`
	CacheFileName                string   `json:"cachefilename,omitempty"`
	CacheInDisk                  string   `json:"cacheindisk,omitempty"`
	CacheInMemory                string   `json:"cacheinmemory,omitempty"`
	CacheInSecondary             string   `json:"cacheinsecondary,omitempty"`
	CacheResDate                 string   `json:"cacheresdate,omitempty"`
	CacheResHdrSize              int      `json:"cachereshdrsize,omitempty"`
	CacheResLastMod              string   `json:"cachereslastmod,omitempty"`
	CacheResSize                 int      `json:"cacheressize,omitempty"`
	CacheURLs                    string   `json:"cacheurls,omitempty"`
	CEFlags                      int      `json:"ceflags,omitempty"`
	ContentGroup                 string   `json:"contentgroup,omitempty"`
	Count                        float64  `json:"__count,omitempty"`
	DestIPv46                    string   `json:"destipv46,omitempty"`
	DestPort                     int      `json:"destport,omitempty"`
	Flushed                      string   `json:"flushed,omitempty"`
	Group                        string   `json:"group,omitempty"`
	GroupName                    string   `json:"groupname,omitempty"`
	HitParams                    []string `json:"hitparams,omitempty"`
	HitValues                    []string `json:"hitvalues,omitempty"`
	Host                         string   `json:"host,omitempty"`
	HTTPCalloutCell              string   `json:"httpcalloutcell,omitempty"`
	HTTPCalloutName              string   `json:"httpcalloutname,omitempty"`
	HTTPCalloutResult            string   `json:"httpcalloutresult,omitempty"`
	HTTPMethod                   string   `json:"httpmethod,omitempty"`
	HTTPStatus                   int      `json:"httpstatus,omitempty"`
	HTTPStatusOutput             int      `json:"httpstatusoutput,omitempty"`
	IgnoreMarkerObjects          string   `json:"ignoremarkerobjects,omitempty"`
	IncludeNotReadyObjects       string   `json:"includenotreadyobjects,omitempty"`
	Locator                      int      `json:"locator,omitempty"`
	LocatorShow                  int      `json:"locatorshow,omitempty"`
	MarkerReason                 string   `json:"markerreason,omitempty"`
	NextGenAPIResource           string   `json:"_nextgenapiresource,omitempty"`
	NodeID                       int      `json:"nodeid,omitempty"`
	Policy                       int      `json:"policy,omitempty"`
	PolicyName                   string   `json:"policyname,omitempty"`
	Port                         int      `json:"port,omitempty"`
	Prefetch                     string   `json:"prefetch,omitempty"`
	PrefetchPeriod               int      `json:"prefetchperiod,omitempty"`
	PrefetchPeriodMillisec       int      `json:"prefetchperiodmillisec,omitempty"`
	ReturnType                   string   `json:"returntype,omitempty"`
	Rule                         []string `json:"rule,omitempty"`
	SelectorName                 []string `json:"selectorname,omitempty"`
	SelectorValue                []string `json:"selectorvalue,omitempty"`
	ToSecondary                  string   `json:"tosecondary,omitempty"`
	TotalObjs                    int      `json:"totalobjs,omitempty"`
	URL                          string   `json:"url,omitempty"`
	WarnBucketSkip               int      `json:"warnbucketskip,omitempty"`
}

type CachePolicyLabelBinding struct {
	CachePolicyLabelCachePolicyBinding   []any  `json:"cachepolicylabel_cachepolicy_binding,omitempty"`
	CachePolicyLabelPolicyBindingBinding []any  `json:"cachepolicylabel_policybinding_binding,omitempty"`
	LabelName                            string `json:"labelname,omitempty"`
}

type CachePolicyLabelCachePolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CacheContentGroup struct {
	AbsExpiry              []string `json:"absexpiry,omitempty"`
	AbsExpiryGMT           []string `json:"absexpirygmt,omitempty"`
	AlwaysEvalPolicies     string   `json:"alwaysevalpolicies,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Cache304Hits           int      `json:"cache304hits,omitempty"`
	CacheCells             int      `json:"cachecells,omitempty"`
	CacheControl           string   `json:"cachecontrol,omitempty"`
	CacheGroupIncarnation  int      `json:"cachegroupincarnation,omitempty"`
	CacheNon304Hits        int      `json:"cachenon304hits,omitempty"`
	CacheNumInvalPolicy    int      `json:"cachenuminvalpolicy,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	DiskLimit              int      `json:"disklimit,omitempty"`
	ExpireAtLastByte       string   `json:"expireatlastbyte,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Flags                  int      `json:"flags,omitempty"`
	FlashCache             string   `json:"flashcache,omitempty"`
	HeurExpiryParam        int      `json:"heurexpiryparam,omitempty"`
	HitParams              []string `json:"hitparams,omitempty"`
	HitSelector            string   `json:"hitselector,omitempty"`
	Host                   string   `json:"host,omitempty"`
	IgnoreParamValueCase   string   `json:"ignoreparamvaluecase,omitempty"`
	IgnoreReloadReq        string   `json:"ignorereloadreq,omitempty"`
	IgnoreReqCachingHdrs   string   `json:"ignorereqcachinghdrs,omitempty"`
	InsertAge              string   `json:"insertage,omitempty"`
	InsertETag             string   `json:"insertetag,omitempty"`
	InsertVia              string   `json:"insertvia,omitempty"`
	InvalParams            []string `json:"invalparams,omitempty"`
	InvalRestrictedToHost  string   `json:"invalrestrictedtohost,omitempty"`
	InvalSelector          string   `json:"invalselector,omitempty"`
	LazyDNSResolve         string   `json:"lazydnsresolve,omitempty"`
	MarkerCells            int      `json:"markercells,omitempty"`
	MatchCookies           string   `json:"matchcookies,omitempty"`
	MaxResSize             int      `json:"maxressize,omitempty"`
	MemDUsage              int      `json:"memdusage,omitempty"`
	MemLimit               int      `json:"memlimit,omitempty"`
	MemUsage               int      `json:"memusage,omitempty"`
	MinHits                int      `json:"minhits,omitempty"`
	MinResSize             int      `json:"minressize,omitempty"`
	Name                   string   `json:"name,omitempty"`
	NextGenAPIResource     string   `json:"_nextgenapiresource,omitempty"`
	Persist                string   `json:"persist,omitempty"`
	PersistHA              string   `json:"persistha,omitempty"`
	Pinned                 string   `json:"pinned,omitempty"`
	PolicyName             []string `json:"policyname,omitempty"`
	PollEveryTime          string   `json:"polleverytime,omitempty"`
	Prefetch               string   `json:"prefetch,omitempty"`
	PrefetchCur            int      `json:"prefetchcur,omitempty"`
	PrefetchMaxPending     int      `json:"prefetchmaxpending,omitempty"`
	PrefetchPeriod         int      `json:"prefetchperiod,omitempty"`
	PrefetchPeriodMillisec int      `json:"prefetchperiodmillisec,omitempty"`
	Query                  string   `json:"query,omitempty"`
	QuickAbortSize         int      `json:"quickabortsize,omitempty"`
	RelExpiry              int      `json:"relexpiry,omitempty"`
	RelExpiryMillisec      int      `json:"relexpirymillisec,omitempty"`
	RemoveCookies          string   `json:"removecookies,omitempty"`
	SelectorValue          string   `json:"selectorvalue,omitempty"`
	ToSecondary            string   `json:"tosecondary,omitempty"`
	TypeField              string   `json:"type,omitempty"`
	WeakNegRelExpiry       int      `json:"weaknegrelexpiry,omitempty"`
	WeakPosRelExpiry       int      `json:"weakposrelexpiry,omitempty"`
}

type CacheForwardProxy struct {
	Count              float64 `json:"__count,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
}
