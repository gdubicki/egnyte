package egnyte

import (
	"encoding/json"
	"fmt"
	"strconv"
)

var userPaths = map[string]string{
	"USERINFO":    "/pubapi/v1/userinfo",
	"USERDETAILS": "/pubapi/v2/users/"}

// User represents Egnyte user
type User struct {
	ID         int    `json:"id"`
	UserName   string `json:"userName"`
	ExternalID string `json:"externalId"`
	Email      string `json:"email"`
	// TODO: extend User struct
}

// Me returns currently authenticated user
func (u User) Me(c Client) string {
	endpoint := c.Domain + userPaths["USERINFO"]
	resp, err := get(c.AccessToken, endpoint)
	_ = err
	return string(resp) // TODO: return User struct
}

// Details returns info about user found by ID
func (u User) Details(c Client, id int) User {
	endpoint := c.Domain + userPaths["USERDETAILS"] + strconv.Itoa(id)
	resp, err := get(c.AccessToken, endpoint)

	err = json.Unmarshal(resp, &u)
	if err != nil {
		fmt.Println("error:", err)
	}

	return User{
		u.ID,
		u.UserName,
		u.ExternalID,
		u.Email}
}
