package model

// Location represents a location containing a latitude and longitude
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
