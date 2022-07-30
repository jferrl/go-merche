package merche

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"testing"
)

type writer struct {
	data []byte
}

func (w *writer) Write(data []byte) (n int, err error) {
	w.data = data
	return len(data), nil
}

func TestClient_do(t *testing.T) {
	type fakeResponse struct{}

	type args struct {
		v any
	}

	tests := []struct {
		name            string
		mercedesAPIMock *httptest.Server
		args            args
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
			args: args{
				v: &fakeResponse{},
			},
		},
		{
			name:            "mercedes api response: nil decoding",
			mercedesAPIMock: createFakeServer(http.StatusOK, "vehicle_status_get_resources.json"),
			args: args{
				v: nil,
			},
		},
		{
			name:            "mercedes api response: copy to writer",
			mercedesAPIMock: createFakeServer(http.StatusOK, "vehicle_status_get_resources.json"),
			args: args{
				v: new(writer),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, _ := url.Parse(tt.mercedesAPIMock.URL + "/")

			c := NewClient(nil)
			c.BaseURL = url
			req, _ := c.newRequest(context.Background(), http.MethodGet, "", http.NoBody)

			_, err := c.do(req, tt.args.v)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Client.do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func createFakeServer(statusCode int, res string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		http.ServeFile(w, r, filepath.Join("testdata", res))
	}))
}
