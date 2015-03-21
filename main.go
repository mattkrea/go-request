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

func (r *Request) Headers(h map[string]string) *Request {
	r.headers = h

	return r
}

func (r *Request) Do() ([]byte, error) {

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
