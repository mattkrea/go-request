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
	HEAD                  = "HEAD"
	PUT                   = "PUT"
	POST                  = "POST"
	OPTIONS               = "OPTIONS"
)

type Request struct {
	url     string
	headers map[string]string
	body    bytes.Buffer
	method  RequestMethod
}

func New(url string) *Request {

	r := &Request{
		url: url,
	}

	return r
}

func (r *Request) Method(method RequestMethod) *Request {

	r.method = method

	return r
}

func (r *Request) Get(url string) *Request {

	return &Request{
		url:    url,
		method: GET,
	}
}

func (r *Request) Put(url string) *Request {

	return &Request{
		url:    url,
		method: PUT,
	}
}

func (r *Request) Head(url string) *Request {

	return &Request{
		url:    url,
		method: HEAD,
	}
}

func (r *Request) Options(url string) *Request {

	return &Request{
		url:    url,
		method: OPTIONS,
	}
}

func (r *Request) Post(url string) *Request {

	return &Request{
		url:    url,
		method: POST,
	}
}

func (r *Request) Header(key string, value string) *Request {

	r.headers[key] = value

	return r
}

func (r *Request) Headers(h map[string]string) *Request {
	r.headers = h

	return r
}

func (r *Request) Do() ([]byte, error) {

	if r.method == "" {
		r.method = GET
	}

	cli := &http.Client{}
	req, err := http.NewRequest(r.method.string(), r.url, nil)
	if err != nil {
		return nil, err
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
