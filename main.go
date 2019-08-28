package main

const (
	name    = "null-provider"
	version = "1.0.0"
)

type provider string

// Unused method
func main() {}

// Name returns the name of this provider
func (p *provider) Name() string {
	return name
}

// Version returns the version of this provider
func (p *provider) Version() string {
	return version
}

// Action is used to signal this provider to perform an action
func (p *provider) Action(action string) error {
	return nil
}

// SetKey is used to set data in this provider
func (p *provider) SetKey(k string, v string) {

}

// GetKey is used to retrieve data from this provider
func (p *provider) GetKey(k string) (v string) {
	return
}

// Provider exported symbol
var Provider provider
