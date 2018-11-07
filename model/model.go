package model

import (
	fmt "fmt"
	io "io"
	strconv "strconv"
)

type Access struct {
	Code       string     `json:"code"`
	AccessData AccessData `json:"accessData"`
	Error      []Error    `json:"error"`
	CreatedAt  string     `json:"createdAt"`
	UpdatedAt  string     `json:"updatedAt"`
}

func (Access) IsNode() {}

type AccessConnection struct {
	Edges    []AccessEdge `json:"edges"`
	PageInfo PageInfo     `json:"pageInfo"`
}

type AccessData struct {
	Name              string          `json:"name"`
	IsActive          bool            `json:"isActive"`
	Code              string          `json:"code"`
	Supplier          Supplier        `json:"supplier"`
	IsTest            bool            `json:"isTest"`
	User              string          `json:"user"`
	Password          string          `json:"password"`
	Urls              Urls            `json:"urls"`
	Parameters        []Parameter     `json:"parameters"`
	Markets           []string        `json:"markets"`
	RateRules         []RateRulesType `json:"rateRules"`
	Owner             Organization    `json:"owner"`
	IsSchedulerActive bool            `json:"isSchedulerActive"`
	Groups            GroupConnection `json:"groups"`
	// Shared            Access          `json:"shared"` // invalid recursive type
}

func (AccessData) IsEntityData() {}

type AccessEdge struct {
	Cursor string `json:"cursor"`
	Node   Access `json:"node"`
}

type AccessFilter struct {
	AccessID []string `json:"accessID"`
	Group    []string `json:"group"`
	Name     []string `json:"name"`
	Owner    []string `json:"owner"`
	IsActive bool     `json:"isActive"`
}

type AccessInput struct {
	Name              string           `json:"name"`
	IsActive          bool             `json:"isActive"`
	Code              string           `json:"code"`
	Supplier          string           `json:"supplier"`
	IsTest            bool             `json:"isTest"`
	User              string           `json:"user"`
	Password          string           `json:"password"`
	Urls              UrlsInput        `json:"urls"`
	Parameters        []ParameterInput `json:"parameters"`
	Markets           []string         `json:"markets"`
	RateRules         []RateRulesType  `json:"rateRules"`
	Shared            string           `json:"shared"`
	Group             string           `json:"group"`
	Owner             string           `json:"owner"`
	IsSchedulerActive bool             `json:"isSchedulerActive"`
}

type AdminQuery struct {
	Query Query `json:"admin"`
}

type AdminMutation struct {
	Mutation Mutation `json:"admin"`
}

type Mutation struct {
	CreateAccess              Access   `json:"createAccess"`
	UpdateAccess              Access   `json:"updateAccess"`
	GrantAccessToGroup        Access   `json:"grantAccessToGroup"`
	DeleteAccessFromGroup     Access   `json:"deleteAccessFromGroup"`
	GrantSupplierToGroup      Supplier `json:"grantSupplierToGroup"`
	DeleteSupplierFromGroup   Supplier `json:"deleteSupplierFromGroup"`
	GrantClientToGroup        Client   `json:"grantClientToGroup"`
	DeleteClientFromGroup     Client   `json:"deleteClientFromGroup"`
	CreateClient              Client   `json:"createClient"`
	UpdateClient              Client   `json:"updateClient"`
	CreateProfile             Profile  `json:"createProfile"`
	UpdateProfile             Profile  `json:"updateProfile"`
	AddEntitiesToProfile      Profile  `json:"addEntitiesToProfile"`
	RemoveEntitiesFromProfile Profile  `json:"removeEntitiesFromProfile"`
}

type Query struct {
	Accesses     AccessConnection      `json:"accesses"`
	Suppliers    SupplierConnection    `json:"suppliers"`
	Clients      ClientConnection      `json:"clients"`
	ServiceAPI   ServiceApi            `json:"serviceApi"`
	PointsOfSale PointOfSaleConnection `json:"pointsOfSale"`
	Profiles     ProfileConnection     `json:"profiles"`
	Entities     EntityConnection      `json:"entities"`
}

type AdviseMessage interface {
	IsAdviseMessage()
}

type Client struct {
	Code       string     `json:"code"`
	ClientData ClientData `json:"clientData"`
	Error      []Error    `json:"error"`
	CreatedAt  string     `json:"createdAt"`
	UpdatedAt  string     `json:"updatedAt"`
}

func (Client) IsNode() {}

type ClientConnection struct {
	Edges    []ClientEdge `json:"edges"`
	PageInfo PageInfo     `json:"pageInfo"`
}

type ClientData struct {
	Code     string       `json:"code"`
	Name     string       `json:"name"`
	IsActive bool         `json:"isActive"`
	Group    Group        `json:"group"`
	Owner    Organization `json:"owner"`
}

func (ClientData) IsEntityData() {}

type ClientEdge struct {
	Cursor string `json:"cursor"`
	Node   Client `json:"node"`
}

type ClientFilter struct {
	ClientID []string `json:"clientID"`
	Name     []string `json:"name"`
	GroupID  []string `json:"groupID"`
	IsActive bool     `json:"isActive"`
	Owner    []string `json:"owner"`
}

type Context struct {
	Code      string  `json:"code"`
	Error     []Error `json:"error"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

func (Context) IsNode() {}

type CreateClientInput struct {
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
	Group    string `json:"group"`
	Owner    string `json:"owner"`
}

type CreateProfileInput struct {
	Label       string        `json:"label"`
	Group       string        `json:"group"`
	Type        ProfileType   `json:"type"`
	Entities    EntitiesInput `json:"entities"`
	Owner       string        `json:"owner"`
	IsActive    bool          `json:"isActive"`
	IsPublished bool          `json:"isPublished"`
}

type EntitiesInput struct {
	Suppliers []string `json:"suppliers"`
	Clients   []string `json:"clients"`
	Accesses  []string `json:"accesses"`
}

type Entity struct {
	Code          string          `json:"code"`
	Entity        EntityData      `json:"entity"`
	AdviseMessage []AdviseMessage `json:"adviseMessage"`
	CreatedAt     string          `json:"createdAt"`
	UpdatedAt     string          `json:"updatedAt"`
}

func (Entity) IsNode() {}

type EntityConnection struct {
	Edges    []EntityEdge `json:"edges"`
	PageInfo PageInfo     `json:"pageInfo"`
}

type EntityData interface {
	IsEntityData()
}

type EntityEdge struct {
	Cursor string `json:"cursor"`
	Node   Entity `json:"node"`
}

type EntityFilter struct {
	Codes    []string `json:"codes"`
	Groups   []string `json:"groups"`
	Owner    []string `json:"owner"`
	IsActive bool     `json:"isActive"`
}

type Error struct {
	Code        string `json:"code"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type ExternalMessage struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Group struct {
	Code      string  `json:"code"`
	Error     []Error `json:"error"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

func (Group) IsNode() {}

type GroupConnection struct {
	Edges    []GroupEdge `json:"edges"`
	PageInfo PageInfo    `json:"pageInfo"`
}

type GroupEdge struct {
	Cursor string `json:"cursor"`
	Node   Group  `json:"node"`
}

type GroupInput struct {
	ID     string   `json:"id"`
	Groups []string `json:"groups"`
}

type Node interface {
	IsNode()
}

type Organization struct {
	Code      string  `json:"code"`
	Error     []Error `json:"error"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

func (Organization) IsNode() {}

type PageInfo struct {
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
}

type Parameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ParameterInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PointOfSale struct {
	Code            string          `json:"code"`
	PointOfSaleData PointOfSaleData `json:"pointOfSaleData"`
	AdviseMessage   []AdviseMessage `json:"adviseMessage"`
	CreatedAt       string          `json:"createdAt"`
	UpdatedAt       string          `json:"updatedAt"`
}

func (PointOfSale) IsNode() {}

type PointOfSaleConnection struct {
	Edges    []PointOfSaleEdge `json:"edges"`
	PageInfo PageInfo          `json:"pageInfo"`
}

type PointOfSaleData struct {
	Code  string       `json:"code"`
	Name  string       `json:"name"`
	Group string       `json:"group"`
	Owner Organization `json:"owner"`
}

type PointOfSaleEdge struct {
	Cursor string      `json:"cursor"`
	Node   PointOfSale `json:"node"`
}

type PointOfSaleFilter struct {
	Code    []string `json:"code"`
	GroupID []string `json:"groupID"`
	Owner   []string `json:"owner"`
}

type Profile struct {
	Code          string          `json:"code"`
	ProfileData   ProfileData     `json:"profileData"`
	AdviseMessage []AdviseMessage `json:"adviseMessage"`
	CreatedAt     string          `json:"createdAt"`
	UpdatedAt     string          `json:"updatedAt"`
}

func (Profile) IsNode() {}

type ProfileConnection struct {
	Edges    []ProfileEdge `json:"edges"`
	PageInfo PageInfo      `json:"pageInfo"`
}

type ProfileData struct {
	Code        string           `json:"code"`
	Label       string           `json:"label"`
	Type        ProfileType      `json:"type"`
	Entities    EntityConnection `json:"entities"`
	Group       Group            `json:"group"`
	Owner       Organization     `json:"owner"`
	IsActive    bool             `json:"isActive"`
	IsPublished bool             `json:"isPublished"`
}

type ProfileEdge struct {
	Cursor string  `json:"cursor"`
	Node   Profile `json:"node"`
}

type ProfileFilter struct {
	ProfileIDs []string `json:"profileIDs"`
	GroupIDs   []string `json:"groupIDs"`
	Owner      []string `json:"owner"`
}

type ProviderData struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
	IsPublic bool   `json:"isPublic"`
}

type ServiceApi struct {
	Code       string             `json:"code"`
	Name       string             `json:"name"`
	Operations []ServiceOperation `json:"operations"`
	Error      []Error            `json:"error"`
}

type ServiceApiFilter struct {
	APIName       string `json:"ApiName"`
	OperationName string `json:"OperationName"`
	OperationType string `json:"OperationType"`
}

type ServiceOperation struct {
	Code            string `json:"code"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	TravelOperation int    `json:"travelOperation"`
}

type Supplier struct {
	Code         string       `json:"code"`
	SupplierData SupplierData `json:"supplierData"`
	Error        []Error      `json:"error"`
	CreatedAt    string       `json:"createdAt"`
	UpdatedAt    string       `json:"updatedAt"`
}

func (Supplier) IsNode() {}

type SupplierConnection struct {
	Edges    []SupplierEdge `json:"edges"`
	PageInfo PageInfo       `json:"pageInfo"`
}

type SupplierData struct {
	Code          string           `json:"code"`
	Name          string           `json:"name"`
	IsActive      bool             `json:"isActive"`
	Provider      ProviderData     `json:"provider"`
	Context       string           `json:"context"`
	ServiceAPI    int              `json:"serviceApi"`
	SupplierGroup string           `json:"supplierGroup"`
	Accesses      AccessConnection `json:"accesses"`
	Owner         Organization     `json:"owner"`
	Groups        GroupConnection  `json:"groups"`
}

func (SupplierData) IsEntityData() {}

type SupplierEdge struct {
	Cursor string   `json:"cursor"`
	Node   Supplier `json:"node"`
}

type SupplierFilter struct {
	SupplierID []string `json:"supplierID"`
	AccessID   []string `json:"accessID"`
	GroupID    []string `json:"groupID"`
	IsActive   bool     `json:"isActive"`
	ServiceAPI []int    `json:"serviceAPI"`
	Owner      []string `json:"owner"`
}

type SupplierGroup struct {
	GroupCode string `json:"groupCode"`
	IsActive  bool   `json:"isActive"`
}

type UpdateClientInput struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	IsActive bool   `json:"isActive"`
	Owner    string `json:"owner"`
}

type UpdateEntitiesInput struct {
	Code     string        `json:"code"`
	Entities EntitiesInput `json:"entities"`
}

type UpdateProfileInput struct {
	Code        string      `json:"code"`
	Label       string      `json:"label"`
	Type        ProfileType `json:"type"`
	Owner       string      `json:"owner"`
	IsActive    bool        `json:"isActive"`
	IsPublished bool        `json:"isPublished"`
}

type Urls struct {
	Search  string `json:"search"`
	Quote   string `json:"quote"`
	Book    string `json:"book"`
	Generic string `json:"generic"`
}

type UrlsInput struct {
	Search  string `json:"search"`
	Quote   string `json:"quote"`
	Book    string `json:"book"`
	Generic string `json:"generic"`
}

type AdviseMessageLevel string

const (
	AdviseMessageLevelWarn  AdviseMessageLevel = "WARN"
	AdviseMessageLevelError AdviseMessageLevel = "ERROR"
	AdviseMessageLevelInfo  AdviseMessageLevel = "INFO"
)

func (e AdviseMessageLevel) IsValid() bool {
	switch e {
	case AdviseMessageLevelWarn, AdviseMessageLevelError, AdviseMessageLevelInfo:
		return true
	}
	return false
}

func (e AdviseMessageLevel) String() string {
	return string(e)
}

func (e AdviseMessageLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	e = AdviseMessageLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AdviseMessageLevel", str)
	}
	return nil
}

func (e AdviseMessageLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProfileType string

const (
	ProfileTypeBuyer  ProfileType = "BUYER"
	ProfileTypeSeller ProfileType = "SELLER"
)

func (e ProfileType) IsValid() bool {
	switch e {
	case ProfileTypeBuyer, ProfileTypeSeller:
		return true
	}
	return false
}

func (e ProfileType) String() string {
	return string(e)
}

func (e ProfileType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	e = ProfileType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProfileType", str)
	}
	return nil
}

func (e ProfileType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RateRulesType string

const (
	RateRulesTypePackage          RateRulesType = "PACKAGE"
	RateRulesTypeOlder55          RateRulesType = "OLDER55"
	RateRulesTypeOlder60          RateRulesType = "OLDER60"
	RateRulesTypeOlder65          RateRulesType = "OLDER65"
	RateRulesTypeCanaryResident   RateRulesType = "CANARY_RESIDENT"
	RateRulesTypeBalearicResident RateRulesType = "BALEARIC_RESIDENT"
	RateRulesTypeLargeFamily      RateRulesType = "LARGE_FAMILY"
	RateRulesTypeHoneymoon        RateRulesType = "HONEYMOON"
	RateRulesTypePublicServant    RateRulesType = "PUBLIC_SERVANT"
	RateRulesTypeUnemployed       RateRulesType = "UNEMPLOYED"
	RateRulesTypeNormal           RateRulesType = "NORMAL"
	RateRulesTypeNonRefundable    RateRulesType = "NON_REFUNDABLE"
)

func (e RateRulesType) IsValid() bool {
	switch e {
	case RateRulesTypePackage, RateRulesTypeOlder55, RateRulesTypeOlder60, RateRulesTypeOlder65, RateRulesTypeCanaryResident, RateRulesTypeBalearicResident, RateRulesTypeLargeFamily, RateRulesTypeHoneymoon, RateRulesTypePublicServant, RateRulesTypeUnemployed, RateRulesTypeNormal, RateRulesTypeNonRefundable:
		return true
	}
	return false
}

func (e RateRulesType) String() string {
	return string(e)
}

func (e RateRulesType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	e = RateRulesType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RateRulesType", str)
	}
	return nil
}

func (e RateRulesType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
