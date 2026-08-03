package sys11dbaassdk

import (
	"context"
	"net/http"

	databasev1 "github.com/syseleven/sys11dbaas-sdk/database/v1"
	databasev2 "github.com/syseleven/sys11dbaas-sdk/database/v2"
)

type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	server string

	v1Client  *databasev1.TypedClient
	v1Options []databasev1.ClientOption
	v2Client  *databasev2.TypedClient
	v2Options []databasev2.ClientOption
}

func NewClient(server string, options ...ClientOption) (*Client, error) {
	c := &Client{
		server:    server,
		v1Options: make([]databasev1.ClientOption, 0),
		v2Options: make([]databasev2.ClientOption, 0),
	}

	for _, option := range options {
		option(c)
	}

	var err error
	c.v1Client, err = databasev1.NewTypedClient(c.server, c.v1Options...)
	if err != nil {
		return nil, err
	}

	c.v2Client, err = databasev2.NewTypedClient(c.server, c.v2Options...)
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
		fn := func(ctx context.Context, req *http.Request) error {
			req.Header.Add("x-s11-api-key", apiKey)
			return nil
		}
		c.v1Options = append(c.v1Options, databasev1.WithRequestEditorFn(fn))
		c.v2Options = append(c.v2Options, databasev2.WithRequestEditorFn(fn))
		return nil
	}
}

func WithServiceAccount(token string) ClientOption {
	return func(c *Client) error {
		fn := func(ctx context.Context, req *http.Request) error {
			req.Header.Add("Authorization", "Bearer "+token)
			return nil
		}
		c.v1Options = append(c.v1Options, databasev1.WithRequestEditorFn(fn))
		c.v2Options = append(c.v2Options, databasev2.WithRequestEditorFn(fn))
		return nil
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		fn := func(ctx context.Context, req *http.Request) error {
			req.Header.Add("User-Agent", userAgent)
			return nil
		}
		c.v1Options = append(c.v1Options, databasev1.WithRequestEditorFn(fn))
		c.v2Options = append(c.v2Options, databasev2.WithRequestEditorFn(fn))
		return nil
	}
}

func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.v1Options = append(c.v1Options, databasev1.WithHTTPClient(doer))
		c.v2Options = append(c.v2Options, databasev2.WithHTTPClient(doer))
		return nil
	}
}
