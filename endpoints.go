package cloudElements

import (
	"net/url"
	"path"
)

const (
	baseURL = "console.cloud-elements.com/elements/api-v2"
)

func folderURL() url.URL {
	return url.URL{
		Path:   path.Join(baseURL, "/hubs/documents/folders/contents"),
		Scheme: "https",
	}
}
