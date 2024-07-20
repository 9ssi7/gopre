package server

// Server is the interface that must be implemented by a server
// It is used to listen for requests
// The Listen method is used to listen for requests
// The Listen method returns an error
// The error is nil if the server is listening successfully
type Listener interface {
	// Listen is used to listen for requests
	// The Listen method returns an error
	// The error is nil if the server is listening successfully
	Listen() error
}
