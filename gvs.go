package main

import (
	"flag"
	"fmt"

	"github.com/peterducai/gvs/cmd/gvs"
)

func main() {
	fmt.Println("gvs")

	var name string
	flag.StringVar(&name, "opt", "", "Usage")

	flag.Parse()

	fmt.Println(name)

	gvs.ConfigHealthCheck()

}
