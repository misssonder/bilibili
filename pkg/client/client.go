package client

import (
	"io"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	cookie     []string
}

func (client *Client) newCookieRequest(method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for _, c := range client.cookie {
		request.Header.Add("Cookie", c)
	}
	return request, nil
}

func (client *Client) SetCookie(cookie []string) {
	client.cookie = cookie
}
