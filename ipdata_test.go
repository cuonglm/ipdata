package ipdata

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setUp() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
}

func tearDown() {
	server.Close()
}

func assertEqual(t *testing.T, result interface{}, expect interface{}) {
	if result != expect {
		t.Fatalf("Expect (Value: %v) (Type: %T) - Got (Value: %v) (Type: %T)", expect, expect, result, result)
	}
}

const fakeResponse = `
{
    "ip": "foo",
    "city": "bar",
    "region": "baz",
    "country_name": "",
    "country_code": "",
    "continent_name": "",
    "continent_code": "",
    "latitude": 37.751,
    "longitude": -97.822,
    "asn": "",
    "organisation": "",
    "postal": "",
    "currency": "USD",
    "currency_symbol": "$",
    "calling_code": "1",
    "flag": "https://ipdata.co/flags/us.png",
    "time_zone": ""
}
`

func TestLookup(t *testing.T) {
	setUp()
	defer tearDown()

	mux.HandleFunc("/foo/en", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		w.Header().Add("Content-Type", acceptContentType)
		_, _ = w.Write([]byte(fakeResponse))
	})

	idc := NewClient(WithURL(server.URL))
	r, err := idc.Lookup("foo")
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, r.IP, "foo")
	assertEqual(t, r.City, "bar")
	assertEqual(t, r.Region, "baz")
}

func TestWithAPIKey(t *testing.T) {
	apiKey := "foo"
	idc := NewClient(WithAPIKey(apiKey))
	assertEqual(t, idc.APIKey, apiKey)
}

func TestWithLanguage(t *testing.T) {
	lang := "vi"
	idc := NewClient(WithLanguage(lang))
	assertEqual(t, idc.Language, lang)
}

func TestWithURL(t *testing.T) {
	url := "http://localhost"
	idc := NewClient(WithURL(url))
	assertEqual(t, idc.url, url)
}
