package egnyte

import "strings"

// TODO: tests
func sanitizeDomain(domain string) string {
	path := ".egnyte.com"
	protocol := "https://"

	if !strings.Contains(domain, path) {
		domain = domain + path
	}

	if !strings.Contains(domain, protocol) {
		domain = protocol + domain
	}

	// check domain validity
	// GET <domain>/rest/public/1.0/env-pub --> return 200

	return domain
}
