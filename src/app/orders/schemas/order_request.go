package schemas

type (
	OrderRequest struct {
		Search    string `json:"search,omitempty"`
		StartDate string `json:"start_date,omitempty"`
		EndDate   string `json:"end_date,omitempty"`
	}
)
