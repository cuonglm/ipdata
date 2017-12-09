package ipdata

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	apiURL            = "https://api.ipdata.co"
	acceptContentType = "application/json"
	userAgent         = "goipdata/0.1"
)

// Client represents ipdata client to interact with server
type Client struct {
	APIKey     string
	Language   string
	httpClient *http.Client
	url        string
}

// WithLanguage specifies the language
func WithLanguage(lang string) Option {
	return func(i *Client) {
		i.Language = lang
	}
}

// WithAPIKey specifies api key to use
func WithAPIKey(apiKey string) Option {
	return func(i *Client) {
		i.APIKey = apiKey
	}
}

// WithURL specifies server api url
func WithURL(url string) Option {
	return func(i *Client) {
		i.url = url
	}
}

// ResponseData represents data response from ip data server
type ResponseData struct {
	IP             string  `json:"ip"`
	City           string  `json:"city"`
	Region         string  `json:"region"`
	CountryName    string  `json:"country_name"`
	CountryCode    string  `json:"country_code"`
	ContinentName  string  `json:"continent_name"`
	ContinentCode  string  `json:"continent_code"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	ASN            string  `json:"asn"`
	Organisation   string  `json:"organisation"`
	Postal         string  `json:"postal"`
	Currency       string  `json:"currency"`
	CurrencySymbol string  `json:"currency_symbol"`
	CallingCode    string  `json:"calling_code"`
	Flag           string  `json:"flag"`
	TimeZone       string  `json:"time_zone"`
}

// Option configures how we construct Client
type Option func(*Client)

// NewClient returns new Client instance with given options
func NewClient(options ...Option) *Client {
	i := &Client{
		httpClient: http.DefaultClient,
		url:        apiURL,
		Language:   "en",
	}
	for _, option := range options {
		option(i)
	}

	return i
}

// Lookup gathers information about given ip
func (c *Client) Lookup(ip string) (*ResponseData, error) {
	req, _ := http.NewRequest("GET", strings.Join([]string{c.url, ip, c.Language}, "/"), nil)
	req.Header.Add("Accept", acceptContentType)
	req.Header.Add("user-agent", userAgent)
	if c.APIKey != "" {
		req.Header.Add("api-key", c.APIKey)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	r := &ResponseData{}
	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		return nil, err
	}

	return r, nil
}
