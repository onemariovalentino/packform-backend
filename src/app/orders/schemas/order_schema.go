package schemas

type (
	OrderRequest struct {
		Search        string `json:"search"`
		StartDate     string `json:"start_date"`
		EndDate       string `json:"end_date"`
		SortDirection string `json:"sort_direction"`
		Page          int    `json:"page"`
		PerPage       int    `json:"per_page"`
	}
)
