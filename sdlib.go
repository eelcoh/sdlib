package sdlib

// Path ...
type Path struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

// Manifest ....
type Manifest struct {
	Paths    []*Path `json:"endpoints"`
	Hostname string  `json:"hostname"`
	IP       string  `json:"ip"`
	Version  string  `json:"version"`
}

// Endpoint
type Endpoint struct {
	Instance string `json:"instance"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Service  string `json:"service"`
	Method   string `json:"method"`
	Path     string `json:"path"`
}
