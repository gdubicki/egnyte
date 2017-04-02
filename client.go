package egnyte

// Client represents Egnyte client
type Client struct {
	AccessToken string
	Domain      string
	User
}

func getClient(token, domain string) Client {
	egnyteDomain := sanitizeDomain(domain)
	return Client{
		token,
		egnyteDomain,
		User{}}
}
