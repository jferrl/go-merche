package merche

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_do(t *testing.T) {
	type FakeResponse struct{}

	tests := []struct {
		name            string
		mercedesAPIMock *httptest.Server
		want            *Response
		wantErr         error
	}{
		{
			name:            "auth error response",
			mercedesAPIMock: createFakeServer(http.StatusUnauthorized, "auth_error.json"),
			wantErr: &UnauthorizedError{
				ErrorMessage: "Unauthorized",
				StatusCode:   "401",
				Message:      "Token invalid: Not active",
			},
		},
		{
			name:            "exve error response",
			mercedesAPIMock: createFakeServer(http.StatusBadRequest, "exve_error.json"),
			wantErr: &ExVeError{
				ExveErrorID:  "Id",
				ExveErrorMsg: "Msg",
				ExveErrorRef: "Ref",
			},
		},
		{
			name:            "api error",
			mercedesAPIMock: createFakeServer(http.StatusNotFound, ""),
			wantErr: &MercedesAPIError{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name:            "mercedes not content response",
			mercedesAPIMock: createFakeServer(http.StatusNoContent, ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, _ := url.Parse(tt.mercedesAPIMock.URL + "/")

			c := NewClient(nil)
			c.BaseURL = url
			req, _ := c.newRequest(context.Background(), http.MethodGet, "", http.NoBody)

			var fr FakeResponse

			_, err := c.do(req, &fr)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Client.do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
