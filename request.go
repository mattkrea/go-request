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

func Put(url string, body interface{}) *Request {

	requestBody, _ := json.Marshal(body)

	return &Request{
		url:    url,
		method: PUT,
		body:   bytes.NewBuffer(requestBody),
	}
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

func Post(url string, body interface{}) *Request {

	requestBody, _ := json.Marshal(body)

	return &Request{
		url:    url,
		method: POST,
		body:   bytes.NewBuffer(requestBody),
	}
}

func (r *Request) Header(key string, value string) *Request {

	r.headers[key] = value

	return r
}

func (r *Request) Headers(h map[string]string) *Request {

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
	req, err := http.NewRequest(r.method.string(), r.url, r.body)
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
