package domain

// AccessData json response
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
	Owner             Owner    `json:"owner"`
	Supplier          Supplier `json:"supplier"`

	Urls struct {
		Search  string `json:"search"`
		Quote   string `json:"quote"`
		Book    string `json:"book"`
		Generic string `json:"generic"`
	} `json:"urls"`

	Parameters []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"parameters"`
}
