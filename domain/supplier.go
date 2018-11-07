package domain

// Supplier json response
type Supplier struct {
	Code         string `json:"code"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	SupplierData struct {
		Code          string `json:"code"`
		Name          string `json:"name"`
		IsActive      bool   `json:"isActive"`
		Context       string `json:"context"`
		ServiceAPI    int    `json:"serviceApi"`
		SupplierGroup string `json:"supplierGroup"`
	} `json:"supplierData"`
}
