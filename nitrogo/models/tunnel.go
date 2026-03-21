package models

// tunnel configuration structs
type TunnelTrafficPolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	ClientTransactions int      `json:"clienttransactions,omitempty"`
	ClientTTLB         int      `json:"clientttlb,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	ExpressionType     string   `json:"expressiontype,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	IsDefault          bool     `json:"isdefault,omitempty"`
	LogAction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	RxBytes            int      `json:"rxbytes,omitempty"`
	ServerTransactions int      `json:"servertransactions,omitempty"`
	ServerTTLB         int      `json:"serverttlb,omitempty"`
	TxBytes            int      `json:"txbytes,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type TunnelGlobalBinding struct {
	TunnelGlobalTunnelTrafficPolicyBinding []any `json:"tunnelglobal_tunneltrafficpolicy_binding,omitempty"`
}

type TunnelTrafficPolicyBinding struct {
	Name                                   string `json:"name,omitempty"`
	TunnelTrafficPolicyTunnelGlobalBinding []any  `json:"tunneltrafficpolicy_tunnelglobal_binding,omitempty"`
}

type TunnelTrafficPolicyTunnelGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type TunnelGlobalTunnelTrafficPolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	GlobalBindType         string   `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	NumPol                 int      `json:"numpol,omitempty"`
	PolicyName             string   `json:"policyname,omitempty"`
	PolicyType             string   `json:"policytype,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	State                  string   `json:"state,omitempty"`
	TypeField              string   `json:"type,omitempty"`
}
