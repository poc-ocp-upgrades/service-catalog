package brokerapi

type Service struct {
	Name		string		`json:"name"`
	ID		string		`json:"id"`
	Description	string		`json:"description"`
	Tags		[]string	`json:"tags,omitempty"`
	Requires	[]string	`json:"requires,omitempty"`
	Bindable	bool		`json:"bindable"`
	Metadata	interface{}	`json:"metadata,omitempty"`
	DashboardClient	interface{}	`json:"dashboard_client"`
	PlanUpdateable	bool		`json:"plan_updateable,omitempty"`
	Plans		[]ServicePlan	`json:"plans"`
}
