package cloudElements

// Client is the top level connector to Cloud Elements
type Client struct {
	credentials Credentials
	Folders     Folders
}

// NewClient will return a new instance of Client
func NewClient(cr Credentials) *Client {
	c := Client{
		credentials: cr,
	}

	c.Folders = Folders{client: &c}

	return &c
}
