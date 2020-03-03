package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

// These be the flags ya looking for
var (
	user string
)

func main() {
	// parsing ya flags
	flag.Parse()

	//if the user does not supply the necessary flags, just print usage
	if flag.NFlag() == 0 {
		printUsage()
	}

	// Separating the users out via a split on the comma
	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	fmt.Println("")
	for _, u := range users {
		result := getUsers(u)
		fmt.Printf("Username: %s ", result.Login)
		fmt.Printf("Name: %s ", result.Name)
		fmt.Printf("Email: %s ", result.Email)
		fmt.Printf("Bio: %s ", result.Bio)
		fmt.Printf("Public Repo: %i ", result.PublicRepos)
		fmt.Println("")
	}
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
