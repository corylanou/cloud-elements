package cloudElements

import (
	"net/url"
	"path"
)

const (
	baseURL = "console.cloud-elements.com/elements/api-v2"
)

func filesURL() url.URL {
	return url.URL{
		Path:   path.Join(baseURL, "/hubs/documents/files"),
		Scheme: "https",
	}
}
func foldersURL() url.URL {
	return url.URL{
		Path:   path.Join(baseURL, "/hubs/documents/folders"),
		Scheme: "https",
	}
}

func folderContentsURL() url.URL {
	return url.URL{
		Path:   path.Join(baseURL, "/hubs/documents/folders/contents"),
		Scheme: "https",
	}
}
