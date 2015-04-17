# go-request
WIP: Goal is to provide a chainable API for HTTP requests

### Example

```go
import "github.com/mattkrea/go-request"

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

}
```