package main

import "github.com/talkops/sdk-go"

func main() {
	talkops.NewExtension().
		SetName("My Awesome Extension").
		Start()
}
