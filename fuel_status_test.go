package merche

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuelStatusService_GetFuelStatus(t *testing.T) {
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
		want    []*FuelStatus
		wantErr bool
	}{
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
			name: "get containers status",
			fields: fields{
				mercedesAPIMock: createFakeServer(http.StatusOK, "fuel_status_get_containers.json"),
			},
			args: args{
				ctx: ctx,
				opts: &Options{
					VehicleID: fakeVehicleID,
				},
			},
			want: []*FuelStatus{
				{
					RangeLiquid: &Resource{
						Value:     String("1648"),
						Timestamp: Int64(1541406596000),
					},
				},
				{
					TankLevelPercent: &Resource{
						Value:     String("84"),
						Timestamp: Int64(1541233886000),
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

			got, _, err := c.FuelStatus.GetFuelStatus(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("FuelStatus.GetFuelStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "FuelStatus.GetFuelStatus() got = %v, want %v", got, tt.want)
		})
	}
}
