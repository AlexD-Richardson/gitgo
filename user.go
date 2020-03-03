package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	apiURL       = "https://api.github.com"
	userEndpoint = "/users/"
)

// User struct formulated from the GitHub API: https://developer.github.com/v3/users/
type User struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	Name              string      `json:"name"`
	Company           string      `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             string      `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               string      `json:"bio"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

func getUsers(name string) User {
	// Making GET request to GitHub API for the user(s)
	resp, err := http.Get(apiURL + userEndpoint + name)

	// if an err was passed back from GET request then log the err and exit with Fatalf
	if err != nil {

		log.Fatalf("Slowly back away from the computer... Jk GET request messed up: %s\n", err)
	}

	// Deferring the closing of the resp body till the top-level func returns
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("MALFUNCTION!... jk there was an error reading your data: %s\n", err)
	}

	var user User
	json.Unmarshal(body, &user)

	if (User{}) == user {
		fmt.Println("I don't know who you are looking for but they don't exist")
		os.Exit(1)
		return user
	}

	return user
}
