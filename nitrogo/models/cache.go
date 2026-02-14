// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Cacheglobalbinding struct {
}

type Cacheglobalcachepolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Flowtype               int    `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Precededefrules        string `json:"precededefrules,omitempty"`
}

type Cacheobject struct {
	Url                          string `json:"url,omitempty"`
	Locator                      int    `json:"locator,omitempty"`
	Httpstatus                   int    `json:"httpstatus,omitempty"`
	Host                         string `json:"host,omitempty"`
	Port                         int    `json:"port,omitempty"`
	Groupname                    string `json:"groupname,omitempty"`
	Httpmethod                   string `json:"httpmethod,omitempty"`
	Group                        string `json:"group,omitempty"`
	Ignoremarkerobjects          string `json:"ignoremarkerobjects,omitempty"`
	Includenotreadyobjects       string `json:"includenotreadyobjects,omitempty"`
	Nodeid                       int    `json:"nodeid,omitempty"`
	Tosecondary                  string `json:"tosecondary,omitempty"`
	Cacheressize                 string `json:"cacheressize,omitempty"`
	Cachereshdrsize              string `json:"cachereshdrsize,omitempty"`
	Cacheetag                    string `json:"cacheetag,omitempty"`
	Httpstatusoutput             string `json:"httpstatusoutput,omitempty"`
	Cachereslastmod              string `json:"cachereslastmod,omitempty"`
	Cachecontrol                 string `json:"cachecontrol,omitempty"`
	Cacheresdate                 string `json:"cacheresdate,omitempty"`
	Contentgroup                 string `json:"contentgroup,omitempty"`
	Destipv46                    string `json:"destipv46,omitempty"`
	Destport                     string `json:"destport,omitempty"`
	Cachecellcomplex             string `json:"cachecellcomplex,omitempty"`
	Hitparams                    string `json:"hitparams,omitempty"`
	Hitvalues                    string `json:"hitvalues,omitempty"`
	Cachecellreqtime             string `json:"cachecellreqtime,omitempty"`
	Cachecellrestime             string `json:"cachecellrestime,omitempty"`
	Cachecurage                  string `json:"cachecurage,omitempty"`
	Cachecellexpires             string `json:"cachecellexpires,omitempty"`
	Cachecellexpiresmillisec     string `json:"cachecellexpiresmillisec,omitempty"`
	Flushed                      string `json:"flushed,omitempty"`
	Prefetch                     string `json:"prefetch,omitempty"`
	Prefetchperiod               string `json:"prefetchperiod,omitempty"`
	Prefetchperiodmillisec       string `json:"prefetchperiodmillisec,omitempty"`
	Cachecellcurreaders          string `json:"cachecellcurreaders,omitempty"`
	Cachecellcurmisses           string `json:"cachecellcurmisses,omitempty"`
	Cachecellhits                string `json:"cachecellhits,omitempty"`
	Cachecellmisses              string `json:"cachecellmisses,omitempty"`
	Cachecelldhits               string `json:"cachecelldhits,omitempty"`
	Cachecellcompressionformat   string `json:"cachecellcompressionformat,omitempty"`
	Cachecellappfwmetadataexists string `json:"cachecellappfwmetadataexists,omitempty"`
	Cachecellhttp11              string `json:"cachecellhttp11,omitempty"`
	Cachecellweaketag            string `json:"cachecellweaketag,omitempty"`
	Cachecellresbadsize          string `json:"cachecellresbadsize,omitempty"`
	Markerreason                 string `json:"markerreason,omitempty"`
	Cachecellpolleverytime       string `json:"cachecellpolleverytime,omitempty"`
	Cachecelletaginserted        string `json:"cachecelletaginserted,omitempty"`
	Cachecellreadywithlastbyte   string `json:"cachecellreadywithlastbyte,omitempty"`
	Cacheinmemory                string `json:"cacheinmemory,omitempty"`
	Cacheindisk                  string `json:"cacheindisk,omitempty"`
	Cacheinsecondary             string `json:"cacheinsecondary,omitempty"`
	Cachedirname                 string `json:"cachedirname,omitempty"`
	Cachefilename                string `json:"cachefilename,omitempty"`
	Cachecelldestipverified      string `json:"cachecelldestipverified,omitempty"`
	Cachecellfwpxyobj            string `json:"cachecellfwpxyobj,omitempty"`
	Cachecellbasefile            string `json:"cachecellbasefile,omitempty"`
	Cachecellminhitflag          string `json:"cachecellminhitflag,omitempty"`
	Cachecellminhit              string `json:"cachecellminhit,omitempty"`
	Policy                       string `json:"policy,omitempty"`
	Policyname                   string `json:"policyname,omitempty"`
	Selectorname                 string `json:"selectorname,omitempty"`
	Rule                         string `json:"rule,omitempty"`
	Selectorvalue                string `json:"selectorvalue,omitempty"`
	Cacheurls                    string `json:"cacheurls,omitempty"`
	Warnbucketskip               string `json:"warnbucketskip,omitempty"`
	Totalobjs                    string `json:"totalobjs,omitempty"`
	Httpcalloutcell              string `json:"httpcalloutcell,omitempty"`
	Httpcalloutname              string `json:"httpcalloutname,omitempty"`
	Returntype                   string `json:"returntype,omitempty"`
	Httpcalloutresult            string `json:"httpcalloutresult,omitempty"`
	Locatorshow                  string `json:"locatorshow,omitempty"`
	Ceflags                      string `json:"ceflags,omitempty"`
	Nextgenapiresource           string `json:"_nextgenapiresource,omitempty"`
}

type Cacheparameter struct {
	Memlimit            int    `json:"memlimit,omitempty"`
	Via                 string `json:"via,omitempty"`
	Verifyusing         string `json:"verifyusing,omitempty"`
	Maxpostlen          int    `json:"maxpostlen"`
	Prefetchmaxpending  int    `json:"prefetchmaxpending,omitempty"`
	Enablebypass        string `json:"enablebypass,omitempty"`
	Undefaction         string `json:"undefaction,omitempty"`
	Enablehaobjpersist  string `json:"enablehaobjpersist,omitempty"`
	Cacheevictionpolicy string `json:"cacheevictionpolicy,omitempty"`
	Disklimit           string `json:"disklimit,omitempty"`
	Maxdisklimit        string `json:"maxdisklimit,omitempty"`
	Memlimitactive      string `json:"memlimitactive,omitempty"`
	Maxmemlimit         string `json:"maxmemlimit,omitempty"`
	Prefetchcur         string `json:"prefetchcur,omitempty"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
}

type Cachepolicy struct {
	Policyname         string   `json:"policyname,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Action             string   `json:"action,omitempty"`
	Storeingroup       string   `json:"storeingroup,omitempty"`
	Invalgroups        []string `json:"invalgroups,omitempty"`
	Invalobjects       []string `json:"invalobjects,omitempty"`
	Undefaction        string   `json:"undefaction,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Hits               string   `json:"hits,omitempty"`
	Undefhits          string   `json:"undefhits,omitempty"`
	Flags              string   `json:"flags,omitempty"`
	Builtin            string   `json:"builtin,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Cachepolicybinding struct {
	Policyname string `json:"policyname,omitempty"`
}

type Cachepolicycacheglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
}

type Cachepolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}

type Cachecontentgroup struct {
	Name                   string   `json:"name,omitempty"`
	Weakposrelexpiry       int      `json:"weakposrelexpiry"`
	Heurexpiryparam        int      `json:"heurexpiryparam"`
	Relexpiry              int      `json:"relexpiry"`
	Relexpirymillisec      int      `json:"relexpirymillisec"`
	Absexpiry              []string `json:"absexpiry,omitempty"`
	Absexpirygmt           []string `json:"absexpirygmt,omitempty"`
	Weaknegrelexpiry       int      `json:"weaknegrelexpiry"`
	Hitparams              []string `json:"hitparams,omitempty"`
	Invalparams            []string `json:"invalparams,omitempty"`
	Ignoreparamvaluecase   string   `json:"ignoreparamvaluecase,omitempty"`
	Matchcookies           string   `json:"matchcookies,omitempty"`
	Invalrestrictedtohost  string   `json:"invalrestrictedtohost,omitempty"`
	Polleverytime          string   `json:"polleverytime,omitempty"`
	Ignorereloadreq        string   `json:"ignorereloadreq,omitempty"`
	Removecookies          string   `json:"removecookies,omitempty"`
	Prefetch               string   `json:"prefetch,omitempty"`
	Prefetchperiod         int      `json:"prefetchperiod"`
	Prefetchperiodmillisec int      `json:"prefetchperiodmillisec"`
	Prefetchmaxpending     int      `json:"prefetchmaxpending"`
	Flashcache             string   `json:"flashcache,omitempty"`
	Expireatlastbyte       string   `json:"expireatlastbyte,omitempty"`
	Insertvia              string   `json:"insertvia,omitempty"`
	Insertage              string   `json:"insertage,omitempty"`
	Insertetag             string   `json:"insertetag,omitempty"`
	Cachecontrol           string   `json:"cachecontrol,omitempty"`
	Quickabortsize         int      `json:"quickabortsize"`
	Minressize             int      `json:"minressize"`
	Maxressize             int      `json:"maxressize,omitempty"`
	Memlimit               int      `json:"memlimit,omitempty"`
	Ignorereqcachinghdrs   string   `json:"ignorereqcachinghdrs,omitempty"`
	Minhits                int      `json:"minhits"`
	Alwaysevalpolicies     string   `json:"alwaysevalpolicies,omitempty"`
	Persistha              string   `json:"persistha,omitempty"`
	Pinned                 string   `json:"pinned,omitempty"`
	Lazydnsresolve         string   `json:"lazydnsresolve,omitempty"`
	Hitselector            string   `json:"hitselector,omitempty"`
	Invalselector          string   `json:"invalselector,omitempty"`
	Type                   string   `json:"type,omitempty"`
	Query                  string   `json:"query,omitempty"`
	Host                   string   `json:"host,omitempty"`
	Selectorvalue          string   `json:"selectorvalue,omitempty"`
	Tosecondary            string   `json:"tosecondary,omitempty"`
	Flags                  string   `json:"flags,omitempty"`
	Prefetchcur            string   `json:"prefetchcur,omitempty"`
	Memusage               string   `json:"memusage,omitempty"`
	Memdusage              string   `json:"memdusage,omitempty"`
	Disklimit              string   `json:"disklimit,omitempty"`
	Cachenon304hits        string   `json:"cachenon304hits,omitempty"`
	Cache304hits           string   `json:"cache304hits,omitempty"`
	Cachecells             string   `json:"cachecells,omitempty"`
	Cachegroupincarnation  string   `json:"cachegroupincarnation,omitempty"`
	Persist                string   `json:"persist,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Cachenuminvalpolicy    string   `json:"cachenuminvalpolicy,omitempty"`
	Markercells            string   `json:"markercells,omitempty"`
	Builtin                string   `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
}

type Cacheforwardproxy struct {
	Ipaddress          string `json:"ipaddress,omitempty"`
	Port               int    `json:"port,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Cachepolicycsvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
}

type Cachepolicylbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
}

type Cachepolicypolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
}

type Cachepolicylabel struct {
	Labelname              string `json:"labelname,omitempty"`
	Evaluates              string `json:"evaluates,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Priority               string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Flowtype               string `json:"flowtype,omitempty"`
	Builtin                string `json:"builtin,omitempty"`
	Feature                string `json:"feature,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Cachepolicylabelpolicybindingbinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Cacheselector struct {
	Selectorname       string   `json:"selectorname,omitempty"`
	Rule               []string `json:"rule,omitempty"`
	Flags              string   `json:"flags,omitempty"`
	Builtin            string   `json:"builtin,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Cacheglobalpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Numpol                 uint32 `json:"numpol,omitempty"`
	Flowtype               uint32 `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Precededefrules        string `json:"precededefrules,omitempty"`
}

type Cachepolicycachepolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
}

type Cachepolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
}

type Cachepolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
}

type Cachepolicylabelcachepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Cachepolicylabelpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}
