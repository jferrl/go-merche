package merche

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
