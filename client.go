package sys11dbaassdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type AuthMode string

const (
	AuthModeApiKey AuthMode = "apikey"
	AuthModeSSO    AuthMode = "sso"
)

type errorMsg struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type Client struct {
	baseUrl  string
	apiKey   string
	user     string
	client   *http.Client
	agent    string
	authMode AuthMode
}

func NewClient(baseurl, apikey, agent string, authMode AuthMode) (*Client, error) {
	client := &Client{
		baseUrl:  baseurl,
		apiKey:   apikey,
		user:     apikey,
		agent:    agent,
		authMode: authMode,
	}
	client.client = &http.Client{
		Timeout: 60 * time.Second,
	}
	return client, nil
}

func (c *Client) get(path string, verbose bool) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseUrl+path, nil)
	if err != nil {
		return nil, err
	}
	return c.doReq(req, verbose)
}

func (c *Client) delete(path string, verbose bool) ([]byte, error) {
	req, err := http.NewRequest(http.MethodDelete, c.baseUrl+path, nil)
	if err != nil {
		return nil, err
	}
	return c.doReq(req, verbose)
}

func (c *Client) post(path string, data []byte, verbose bool) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, c.baseUrl+path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return c.doReq(req, verbose)
}

func (c *Client) patch(path string, data []byte, verbose bool) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPatch, c.baseUrl+path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return c.doReq(req, verbose)
}

func (c *Client) doReq(req *http.Request, verbose bool) ([]byte, error) {
	req.Header.Add("Accept", "application/json")
	if c.authMode == AuthModeApiKey {
		req.Header.Add("x-s11-api-key", c.apiKey)
	} else if c.authMode == AuthModeSSO {
		req.Header.Add("Authorization", "Bearer "+c.apiKey)
	}
	req.Header.Add("User-Agent", c.agent)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if verbose {
		fmt.Printf("status: %s\nraw response:\n%s\n", resp.Status, respBody)
	}
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		if resp.StatusCode == http.StatusUnauthorized {
			return nil, errors.New("authentication failed")
		}
		e := &errorMsg{}
		err := json.Unmarshal(respBody, e)
		if err != nil {
			return nil, err
		}
		return respBody, errors.New(e.Message)
	}
	return respBody, nil
}
