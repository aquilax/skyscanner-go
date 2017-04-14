package skyscanner

import (
	"net/url"
	"testing"
)

func Test_getCacheQuotesURL(t *testing.T) {
	testURL, _ := url.Parse("http://example.com/test/v1/")
	type args struct {
		u      *url.URL
		apiKey string
		req    *CachedQuotesRequest
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
				apiKey: "123",
				req: &CachedQuotesRequest{
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
				apiKey: "1234",
				req: &CachedQuotesRequest{
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
			if got := getCacheQuotesURL(tt.args.u, tt.args.apiKey, tt.args.req); got != tt.want {
				t.Errorf("getCacheQuotesURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
