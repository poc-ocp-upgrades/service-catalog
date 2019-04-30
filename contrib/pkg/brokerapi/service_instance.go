package brokerapi

type ServiceInstance struct {
	ID			string			`json:"id"`
	DashboardURL		string			`json:"dashboard_url"`
	InternalID		string			`json:"internal_id,omitempty"`
	ServiceID		string			`json:"service_id"`
	PlanID			string			`json:"plan_id"`
	OrganizationGUID	string			`json:"organization_guid"`
	SpaceGUID		string			`json:"space_guid"`
	LastOperation		*LastOperationResponse	`json:"last_operation,omitempty"`
	Parameters		map[string]interface{}	`json:"parameters,omitempty"`
}
type CreateServiceInstanceRequest struct {
	OrgID			string			`json:"organization_guid,omitempty"`
	PlanID			string			`json:"plan_id,omitempty"`
	ServiceID		string			`json:"service_id,omitempty"`
	SpaceID			string			`json:"space_guid,omitempty"`
	Parameters		map[string]interface{}	`json:"parameters,omitempty"`
	AcceptsIncomplete	bool			`json:"accepts_incomplete,omitempty"`
	ContextProfile		ContextProfile		`json:"context,omitempty"`
}

const ContextProfilePlatformKubernetes string = "kubernetes"

type ContextProfile struct {
	Platform	string	`json:"platform,omitempty"`
	Namespace	string	`json:"namespace,omitempty"`
}
type CreateServiceInstanceResponse struct {
	DashboardURL	string	`json:"dashboard_url,omitempty"`
	Operation	string	`json:"operation,omitempty"`
}
type UpdateServiceInstanceRequest struct {
	ContextProfile	ContextProfile		`json:"context,omitempty"`
	ServiceID	string			`json:"service_id,omitempty"`
	PlanID		string			`json:"plan_id,omitempty"`
	Parameters	map[string]interface{}	`json:"parameters,omitempty"`
	PreviousValues	PreviousValues		`json:"previous_values,omitempty"`
}
type PreviousValues struct {
	OrgID		string	`json:"organization_guid,omitempty"`
	PlanID		string	`json:"plan_id,omitempty"`
	ServiceID	string	`json:"service_id,omitempty"`
	SpaceID		string	`json:"space_guid,omitempty"`
}
type UpdateServiceInstanceResponse struct {
	DashboardURL	string	`json:"dashboard_url,omitempty"`
	Operation	string	`json:"operation,omitempty"`
}
type DeleteServiceInstanceRequest struct {
	ServiceID		string	`json:"service_id"`
	PlanID			string	`json:"plan_id"`
	AcceptsIncomplete	bool	`json:"accepts_incomplete,omitempty"`
}
type DeleteServiceInstanceResponse struct {
	Operation string `json:"operation,omitempty"`
}
type LastOperationRequest struct {
	ServiceID	string	`json:"service_id,omitempty"`
	PlanID		string	`json:"plan_id,omitempty"`
	Operation	string	`json:"operation,omitempty"`
}
type LastOperationResponse struct {
	State		string	`json:"state"`
	Description	string	`json:"description,omitempty"`
}

const (
	StateInProgress	= "in progress"
	StateSucceeded	= "succeeded"
	StateFailed	= "failed"
)
