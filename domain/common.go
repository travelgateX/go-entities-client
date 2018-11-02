package domain

type Response struct {
	Data  Data    `json:"data"`
	Error []Error `json:"error"`
}

type Data struct {
	Admin Admin `json:"admin"`
}

type Admin struct {
	Accesses Edges `json:"accesses"`
}

type Edges struct {
	Edges []Edge `json:"edges"`
}

type Edge struct {
	Node Node `json:"node"`
}

type Node struct {
	Code       string     `json:"code"`
	AccessData AccessData `json:"accessData"`
	Errors     []Error    `json:"error"`
	CreatedAt  string     `json:"createdAt"`
	UpdatedAt  string     `json:"updatedAt"`
}

type PageInfo struct {
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

type Error []struct {
	Code        string `json:"code"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
