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
	
	err := request.New("https://api.github.com/users/mattkrea/repos").Method("GET").Map(&repos)
	if err != nil {
		// handle err
	}

}
```