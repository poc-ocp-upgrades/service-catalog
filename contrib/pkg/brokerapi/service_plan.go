package brokerapi

type ServicePlan struct {
	Name		string		`json:"name"`
	ID		string		`json:"id"`
	Description	string		`json:"description"`
	Metadata	interface{}	`json:"metadata,omitempty"`
	Free		bool		`json:"free,omitempty"`
	Bindable	*bool		`json:"bindable,omitempty"`
	Schemas		*Schemas	`json:"schemas,omitempty"`
}
