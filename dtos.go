package gopiperun

import "github.com/armando-couto/goutils"

type OportunityResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    []Oportunity `json:"data"`
	Meta    struct {
		Total       int `json:"total"`
		Count       int `json:"count"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
		Links       struct {
			Next string `json:"next"`
		} `json:"links"`
	} `json:"meta"`
}

type OportunityCreateResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    Oportunity `json:"data"`
	Meta    struct {
		Total       int `json:"total"`
		Count       int `json:"count"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
		Links       struct {
			Next string `json:"next"`
		} `json:"links"`
	} `json:"meta"`
}

type Oportunity struct {
	ID                 int                     `json:"id"`
	AccountID          int                     `json:"account_id"`
	Hash               string                  `json:"hash"`
	RdstationReference string                  `json:"rdstation_reference"`
	DataCongelamento   string                  `json:"data_congelamento"`
	TypeReference      int                     `json:"type_reference"`
	Reference          string                  `json:"reference"`
	Temperature        int                     `json:"temperature"`
	Probability        int                     `json:"probability"`
	PipelineID         int                     `json:"pipeline_id"`
	OwnerID            int                     `json:"owner_id"`
	StageID            int                     `json:"stage_id"`
	PersonID           int                     `json:"person_id"`
	CompanyID          int                     `json:"company_id"`
	LostReasonID       int                     `json:"lost_reason_id"`
	OriginID           int                     `json:"origin_id"`
	StartedInStageID   int                     `json:"started_in_stage_id"`
	CityID             int                     `json:"city_id"`
	CreatedAt          string                  `json:"created_at"`
	Title              string                  `json:"title"`
	Description        string                  `json:"description"`
	Observation        string                  `json:"observation"`
	Status             int                     `json:"status"`
	ClosedAt           string                  `json:"closed_at"`
	ReasonClose        string                  `json:"reason_close"`
	Deleted            int                     `json:"deleted"`
	Freezed            int                     `json:"freezed"`
	Value              goutils.KeepZero        `json:"value"`
	Order              int                     `json:"order"`
	UpdatedAt          string                  `json:"updated_at"`
	LastStageUpdatedAt string                  `json:"last_stage_updated_at"`
	ValueMrr           goutils.KeepZero        `json:"value_mrr"`
	ProbablyClosedAt   string                  `json:"probably_closed_at"`
	LastContactAt      string                  `json:"last_contact_at"`
	StageChangedAt     string                  `json:"stage_changed_at"`
	CustomFields       []OportunityCustomField `json:"custom_fields"`
}

type OportunityCustomField struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	Hash            string      `json:"hash"`
	Type            int         `json:"type"`
	Belongs         int         `json:"belongs"`
	Values          []string    `json:"values"`
	RawValue        string      `json:"raw_value"`
	Formula         interface{} `json:"formula"`
	OutputType      interface{} `json:"output_type"`
	DecimalPlaces   int         `json:"decimal_places"`
	AllowNegative   interface{} `json:"allow_negative"`
	CurrencyID      interface{} `json:"currency_id"`
	ThousandSep     int         `json:"thousand_sep"`
	Options         []string    `json:"options"`
	SelectedOptions []string    `json:"selected_options"`
}

type CompanyResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Company `json:"data"`
	Meta    struct {
		Total       int `json:"total"`
		Count       int `json:"count"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
	} `json:"meta"`
}

type CompanyCreateResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    Company `json:"data"`
	Meta    struct {
		Total       int `json:"total"`
		Count       int `json:"count"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
		Links       struct {
			Next string `json:"next"`
		} `json:"links"`
	} `json:"meta"`
}

type Company struct {
	ID                  int         `json:"id"`
	AccountID           int         `json:"account_id"`
	CnaeID              int         `json:"cnae_id"`
	CityID              int         `json:"city_id"`
	ManagerID           int         `json:"manager_id"`
	SegmentID           interface{} `json:"segment_id"`
	Name                string      `json:"name"`
	Cnpj                string      `json:"cnpj"`
	Website             interface{} `json:"website"`
	EmailNf             interface{} `json:"email_nf"`
	Logo                interface{} `json:"logo"`
	Hash                string      `json:"hash"`
	Observation         interface{} `json:"observation"`
	AddressPostalCode   string      `json:"address_postal_code"`
	Address             string      `json:"address"`
	AddressNumber       string      `json:"address_number"`
	AddressComplement   interface{} `json:"address_complement"`
	CompanyName         string      `json:"company_name"`
	Ie                  interface{} `json:"ie"`
	District            string      `json:"district"`
	Country             interface{} `json:"country"`
	Facebook            interface{} `json:"facebook"`
	Linkedin            interface{} `json:"linkedin"`
	Cep                 string      `json:"cep"`
	CustomerAt          interface{} `json:"customer_at"`
	FoundationAt        string      `json:"foundation_at"`
	CreatedAt           string      `json:"created_at"`
	CompanyType         interface{} `json:"company_type"`
	CompanyStatus       interface{} `json:"company_status"`
	CompanySituation    interface{} `json:"company_situation"`
	StatusTouch         interface{} `json:"status_touch"`
	NpsScore            interface{} `json:"nps_score"`
	SocialCapital       int         `json:"social_capital"`
	Lat                 interface{} `json:"lat"`
	Lng                 interface{} `json:"lng"`
	Status              bool        `json:"status"`
	ExternalCode        interface{} `json:"external_code"`
	ForeignContact      int         `json:"foreign_contact"`
	Size                interface{} `json:"size"`
	UpdatedAt           string      `json:"updated_at"`
	IsBrand             bool        `json:"is_brand"`
	IsSupplier          bool        `json:"is_supplier"`
	IsClient            bool        `json:"is_client"`
	IsCarrier           bool        `json:"is_carrier"`
	IsFranchise         bool        `json:"is_franchise"`
	IsChannel           bool        `json:"is_channel"`
	OwnerID             interface{} `json:"owner_id"`
	ManagerFieldSalesID interface{} `json:"manager_field_sales_id"`
	IsDistributor       bool        `json:"is_distributor"`
	IsManufacturer      bool        `json:"is_manufacturer"`
	IsPartner           bool        `json:"is_partner"`
}

type PersonResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    []Person `json:"data"`
	Meta    struct {
		Total       int `json:"total"`
		Count       int `json:"count"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
	} `json:"meta"`
}

type PersonCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Person `json:"data"`
	Meta    struct {
		Total       int `json:"total"`
		Count       int `json:"count"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
		Links       struct {
			Next string `json:"next"`
		} `json:"links"`
	} `json:"meta"`
}

type Person struct {
	ID                int         `json:"id"`
	AccountID         int         `json:"account_id"`
	CompanyID         int         `json:"company_id"`
	CityID            int         `json:"city_id"`
	OwnerID           interface{} `json:"owner_id"`
	ManagerID         interface{} `json:"manager_id"`
	Cpf               string      `json:"cpf"`
	Avatar            interface{} `json:"avatar"`
	Name              string      `json:"name"`
	Hash              string      `json:"hash"`
	Website           interface{} `json:"website"`
	JobTitle          interface{} `json:"job_title"`
	Gender            interface{} `json:"gender"`
	BirthDay          interface{} `json:"birth_day"`
	Observation       string      `json:"observation"`
	CustomerAt        interface{} `json:"customer_at"`
	Facebook          interface{} `json:"facebook"`
	Linkedin          interface{} `json:"linkedin"`
	AddressPostalCode string      `json:"address_postal_code"`
	Address           string      `json:"address"`
	AddressNumber     string      `json:"address_number"`
	AddressComplement string      `json:"address_complement"`
	District          string      `json:"district"`
	CreatedAt         string      `json:"created_at"`
	Lat               interface{} `json:"lat"`
	Lng               interface{} `json:"lng"`
	ExternalCode      interface{} `json:"external_code"`
	Status            bool        `json:"status"`
	ForeignContact    int         `json:"foreign_contact"`
	UpdatedAt         string      `json:"updated_at"`
}

type CityResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []City `json:"data"`
	Meta    struct {
		Total       int           `json:"total"`
		Count       int           `json:"count"`
		PerPage     int           `json:"per_page"`
		CurrentPage int           `json:"current_page"`
		TotalPages  int           `json:"total_pages"`
		Links       []interface{} `json:"links"`
	} `json:"meta"`
}

type City struct {
	ID   int    `json:"id"`
	Uf   string `json:"uf"`
	Name string `json:"name"`
}

type UserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []User `json:"data"`
	Meta    struct {
		Total       int           `json:"total"`
		Count       int           `json:"count"`
		PerPage     int           `json:"per_page"`
		CurrentPage int           `json:"current_page"`
		TotalPages  int           `json:"total_pages"`
		Links       []interface{} `json:"links"`
	} `json:"meta"`
}

type User struct {
	ID                   int         `json:"id"`
	Email                string      `json:"email"`
	Cpf                  string      `json:"cpf"`
	Avatar               string      `json:"avatar"`
	Hash                 string      `json:"hash"`
	SessionLifetime      int         `json:"session_lifetime"`
	OutlookToken         interface{} `json:"outlook_token"`
	OutlookCalendar      interface{} `json:"outlook_calendar"`
	OutlookCalendarEmail interface{} `json:"outlook_calendar_email"`
	SessionID            string      `json:"session_id"`
	TimeLoginStart       interface{} `json:"time_login_start"`
	TimeLoginEnd         interface{} `json:"time_login_end"`
	DaysLoginAllowed     interface{} `json:"days_login_allowed"`
	AccountID            int         `json:"account_id"`
	PipelineID           int         `json:"pipeline_id"`
	RegionID             interface{} `json:"region_id"`
	Name                 string      `json:"name"`
	Gender               interface{} `json:"gender"`
	BirthDay             string      `json:"birth_day"`
	Active               int         `json:"active"`
	Signature            interface{} `json:"signature"`
	URLPublicCalendar    interface{} `json:"url_public_calendar"`
	Telephone            string      `json:"telephone"`
	Cellphone            string      `json:"cellphone"`
	SelectedEmail        interface{} `json:"selected_email"`
	ValidatedEmail       interface{} `json:"validated_email"`
	PipelineView         int         `json:"pipeline_view"`
	OnlyYoursDeals       int         `json:"only_yours_deals"`
	LastLoginAt          string      `json:"last_login_at"`
	CreatedAt            string      `json:"created_at"`
	InboundNotify        int         `json:"inbound_notify"`
	ScheduleActive       int         `json:"schedule_active"`
	ScheduleDuration     int         `json:"schedule_duration"`
	Permission           string      `json:"permission"`
	CreatedBy            int         `json:"created_by"`
	TimezoneActive       int         `json:"timezone_active"`
	Timezone             interface{} `json:"timezone"`
	AppearInURL          interface{} `json:"appear_in_url"`
	NameTag              string      `json:"name_tag"`
	GoogleToken          interface{} `json:"google_token"`
	GoogleCalendar       interface{} `json:"google_calendar"`
	AllowsSupportChat    bool        `json:"allows_support_chat"`
}

type FileCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    File   `json:"data"`
}

type File struct {
	ID              int    `json:"id"`
	AccountID       int    `json:"account_id"`
	UserId          int    `json:"user_id"`
	DealId          string `json:"deal_id"`
	EmailId         int    `json:"email_id"`
	TemplateEmailId int    `json:"template_email_id"`
	ActivityId      int    `json:"activity_id"`
	CallId          int    `json:"call_id"`
	Name            string `json:"name"`
	Url             string `json:"url"`
	UrlAws          string `json:"url_aws"`
	Format          string `json:"format"`
	Description     string `json:"description"`
	Hash            string `json:"hash"`
	CreatedAt       string `json:"created_at"`
	PublicFormFile  string `json:"public_form_file"`
	Size            int    `json:"size"`
}
