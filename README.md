gorest
======

A restful framework in Golang.

#Installation
`go get -u "github.com/elgs/gorest"`

#Sample Code
```go
package main

import (
	"github.com/elgs/gorest"
)

func main() {
	ds := "user:pass@tcp(host:3306)/test.test"
	dbo := &gorest.MySqlDataOperator{Ds: ds}
	r := &gorest.Gorest{
		Port:      8080,
		UrlPrefix: "api",
		Dbo:       dbo}
	r.Serve()
}

```
where in the test.test table:
```
ID	NAME		CREATE_TIME

0	Alicia	2014-08-18 18:43:38
1	Brian	2014-08-18 18:43:38
2	Chloe	2014-08-18 18:43:38
4	Bianca	2014-08-18 18:43:38
5	Leo		2014-08-18 18:43:38
6	Joy		2014-08-18 18:43:38
7	Samuel	2014-08-18 18:43:38
```

#Client
`curl http://localhost:8080/api/test.test`
outputs:
```json
$ curl http://localhost:8080/api/test.test
{"data":[{"create_time":"2014-08-18 18:43:38","id":"0","name":"Alicia"},{"create_time":"2014-08-18 18:43:38","id":"1","name":"Brian"},{"create_time":"2014-08-18 18:43:38","id":"2","name":"Chloe"},{"create_time":"2014-08-18 18:43:38","id":"4","name":"Bianca"},{"create_time":"2014-08-18 18:43:38","id":"5","name":"Leo"},{"create_time":"2014-08-18 18:43:38","id":"6","name":"Joy"},{"create_time":"2014-08-18 18:43:38","id":"7","name":"Samuel"}],"total":-1}elgss-iMac:gorest
```

#API
TODO