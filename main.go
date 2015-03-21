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
	URL     string
	Headers map[string]string
	Body    []byte
	Method  RequestMethod
}

func New(url string) *Request {

	r & Request{
		URL:  url,
		Body: nil,
	}

	return r
}

func (r *Request) Method(method RequestMethod) *Request {

	r.Method = method

	return r
}

func (r *Request) Headers(h map[string]string) *Request {
	r.Headers = h

	return r
}

func (r *Request) Map(i interface{}) error {

}

func (r *Request) Do() ([]byte, error) {

	cli := &http.Client{}
	req, err := http.NewRequest(r.Method, r.URL, r.Body)
	if err != nil {
		return nil, error
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, error
	}

	return body, nil
}
