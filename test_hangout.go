package main

import (
    "fmt"
    "context"
//     "google.golang.org/api/option"
    "google.golang.org/api/chat/v1"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

func main() {
    var scopes []string
    scopes = append(scopes, "https://www.googleapis.com/auth/userinfo.email")
    var config = &oauth2.Config{
        ClientID:     "459361267143-35gk0cefv98lg7t2qjsh8as5ioc2004t.apps.googleusercontent.com",
        ClientSecret: "58OmcPyh-2su8ebVtrCPaqta",
        Endpoint:     google.Endpoint,
        Scopes:       scopes,
        RedirectURL:  "https://mm.jtg.tools",
    }
    ctx := context.Background()
//     token, err := config.Exchange(ctx, "online")
//     chatService, err := chat.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))



    url := config.AuthCodeURL("state")
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Handle the exchange code to initiate a transport.
	tok, err := config.Exchange(ctx, "authorization-code")
	if err != nil {
		fmt.Printf("Error appeared: ", err)
	}
	client := config.Client(oauth2.NoContext, tok)
    chatService, _ := chat.New(client)
    spacesListCall := chat.NewSpacesService(chatService).List()
    resp, err := spacesListCall.Do()
    if resp != nil {
        print(resp)
    }
    if err != nil {
        fmt.Printf("Error appeared: ", err)
    } else {
        print(resp)
    }
}
