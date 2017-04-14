package skyscanner_test

import (
	"os"
	"testing"

	"net/http"

	"net"

	"github.com/aquilax/skyscanner-go"
)

var apiKey string

func TestMain(m *testing.M) {
	apiKey = os.Getenv("API_KEY")
	if apiKey != "" {
		os.Exit(m.Run())
	}
	os.Exit(0)
}

func TestLibrary(t *testing.T) {
	ss := skyscanner.New(http.DefaultTransport, apiKey)
	lr := &skyscanner.LiveRequest{
		Currency:         "SEK",
		Locale:           "en_GB",
		OriginPlace:      "sto",
		DestinationPlace: "anywhere",
		OutboundDate:     skyscanner.PartialDate("2017-05-05"),
		CabinClass:       skyscanner.CabinClassEconomy,
		Adults:           1,
		Children:         0,
		Infants:          0,
	}
	location, err := ss.CreateLiveSession(net.ParseIP("127.0.0.1"), lr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("location")
	t.Log(location)
}
