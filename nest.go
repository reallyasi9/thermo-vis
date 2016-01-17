package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	// "io/ioutil"
	"os"
)

const tokenFileName string = ".nest_access_token"

func main() {

	clientSecret, found := os.LookupEnv("NEST_CLIENT_SECRET")
	if !found {
		fmt.Println("Add the client secret to the environment variable NEST_CLIENT_SECRET and try again")
		os.Exit(-1)
	}

	tokenURL := fmt.Sprintf("https://api.home.nest.com/oauth2/access_token?client_id=811ab8b7-b82d-4e38-bb81-4eb15f7af940&client_secret=%s&grant_type=authorization_code", clientSecret)

	conf := &oauth2.Config{
		ClientID:     "811ab8b7-b82d-4e38-bb81-4eb15f7af940",
		ClientSecret: clientSecret,
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://home.nest.com/login/oauth2",
			TokenURL: tokenURL,
		},
	}

	// Determine if there is a saved token and try to use that
	tok := new(oauth2.Token)
	tokenFile, err := os.Open(tokenFileName)
	if err == nil {
		// Attempt to use this token
		decoder := json.NewDecoder(tokenFile)
		err = decoder.Decode(tok)
	}

	if err != nil {
		fmt.Println(err)
		// Redirect user to consent page to ask for permission.
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		fmt.Printf("Visit the URL for the auth dialog: %v\n\nEnter the code and press Enter: ", url)

		// Use the authorization code that is pushed to the redirect URL.
		// NewTransportWithCode will do the handshake to retrieve
		// an access token and initiate a Transport that is
		// authorized and authenticated by the retrieved token.
		var code string
		if _, err = fmt.Scan(&code); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		tok, err = conf.Exchange(oauth2.NoContext, code)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		// Save this token for later
		tokenFile, err = os.Create(tokenFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		encoder := json.NewEncoder(tokenFile)
		err = encoder.Encode(tok)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
	tokenFile.Close()

	client := conf.Client(oauth2.NoContext, tok)
	resp, err := client.Get("https://developer-api.nest.com/structures")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var structures map[string]structure
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&structures)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	resp.Body.Close()

	//fmt.Printf("Structures: %+v\n", structures)

	resp, err = client.Get("https://developer-api.nest.com/devices")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//
	// contents, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("%s\n", string(contents))
	// resp.Body.Close()
	// resp, _ = client.Get("https://developer-api.nest.com/devices")

	var devices devices
	decoder = json.NewDecoder(resp.Body)
	err = decoder.Decode(&devices)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	resp.Body.Close()

	//fmt.Printf("Devices: %+v\n", devices)

}
