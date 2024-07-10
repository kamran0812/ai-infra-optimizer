package cloud

import "time"

// ResourceUsage represents the usage data for a single cloud resource
type ResourceUsage struct {
	Provider   string
	ResourceID string
	Type       string
	CPU        float64
	Memory     float64
	Timestamp  time.Time
}

// Provider interface defines the methods that each cloud provider must implement
type Provider interface {
	GetName() string
	GetResourceUsage() ([]ResourceUsage, error)
}