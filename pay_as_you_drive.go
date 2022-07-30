package merche

import (
	"context"
	"fmt"
	"net/http"
)

// VehicleStatus defines the response from VehicleStatus.
type PayAsYouDriveStatus struct {
	Odo *Resource `json:"odo,omitempty"`
}

// PayAsYouDriveService handles communication with vehicle status related
// methods of the Mercedes API.

// The Pay As You Drive Insurance API provide odometer status.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/vehicle_status/docs
type PayAsYouDriveService service

// GetPayAsYouDriveStatus gets containers resource of the Pay as you drive API
// to get the values of all resources that are available for readout.
// The response contains the available resource values and the corresponding
// readout timestamp for the corresponding car.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/pay_as_you_drive_insurance/docs#_3_get_all_values_of_the_pay_as_you_drive_insurance_api
func (s *PayAsYouDriveService) GetPayAsYouDriveStatus(ctx context.Context, opts *Options) ([]*PayAsYouDriveStatus, *Response, error) {
	path := fmt.Sprintf("%v/%v/containers/payasyoudrive", apiPathPrefix, opts.VehicleID)

	req, err := s.client.newRequest(ctx, http.MethodGet, path, http.NoBody)
	if err != nil {
		return nil, nil, err
	}

	var status []*PayAsYouDriveStatus
	resp, err := s.client.do(req, &status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}
