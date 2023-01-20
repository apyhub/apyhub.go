package helper

// commons request and response structure
type Request struct {
	Url string `json:"url"`
}

type Response struct {
	Data string `json:"data"`
}

type WatermarkRequest struct {
	ImageUrl     string
	WatermarkUrl string
}

type Watermark struct {
	Image     interface{}
	Watermark interface{}
}

type ErrJsn struct {
	Error string `json:"error"`
}

type WrongJsn struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// Country Code
type CountryResponse struct {
	Data Country `json:"data"`
}

type CountryListResponse struct {
	Data []Country `json:"data"`
}

type Country struct {
	Value        string   `json:"value"`
	Key          string   `json:"key"`
	Cca3         string   `json:"cca3"`
	Emoji        string   `json:"emoji"`
	Calling_code []int    `json:"calling_code"`
	Subdivision  []subdiv `json:"subdivision"`
}

type subdiv struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// currency struct
type CurrencyConverter struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Date   string `json:"date,omitempty"`
}

type CurrencyResponse struct {
	Data float64 `json:"data"`
}

type CurrencyList struct {
	Data []Currency `json:"data"`
}

type Currency struct {
	Emoji string `json:"emoji"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Timezone struct
type TimezoneResponse struct {
	Data []Timezone `json:"data"`
}

type Timezone struct {
	Key          string   `json:"key"`
	Value        string   `json:"value"`
	Abbreviation []string `json:"abbreviation"`
	UTC_time     string   `json:"utc_time"`
}

// Fuzzy search
type FuzzyRequest struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type FuzzyResponse struct {
	Data []string `json:"data"`
}

// Metadata struct
type MetaResponse struct {
	Data map[string]interface{} `json:"data"`
}

// Charts struct
type Chart struct {
	Title string   `json:"title"`
	Theme string   `json:"theme"`
	Data  []Values `json:"data"`
}

type StackedBarChart struct {
	Title string         `json:"title"`
	Theme string         `json:"theme"`
	Data  []StackedValue `json:"data"`
}

type StackedValue struct {
	Name   string   `json:"name"`
	Values []Values `json:"values"`
}

type Values struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

type CalenderEvent struct {
	Summary         string    `json:"summary"`
	Description     string    `json:"description"`
	OrganizerEmail  string    `json:"organizer_email"`
	AttendeesEmails []string  `json:"attendees_emails"`
	Location        string    `json:"location"`
	TimeZone        string    `json:"time_zone"`
	MeetingDate     string    `json:"meeting_date"`
	StartTime       string    `json:"start_time"`
	EndTime         string    `json:"end_time"`
	Reccuring       bool      `json:"recurring"`
	Recurrence      Reccurent `json:"recurrence"`
}

type Reccurent struct {
	Frequency string `json:"frequency"`
	Interval  int    `json:"interval"`
	Count     int    `json:"count"`
}

type ZipSecure struct {
	Urls     []string `json:"urls"`
	Password string   `json:"password"`
}

type Zip struct {
	Urls []string `json:"urls"`
}

type ContentRequest struct {
	Content string `json:"content"`
}

type EmailRequest struct {
	Email string
}

type PostRequest struct {
	Postcode string
}

type VatRequest struct {
	Vat string
}
