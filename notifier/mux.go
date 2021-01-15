package notifier

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter initializes the notification server
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	return router
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {

}
