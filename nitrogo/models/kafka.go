package models

// kafka configuration structs
type KafkaCluster struct {
	ActiveSvc          int     `json:"activesvc,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NumTopics          int     `json:"numtopics,omitempty"`
	TopicName          string  `json:"topicname,omitempty"`
	TotalSvc           int     `json:"totalsvc,omitempty"`
}

type KafkaClusterBinding struct {
	KafkaClusterServiceGroupBinding []any  `json:"kafkacluster_servicegroup_binding,omitempty"`
	Name                            string `json:"name,omitempty"`
}

type KafkaClusterServiceGroupBinding struct {
	Name             string `json:"name,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
}
