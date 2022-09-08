package merche

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayAsYouDriveService_GetPayAsYouDriveStatus(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		mercedesAPIMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts *Options
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*PayAsYouDriveStatus
		wantErr bool
	}{
		{
			name: "nil context error",
			fields: fields{
				mercedesAPIMock: createFakeServer(http.StatusOK, ""),
			},
			args: args{
				ctx: nil,
				opts: &Options{
					VehicleID: fakeVehicleID,
				},
			},
			wantErr: true,
		},
		{
			name: "decoding response error",
			fields: fields{
				mercedesAPIMock: createFakeServer(http.StatusOK, "invalid_response"),
			},
			args: args{
				ctx: ctx,
				opts: &Options{
					VehicleID: fakeVehicleID,
				},
			},
			wantErr: true,
		},
		{
			name: "get containers",
			fields: fields{
				mercedesAPIMock: createFakeServer(http.StatusOK, "pay_as_you_drive_get_containers.json"),
			},
			args: args{
				ctx: ctx,
				opts: &Options{
					VehicleID: fakeVehicleID,
				},
			},
			want: []*PayAsYouDriveStatus{
				{
					Odo: &Resource{
						Value:     String("319947"),
						Timestamp: Int64(1541749824000),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.mercedesAPIMock.URL + "/")

			c := NewClient(tt.fields.mercedesAPIMock.Client())
			c.BaseURL = baseURL

			got, _, err := c.PayAsYouDrive.GetPayAsYouDriveStatus(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("PayAsYouDrive.GetPayAsYouDriveStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "PayAsYouDrive.GetPayAsYouDriveStatus() got = %v, want %v", got, tt.want)
		})
	}
}
