package skyscanner

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const createSessionURLPath = "pricing"
const createSessionURLVersion = "v1.0"

type CabinClass string

const (
	CabinClassAny            CabinClass = "any"
	CabinClassEconomy        CabinClass = "economy"
	CabinClassPremiumEconomy CabinClass = "premiumeconomy"
	CabinClassBusiness       CabinClass = "business"
	CabinClassFirst          CabinClass = "first"
)

type LiveRequest struct {
	Currency         string
	Locale           string
	OriginPlace      string
	DestinationPlace string
	OutboundDate     PartialDate
	InboundDate      *PartialDate
	CabinClass       CabinClass
	Adults           int
	Children         int
	Infants          int
}

func getLiveSessionPostData(apiKey string, lr *LiveRequest) string {
	data := url.Values{}
	data.Add("currency", lr.Currency)
	data.Add("locale", lr.Locale)
	data.Add("originPlace", lr.OriginPlace)
	data.Add("destinationPlace", lr.DestinationPlace)
	data.Add("outboundDate", lr.OutboundDate.String())
	if lr.InboundDate != nil {
		data.Add("outboundDate", lr.OutboundDate.String())
	}
	if lr.CabinClass != CabinClassAny {
		data.Add("cabinClass", string(lr.CabinClass))
	}
	data.Add("adults", strconv.Itoa(lr.Adults))
	data.Add("children", strconv.Itoa(lr.Children))
	data.Add("infants", strconv.Itoa(lr.Infants))
	data.Add("apiKey", apiKey)
	return data.Encode()
}

func (ss *SkyScanner) CreateLiveSession(ip net.IP, lr *LiveRequest) (string, error) {
	rru, _ := url.Parse(strings.Join([]string{
		createSessionURLPath,
		createSessionURLVersion,
	}, "/"))
	url := ss.u.ResolveReference(rru)
	requestData := getLiveSessionPostData(ss.apiKey, lr)
	r, err := http.NewRequest("POST", url.String(), bytes.NewBufferString(requestData))
	if err != nil {
		return "", err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Accept", "application/json")
	ss.client.Do(r)

	var resp *http.Response
	if resp, err = ss.client.Do(r); err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	location := resp.Header.Get("Location")
	return location, nil
}
