// handler
package gorest

import (
	"net/http"
)

var handlerRegistry = make(map[string]func(dbo DataOperator, gr *Gorest) func(w http.ResponseWriter, r *http.Request))

func RegisterHandler(id string, handler func(dbo DataOperator, gr *Gorest) func(w http.ResponseWriter, r *http.Request)) {
	handlerRegistry[id] = handler
}

func GetHandler(id string) func(dbo DataOperator, gr *Gorest) func(w http.ResponseWriter, r *http.Request) {
	return handlerRegistry[id]
}

func RegisterHttpHandlers(dbo DataOperator, gr *Gorest) {
	for id, makeHandler := range handlerRegistry {
		RegisterDataOperator(id, makeHandler(dbo, gr))
	}
}
