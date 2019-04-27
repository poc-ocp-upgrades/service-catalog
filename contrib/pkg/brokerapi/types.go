package brokerapi

type Types struct {
	Instance	string	`json:"instance"`
	Binding		string	`json:"binding"`
}

const (
	InstanceType	= "instanceType"
	BindingType	= "bindingType"
)
