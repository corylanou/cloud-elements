package cloudElements

import (
	"net/http"
	"net/url"
)

const (
	GOOGLE_DRIVE = iota
	DROPBOX
)

// CloudFiles defines all methods to access/manipulate the Folders endpoint
type CloudFiles struct {
	client *Client
}

// CloudFile contains all meta data about a file
type CloudFile struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Path         string   `json:"path"`
	CreateDate   string   `json:"createDate"`
	Directory    bool     `json:"directory"`
	ModifiedDate string   `json:"modifiedDate"`
	Size         int      `json:"size"`
	Tags         []string `json:"tags"`
}

// Contents will retrieve the contents of a folder
func (c CloudFiles) Contents(path string, provider int) ([]CloudFile, *Error) {
	if path == "" {
		path = "/"
	}
	query := url.Values{}
	query.Set("path", path)

	u := folderURL()
	u.RawQuery = query.Encode()

	req := newRequest(u, map[string]string{"Authorization": c.client.credentials.Authorization(provider)})

	resp, err := req.do()
	if err != nil {
		return []CloudFile{}, &Error{Message: err.Error()}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []CloudFile{}, parseError(resp.Body)
	}
	var cloudFiles []CloudFile

	if e := parse(resp.Body, &cloudFiles); e != nil {
		return []CloudFile{}, e
	}
	return cloudFiles, nil
}
