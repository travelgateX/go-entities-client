package domain

// Response json response
type Response struct {
	Data  Data    `json:"data"`
	Error []Error `json:"error"`
}

// Data json response
type Data struct {
	Admin Admin `json:"admin"`
}

// Admin json response
type Admin struct {
	Accesses Edges `json:"accesses"`
}

// Edges json response
type Edges struct {
	Edges []Edge `json:"edges"`
}

// Edge json response
type Edge struct {
	Node Node `json:"node"`
}

// Node json response
type Node struct {
	Code       string     `json:"code"`
	AccessData AccessData `json:"accessData"`
	Errors     []Error    `json:"error"`
	CreatedAt  string     `json:"createdAt"`
	UpdatedAt  string     `json:"updatedAt"`
}

// PageInfo json response
type PageInfo struct {
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

// Error json response
type Error []struct {
	Code        string `json:"code"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// Owner json response
type Owner struct {
	Code      string `json:"code"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Error     []struct {
		Code        string `json:"code"`
		Description string `json:"description"`
		Type        string `json:"type"`
	} `json:"error"`
}

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
