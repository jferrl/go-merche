package merche

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fakeVehicleID = "EXVETESTVIN000001"
)

func TestVehicleStatusService_GetAvailableResources(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		mercedesAPIMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts *GetVehicleStatusOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*ResourceMetaInfo
		wantErr bool
	}{
		{
			name: "decoding response error",
			fields: fields{
				mercedesAPIMock: createFakeServer(http.StatusOK, "invalid_response"),
			},
			args: args{
				ctx: ctx,
				opts: &GetVehicleStatusOptions{
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
				opts: &GetVehicleStatusOptions{
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

			got, _, err := c.VehicleStatus.GetAvailableResources(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("VehicleStatusService.GetAvailableResources() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "VehicleStatusService.GetAvailableResources() got = %v, want %v", got, tt.want)
		})
	}
}

func TestVehicleStatusService_GetContainersVehicleStatus(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		mercedesAPIMock *httptest.Server
	}
	type args struct {
		ctx  context.Context
		opts *GetVehicleStatusOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*VehicleStatus
		wantErr bool
	}{
		{
			name: "decoding response error",
			fields: fields{
				mercedesAPIMock: createFakeServer(http.StatusOK, "invalid_response"),
			},
			args: args{
				ctx: ctx,
				opts: &GetVehicleStatusOptions{
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
				opts: &GetVehicleStatusOptions{
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

			got, _, err := c.VehicleStatus.GetContainersVehicleStatus(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("VehicleStatusService.GetContainersVehicleStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "VehicleStatusService.GetContainersVehicleStatus() got = %v, want %v", got, tt.want)
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
