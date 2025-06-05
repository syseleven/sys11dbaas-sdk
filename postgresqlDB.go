package sys11dbaassdk

import (
	"encoding/json"
)

// v1

func (c *Client) GetPostgreSQLDBV1(grq *GetPostgreSQLRequestV1) (*GetPostgreSQLResponseV1, error) {
	path := "/" + grq.Organization + "/" + grq.Project + "/v1/databases/" + grq.UUID
	data, err := c.get(path, grq.Verbose)
	if err != nil {
		return nil, err
	}

	responseData := &GetPostgreSQLResponseV1{}
	err = json.Unmarshal(data, responseData)

	return responseData, err
}

func (c *Client) ListPostgreSQLDBsV1(grq *GetPostgreSQLsRequestV1) ([]*GetPostgreSQLResponseV1, error) {
	path := "/" + grq.Organization + "/" + grq.Project + "/v1/databases"
	data, err := c.get(path, grq.Verbose)
	if err != nil {
		return nil, err
	}

	responseData := make([]*GetPostgreSQLResponseV1, 0)
	err = json.Unmarshal(data, &responseData)

	return responseData, err
}

func (c *Client) CreatePostgreSQLDBV1(crq *CreatePostgreSQLRequestV1) (*CreatePostgreSQLResponseV1, error) {
	path := "/" + crq.Organization + "/" + crq.Project + "/v1/databases"

	data, err := json.Marshal(crq)
	if err != nil {
		return nil, err
	}

	responseData, err := c.post(path, data, crq.Verbose)
	if err != nil {
		return nil, err
	}

	response := &CreatePostgreSQLResponseV1{}
	err = json.Unmarshal(responseData, response)

	return response, err
}

func (c *Client) DeletePostgreSQLDBV1(drq *DeletePostgreSQLRequestV1) (*DeletePostgreSQLResponseV1, error) {
	path := "/" + drq.Organization + "/" + drq.Project + "/v1/databases/" + drq.UUID
	_, err := c.delete(path, drq.Verbose)
	if err != nil {
		return nil, err
	}

	response := &DeletePostgreSQLResponseV1{}

	return response, err
}

func (c *Client) UpdatePostgreSQLDBV1(urq *UpdatePostgreSQLRequestV1) (*UpdatePostgreSQLResponseV1, error) {
	path := "/" + urq.Organization + "/" + urq.Project + "/v1/databases/" + urq.UUID

	data, err := json.Marshal(urq)
	if err != nil {
		return nil, err
	}

	_, err = c.patch(path, data, urq.Verbose)
	if err != nil {
		return nil, err
	}

	response := &UpdatePostgreSQLResponseV1{}

	return response, err
}

// v2

func (c *Client) GetPostgreSQLDBV2(grq *GetPostgreSQLRequestV2) (*GetPostgreSQLResponseV2, error) {
	path := "/" + grq.Organization + "/" + grq.Project + "/v2/databases/" + grq.UUID
	data, err := c.get(path, grq.Verbose)
	if err != nil {
		return nil, err
	}

	responseData := &GetPostgreSQLResponseV2{}
	err = json.Unmarshal(data, responseData)

	return responseData, err
}

func (c *Client) ListPostgreSQLDBsV2(grq *GetPostgreSQLsRequestV2) ([]*GetPostgreSQLResponseV2, error) {
	path := "/" + grq.Organization + "/" + grq.Project + "/v2/databases"
	data, err := c.get(path, grq.Verbose)
	if err != nil {
		return nil, err
	}

	responseData := make([]*GetPostgreSQLResponseV2, 0)
	err = json.Unmarshal(data, &responseData)

	return responseData, err
}

func (c *Client) CreatePostgreSQLDBV2(crq *CreatePostgreSQLRequestV2) (*CreatePostgreSQLResponseV2, error) {
	path := "/" + crq.Organization + "/" + crq.Project + "/v2/databases"

	data, err := json.Marshal(crq)
	if err != nil {
		return nil, err
	}

	responseData, err := c.post(path, data, crq.Verbose)
	if err != nil {
		return nil, err
	}

	response := &CreatePostgreSQLResponseV2{}
	err = json.Unmarshal(responseData, response)

	return response, err
}

func (c *Client) DeletePostgreSQLDBV2(drq *DeletePostgreSQLRequestV2) (*DeletePostgreSQLResponseV2, error) {
	path := "/" + drq.Organization + "/" + drq.Project + "/v2/databases/" + drq.UUID
	_, err := c.delete(path, drq.Verbose)
	if err != nil {
		return nil, err
	}

	response := &DeletePostgreSQLResponseV2{}

	return response, err
}

func (c *Client) UpdatePostgreSQLDBV2(urq *UpdatePostgreSQLRequestV2) (*UpdatePostgreSQLResponseV2, error) {
	path := "/" + urq.Organization + "/" + urq.Project + "/v2/databases/" + urq.UUID

	data, err := json.Marshal(urq)
	if err != nil {
		return nil, err
	}

	_, err = c.patch(path, data, urq.Verbose)
	if err != nil {
		return nil, err
	}

	response := &UpdatePostgreSQLResponseV2{}

	return response, err
}
