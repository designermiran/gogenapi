package main

import "os"

//go:generate go-bindata -o ./gogenapi/bindata.go -pkg gogenapi _templates/...

func main() {
	os.Exit(Run(os.Args[1:]))
}
