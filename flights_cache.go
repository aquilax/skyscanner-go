package skyscanner

import (
	"encoding/json"
	"net/url"
	"strings"
)

const browseQuotesPath = "browsequotes"
const browseQuotesVersion = "v1.0"

type CachedQuotesRequest struct {
	Country             string
	Currency            string
	Locale              string
	OriginPlace         string
	DestinationPlace    string
	OutboundPartialDate PartialDate
	InboundPartialDate  *PartialDate
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

type CacheQuotesResponse struct {
	Quotes     []Quote    `json:"Quotes"`
	Places     []Place    `json:"Places"`
	Carriers   []Carrier  `json:"Carriers"`
	Currencies []Currency `json:"Currencies"`
}

func getCacheQuotesURL(u *url.URL, req *CachedQuotesRequest) string {
	path := []string{
		browseQuotesPath,
		browseQuotesVersion,
		req.Country,
		req.Currency,
		req.Locale,
		req.OriginPlace,
		req.DestinationPlace,
		req.OutboundPartialDate.String(),
	}
	if req.InboundPartialDate != nil {
		path = append(path, req.InboundPartialDate.String())
	}
	u.Path = u.Path + strings.Join(path, "/")
	return u.String()
}

func (ss *SkyScanner) CacheQuotes(req *CachedQuotesRequest) (*CacheQuotesResponse, error) {
	url := getCacheQuotesURL(ss.u, req)
	var content []byte
	var err error
	if content, err = ss.fetchURL(url); err != nil {
		return nil, err
	}
	var resp CacheQuotesResponse
	if err = json.Unmarshal(content, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (ss *SkyScanner) CacheRoutes() {
}

func (ss *SkyScanner) CacheDates() {
}

func (ss *SkyScanner) CacheDatesGrid() {
}
