package rpc

var (
	// Services is a list of exported services
	// registered by the RPC server at runtime
	Services = []interface{}{
		Ping{},
	}
)

// Ping is a basic RPC service for checking the health of a node
type Ping struct{}

// Ping an rpc server
func (p Ping) Ping(data []byte, ack *bool) error {
	*ack = true
	return nil
}
