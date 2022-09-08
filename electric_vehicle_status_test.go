package merche

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElectricVehicleStatusService_GetElectricVehicleStatus(t *testing.T) {
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
		want    []*ElectricVehicleStatus
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
				mercedesAPIMock: createFakeServer(http.StatusOK, "electric_vehicle_status_get_containers.json"),
			},
			args: args{
				ctx: ctx,
				opts: &Options{
					VehicleID: fakeVehicleID,
				},
			},
			want: []*ElectricVehicleStatus{
				{
					Soc: &Resource{
						Value:     String("35"),
						Timestamp: Int64(1541749824000),
					},
				},
				{
					RangeElectric: &Resource{
						Value:     String("1021"),
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

			got, _, err := c.ElectricVehicleStatus.GetElectricVehicleStatus(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("ElectricVehicleStatus.GetElectricVehicleStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "ElectricVehicleStatus.GetElectricVehicleStatus() got = %v, want %v", got, tt.want)
		})
	}
}
