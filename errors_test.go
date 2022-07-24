package merche

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handleResponseError(t *testing.T) {
	tests := []struct {
		name            string
		mercedesAPIMock *httptest.Server
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
			name:            "mercedes api error",
			mercedesAPIMock: createFakeServer(http.StatusNoContent, ""),
			wantErr: &MercedesAPIError{
				StatusCode: http.StatusNoContent,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockedClient := tt.mercedesAPIMock.Client()
			resp, _ := mockedClient.Get(tt.mercedesAPIMock.URL)
			if err := handleResponseError(resp); err.Error() != tt.wantErr.Error() {
				t.Errorf("handleResponseError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
