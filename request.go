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

type Request struct {
	url     string
	headers map[string]string
	body    *bytes.Buffer
	method  RequestMethod
}

func New(url string) *Request {

	r := &Request{
		url:    url,
		method: GET,
	}

	r.headers = make(map[string]string)

	return r
}

func (r *Request) Method(method RequestMethod) *Request {

	r.method = method

	return r
}

func Get(url string) *Request {

	return &Request{
		url:    url,
		method: GET,
	}
}

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

func Head(url string) *Request {

	return &Request{
		url:    url,
		method: HEAD,
	}
}

func Options(url string) *Request {

	return &Request{
		url:    url,
		method: OPTIONS,
	}
}

func (r *Request) Header(key string, value string) *Request {

	if r.headers == nil {
		r.headers = make(map[string]string)
	}

	r.headers[key] = value

	return r
}

func (r *Request) Headers(h map[string]string) *Request {

	if r.headers == nil {
		r.headers = make(map[string]string)
	}

	for k, v := range h {
		r = r.Header(k, v)
	}

	return r
}

func (r *Request) Do() ([]byte, error) {

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
		return nil, err
	}

	for header, value := range r.headers {
		req.Header.Add(header, value)
	}

	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r *Request) Map(i interface{}) error {

	resp, err := r.Do()
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, &i)

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
