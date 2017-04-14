package skyscanner

import (
	"net/url"
	"testing"
)

func Test_getCacheRequestURL(t *testing.T) {
	testURL, _ := url.Parse("http://example.com/test/v1/")
	type args struct {
		u      *url.URL
		path   []string
		apiKey string
		req    *CachedRequest
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test full request",
			args{
				u:      testURL,
				path:   []string{"browsequotes", "v1.0"},
				apiKey: "123",
				req: &CachedRequest{
					"Country",
					"Currency",
					"Locale",
					"OriginPlace",
					"DestinationPlace",
					"OutboundPartialDate",
					NewPartialDate("InboundPartialDate"),
				},
			},
			"http://example.com/test/v1/browsequotes/v1.0/Country/Currency/Locale/OriginPlace/DestinationPlace/OutboundPartialDate/InboundPartialDate?apiKey=123",
		},
		{
			"Test empty Inbound Partual date request",
			args{
				u:      testURL,
				path:   []string{"browsequotes", "v1.0"},
				apiKey: "1234",
				req: &CachedRequest{
					"Country",
					"Currency",
					"Locale",
					"OriginPlace",
					"DestinationPlace",
					"OutboundPartialDate",
					nil,
				},
			},
			"http://example.com/test/v1/browsequotes/v1.0/Country/Currency/Locale/OriginPlace/DestinationPlace/OutboundPartialDate?apiKey=1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCachedRequestURL(tt.args.u, tt.args.path, tt.args.apiKey, tt.args.req); got != tt.want {
				t.Errorf("getCachedRequestURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
