package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"log"
)

func main() {
	opt := option.WithCredentialsFile("./service-account-key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase app: %v", err)
	}

	// Create a Firebase auth client instance
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to create Firebase auth client: %v", err)
	}
}
