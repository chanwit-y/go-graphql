package model

type (
	CustomerRes struct {
		CustomerID  int    `json:"customer_id"`
		Name        string `json:"name"`
		DateOfBirth string `json:"date_of_birth"`
		City        string `json:"city"`
		ZipCode     string `json:"zip_code"`
		Status      int    `json:"status"`
	}
)
