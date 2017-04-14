package skyscanner

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const apiKey = "api-key"

type DummyTransport struct {
	content string
}

func NewDummyTransport(content string) *DummyTransport {
	return &DummyTransport{content}
}

func (dt *DummyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	b := bytes.NewBufferString(dt.content)
	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bufio.NewReader(b)),
	}
	return resp, nil
}

func TestSkyscanner(t *testing.T) {
	Convey("Given Skyscanner", t, func() {
		ss := New(NewDummyTransport(""), apiKey)
		Convey("Skyscanner must not be nil", func() {
			So(ss, ShouldNotBeNil)
		})
	})
}
