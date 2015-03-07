package cloudElements

// Client is the top level connector to Cloud Elements
type Client struct {
	credentials Credentials
}

// NewClient will return a new instance of Client
func NewClient(cr Credentials) *Client {
	return &Client{
		credentials: cr,
	}
}
