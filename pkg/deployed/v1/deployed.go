package v1

// DeployedClient is the client to interact with deployed
type Deployed struct {
}

type ClientOption func(*Deployed)

func NewClient(opts ...ClientOption) *Deployed {
	client := &Deployed{}
	for _, opt := range opts {
		opt(client)
	}

	return client
}
