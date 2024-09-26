package dto

type (
	NewCalenderItemReq struct {
		Title     string `json:"title" validate:"required"`
		StartDate string `json:"startDate" validate:"required"`
		EndDate   string `json:"endDate" validate:"required"`
		Content   string `json:"content" validate:"required"`
	}

	GetCalenderItemReq struct {
		ID string `param:"id" validate:"required"`
	}

	UpdateCalenderItemReq struct {
		ID string `json:"id" validate:"required"`
	}

	DeleteCalenderItemReq struct {
		ID string `json:"id" validate:"required"`
	}
)
