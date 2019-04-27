package brokerapi

type ServiceBinding struct {
	ID			string			`json:"id"`
	ServiceID		string			`json:"service_id"`
	AppID			string			`json:"app_id"`
	ServicePlanID		string			`json:"service_plan_id"`
	PrivateKey		string			`json:"private_key"`
	ServiceInstanceID	string			`json:"service_instance_id"`
	BindResource		map[string]interface{}	`json:"bind_resource,omitempty"`
	Parameters		map[string]interface{}	`json:"parameters,omitempty"`
}
type BindingRequest struct {
	AppGUID		string			`json:"app_guid,omitempty"`
	PlanID		string			`json:"plan_id,omitempty"`
	ServiceID	string			`json:"service_id,omitempty"`
	BindResource	map[string]interface{}	`json:"bind_resource,omitempty"`
	Parameters	map[string]interface{}	`json:"parameters,omitempty"`
}
type CreateServiceBindingResponse struct {
	Credentials Credential `json:"credentials"`
}
type Credential map[string]interface{}
