package dto

type (
	ErrorRes struct {
		IsError    bool        `json:"isError"`
		StatusCode int         `json:"statusCode"`
		Error      interface{} `json:"error"`
	}

	NewCalenderItemRes struct {
		IsError bool   `json:"isError"`
		ID      string `json:"id"`
	}

	GetCalenderItemRes struct {
		IsError bool   `json:"isError"`
		ID      string `json:"id"`
	}

	UpdateCalenderItemRes struct {
		IsError    bool   `json:"isError"`
		StatusCode int    `json:"statusCode"`
		ID         string `json:"id"`
	}

	DeleteCalenderItemRes struct {
		IsError    bool   `json:"isError"`
		StatusCode int    `json:"statusCode"`
		ID         string `json:"id"`
	}
)
