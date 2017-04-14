package skyscanner

import (
	"net/url"
	"testing"
)

func Test_getCacheQuotesURL(t *testing.T) {
	testURL, _ := url.Parse("http://example.com/test/v1?apiKey=123")
	type args struct {
		u   *url.URL
		req *CachedQuotesRequest
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test full request",
			args{
				u: testURL,
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
			"http://example.com/test/v1browsequotes/v1.0/Country/Currency/Locale/OriginPlace/DestinationPlace/OutboundPartialDate/InboundPartialDate?apiKey=123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCacheQuotesURL(tt.args.u, tt.args.req); got != tt.want {
				t.Errorf("getCacheQuotesURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
