package data

type Availabilities struct {
	Region      string       `json:"region"`
	Hardware    string       `json:"availability"`
	Datacenters []Datacenter `json:"datacenters"`
}
