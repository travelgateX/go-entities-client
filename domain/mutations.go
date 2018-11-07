package domain

// GrantAccessToGroup json response
type GrantAccessToGroup struct {
	AccessData interface{} `json:"accessData"`
	Errors     Error       `json:"error"`
}

// DeleteAccessFromGroup json response
type DeleteAccessFromGroup struct {
	AccessData interface{} `json:"accessData"`
	Errors     Error       `json:"error"`
}
