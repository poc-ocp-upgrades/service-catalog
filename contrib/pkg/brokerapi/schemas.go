package brokerapi

type Schemas struct {
	ServiceInstance	*ServiceInstanceSchema	`json:"service_instance,omitempty"`
	ServiceBinding	*ServiceBindingSchema	`json:"service_binding,omitempty"`
}
type ServiceInstanceSchema struct {
	Create	*InputParametersSchema	`json:"create,omitempty"`
	Update	*InputParametersSchema	`json:"update,omitempty"`
}
type ServiceBindingSchema struct {
	Create *RequestResponseSchema `json:"create,omitempty"`
}
type InputParametersSchema struct {
	Parameters interface{} `json:"parameters,omitempty"`
}
type RequestResponseSchema struct {
	InputParametersSchema
	Response	interface{}	`json:"response,omitempty"`
}
