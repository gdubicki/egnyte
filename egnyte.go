package egnyte

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GetAccessToken returns Egnyte access token
func GetAccessToken(domain, clientID, clientSecret, username, password string) string {
	egnyteDomain := sanitizeDomain(domain)
	endpoint := egnyteDomain + "/puboauth/token"
	payload := strings.NewReader(
		fmt.Sprintf("client_id=%s&client_secret=%s&username=%s&password=%s&grant_type=%s",
			clientID,
			clientSecret,
			username,
			password,
			"password"))

	resp, err := post("", endpoint, payload)

	type accessToken struct {
		AccessToken string `json:"access_token"`
	}
	var a accessToken

	err = json.Unmarshal(resp, &a)
	if err != nil {
		fmt.Println("error:", err)
	}
	return a.AccessToken
}

// Connect return Egnyte client
func Connect(token, domain string) Client {
	return getClient(token, domain)
}
