package admin

import (
	"context"
	"log"

	firebase "firebase.google.com/go"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type UserInfo struct {
	Name    string
	Email   string
	Phoneno string
	Address string
}

type UserInfoAPI struct {
	client *firestore.Client
}

func NewUserInfoAPI() *UserInfoAPI {

	u := &UserInfoAPI{}
	opt := option.WithCredentialsFile("./100diary.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	firestoreClient, err := app.Firestore(context.Background())
	if err != nil {
		panic(err)
	}
	u.client = firestoreClient
	return u
}

func (u *UserInfoAPI) GetAll() *UserInfo {
	userInfos := u.client.Collection("UserInfo")
	log.Printf("UserInfos : %v\n", userInfos)
	return nil
}

func (u *UserInfoAPI) WriteUserInfo() {
	doc, wr, err := u.client.Collection("UserInfo").Add(context.Background(), map[string]interface{}{
		"email": "rhkdgus0826@gmail.com",
		"name":  "papayetoo",
		"age":   34,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("UserInfo Craete Time : %v %v\n", doc, wr.UpdateTime)
}
