package models

// cache configuration structs
type CachepolicyCsvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CacheglobalBinding struct {
	CacheglobalCachepolicyBinding []interface{} `json:"cacheglobal_cachepolicy_binding,omitempty"`
}

type CachepolicyCachepolicylabelBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Cachepolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Flags              int      `json:"flags,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Invalgroups        []string `json:"invalgroups,omitempty"`
	Invalobjects       []string `json:"invalobjects,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Policyname         string   `json:"policyname,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Storeingroup       string   `json:"storeingroup,omitempty"`
	Undefaction        string   `json:"undefaction,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type CachepolicyCacheglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CachepolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Cacheparameter struct {
	Cacheevictionpolicy string `json:"cacheevictionpolicy,omitempty"`
	Disklimit           int    `json:"disklimit,omitempty"`
	Enablebypass        string `json:"enablebypass,omitempty"`
	Enablehaobjpersist  string `json:"enablehaobjpersist,omitempty"`
	Maxdisklimit        int    `json:"maxdisklimit,omitempty"`
	Maxmemlimit         int    `json:"maxmemlimit,omitempty"`
	Maxpostlen          int    `json:"maxpostlen,omitempty"`
	Memlimit            int    `json:"memlimit,omitempty"`
	Memlimitactive      int    `json:"memlimitactive,omitempty"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
	Prefetchcur         int    `json:"prefetchcur,omitempty"`
	Prefetchmaxpending  int    `json:"prefetchmaxpending,omitempty"`
	Undefaction         string `json:"undefaction,omitempty"`
	Verifyusing         string `json:"verifyusing,omitempty"`
	Via                 string `json:"via,omitempty"`
}

type CachepolicylabelPolicybindingBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CachepolicyBinding struct {
	CachepolicyCacheglobalBinding      []interface{} `json:"cachepolicy_cacheglobal_binding,omitempty"`
	CachepolicyCachepolicylabelBinding []interface{} `json:"cachepolicy_cachepolicylabel_binding,omitempty"`
	CachepolicyCsvserverBinding        []interface{} `json:"cachepolicy_csvserver_binding,omitempty"`
	CachepolicyLbvserverBinding        []interface{} `json:"cachepolicy_lbvserver_binding,omitempty"`
	Policyname                         string        `json:"policyname,omitempty"`
}

type Cacheselector struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Flags              int      `json:"flags,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               []string `json:"rule,omitempty"`
	Selectorname       string   `json:"selectorname,omitempty"`
}

type Cachepolicylabel struct {
	Builtin                []string `json:"builtin,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Evaluates              string   `json:"evaluates,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Flowtype               int      `json:"flowtype,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Hits                   int      `json:"hits,omitempty"`
	InvokeLabelname        string   `json:"invoke_labelname,omitempty"`
	Labelname              string   `json:"labelname,omitempty"`
	Labeltype              string   `json:"labeltype,omitempty"`
	Newname                string   `json:"newname,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
	Numpol                 int      `json:"numpol,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type CacheglobalCachepolicyBinding struct {
	Flowtype               int    `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Precededefrules        string `json:"precededefrules,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type Cacheobject struct {
	Cachecellappfwmetadataexists string   `json:"cachecellappfwmetadataexists,omitempty"`
	Cachecellbasefile            string   `json:"cachecellbasefile,omitempty"`
	Cachecellcomplex             string   `json:"cachecellcomplex,omitempty"`
	Cachecellcompressionformat   string   `json:"cachecellcompressionformat,omitempty"`
	Cachecellcurmisses           int      `json:"cachecellcurmisses,omitempty"`
	Cachecellcurreaders          int      `json:"cachecellcurreaders,omitempty"`
	Cachecelldestipverified      string   `json:"cachecelldestipverified,omitempty"`
	Cachecelldhits               int      `json:"cachecelldhits,omitempty"`
	Cachecelletaginserted        string   `json:"cachecelletaginserted,omitempty"`
	Cachecellexpires             int      `json:"cachecellexpires,omitempty"`
	Cachecellexpiresmillisec     int      `json:"cachecellexpiresmillisec,omitempty"`
	Cachecellfwpxyobj            string   `json:"cachecellfwpxyobj,omitempty"`
	Cachecellhits                int      `json:"cachecellhits,omitempty"`
	Cachecellhttp11              string   `json:"cachecellhttp11,omitempty"`
	Cachecellminhit              int      `json:"cachecellminhit,omitempty"`
	Cachecellminhitflag          string   `json:"cachecellminhitflag,omitempty"`
	Cachecellmisses              int      `json:"cachecellmisses,omitempty"`
	Cachecellpolleverytime       string   `json:"cachecellpolleverytime,omitempty"`
	Cachecellreadywithlastbyte   string   `json:"cachecellreadywithlastbyte,omitempty"`
	Cachecellreqtime             int      `json:"cachecellreqtime,omitempty"`
	Cachecellresbadsize          string   `json:"cachecellresbadsize,omitempty"`
	Cachecellrestime             int      `json:"cachecellrestime,omitempty"`
	Cachecellweaketag            string   `json:"cachecellweaketag,omitempty"`
	Cachecontrol                 string   `json:"cachecontrol,omitempty"`
	Cachecurage                  int      `json:"cachecurage,omitempty"`
	Cachedirname                 string   `json:"cachedirname,omitempty"`
	Cacheetag                    string   `json:"cacheetag,omitempty"`
	Cachefilename                string   `json:"cachefilename,omitempty"`
	Cacheindisk                  string   `json:"cacheindisk,omitempty"`
	Cacheinmemory                string   `json:"cacheinmemory,omitempty"`
	Cacheinsecondary             string   `json:"cacheinsecondary,omitempty"`
	Cacheresdate                 string   `json:"cacheresdate,omitempty"`
	Cachereshdrsize              int      `json:"cachereshdrsize,omitempty"`
	Cachereslastmod              string   `json:"cachereslastmod,omitempty"`
	Cacheressize                 int      `json:"cacheressize,omitempty"`
	Cacheurls                    string   `json:"cacheurls,omitempty"`
	Ceflags                      int      `json:"ceflags,omitempty"`
	Contentgroup                 string   `json:"contentgroup,omitempty"`
	Count                        float64  `json:"__count,omitempty"`
	Destipv46                    string   `json:"destipv46,omitempty"`
	Destport                     int      `json:"destport,omitempty"`
	Flushed                      string   `json:"flushed,omitempty"`
	Group                        string   `json:"group,omitempty"`
	Groupname                    string   `json:"groupname,omitempty"`
	Hitparams                    []string `json:"hitparams,omitempty"`
	Hitvalues                    []string `json:"hitvalues,omitempty"`
	Host                         string   `json:"host,omitempty"`
	Httpcalloutcell              string   `json:"httpcalloutcell,omitempty"`
	Httpcalloutname              string   `json:"httpcalloutname,omitempty"`
	Httpcalloutresult            string   `json:"httpcalloutresult,omitempty"`
	Httpmethod                   string   `json:"httpmethod,omitempty"`
	Httpstatus                   int      `json:"httpstatus,omitempty"`
	Httpstatusoutput             int      `json:"httpstatusoutput,omitempty"`
	Ignoremarkerobjects          string   `json:"ignoremarkerobjects,omitempty"`
	Includenotreadyobjects       string   `json:"includenotreadyobjects,omitempty"`
	Locator                      int      `json:"locator,omitempty"`
	Locatorshow                  int      `json:"locatorshow,omitempty"`
	Markerreason                 string   `json:"markerreason,omitempty"`
	Nextgenapiresource           string   `json:"_nextgenapiresource,omitempty"`
	Nodeid                       int      `json:"nodeid,omitempty"`
	Policy                       int      `json:"policy,omitempty"`
	Policyname                   string   `json:"policyname,omitempty"`
	Port                         int      `json:"port,omitempty"`
	Prefetch                     string   `json:"prefetch,omitempty"`
	Prefetchperiod               int      `json:"prefetchperiod,omitempty"`
	Prefetchperiodmillisec       int      `json:"prefetchperiodmillisec,omitempty"`
	Returntype                   string   `json:"returntype,omitempty"`
	Rule                         []string `json:"rule,omitempty"`
	Selectorname                 []string `json:"selectorname,omitempty"`
	Selectorvalue                []string `json:"selectorvalue,omitempty"`
	Tosecondary                  string   `json:"tosecondary,omitempty"`
	Totalobjs                    int      `json:"totalobjs,omitempty"`
	Url                          string   `json:"url,omitempty"`
	Warnbucketskip               int      `json:"warnbucketskip,omitempty"`
}

type CachepolicylabelBinding struct {
	CachepolicylabelCachepolicyBinding   []interface{} `json:"cachepolicylabel_cachepolicy_binding,omitempty"`
	CachepolicylabelPolicybindingBinding []interface{} `json:"cachepolicylabel_policybinding_binding,omitempty"`
	Labelname                            string        `json:"labelname,omitempty"`
}

type CachepolicylabelCachepolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Cachecontentgroup struct {
	Absexpiry              []string `json:"absexpiry,omitempty"`
	Absexpirygmt           []string `json:"absexpirygmt,omitempty"`
	Alwaysevalpolicies     string   `json:"alwaysevalpolicies,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Cache304hits           int      `json:"cache304hits,omitempty"`
	Cachecells             int      `json:"cachecells,omitempty"`
	Cachecontrol           string   `json:"cachecontrol,omitempty"`
	Cachegroupincarnation  int      `json:"cachegroupincarnation,omitempty"`
	Cachenon304hits        int      `json:"cachenon304hits,omitempty"`
	Cachenuminvalpolicy    int      `json:"cachenuminvalpolicy,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Disklimit              int      `json:"disklimit,omitempty"`
	Expireatlastbyte       string   `json:"expireatlastbyte,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Flags                  int      `json:"flags,omitempty"`
	Flashcache             string   `json:"flashcache,omitempty"`
	Heurexpiryparam        int      `json:"heurexpiryparam,omitempty"`
	Hitparams              []string `json:"hitparams,omitempty"`
	Hitselector            string   `json:"hitselector,omitempty"`
	Host                   string   `json:"host,omitempty"`
	Ignoreparamvaluecase   string   `json:"ignoreparamvaluecase,omitempty"`
	Ignorereloadreq        string   `json:"ignorereloadreq,omitempty"`
	Ignorereqcachinghdrs   string   `json:"ignorereqcachinghdrs,omitempty"`
	Insertage              string   `json:"insertage,omitempty"`
	Insertetag             string   `json:"insertetag,omitempty"`
	Insertvia              string   `json:"insertvia,omitempty"`
	Invalparams            []string `json:"invalparams,omitempty"`
	Invalrestrictedtohost  string   `json:"invalrestrictedtohost,omitempty"`
	Invalselector          string   `json:"invalselector,omitempty"`
	Lazydnsresolve         string   `json:"lazydnsresolve,omitempty"`
	Markercells            int      `json:"markercells,omitempty"`
	Matchcookies           string   `json:"matchcookies,omitempty"`
	Maxressize             int      `json:"maxressize,omitempty"`
	Memdusage              int      `json:"memdusage,omitempty"`
	Memlimit               int      `json:"memlimit,omitempty"`
	Memusage               int      `json:"memusage,omitempty"`
	Minhits                int      `json:"minhits,omitempty"`
	Minressize             int      `json:"minressize,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
	Persist                string   `json:"persist,omitempty"`
	Persistha              string   `json:"persistha,omitempty"`
	Pinned                 string   `json:"pinned,omitempty"`
	Policyname             []string `json:"policyname,omitempty"`
	Polleverytime          string   `json:"polleverytime,omitempty"`
	Prefetch               string   `json:"prefetch,omitempty"`
	Prefetchcur            int      `json:"prefetchcur,omitempty"`
	Prefetchmaxpending     int      `json:"prefetchmaxpending,omitempty"`
	Prefetchperiod         int      `json:"prefetchperiod,omitempty"`
	Prefetchperiodmillisec int      `json:"prefetchperiodmillisec,omitempty"`
	Query                  string   `json:"query,omitempty"`
	Quickabortsize         int      `json:"quickabortsize,omitempty"`
	Relexpiry              int      `json:"relexpiry,omitempty"`
	Relexpirymillisec      int      `json:"relexpirymillisec,omitempty"`
	Removecookies          string   `json:"removecookies,omitempty"`
	Selectorvalue          string   `json:"selectorvalue,omitempty"`
	Tosecondary            string   `json:"tosecondary,omitempty"`
	TypeField              string   `json:"type,omitempty"`
	Weaknegrelexpiry       int      `json:"weaknegrelexpiry,omitempty"`
	Weakposrelexpiry       int      `json:"weakposrelexpiry,omitempty"`
}

type Cacheforwardproxy struct {
	Count              float64 `json:"__count,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
}
