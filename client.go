package cloudElements

// Client is the top level connector to Cloud Elements
type Client struct {
	credentials Credentials
	CloudFiles  CloudFiles
}

// NewClient will return a new instance of Client
func NewClient(cr Credentials) *Client {
	c := Client{
		credentials: cr,
	}

	c.CloudFiles = CloudFiles{client: &c}

	return &c
}
