gorest
======

A restful framework in Golang.

#Installation
`go get -u github.com/elgs/gorest`

#Sample Code
```go
package main

import (
	"github.com/elgs/gorest"
)

func main() {
	ds := "user:pass@tcp(host:3306)/test?loc=Singapore"
	dbo := &gorest.MySqlDataOperator{Ds: ds}
	r := &gorest.Gorest{
		Port:      8080,
		UrlPrefix: "api",
		Dbo:       dbo}
	r.Serve()
}

```

#Client
`curl http://localhost:8080/api/test.test`