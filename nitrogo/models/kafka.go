package models

// kafka configuration structs
type Kafkacluster struct {
	Activesvc          int     `json:"activesvc,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Numtopics          int     `json:"numtopics,omitempty"`
	Topicname          string  `json:"topicname,omitempty"`
	Totalsvc           int     `json:"totalsvc,omitempty"`
}

type KafkaclusterBinding struct {
	KafkaclusterServicegroupBinding []interface{} `json:"kafkacluster_servicegroup_binding,omitempty"`
	Name                            string        `json:"name,omitempty"`
}

type KafkaclusterServicegroupBinding struct {
	Name             string `json:"name,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}
