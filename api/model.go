package api

// Volume is a struct representation of the Ubuntu desktop's current volume
type Volume struct {
	Value string `json:"value"`
	Muted bool   `json:"muted"`
}
