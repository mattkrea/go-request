package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RequestMethod string

const (
	GET     RequestMethod = "GET"
	HEAD    RequestMethod = "HEAD"
	PUT     RequestMethod = "PUT"
	POST    RequestMethod = "POST"
	OPTIONS RequestMethod = "OPTIONS"
	DELETE  RequestMethod = "DELETE"
)

// Request is the base struct for all methods
type Request struct {
	url     string
	headers map[string]string
	body    *bytes.Buffer
	method  RequestMethod
}

type Response struct {
	Status  int
	Headers *http.Header
	Bytes   []byte
}

// New initializes a request (defaulting to GET)
func New(url string) *Request {

	r := &Request{
		url:    url,
		method: GET,
	}

	r.headers = make(map[string]string)

	return r
}

// Method allows you to set a request method in the case
// that you choose to use a method outside of those provided
// with shorthand methods such as Get(), Post(), etc
func (r *Request) Method(method RequestMethod) *Request {

	r.method = method

	return r
}

// Get is shorthand for New().Method(GET)
func Get(url string) *Request {

	return &Request{
		url:    url,
		method: GET,
	}
}

// Post is shorthand for New().Method(POST)
// You may also provide a request body (struct) that
// will be serialized into JSON and sent with the request
func Post(url string, body interface{}) *Request {

	r := &Request{
		url:    url,
		method: POST,
	}

	if r.headers == nil {
		r.headers = make(map[string]string)
	}

	if body != nil {

		requestBody, _ := json.Marshal(body)

		r.body = bytes.NewBuffer(requestBody)
		r.headers["Content-Type"] = "application/json"
	}

	return r
}

// Put is shorthand for New().Method(PUT)
// Similar to Post() you can provide a request body
func Put(url string, body interface{}) *Request {

	r := &Request{
		url:    url,
		method: PUT,
	}

	if r.headers == nil {
		r.headers = make(map[string]string)
	}

	if body != nil {

		requestBody, _ := json.Marshal(body)

		r.body = bytes.NewBuffer(requestBody)
		r.headers["Content-Type"] = "application/json"
	}

	return r
}

// Head is shorthand for New().Method(HEAD)
func Head(url string) *Request {

	return &Request{
		url:    url,
		method: HEAD,
	}
}

// Options is shorthand for New().Method(OPTION)
func Options(url string) *Request {

	return &Request{
		url:    url,
		method: OPTIONS,
	}
}

// Header sets a request header
func (r *Request) Header(key string, value string) *Request {

	if r.headers == nil {
		r.headers = make(map[string]string)
	}

	r.headers[key] = value

	return r
}

// Headers is a shortcut to 'Header' allowing you to set
// multiple in a single method call
func (r *Request) Headers(h map[string]string) *Request {

	if r.headers == nil {
		r.headers = make(map[string]string)
	}

	for k, v := range h {
		r = r.Header(k, v)
	}

	return r
}

// Do provides a simple way to finalize a request with all
// previously provided settings and returns a byte slice or error
func (r *Request) Do() (*Response, error) {

	if r.method == "" {
		r.method = GET
	}

	cli := &http.Client{}

	var req *http.Request
	var err error

	if r.body != nil {
		req, err = http.NewRequest(r.method.string(), r.url, r.body)
	} else {
		req, err = http.NewRequest(r.method.string(), r.url, nil)
	}

	if err != nil {
		return &Response{}, err
	}

	for header, value := range r.headers {
		req.Header.Add(header, value)
	}

	res, err := cli.Do(req)
	if err != nil {
		return &Response{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &Response{}, err
	}

	return &Response{
		Status:  res.StatusCode,
		Headers: &res.Header,
		Bytes:   body,
	}, nil
}

// Map allows you to provide a struct that a response
// body will be deserialized directly into rather than
// having to call request.Do() and handle the JSON yourself
func (r *Request) Map(i interface{}) error {

	resp, err := r.Do()
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Bytes, &i)

	return err
}

func (m RequestMethod) string() string {
	switch m {
	case HEAD:
		return "HEAD"
	case PUT:
		return "PUT"
	case POST:
		return "POST"
	case OPTIONS:
		return "OPTIONS"
	}

	return "GET"
}
