package main

import (
	flag "github.com/ogier/pflag"
)

// flags
var (
	user string
)

func main() {
	flag.Parse()
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
