package merche

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fakeVehicleID = "EXVETESTVIN000001"
)

func TestVehicleStatusService_GetVehicleStatus(t *testing.T) {
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
		want    []*VehicleStatus
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
			name: "get containers status",
			fields: fields{
				mercedesAPIMock: createFakeServer(http.StatusOK, "vehicle_status_get_containers.json"),
			},
			args: args{
				ctx: ctx,
				opts: &Options{
					VehicleID: fakeVehicleID,
				},
			},
			want: []*VehicleStatus{
				{
					Doorlockstatusdecklid: &Resource{
						Value:     String("true"),
						Timestamp: Int64(1541406596000),
					},
				},
				{
					Doorstatusfrontleft: &Resource{
						Value:     String("false"),
						Timestamp: Int64(1541751294000),
					},
				},
				{
					Doorstatusfrontright: &Resource{
						Value:     String("true"),
						Timestamp: Int64(1541751278000),
					},
				},
				{
					Doorstatusrearleft: &Resource{
						Value:     String("false"),
						Timestamp: Int64(1541751294000),
					},
				},
				{
					Doorstatusrearright: &Resource{
						Value:     String("false"),
						Timestamp: Int64(1541751263000),
					},
				},
				{
					InteriorLightsFront: &Resource{
						Value:     String("false"),
						Timestamp: Int64(1541751263000),
					},
				},
				{
					InteriorLightsRear: &Resource{
						Value:     String("false"),
						Timestamp: Int64(1541751263000),
					},
				},
				{
					Lightswitchposition: &Resource{
						Value:     String("2"),
						Timestamp: Int64(1541751263000),
					},
				},
				{
					ReadingLampFrontLeft: &Resource{
						Value:     String("false"),
						Timestamp: Int64(1541751263000),
					},
				},
				{
					ReadingLampFrontRight: &Resource{
						Value:     String("false"),
						Timestamp: Int64(1541751263000),
					},
				},
				{
					Rooftopstatus: &Resource{
						Value:     String("1"),
						Timestamp: Int64(1541751263000),
					},
				},
				{
					Sunroofstatus: &Resource{
						Value:     String("1"),
						Timestamp: Int64(1541751044000),
					},
				},
				{
					Windowstatusfrontleft: &Resource{
						Value:     String("3"),
						Timestamp: Int64(1541751248000),
					},
				},
				{
					Windowstatusfrontright: &Resource{
						Value:     String("2"),
						Timestamp: Int64(1541751263000),
					},
				},
				{
					Windowstatusrearleft: &Resource{
						Value:     String("5"),
						Timestamp: Int64(1541751073000),
					},
				},
				{
					Windowstatusrearright: &Resource{
						Value:     String("4"),
						Timestamp: Int64(1541751161000),
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

			got, _, err := c.VehicleStatus.GetVehicleStatus(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("VehicleStatus.GetVehicleStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "VehicleStatus.GetVehicleStatus() got = %v, want %v", got, tt.want)
		})
	}
}
