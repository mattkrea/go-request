# go-request
Goal is to provide a chainable API for HTTP requests to get rid of
some of the common boilerplate.

[![GoDoc](https://godoc.org/github.com/mattkrea/go-request?status.svg)](https://godoc.org/github.com/mattkrea/go-request) [![Build Status](https://travis-ci.org/mattkrea/go-request.svg?branch=master)](https://travis-ci.org/mattkrea/go-request) [![Coverage Status](https://coveralls.io/repos/mattkrea/go-request/badge.svg)](https://coveralls.io/r/mattkrea/go-request)

### Example

```go
import "github.com/mattkrea/go-request"
import "log"

func main() {

	type Repo struct {
		ID int64 `json:"id"`
		Name string `json:"full_name"`
	}

	var repos []Repo
	
	// Use the full API
	r := request.New("https://api.github.com/users/mattkrea/repos")
	r.Method("GET")
	err := r.Map(&repos)
	if err != nil {
		// handle err
	}

	// Or you can use the shorthand
	err := request.Get("https://api/github.com/users/mattkrea/repos").Map(&repos)
	if err != nil {
		// handle err
	}

	// Want raw data back?
	res, err := request.Get("http://httpbin.org/get").Do()
	if err != nil {
		// handle err
	} else if res.Status != 200 {
		// handle non-success status code
	}

	// Or how about async?
	responseQueue, errorQueue := request.Get("http://httpbin.org/get").DoAsync()
	if err := <- errorQueue; err != nil {
		// handle err
	}

	response := <- responseQueue

	log.Printf("Raw: %v", response.Bytes)
}
```