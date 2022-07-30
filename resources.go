package merche

import (
	"context"
	"fmt"
	"net/http"
)

const apiPathPrefix = "vehicledata/v2/vehicles"

// ResourcesService handles communication with vehicle available resources.
type ResourcesService service

// GetAvailableResources gets resources of the Mercedes API.
// It gets the resources that are available for readout.
//
// Mercedes API docs:
// https://developer.mercedes-benz.com/products/vehicle_status/docs#_3_get_all_values_of_the_vehicle_status_api
// https://developer.mercedes-benz.com/products/fuel_status/docs#_1_get_the_available_resources_that_can_be_read_out
func (s *ResourcesService) GetAvailableResources(ctx context.Context, opts *Options) ([]*ResourceMetaInfo, *Response, error) {
	path := fmt.Sprintf("%v/%v/resources", apiPathPrefix, opts.VehicleID)

	req, err := s.client.newRequest(ctx, http.MethodGet, path, http.NoBody)
	if err != nil {
		return nil, nil, err
	}

	var resources []*ResourceMetaInfo
	resp, err := s.client.do(req, &resources)
	if err != nil {
		return nil, resp, err
	}

	return resources, resp, nil
}
