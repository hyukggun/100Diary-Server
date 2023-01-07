package main

import (
	"100diary/admin"
)

func main() {
	api := admin.NewUserInfoAPI()
	api.GetAll()
	api.WriteUserInfo()
}
