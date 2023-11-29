package sys11dbaassdk

import (
	"encoding/json"
)

func (c *Client) GetPostgreSQLDB(grq *GetPostgreSQLRequest) (*GetPostgreSQLResponse, error) {
	path := "/" + grq.Organization + "/" + grq.Project + "/v1/databases/" + grq.UUID
	data, err := c.get(path)
	if err != nil {
		return nil, err
	}

	responseData := &GetPostgreSQLResponse{}
	err = json.Unmarshal(data, responseData)

	return responseData, err
}

func (c *Client) ListPostgreSQLDBs(grq *GetPostgreSQLsRequest) ([]*GetPostgreSQLResponse, error) {
	path := "/" + grq.Organization + "/" + grq.Project + "/v1/databases"
	data, err := c.get(path)
	if err != nil {
		return nil, err
	}

	responseData := make([]*GetPostgreSQLResponse, 0)
	err = json.Unmarshal(data, &responseData)

	return responseData, err
}

func (c *Client) CreatePostgreSQLDB(crq *CreatePostgreSQLRequest) (*CreatePostgreSQLResponse, error) {
	path := "/" + crq.Organization + "/" + crq.Project + "/v1/databases"

	data, err := json.Marshal(crq)
	if err != nil {
		return nil, err
	}

	responseData, err := c.post(path, data)
	if err != nil {
		return nil, err
	}

	response := &CreatePostgreSQLResponse{}
	err = json.Unmarshal(responseData, response)

	return response, err
}

func (c *Client) DeletePostgreSQLDB(drq *DeletePostgreSQLRequest) (*DeletePostgreSQLResponse, error) {
	path := "/" + drq.Organization + "/" + drq.Project + "/v1/databases/" + drq.UUID
	_, err := c.delete(path)
	if err != nil {
		return nil, err
	}

	response := &DeletePostgreSQLResponse{}

	return response, err
}

func (c *Client) UpdatePostgreSQLDB(urq *UpdatePostgreSQLRequest) (*UpdatePostgreSQLResponse, error) {
	path := "/" + urq.Organization + "/" + urq.Project + "/v1/databases/" + urq.UUID

	data, err := json.Marshal(urq)
	if err != nil {
		return nil, err
	}

	_, err = c.patch(path, data)
	if err != nil {
		return nil, err
	}

	response := &UpdatePostgreSQLResponse{}

	return response, err
}
