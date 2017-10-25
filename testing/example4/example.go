package example4

import (
	"log"
	"net/http"
	"myproject/gotraining-test/testing/example4/handlers"
)

func main() {
	handlers.Routes()

	log.Println("listener: started : Listening on port 4000")
	http.ListenAndServe(":4000", nil)
}
