package cloudElements

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
)

type Provider int

func (p Provider) String() string {
	switch p {
	case GOOGLE_DRIVE:
		return "Google Drive"
	case DROPBOX:
		return "DropBox"
	default:
		return ""
	}
}

const (
	GOOGLE_DRIVE Provider = iota
	DROPBOX
)

// Folders defines all methods to access/manipulate the Folders endpoint
type Folders struct {
	client *Client
}

// Files is a slice of File
type Files []File

// CloudFile contains all meta data about a file
type File struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Path         string   `json:"path"`
	CreateDate   string   `json:"createDate"`
	Directory    bool     `json:"directory"`
	ModifiedDate string   `json:"modifiedDate"`
	Size         int      `json:"size"`
	Tags         []string `json:"tags"`
}

// Find a file by name in a slice of files
func (f *Files) FindByName(name string) *File {
	for _, file := range *f {
		if file.Name == name {
			return &file
		}
	}
	return nil
}

// Contents will retrieve the contents of a folder
func (f Folders) Contents(folderPath string, provider Provider) (Files, *Error) {
	folderPath = path.Join("/", folderPath)
	query := url.Values{}
	query.Set("path", folderPath)

	u := folderContentsURL()
	u.RawQuery = query.Encode()

	req := newRequest(u, map[string]string{"Authorization": f.client.credentials.Authorization(provider)})

	resp, err := req.do("GET", nil)
	if err != nil {
		return Files{}, &Error{Message: err.Error()}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		e := parseError(resp.Body)
		e.StatusCode = resp.StatusCode
		return Files{}, e
	}
	var files Files

	if e := parse(resp.Body, &files); e != nil {
		return Files{}, e
	}
	return files, nil
}

// CreateFolder will create a folder
func (f Folders) CreateFolder(folderPath string, provider Provider) (*File, *Error) {
	folderPath = path.Join("/", folderPath)
	name := path.Base(folderPath)
	query := url.Values{}
	query.Set("path", folderPath)

	u := foldersURL()
	u.RawQuery = query.Encode()

	file := File{Name: name, Path: folderPath, Directory: true}
	b, err := json.Marshal(&file)
	if err != nil {
		return nil, &Error{Message: err.Error()}
	}
	headers := map[string]string{}
	headers["Authorization"] = f.client.credentials.Authorization(provider)
	headers["Content-Type"] = "application/json"

	req := newRequest(u, headers)

	buf := bytes.NewBuffer(b)
	resp, err := req.do("POST", buf)
	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		e := parseError(resp.Body)
		e.StatusCode = resp.StatusCode
		return nil, e
	}

	if e := parse(resp.Body, &file); e != nil {
		return nil, e
	}
	return &file, nil
}

// CreateFile will create a file
func (f Folders) CreateFile(fileInfo File, overwrite bool, reader io.ReadCloser, provider Provider) (*File, *Error) {
	defer reader.Close()
	filePath := path.Join("/", fileInfo.Path, fileInfo.Name)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return nil, &Error{Message: err.Error()}
	}
	_, err = io.Copy(part, reader)
	writer.WriteField("path", filePath)
	err = writer.Close()
	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	query := url.Values{}
	query.Set("path", filePath)
	if overwrite {
		query.Set("overwrite", "true")
	}

	u := filesURL()
	u.RawQuery = query.Encode()

	headers := map[string]string{}
	headers["Authorization"] = f.client.credentials.Authorization(provider)
	headers["Content-Type"] = "multipart/form-data"

	req := newRequest(u, headers)

	resp, err := req.do("POST", body)
	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		e := parseError(resp.Body)
		e.StatusCode = resp.StatusCode
		return nil, e
	}

	var file File
	if e := parse(resp.Body, &file); e != nil {
		return nil, e
	}
	return &file, nil
}

// GetFile will retrieve the file and return it as a reader
func (f Folders) GetFileById(id string, provider Provider) (io.ReadCloser, *Error) {
	u := filesURL()
	u.Path = path.Join(u.Path, id)

	req := newRequest(u, map[string]string{"Authorization": f.client.credentials.Authorization(provider)})

	resp, err := req.do("GET", nil)
	if err != nil {
		return nil, &Error{Message: err.Error()}
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		e := parseError(resp.Body)
		e.StatusCode = resp.StatusCode
		return nil, e
	}
	return resp.Body, nil
}
