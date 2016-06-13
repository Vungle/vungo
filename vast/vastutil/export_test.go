package vastutil

import "net/http"

// UpdateDefaultUnwrapClient method updates the unwrap HTTP client to a given http.Client object.
func UpdateDefaultUnwrapClient(c *http.Client) {
	defaultUnwrapClient = c
}

// RevertDefaultUnwrapClient method reverts unwrap HTTP client back to the default one.
func RevertDefaultUnwrapClient() {
	defaultUnwrapClient = http.DefaultClient
}
