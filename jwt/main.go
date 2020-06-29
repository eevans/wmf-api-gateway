package main

import (
	"fmt"
	jose "github.com/square/go-jose"
	jwt "github.com/square/go-jose/jwt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type RateLimitType struct {
	RequestsPerUnit int64	`json:"requests_per_unit"`
	Unit 			string 	`json:"unit"`
}

type customClaims struct {
	ClientID  string		`json:"cid"`
	RateLimit RateLimitType	`json:"ratelimit"`
}

func main() {
	// Accepts a single argument for an RSA256 JWK private signing key.
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "%s <private_key.json> <number_of_requests_per_minute>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading %s: %s\n", os.Args[1], err)
	}

	requestPerMinute, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatalf("Error: #{os.Args[2]} is not a number")
	}

	private := jose.JSONWebKey{}

	// Unmarshal the key from JSON
	if err := private.UnmarshalJSON(f); err != nil {
		log.Fatalf("Unable to unmarshal JSON private key (jwk): %s\n", err)
	}

	// Instantiate an RSA256 signer
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: private.Key}, nil)
	if err != nil {
		log.Fatalf("Error creating new Signer: %s\n", err)
	}

	now := jwt.NewNumericDate(time.Now())

	// Set JWTs to expire one hour from generation
	duration, err := time.ParseDuration("1h")
	if err != nil {
		log.Fatalf("Unable to parse duration for expiry: %s\n", err)
	}
	expiry := jwt.NewNumericDate(time.Now().Add(duration))

	cl := jwt.Claims{
		Issuer:    "http://dev.wikipedia.org",
		Audience:  jwt.Audience{"core", "wikifeeds"},
		NotBefore: now,
		IssuedAt:  now,
		Expiry:    expiry,
	}

	customCl := customClaims{"b6123f22-79b6-11ea-8bde-c77510c60c52",
		RateLimitType{requestPerMinute, "MINUTE"}}

	// Sign and encode the JWT
	raw, err := jwt.Signed(signer).Claims(cl).Claims(customCl).CompactSerialize()
	if err != nil {
		log.Fatalf("Error signing and serializing JWT: %s\n", err)
	}

	fmt.Println(raw)
}
