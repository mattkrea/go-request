package request

import (
	"testing"
)

const url string = "http://httpbin.org/get"
const postURL string = "http://httpbin.org/post"

type httpbinResponseStruct struct {
	Args    map[string]interface{} `json:"args"`
	Headers map[string]interface{} `json:"headers"`
	Origin  string                 `json:"origin"`
	URL     string                 `json:"url"`
	JSON    map[string]string      `json:"json"`
}

type httpbinRequest struct {
	Property        string `json:"property"`
	AnotherProperty string `json:"anotherProperty"`
}

func TestAssignURL(t *testing.T) {

	r := New(url)

	if r.url != url {
		t.FailNow()
	}
}

func TestAssignGetMethod(t *testing.T) {

	method := GET

	r := New(url).Method(method)

	if r.method != method {
		t.FailNow()
	}
}

func TestAssignPutMethod(t *testing.T) {

	method := PUT

	r := New(url).Method(method)

	if r.method != method {
		t.FailNow()
	}
}

func TestAssignPostMethod(t *testing.T) {

	method := POST

	r := New(url).Method(method)

	if r.method != method {
		t.FailNow()
	}
}

func TestAssignDeleteMethod(t *testing.T) {

	method := DELETE

	r := New(url).Method(method)

	if r.method != method {
		t.FailNow()
	}
}

func TestAssignHeadMethod(t *testing.T) {

	method := HEAD

	r := New(url).Method(method)

	if r.method != method {
		t.FailNow()
	}
}

func TestAssignOptionsMethod(t *testing.T) {

	method := OPTIONS

	r := New(url).Method(method)

	if r.method != method {
		t.FailNow()
	}
}

func TestAssignIndividualHeader(t *testing.T) {

	r := New(url).Method(GET).Header("X-Test-Key", "Test Value")

	if r.headers["X-Test-Key"] != "Test Value" {
		t.FailNow()
	}
}

func TestAssignHeaderMap(t *testing.T) {

	// Assign an individual header to make sure
	// the header map does not override existing headers
	r := New(url).Method(GET).Header("X-Test-Key", "Test Value")

	r.Headers(map[string]string{"Another-Test": "Test2"})

	if r.headers["X-Test-Key"] != "Test Value" {
		t.FailNow()
	}

	if r.headers["Another-Test"] != "Test2" {
		t.FailNow()
	}
}

func TestGetRequestWithDo(t *testing.T) {

	_, err := New(url).Do()
	if err != nil {
		t.FailNow()
	}
}

func TestGetRequestForStatus(t *testing.T) {

	res, err := Get(url).Do()
	if err != nil {
		t.FailNow()
	}

	if res.Status != 200 {
		t.FailNow()
	}
}

func TestGetRequestAsync(t *testing.T) {

	response := make(chan *Response)
	e := make(chan error)

	Get(url).DoAsync(response, e)

	if err := <-e; err != nil {
		t.FailNow()
	}

	res := <-response
	if res.Status != 200 {
		t.FailNow()
	}
}

func TestGetRequestWithMap(t *testing.T) {

	var result httpbinResponseStruct

	err := New(url).Map(&result)
	if err != nil {
		t.FailNow()
	}

	if result.URL != url {
		t.FailNow()
	}
}

func TestPostRequestWithoutPayload(t *testing.T) {

	var result httpbinResponseStruct

	err := Post("http://httpbin.org/post", nil).Map(&result)
	if err != nil {
		t.FailNow()
	}

	if result.URL != "http://httpbin.org/post" {
		t.FailNow()
	}
}

func TestPostRequestWithPayload(t *testing.T) {

	var result httpbinResponseStruct

	body := &httpbinRequest{
		Property:        "Foo",
		AnotherProperty: "Bar",
	}

	err := Post("http://httpbin.org/post", body).Map(&result)
	if err != nil {
		t.FailNow()
	}

	if result.URL != "http://httpbin.org/post" {
		t.FailNow()
	}

	if result.JSON["anotherProperty"] != body.AnotherProperty {
		t.FailNow()
	}
}
