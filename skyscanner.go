package skyscanner

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Version is the current library version
const Version = "0.1"

const APIBase = "http://partners.api.skyscanner.net/apiservices/"

// SkyScanner is the main library structure
type SkyScanner struct {
	client *http.Client
	apiKey string
	u      *url.URL
}

// New creates new SkyScanner API
func New(rt http.RoundTripper, apiKey string) *SkyScanner {
	u, _ := url.Parse(APIBase)
	return &SkyScanner{
		client: &http.Client{
			Transport: rt,
		},
		apiKey: apiKey,
		u:      u,
	}
}

func (ss *SkyScanner) fetchURL(url string) ([]byte, error) {
	var err error
	var resp *http.Response
	if resp, err = ss.client.Get(url); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
