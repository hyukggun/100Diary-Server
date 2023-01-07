package admin

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type User struct {
	Email    string
	Password string
}

type AuthManager struct {
	client *auth.Client
}

func NewAuthManager() *AuthManager {

	authManager := &AuthManager{}

	opt := option.WithCredentialsFile("./100diary.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	authClient, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	authManager.client = authClient
	return authManager
}

func (a AuthManager) CreateUserByEmail(email, password string) bool {
	newUser := &auth.UserToCreate{}
	record, err := a.client.CreateUser(context.Background(), newUser.Email(email).Password(password))
	if err != nil {
		return false
	}
	log.Println("User Create Success !! ", record)
	return true
}

func (a AuthManager) GetUserByEmail(email string) bool {
	record, err := a.client.GetUserByEmail(context.Background(), email)
	if err != nil {
		log.Printf("GetUserByEmail Occured Error %v\n", err)
		return false
	}
	log.Printf("User Record is %v", record)
	return true
}
