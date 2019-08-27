package main

var (
	name    = "null-provider"
	version = "1.0.0"
)

// Name returns the name of this plugin
func Name() string {
	return name
}

// Version returns the version of this plugin
func Version() string {
	return version
}

// Start signal
func Start() error {
	return nil
}

// Stop signal
func Stop() error {
	return nil
}

// Terminate signal
func Terminate() error {
	return nil
}

// SetKey ...
func SetKey(k string, v string) {

}

// GetKey ...
func GetKey(k string) (v string) {
	return
}
