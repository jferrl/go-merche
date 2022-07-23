package merche

import "net/http"

// Response is a Mercedes API response. This wraps the standard http.Response
// returned from Mercedes.
type Response struct {
	*http.Response
}
