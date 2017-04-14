package skyscanner

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

const browseQuotesPath = "browsequotes"
const browseQuotesVersion = "v1.0"

const browseRoutesPath = "browseroutes"
const browseRoutesVersion = "v1.0"

const browseDatesPath = "browsedates"
const browseDatesVersion = "v1.0"

const browseDatesGridPath = "browsegrid"
const browseDatesGridVersion = "v1.0"

type CachedRequest struct {
	Country             string
	Currency            string
	Locale              string
	OriginPlace         string
	DestinationPlace    string
	OutboundPartialDate PartialDate
	InboundPartialDate  *PartialDate
}

func getCachedRequestURL(u *url.URL, path []string, apiKey string, req *CachedRequest) string {
	path = append(path, []string{
		req.Country,
		req.Currency,
		req.Locale,
		req.OriginPlace,
		req.DestinationPlace,
		req.OutboundPartialDate.String(),
	}...)
	if req.InboundPartialDate != nil {
		path = append(path, req.InboundPartialDate.String())
	}
	up, _ := url.Parse(strings.Join(path, "/"))
	q := up.Query()
	q.Set("apiKey", apiKey)
	up.RawQuery = q.Encode()
	return u.ResolveReference(up).String()
}

type Quote struct {
	QuoteID     int  `json:"QuoteId"`
	MinPrice    int  `json:"MinPrice"`
	Direct      bool `json:"Direct"`
	OutboundLeg struct {
		CarrierIds    []int  `json:"CarrierIds"`
		OriginID      int    `json:"OriginId"`
		DestinationID int    `json:"DestinationId"`
		DepartureDate string `json:"DepartureDate"`
	} `json:"OutboundLeg"`
	InboundLeg struct {
		CarrierIds    []int  `json:"CarrierIds"`
		OriginID      int    `json:"OriginId"`
		DestinationID int    `json:"DestinationId"`
		DepartureDate string `json:"DepartureDate"`
	} `json:"InboundLeg"`
	QuoteDateTime string `json:"QuoteDateTime"`
}

type Place struct {
	PlaceID        int    `json:"PlaceId"`
	Name           string `json:"Name"`
	Type           string `json:"Type"`
	SkyscannerCode string `json:"SkyscannerCode"`
	IataCode       string `json:"IataCode,omitempty"`
	CityName       string `json:"CityName,omitempty"`
	CityID         string `json:"CityId,omitempty"`
	CountryName    string `json:"CountryName,omitempty"`
}

type Carrier struct {
	CarrierID int    `json:"CarrierId"`
	Name      string `json:"Name"`
}

type Currency struct {
	Code                        string `json:"Code"`
	Symbol                      string `json:"Symbol"`
	ThousandsSeparator          string `json:"ThousandsSeparator"`
	DecimalSeparator            string `json:"DecimalSeparator"`
	SymbolOnLeft                bool   `json:"SymbolOnLeft"`
	SpaceBetweenAmountAndSymbol bool   `json:"SpaceBetweenAmountAndSymbol"`
	RoundingCoefficient         int    `json:"RoundingCoefficient"`
	DecimalDigits               int    `json:"DecimalDigits"`
}

type CachedQuotesResponse struct {
	Quotes     []Quote    `json:"Quotes"`
	Places     []Place    `json:"Places"`
	Carriers   []Carrier  `json:"Carriers"`
	Currencies []Currency `json:"Currencies"`
}

// GetCachedQuotes retrieves the cheapest quotes from cache prices.
// ref: https://skyscanner.github.io/slate/#browse-quotes
func (ss *SkyScanner) GetCachedQuotes(req *CachedRequest) (*CachedQuotesResponse, error) {
	url := getCachedRequestURL(ss.u, []string{
		browseQuotesPath,
		browseQuotesVersion,
	}, ss.apiKey, req)
	var content []byte
	var err error
	if content, err = ss.fetchURL(url); err != nil {
		return nil, err
	}
	var resp CachedQuotesResponse
	if err = json.Unmarshal(content, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type Route struct {
	OriginID      int    `json:"OriginId"`
	DestinationID int    `json:"DestinationId"`
	QuoteIds      []int  `json:"QuoteIds,omitempty"`
	Price         int    `json:"Price,omitempty"`
	QuoteDateTime string `json:"QuoteDateTime,omitempty"`
}

type CachedRoutesResponse struct {
	Routes     []Route    `json:"Routes"`
	Quotes     []Quote    `json:"Quotes"`
	Places     []Place    `json:"Places"`
	Carriers   []Carrier  `json:"Carriers"`
	Currencies []Currency `json:"Currencies"`
}

// GetCachedRoutes Retrieves the cheapest routes from cache prices.
//
// ref: https://skyscanner.github.io/slate/#browse-routes
func (ss *SkyScanner) GetCachedRoutes(req *CachedRequest) (*CachedRoutesResponse, error) {
	url := getCachedRequestURL(ss.u, []string{
		browseRoutesPath,
		browseRoutesVersion,
	}, ss.apiKey, req)
	var content []byte
	var err error
	if content, err = ss.fetchURL(url); err != nil {
		return nil, err
	}
	var resp CachedRoutesResponse
	if err = json.Unmarshal(content, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CachedDatesResponse struct{}

// GetCachedDates retrieves the cheapest dates for a given route from cache.
//
// Deprecated: The API currently refers to BrowseRoutes instead
//
// ref: https://skyscanner.github.io/slate/#browse-dates
func (ss *SkyScanner) GetCachedDates(req *CachedRequest) (*CachedDatesResponse, error) {
	return nil, fmt.Errorf("For this query please use the following service [BrowseRoutes]")
	url := getCachedRequestURL(ss.u, []string{
		browseDatesPath,
		browseDatesVersion,
	}, ss.apiKey, req)
	var content []byte
	var err error
	if content, err = ss.fetchURL(url); err != nil {
		return nil, err
	}
	var resp CachedDatesResponse
	if err = json.Unmarshal(content, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type CachedDatesGridResponse struct{}

// GetCachedDatesGrid retrieves the cheapest dates for a given route from cache,
// with the results formatted as a two-dimensional array to be easily displayed as a calendar.
//
// Deprecated: The API currently refers to BrowseRoutes instead
//
// ref: https://skyscanner.github.io/slate/#browse-dates-grid
func (ss *SkyScanner) GetCachedDatesGrid(req *CachedRequest) (*CachedDatesGridResponse, error) {
	return nil, fmt.Errorf("For this query please use the following service [BrowseRoutes]")
	url := getCachedRequestURL(ss.u, []string{
		browseDatesGridPath,
		browseDatesGridVersion,
	}, ss.apiKey, req)
	var content []byte
	var err error
	if content, err = ss.fetchURL(url); err != nil {
		return nil, err
	}
	var resp CachedDatesGridResponse
	if err = json.Unmarshal(content, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
