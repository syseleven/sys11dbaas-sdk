package sys11dbaassdk

import (
	"context"
	"net/http"

	databasev1 "github.com/syseleven/sys11dbaas-sdk/database/v1"
	databasev2 "github.com/syseleven/sys11dbaas-sdk/database/v2"
)

type Client struct {
	server         string
	requestEditors []func(context.Context, *http.Request) error

	v1Client *databasev1.TypedClient
	v2Client *databasev2.TypedClient
}

func NewClient(server string, options ...ClientOption) (*Client, error) {
	c := &Client{
		server: server,
	}

	for _, option := range options {
		option(c)
	}

	v1options := make([]databasev1.ClientOption, 0)
	for _, requestEditor := range c.requestEditors {
		v1options = append(v1options, databasev1.WithRequestEditorFn(requestEditor))
	}

	var err error
	c.v1Client, err = databasev1.NewTypedClient(c.server, v1options...)
	if err != nil {
		return nil, err
	}

	v2options := make([]databasev2.ClientOption, 0)
	for _, requestEditor := range c.requestEditors {
		v2options = append(v2options, databasev2.WithRequestEditorFn(requestEditor))
	}

	c.v2Client, err = databasev2.NewTypedClient(c.server, v2options...)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) V1() *databasev1.TypedClient {
	return c.v1Client
}

func (c *Client) V2() *databasev2.TypedClient {
	return c.v2Client
}

type ClientOption func(*Client) error

func WithApiKey(apiKey string) ClientOption {
	return func(c *Client) error {
		c.requestEditors = append(c.requestEditors, func(ctx context.Context, req *http.Request) error {
			req.Header.Add("x-s11-api-key", apiKey)
			return nil
		})
		return nil
	}
}

func WithServiceAccount(token string) ClientOption {
	return func(c *Client) error {
		c.requestEditors = append(c.requestEditors, func(ctx context.Context, req *http.Request) error {
			req.Header.Add("Authorization", "Bearer "+token)
			return nil
		})
		return nil
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.requestEditors = append(c.requestEditors, func(ctx context.Context, req *http.Request) error {
			req.Header.Add("User-Agent", userAgent)
			return nil
		})
		return nil
	}
}
