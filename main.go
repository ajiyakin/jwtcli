package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/logrusorgru/aurora"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "help" {
		printHelp()
		os.Exit(0)
	}
	if len(os.Args) < 2 {
		fmt.Printf("\n%s\n\n", aurora.Red("error: JWT token is required"))
		printHelp()
		os.Exit(5)
	}
	jwtToken := os.Args[1]

	tokenSegments := strings.Split(jwtToken, ".")
	if len(tokenSegments) < 3 {
		fmt.Printf("\n%s\n\n", aurora.Red("error: invalid JWT token"))
		printHelp()
		os.Exit(5)
	}

	fmt.Printf("\n-------------------------- HEADER --------------------------------\n")
	header, err := jwt.DecodeSegment(tokenSegments[0])
	if err != nil {
		fmt.Printf(aurora.Sprintf(aurora.Red("unable to decode token payload: %v\n"), err))
		os.Exit(5)
	}
	headerClaims, err := decode(string(header))
	if err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("unable to decode JWT header: %v"), err))
		os.Exit(5)
	}

	if err := encode(headerClaims); err != nil {
		fmt.Println("unable to encode JSON JWT header: ", err)
		os.Exit(5)
	}

	fmt.Printf("\n-------------------------- PAYLOAD -------------------------------\n")
	payload, err := jwt.DecodeSegment(tokenSegments[1])
	if err != nil {
		fmt.Printf(aurora.Sprintf(aurora.Red("unable to decode token payload: %v\n"), err))
		os.Exit(5)
	}

	payloadClaims, err := decode(string(payload))
	if err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("unable to decode JWT payload: %v\n"), err))
		os.Exit(5)
	}

	if err := encode(payloadClaims); err != nil {
		fmt.Println(aurora.Sprintf(aurora.Red("unable to encode JSON JWT payload: %v\n"), err))
		os.Exit(5)
	}

	// TODO Handle for given signature to also do verify the token
	//      Maybe the format will looks like this: ./jwt JWT_TOKEN [signature]

	fmt.Printf("\n-------------------------- SIGNATURE -----------------------------\n")
	fmt.Printf("%s\n\n", tokenSegments[2])
}

func printHelp() {
	progName := os.Args[0]
	fmt.Printf(aurora.Sprintf(aurora.Green(`App to decode information within a JWT token
Usage:
  %s JWT_TOKEN
example:
  %s eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU
`+"\n"), progName, progName))
}

func decode(tokenElement string) (claims map[string]interface{}, err error) {
	err = json.NewDecoder(strings.NewReader(tokenElement)).Decode(&claims)
	return
}

func encode(claims map[string]interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	return encoder.Encode(claims)
}
