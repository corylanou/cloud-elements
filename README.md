# Cloud Elements API

This library is intended to be used by the Cloud Elements API.

Currently work on progress.

##PACKAGE DOCUMENTATION

###package cloudElements
```go
import "github.com/corylanou/cloud-elements"
```


### TYPES

```go
type Client struct {
    Folders Folders
    // contains filtered or unexported fields
}
```
Client is the top level connector to Cloud Elements

```go
func NewClient(cr Credentials) *Client
```
NewClient will return a new instance of Client

```go
type Credentials struct {
    Elements     map[Provider]string
    Organization string
    User         string // User Secret
}
```
Credentials stores all token information to authenticate to Cloud
Elements, as well as each element

```go
func (c Credentials) Authorization(element Provider) string
```
Authorization is formatted as: User <secret>, Organization <token>,
Element <token>

```go
type Error struct {
    Message    string `json:"message"`
    RequestId  string `json:"requestId"`
    StatusCode int
}
```
Error is a generic error returned from Cloud Elements, or this client
library

```go
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
```
CloudFile contains all meta data about a file

```go
type Files []File
```
Files is a slice of File

```go
func (f *Files) FindByName(name string) *File
```
Find a file by name in a slice of files

```go
type Folders struct {
    // contains filtered or unexported fields
}
```
Folders defines all methods to access/manipulate the Folders endpoint

```go
func (f Folders) Contents(folderPath string, provider Provider) (Files, *Error)
```
Contents will retrieve the contents of a folder

```go
func (f Folders) CreateFile(fileInfo File, overwrite bool, reader io.ReadCloser, provider Provider) (*File, *Error)
```
CreateFile will create a file

```go
func (f Folders) CreateFolder(folderPath string, provider Provider) (*File, *Error)
```
CreateFolder will create a folder

```go
func (f Folders) GetFileById(id string, provider Provider) (io.ReadCloser, *Error)
```
GetFile will retrieve the file and return it as a reader

```go
type Provider int

const (
    GOOGLE_DRIVE Provider = iota
    DROPBOX
)

func (p Provider) String() string
```


