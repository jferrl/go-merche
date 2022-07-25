package merche

import (
	"context"
	"fmt"
	"net/http"
)

const vehicleStatusPathPrefix = "vehicledata/v2/vehicles"

// ResourceMetaInfo struct for ResourceMetaInfo.
type ResourceMetaInfo struct {
	Href    *string `json:"href,omitempty"`
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// Resource struct for Resource.
type Resource struct {
	Timestamp *int64  `json:"timestamp,omitempty"`
	Value     *string `json:"value,omitempty"`
}

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

// GetVehicleStatusOptions .
type GetVehicleStatusOptions struct {
	VehicleID string
}

// VehicleStatusService handles communication with vehicle status related
// methods of the Mercedes API.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/vehicle_status/docs#_2_get_the_value_of_a_specific_resource
type VehicleStatusService service

// GetAvailableResources gets resources of the Vehicle Status API
// to get the resources that are available for readout.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/vehicle_status/docs#_3_get_all_values_of_the_vehicle_status_api
func (s *VehicleStatusService) GetAvailableResources(ctx context.Context, opts *GetVehicleStatusOptions) ([]*ResourceMetaInfo, *Response, error) {
	path := fmt.Sprintf("%v/%v/resources", vehicleStatusPathPrefix, opts.VehicleID)

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

// GetVehicleStatus gets containers resource of the Vehicle Status API
// to get the values of all resources that are available for readout.
// The response contains the available resource values and the corresponding
// readout timestamp for the corresponding car.
//
// Mercedes API docs: https://developer.mercedes-benz.com/products/vehicle_status/docs#_3_get_all_values_of_the_vehicle_status_api
func (s *VehicleStatusService) GetContainersVehicleStatus(ctx context.Context, opts *GetVehicleStatusOptions) ([]*VehicleStatus, *Response, error) {
	path := fmt.Sprintf("%v/%v/containers/vehiclestatus", vehicleStatusPathPrefix, opts.VehicleID)

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
