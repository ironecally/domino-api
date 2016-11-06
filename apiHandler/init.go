package apiHandler

import (
	"log"

	"github.com/julienschmidt/httprouter"
	grace "gopkg.in/paytm/grace.v1"
)

func InitAPI() {

	router := httprouter.New()

	log.Fatal(grace.Serve("9000", router))
}
