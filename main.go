package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/picatz/hunter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading ENV")
	}
	key := os.Getenv("HUNTER_API_KEY")

	// create a new client using the HUNTER_API_KEY environment variable
	// and the default net/http client
	client := hunter.New(key, hunter.UseDefaultHTTPClient)
	reader := bufio.NewReader(os.Stdin)
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	// verify email
	result, err := client.VerifyEmail(hunter.Params{
		"email": email,
	})
	// handle error
	if err != nil {
		panic(err)
	}
	// do something with the result data
	fmt.Println(result.Data.Score)
	fmt.Println(result.Data.Result)
}
