package merche

import (
	"context"
	"fmt"
	"net/http"
)

type ElectricVehicleStatus struct {
	Soc           *Resource `json:"soc,omitempty"`
	RangeElectric *Resource `json:"rangeelectric,omitempty"`
}

// ElectricVehicleStatusService handles communication with electric vehicle status related
// methods of the Mercedes API.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/electric_vehicle_status/docs
type ElectricVehicleStatusService service

// GetElectricVehicleStatus gets containers resource of the Electric Vehicle Status API
// to get the values of all resources that are available for readout.
// The response contains the available resource values and the corresponding
// readout timestamp for the corresponding car.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/electric_vehicle_status/specifications/electric_vehicle_status_api
func (s *ElectricVehicleStatusService) GetElectricVehicleStatus(ctx context.Context, opts *Options) ([]*ElectricVehicleStatus, *Response, error) {
	path := fmt.Sprintf("%v/%v/containers/electricvehicle", apiPathPrefix, opts.VehicleID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, http.NoBody)
	if err != nil {
		return nil, nil, err
	}

	var status []*ElectricVehicleStatus
	resp, err := s.client.Do(req, &status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}
