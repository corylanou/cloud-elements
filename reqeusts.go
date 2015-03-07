package cloudElements

import (
	"net/http"
	"net/url"
)

func newRequest(u url.URL, headers map[string]string) *request {
	return &request{
		u:       u,
		headers: headers,
	}
}

type request struct {
	u       url.URL
	headers map[string]string
}

func (r *request) do() (*http.Response, error) {
	client := http.Client{}
	req, _ := http.NewRequest("GET", r.u.String(), nil)
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}

	return client.Do(req)
}
