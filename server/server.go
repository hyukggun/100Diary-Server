package server

import (
	"100diary/admin"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ServerStartHandler struct{}

func (s ServerStartHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Server Start")
}

func RunServer() {
	log.Printf("Run Server \n")
	authManager := admin.NewAuthManager()
	http.HandleFunc("/user", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request Method : %v\n", req.Method)
		if req.Method == "GET" {
			email := req.URL.Query().Get("email")
			w.Write([]byte(fmt.Sprintf("%s User has an account : %v\n", email, authManager.GetUserByEmail(email))))
		}
		if req.Method == "POST" {
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				return
			}
			user := &admin.User{}
			if err := json.Unmarshal(body, user); err != nil {
				log.Printf("JSON Decoding Error", err)
				return
			}
			isAlreadyHaveAccount := authManager.GetUserByEmail(user.Email)
			switch {
			case isAlreadyHaveAccount == true:
				w.Write([]byte("Already Have account"))
			default:
				result := authManager.CreateUserByEmail(user.Email, user.Password)
				if result {
					w.Write([]byte("Create User result : Success!!"))
				} else {
					w.Write([]byte("Create User result : Failed!!"))
				}
			}

		}
	})
	// startHandler := func(w http.ResponseWriter, _ *http.Request) {
	// 	log.Printf("Start server !!")
	// }

	http.ListenAndServe(":8000", nil)
}
