package domain

// Response json response
type Response struct {
	Data struct {
		Admin struct {
			DeleteAccessFromGroup struct {
				AccessData struct {
					Name              string `json:"name"`
					Code              string `json:"code"`
					IsActive          bool   `json:"isActive"`
					IsTest            bool   `json:"isTest"`
					IsSchedulerActive bool   `json:"isSchedulerActive"`
					Owner             struct {
						Code string `json:"code"`
					} `json:"owner"`
					Supplier interface{} `json:"supplier"`
				} `json:"accessData"`
				Error interface{} `json:"error"`
			} `json:"deleteAccessFromGroup"`
		} `json:"admin"`
	} `json:"data"`
	Errors []struct {
		Message string   `json:"message"`
		Path    []string `json:"path"`
	} `json:"errors"`
}

// Data json response
type Data struct {
	Admin Admin `json:"admin"`
}

// Admin json response
type Admin struct {
	// Querys
	Accesses Edges `json:"accesses"`

	// Mutations
	GrantAccessToGroup    GrantAccessToGroup    `json:"grantAccessToGroup"`
	DeleteAccessFromGroup DeleteAccessFromGroup `json:"deleteAccessFromGroup"`
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
	Errors     Error      `json:"error"`
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

// DataError json response
type DataError []struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}

// Owner json response
type Owner struct {
	Code      string `json:"code"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Errors    Error  `json:"error"`
}
