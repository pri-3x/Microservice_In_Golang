package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// This function is responsible for generating a JSON Web Token (JWT) and returning it as a string.
func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "priyanshurajput"
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(MySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

// this func is reponsible for  for generating a JWT token by calling the GetJWT function,
// printing the token to the console, and sending the token as the response to the client's request.
func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	fmt.Println(validToken)
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	fmt.Fprintf(w, string(validToken))
}

// this func will handle the HTTP requests.
func handleRequests() {
	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
