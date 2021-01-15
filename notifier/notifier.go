package notifier

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-toast/toast"
	"github.com/gorilla/mux"
)

// Notify sends a desktop notification
func Notify(appID, title, message string, options map[string]string) error {
	if len(options) > 4 {
		return errors.New("max 4 options supported")
	}

	cb := make(chan string)

	router := mux.NewRouter()
	router.HandleFunc("/callback", func(rw http.ResponseWriter, r *http.Request) {
		vals := r.URL.Query()
		option, ok := vals["option"]
		if !ok || len(option) != 1 {
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			cb <- option[0]
			rw.WriteHeader(http.StatusAccepted)
		}
	})
	server := http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	actions := []toast.Action{}
	for slug, text := range options {
		actions = append(actions, toast.Action{
			Type:      "protocol",
			Label:     text,
			Arguments: "http://localhost:8080/callback?option=" + slug,
		})
	}
	notification := toast.Notification{
		AppID:   appID,
		Title:   title,
		Message: message,
		Actions: actions,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	err := notification.Push()

	select {
	case option := <-cb:
		fmt.Printf("recieved %s", option)
	case <-time.After(time.Second * 15):
		fmt.Println("timeout exceeded")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	server.Shutdown(ctx)

	return err
}
