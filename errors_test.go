package merche

import (
	"net/http"
	"testing"
)

func Test_handleResponseError(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handleResponseError(tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("handleResponseError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
