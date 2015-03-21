package main

import (
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
	body    []byte
	method  RequestMethod
}

func New(url string) *Request {

	r := &Request{
		url:  url,
		body: nil,
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

func (r *Request) Map(i interface{}) error {
	return nil
}

func (r *Request) Do() ([]byte, error) {

	cli := &http.Client{}
	req, err := http.NewRequest(r.method.string(), r.url, r.body)
	if err != nil {
		return nil, error
	}

	body, err := ioutil.ReadAll(req.body)
	if err != nil {
		return nil, error
	}

	return body, nil
}

func (m RequestMethod) string() string {
	switch m {
	case GET:
		return "GET"
	case HEAD:
		return "HEAD"
	case PUT:
		return "PUT"
	case POST:
		return "POST"
	case OPTIONS:
		return "OPTIONS"
	}
}
