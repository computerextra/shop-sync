package main

import (
	"fmt"
	"os"

	authentication "github.com/computerextra/shop-sync/Authentication"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	var cred authentication.ClientCredentials
	cred.ClientId = os.Getenv("CLIENT_ID")
	cred.ClientSecret = os.Getenv("CLIENT_SECRET")
	cred.GrantType = "client_credentials"
	cred.ApiUrl = os.Getenv("ADMIN_URL")

	client := cred.Login()

	res, err := client.Get("/products")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", res)
}
