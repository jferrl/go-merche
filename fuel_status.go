package merche

import (
	"context"
	"fmt"
	"net/http"
)

// VehicleStatus defiles the response from VehicleStatus.
type FuelStatus struct {
	RangeLiquid      *Resource `json:"rangeliquid,omitempty"`
	TankLevelPercent *Resource `json:"tanklevelpercent,omitempty"`
}

// FuelStatusService handles communication with fuel status related
// methods of the Mercedes API.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/fuel_status/docs
type FuelStatusService service

// GetFuelStatus gets containers resource of the Fuel Status API
// to get the values of all resources that are available for readout.
// The response contains the available resource values and the corresponding
// readout timestamp for the corresponding car.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/fuel_status/docs#_3_get_all_values_of_the_fuel_status_api
func (s *FuelStatusService) GetFuelStatus(ctx context.Context, opts *Options) ([]*FuelStatus, *Response, error) {
	path := fmt.Sprintf("%v/%v/containers/fuelstatus", apiPathPrefix, opts.VehicleID)

	req, err := s.client.newRequest(ctx, http.MethodGet, path, http.NoBody)
	if err != nil {
		return nil, nil, err
	}

	var status []*FuelStatus
	resp, err := s.client.do(req, &status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}
