package main

import (
	"github.com/minormending/go-windows-toast/notifier"
)

func main() {

	notifier.Notify("com.kevin.app", "Hello", "World", map[string]string{
		"slug_one": "Hello One",
		"slug_two": "Hello two",
	})
}
