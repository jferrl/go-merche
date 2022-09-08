package merche

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVehicleStatusService_GetAvailableResources(t *testing.T) {
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
		want    []*ResourceMetaInfo
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
			name: "get available resources",
			fields: fields{
				mercedesAPIMock: createFakeServer(http.StatusOK, "vehicle_status_get_resources.json"),
			},
			args: args{
				ctx: ctx,
				opts: &Options{
					VehicleID: fakeVehicleID,
				},
			},
			want: []*ResourceMetaInfo{
				{
					Name:    String("doorlockstatusdecklid"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/doorlockstatusdecklid"),
				},
				{
					Name:    String("doorstatusfrontleft"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/doorstatusfrontleft"),
				},
				{
					Name:    String("doorstatusfrontright"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/doorstatusfrontright"),
				},
				{
					Name:    String("doorstatusrearleft"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/doorstatusrearleft"),
				},
				{
					Name:    String("doorstatusrearright"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/doorstatusrearright"),
				},
				{
					Name:    String("interiorLightsFront"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/interiorLightsFront"),
				},
				{
					Name:    String("interiorLightsRear"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/interiorLightsRear"),
				},
				{
					Name:    String("lightswitchposition"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/lightswitchposition"),
				},
				{
					Name:    String("readingLampFrontLeft"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/readingLampFrontLeft"),
				},
				{
					Name:    String("readingLampFrontRight"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/readingLampFrontRight"),
				},
				{
					Name:    String("rooftopstatus"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/rooftopstatus"),
				},
				{
					Name:    String("sunroofstatus"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/sunroofstatus"),
				},
				{
					Name:    String("windowstatusfrontleft"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/windowstatusfrontleft"),
				},
				{
					Name:    String("windowstatusfrontright"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/windowstatusfrontright"),
				},
				{
					Name:    String("windowstatusrearleft"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/windowstatusrearleft"),
				},
				{
					Name:    String("windowstatusrearright"),
					Version: String("1.0"),
					Href:    String("/vehicles/EXVETESTVIN000001/resources/windowstatusrearright"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse(tt.fields.mercedesAPIMock.URL + "/")

			c := NewClient(tt.fields.mercedesAPIMock.Client())
			c.BaseURL = baseURL

			got, _, err := c.Resources.GetAvailableResources(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resources.GetAvailableResources() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "Resources.GetAvailableResources() got = %v, want %v", got, tt.want)
		})
	}
}
