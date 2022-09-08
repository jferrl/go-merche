package merche

import (
	"context"
	"fmt"
	"net/http"
)

// VehicleLockStatus defines the response from VehicleLockStatus.
type VehicleLockStatus struct {
	Doorlockstatusvehicle *Resource `json:"doorlockstatusvehicle,omitempty"`
	Doorlockstatusdecklid *Resource `json:"doorlockstatusdecklid,omitempty"`
	Doorlockstatusgas     *Resource `json:"doorlockstatusgas,omitempty"`
	PositionHeading       *Resource `json:"positionHeading,omitempty"`
}

// VehicleLockStatusService handles communication with vehicle lock status related
// methods of the Mercedes API.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/vehicle_lock_status/docs
type VehicleLockStatusService service

// GetVehicleLockStatus gets containers resource of the Vehicle lock tatus API
// to get the values of all resources that are available for readout.
// The response contains the available resource values and the corresponding
// readout timestamp for the corresponding car.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/vehicle_lock_status/docs#_3_get_all_values_of_the_vehicle_lock_status_api
func (s *VehicleLockStatusService) GetVehicleLockStatus(ctx context.Context, opts *Options) ([]*VehicleLockStatus, *Response, error) {
	path := fmt.Sprintf("%v/%v/containers/vehiclelockstatus", apiPathPrefix, opts.VehicleID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, http.NoBody)
	if err != nil {
		return nil, nil, err
	}

	var status []*VehicleLockStatus
	resp, err := s.client.Do(req, &status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}
