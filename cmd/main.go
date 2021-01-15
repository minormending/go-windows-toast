package main

import "github.com/minormending/go-windows-toast/notifier"

func main() {
	notifier.Notify("com.kevin.app", "Hello", "World")
}
