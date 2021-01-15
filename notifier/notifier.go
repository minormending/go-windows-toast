package notifier

import (
	"github.com/go-toast/toast"
)

// Notify sends a desktop notification
func Notify(appID, title, message string) error {
	notification := toast.Notification{
		AppID:   appID,
		Title:   title,
		Message: message,
		Actions: []toast.Action{
			{Type: "protocol", Label: "I'm a button", Arguments: "https://google.com"},
			{Type: "protocol", Label: "Me too!", Arguments: "https://twitter.com"},
		},
	}
	err := notification.Push()
	return err
}
