package models

// app configuration structs
type Application struct {
	AppName             string `json:"appname,omitempty"`
	AppTemplateFileName string `json:"apptemplatefilename,omitempty"`
	DeploymentFileName  string `json:"deploymentfilename,omitempty"`
}
