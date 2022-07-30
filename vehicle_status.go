package merche

import (
	"context"
	"fmt"
	"net/http"
)

// VehicleStatus defiles the response from VehicleStatus.
type VehicleStatus struct {
	Doorlockstatusdecklid  *Resource `json:"doorlockstatusdecklid,omitempty"`
	Doorstatusfrontleft    *Resource `json:"doorstatusfrontleft,omitempty"`
	Doorstatusfrontright   *Resource `json:"doorstatusfrontright,omitempty"`
	Doorstatusrearleft     *Resource `json:"doorstatusrearleft,omitempty"`
	Doorstatusrearright    *Resource `json:"doorstatusrearright,omitempty"`
	InteriorLightsFront    *Resource `json:"interiorLightsFront,omitempty"`
	InteriorLightsRear     *Resource `json:"interiorLightsRear,omitempty"`
	Lightswitchposition    *Resource `json:"lightswitchposition,omitempty"`
	ReadingLampFrontLeft   *Resource `json:"readingLampFrontLeft,omitempty"`
	ReadingLampFrontRight  *Resource `json:"readingLampFrontRight,omitempty"`
	Rooftopstatus          *Resource `json:"rooftopstatus,omitempty"`
	Sunroofstatus          *Resource `json:"sunroofstatus,omitempty"`
	Windowstatusfrontleft  *Resource `json:"windowstatusfrontleft,omitempty"`
	Windowstatusfrontright *Resource `json:"windowstatusfrontright,omitempty"`
	Windowstatusrearleft   *Resource `json:"windowstatusrearleft,omitempty"`
	Windowstatusrearright  *Resource `json:"windowstatusrearright,omitempty"`
}

// VehicleStatusService handles communication with vehicle status related
// methods of the Mercedes API.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/vehicle_status/docs
type VehicleStatusService service

// GetVehicleStatus gets containers resource of the Vehicle Status API
// to get the values of all resources that are available for readout.
// The response contains the available resource values and the corresponding
// readout timestamp for the corresponding car.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/vehicle_status/docs#_3_get_all_values_of_the_vehicle_status_api
func (s *VehicleStatusService) GetVehicleStatus(ctx context.Context, opts *Options) ([]*VehicleStatus, *Response, error) {
	path := fmt.Sprintf("%v/%v/containers/vehiclestatus", apiPathPrefix, opts.VehicleID)

	req, err := s.client.newRequest(ctx, http.MethodGet, path, http.NoBody)
	if err != nil {
		return nil, nil, err
	}

	var status []*VehicleStatus
	resp, err := s.client.do(req, &status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}
