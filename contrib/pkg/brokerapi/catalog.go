package brokerapi

type Catalog struct {
	Services []*Service `json:"services"`
}
