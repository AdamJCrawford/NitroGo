package models

// tunnel configuration structs
type Tunneltrafficpolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Clienttransactions int      `json:"clienttransactions,omitempty"`
	Clientttlb         int      `json:"clientttlb,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Expressiontype     string   `json:"expressiontype,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Isdefault          bool     `json:"isdefault,omitempty"`
	Logaction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Rxbytes            int      `json:"rxbytes,omitempty"`
	Servertransactions int      `json:"servertransactions,omitempty"`
	Serverttlb         int      `json:"serverttlb,omitempty"`
	Txbytes            int      `json:"txbytes,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type TunnelglobalBinding struct {
	TunnelglobalTunneltrafficpolicyBinding []interface{} `json:"tunnelglobal_tunneltrafficpolicy_binding,omitempty"`
}

type TunneltrafficpolicyBinding struct {
	Name                                   string        `json:"name,omitempty"`
	TunneltrafficpolicyTunnelglobalBinding []interface{} `json:"tunneltrafficpolicy_tunnelglobal_binding,omitempty"`
}

type TunneltrafficpolicyTunnelglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type TunnelglobalTunneltrafficpolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Numpol                 int      `json:"numpol,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Policytype             string   `json:"policytype,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	State                  string   `json:"state,omitempty"`
	TypeField              string   `json:"type,omitempty"`
}
