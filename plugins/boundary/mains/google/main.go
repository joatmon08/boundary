package main

import (
	"fmt"
	"os"

	hp "github.com/hashicorp/boundary/sdk/plugins"
	google "github.com/joatmon08/boundary-plugin-google/plugin"
)

func main() {
	if err := hp.ServePlugin(new(google.GooglePlugin)); err != nil {
		fmt.Println("Error serving plugin", err)
		os.Exit(1)
	}
	os.Exit(0)
}
