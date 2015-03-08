# cloudElements
--
    import "github.com/corylanou/cloud-elements"


## Usage

#### type Client

```go
type Client struct {
}
```

Client is the top level connector to Cloud Elements

#### func  NewClient

```go
func NewClient(cr Credentials) *Client
```
NewClient will return a new instance of Client

#### type Credentials

```go
type Credentials struct {
	Elements     map[Provider]string
	Organization string
	User         string // User Secret
}
```

Credentials stores all token information to authenticate to Cloud Elements, as
well as each element

#### func (Credentials) Authorization

```go
func (c Credentials) Authorization(element Provider) string
```
Authorization is formatted as: User <secret>, Organization <token>, Element
<token>

#### type Error

```go
type Error struct {
	Message    string `json:"message"`
	RequestId  string `json:"requestId"`
	StatusCode int
}
```

Error is a generic error returned from Cloud Elements, or this client libarary

#### type File

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

#### type Files

```go
type Files []File
```

Files is a slice of File

#### func (*Files) FindByName

```go
func (f *Files) FindByName(name string) *File
```
Find a file by name in a slice of files

#### type Folders

```go
type Folders struct {
}
```

Folders defines all methods to access/manipulate the Folders endpoint

#### func  NewFolders

```go
func NewFolders(c *Client) *Folders
```
NewFolders returns a folders object to manipulate the Docuemnts endpoints

#### func (Folders) Contents

```go
func (f Folders) Contents(folderPath string, provider Provider) (Files, *Error)
```
Contents will retrieve the contents of a folder

#### func (Folders) CreateFile

```go
func (f Folders) CreateFile(fileInfo File, overwrite bool, reader io.ReadCloser, provider Provider) (*File, *Error)
```
CreateFile will create a file

#### func (Folders) CreateFolder

```go
func (f Folders) CreateFolder(folderPath string, provider Provider) (*File, *Error)
```
CreateFolder will create a folder

#### func (Folders) GetFileById

```go
func (f Folders) GetFileById(id string, provider Provider) (io.ReadCloser, *Error)
```
GetFile will retrieve the file and return it as a reader

#### type Provider

```go
type Provider int
```


```go
const (
	GOOGLE_DRIVE Provider = iota
	DROPBOX
)
```

#### func (Provider) String

```go
func (p Provider) String() string
```
