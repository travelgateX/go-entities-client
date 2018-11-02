package domain

type Accesses struct {
	Edges []struct {
		Node struct {
			Code       string     `json:"code"`
			CreatedAt  string     `json:"createdAt"`
			UpdatedAt  string     `json:"updatedAt"`
			AccessData AccessData `json:"accessData"`
			Error      []Error    `json:"errors"`
		} `json:"node"`
	} `json:"edges"`
	PageInfo PageInfo `json:"pageInfo"`
}

type AccessData struct {
	Name              string   `json:"name"`
	IsActive          bool     `json:"isActive"`
	Code              string   `json:"code"`
	IsTest            bool     `json:"isTest"`
	User              string   `json:"user"`
	Password          string   `json:"password"`
	Markets           []string `json:"markets"`
	RateRules         []string `json:"rateRules"`
	IsSchedulerActive bool     `json:"isSchedulerActive"`
	Urls              struct {
		Search  string `json:"search"`
		Quote   string `json:"quote"`
		Book    string `json:"book"`
		Generic string `json:"generic"`
	} `json:"urls"`
	Parameters []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"parameters"`
	Owner struct {
		Code      string `json:"code"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
		Error     []struct {
			Code        string `json:"code"`
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"error"`
	} `json:"owner"`
	Supplier struct {
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
	} `json:"supplier"`
}
